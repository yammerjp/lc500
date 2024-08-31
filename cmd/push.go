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
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalf("Failed to get file: %v", err)
		}
		if file == "" {
			log.Fatalf("File is required")
		}
		script, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		hostname, err := cmd.Flags().GetString("hostname")
		if err != nil {
			log.Fatalf("Failed to get hostname: %v", err)
		}
		if hostname == "" {
			log.Fatalf("Hostname is required")
		}

		region, ok := os.LookupEnv("AWS_REGION")
		if !ok {
			log.Fatalf("AWS_REGION is not set")
		}
		bucket, ok := os.LookupEnv("AWS_BUCKET")
		if !ok {
			log.Fatalf("AWS_BUCKET is not set")
		}
		endpoint, ok := os.LookupEnv("AWS_ENDPOINT")
		if !ok {
			log.Fatalf("AWS_ENDPOINT is not set")
		}
		awsConfig, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
		)
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}
		s3Client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(fmt.Sprintf("%s/index.js", hostname)),
			Body:   bytes.NewReader(script),
		})
		if err != nil {
			log.Fatalf("Failed to put object: %v", err)
		}
		fmt.Println("Pushed")
	},
}

func init() {
	scriptCmd.AddCommand(pushCmd)

	pushCmd.PersistentFlags().String("file", "", "Path to the file to push")
	pushCmd.PersistentFlags().String("hostname", "", "Hostname to push")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
