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
	"net/http"
	"os"

	"github.com/spf13/cobra"
	api "github.com/yammerjp/lc500/proto/api/v1"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a VM",
	Long:  `Run a VM`,
	Run: func(cmd *cobra.Command, args []string) {
		client, close, err := NewClient(cmd)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
			os.Exit(1)
		}
		defer close()

		vmid, err := cmd.Flags().GetString("vmid")
		if err != nil {
			log.Fatalf("failed to get vmid: %v", err)
			os.Exit(1)
		}
		vmidfile, err := cmd.Flags().GetString("vmidfile")
		if err != nil {
			log.Fatalf("failed to get vmidfile: %v", err)
			os.Exit(1)
		}
		if vmidfile != "" {
			vmidBytes, err := os.ReadFile(vmidfile)
			if err != nil {
				log.Fatalf("failed to read vmidfile: %v", err)
				os.Exit(1)
			}
			vmid = string(vmidBytes)
		}
		if vmid == "" {
			log.Fatalf("vmid is required")
			os.Exit(1)
		}

		res, err := client.Run(context.Background(), &api.RunRequest{
			Vmid: vmid,
		})
		if err != nil {
			log.Fatalf("failed to run: %v", err)
			os.Exit(1)
		}
		statusMessage := http.StatusText(int(res.HttpResponseStatusCode))
		fmt.Printf("HTTP/1.1 %d %s\n", res.HttpResponseStatusCode, statusMessage)
		for key, header := range res.HttpResponseHeaders {
			fmt.Printf("%s: %s\n", key, header)
		}
		fmt.Println()

		fmt.Println(res.HttpResponseBody)
	},
}

func init() {
	vmCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("target", "t", "localhost:8080", "target address")
	runCmd.Flags().BoolP("insecure", "i", false, "use insecure connection")
	runCmd.Flags().StringP("format", "f", "text", "output format(json or text)")

	runCmd.Flags().String("vmid", "", "vm id")
	runCmd.Flags().String("vmidfile", "", "file to read vm id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
