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

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile a script into a VM",
	Long:  `Compile a script into a VM`,
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

		script, err := cmd.Flags().GetString("script")
		if err != nil {
			log.Fatalf("failed to get script: %v", err)
			os.Exit(1)
		}
		scriptFilePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("failed to get file: %v", err)
			os.Exit(1)
		}
		if script == "" && scriptFilePath == "" {
			log.Fatalf("script or file is required")
			os.Exit(1)
		}
		if script != "" && scriptFilePath != "" {
			log.Fatalf("script and file cannot be used together")
			os.Exit(1)
		}
		if scriptFilePath != "" {
			scriptBytes, err := os.ReadFile(scriptFilePath)
			if err != nil {
				log.Fatalf("failed to read file: %v", err)
				os.Exit(1)
			}
			script = string(scriptBytes)
		}

		_, err = client.Compile(context.Background(), &api.CompileRequest{
			Vmid:   vmid,
			Script: script,
		})
		if err != nil {
			log.Fatalf("failed to compile: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	vmCmd.AddCommand(compileCmd)
	compileCmd.Flags().StringP("target", "t", "localhost:8080", "target address")
	compileCmd.Flags().BoolP("insecure", "i", false, "use insecure connection")
	compileCmd.Flags().StringP("format", "f", "text", "output format(json or text)")

	compileCmd.Flags().String("vmid", "", "vm id")
	compileCmd.Flags().String("vmidfile", "", "file to write vm id")

	compileCmd.Flags().String("script", "", "script to compile")
	compileCmd.Flags().String("file", "", "file to compile")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
