/*
Copyright © 2024 Keisuke Nakayama <me@yammer.jp>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	api "github.com/yammerjp/lc500/proto/api/v1"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new VM (V8 Isolate Instance)",
	Long:  `Initialize a new VM (V8 Isolate Instance)`,
	Run: func(cmd *cobra.Command, args []string) {
		vmidfile, err := cmd.Flags().GetString("vmidfile")
		if err != nil {
			log.Fatalf("failed to get vmidfile: %v", err)
			os.Exit(1)
		}

		client, close, err := NewClient(cmd)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
			os.Exit(1)
		}
		defer close()

		res, err := client.InitVM(context.Background(), &api.InitVMRequest{})
		if err != nil {
			log.Fatalf("failed to init vm: %v", err)
			os.Exit(1)
		}

		if vmidfile != "" {
			err := os.WriteFile(vmidfile, []byte(res.Vmid), 0644)
			if err != nil {
				log.Fatalf("failed to write vmidfile: %v", err)
				os.Exit(1)
			}
		}
		fmt.Println(res.Vmid)
	},
}

func init() {
	vmCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("target", "t", "localhost:8080", "target address")
	initCmd.Flags().BoolP("insecure", "i", false, "use insecure connection")
	initCmd.Flags().StringP("format", "f", "text", "output format(json or text)")
	initCmd.Flags().String("vmidfile", "", "vmid file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
