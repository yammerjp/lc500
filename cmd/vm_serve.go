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
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	api "github.com/yammerjp/lc500/proto/api/v1"
	"github.com/yammerjp/lc500/worker/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// port option

// serveCmd represents the serve command
var vmServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the v8 isolate worker server",
	Long:  `Start the v8 isolate worker server`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

		// port
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()
		api.RegisterWorkerServer(s, server.NewServer())
		if cmd.Flags().Changed("reflection") {
			reflection.Register(s)
		}

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			slog.Error("failed to listen", "error", err)
			os.Exit(1)
		}
		defer listener.Close()

		go func() {
			slog.Info(fmt.Sprintf("start gRPC server on port %d", port))
			if err := s.Serve(listener); err != nil {
				slog.Error("failed to serve", "error", err)
				os.Exit(1)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		slog.Info("stopping gRPC server...")
		s.GracefulStop()
		slog.Info("gRPC server stopped")
	},
}

func init() {
	workerCmd.AddCommand(vmServeCmd)
	vmServeCmd.Flags().IntP("port", "p", 8081, "The port to listen on")
	// reflection option
	vmServeCmd.Flags().BoolP("reflection", "r", false, "Enable reflection")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
