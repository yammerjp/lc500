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
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/yammerjp/lc500/gateway/server"
)

// serveCmd represents the serve command
var gatewayServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the gateway server",
	Long:  `Start the gateway server`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}
		workerTarget, err := cmd.Flags().GetString("worker-target")
		if err != nil {
			panic(err)
		}
		workerInsecure, err := cmd.Flags().GetBool("worker-insecure")
		if err != nil {
			panic(err)
		}
		blueprintEndpoint, err := cmd.Flags().GetString("blueprint-endpoint")
		if err != nil {
			panic(err)
		}

		h, err := server.NewHandler(workerTarget, workerInsecure, blueprintEndpoint)
		if err != nil {
			panic(err)
		}
		defer h.Close()

		err = http.ListenAndServe(fmt.Sprintf(":%d", port), h)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	gatewayCmd.AddCommand(gatewayServeCmd)
	gatewayServeCmd.Flags().IntP("port", "p", 8080, "Port to listen on")

	// worker host worker port worker insecure
	gatewayServeCmd.Flags().StringP("worker-target", "wt", "localhost:8081", "Worker server target")
	gatewayServeCmd.Flags().BoolP("worker-insecure", "wi", false, "Worker server insecure")

	gatewayServeCmd.Flags().StringP("blueprint-endpoint", "be", "http://localhost:8082", "Blueprint server endpoint")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
