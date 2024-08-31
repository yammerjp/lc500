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
	"log"
	"os"

	"github.com/spf13/cobra"
	api "github.com/yammerjp/lc500/proto/api/v1"
)

// disposeCmd represents the dispose command
var disposeCmd = &cobra.Command{
	Use:   "dispose",
	Short: "Dispose a VM",
	Long:  `Dispose a VM`,
	Run: func(cmd *cobra.Command, args []string) {
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
		if vmid == "" && vmidfile != "" {
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

		client, close, err := NewClient(cmd)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
			os.Exit(1)
		}
		defer close()

		_, err = client.DisposeVM(context.Background(), &api.DisposeVMRequest{
			Vmid: vmid,
		})
		if err != nil {
			log.Fatalf("failed to dispose: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	vmCmd.AddCommand(disposeCmd)

	disposeCmd.Flags().String("vmid", "", "VM ID")
	disposeCmd.Flags().String("vmidfile", "", "file to read VM ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disposeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disposeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
