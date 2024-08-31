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
	"strings"

	"github.com/spf13/cobra"
	api "github.com/yammerjp/lc500/proto/api/v1"
)

// setContextCmd represents the setContext command
var setContextCmd = &cobra.Command{
	Use:   "setContext",
	Short: "Set the context of the VM",
	Long:  `Set the context of the VM`,
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

		method, err := cmd.Flags().GetString("method")
		if err != nil {
			log.Fatalf("failed to get method: %v", err)
			os.Exit(1)
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("failed to get url: %v", err)
			os.Exit(1)
		}
		headersSlice, err := cmd.Flags().GetStringSlice("headers")
		if err != nil {
			log.Fatalf("failed to get headers: %v", err)
			os.Exit(1)
		}
		headers := make(map[string]string)
		for _, header := range headersSlice {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) != 2 {
				log.Fatalf("invalid header format: %s", header)
				os.Exit(1)
			}
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		body, err := cmd.Flags().GetString("body")
		if err != nil {
			log.Fatalf("failed to get body: %v", err)
			os.Exit(1)
		}
		additionalContext, err := cmd.Flags().GetString("additionalContext")
		if err != nil {
			log.Fatalf("failed to get additionalContext: %v", err)
			os.Exit(1)
		}

		contextReqest := api.SetContextRequest{
			Vmid: vmid,
		}

		if method != "" && url != "" {
			contextReqest.HttpRequestMethod = method
			contextReqest.HttpRequestUrl = url
			contextReqest.HttpRequestHeaders = headers
			contextReqest.HttpRequestBody = body
		}
		if additionalContext != "" {
			contextReqest.AdditionalContext = additionalContext
		}

		_, err = client.SetContext(context.Background(), &contextReqest)
		if err != nil {
			log.Fatalf("failed to set context: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	vmCmd.AddCommand(setContextCmd)

	setContextCmd.Flags().StringP("target", "t", "localhost:8080", "target address")
	setContextCmd.Flags().BoolP("insecure", "i", false, "use insecure connection")
	setContextCmd.Flags().StringP("format", "f", "text", "output format(json or text)")

	setContextCmd.Flags().String("vmid", "", "vm id")
	setContextCmd.Flags().String("vmidfile", "", "file to read vm id")

	setContextCmd.Flags().String("method", "", "HTTP method")
	setContextCmd.Flags().String("url", "", "HTTP URL")
	setContextCmd.Flags().StringSlice("headers", []string{}, "HTTP headers")
	setContextCmd.Flags().String("body", "", "HTTP body")
	setContextCmd.Flags().String("additionalContext", "", "additional context")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
