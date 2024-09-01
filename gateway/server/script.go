package server

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type ScriptFetcher struct {
	s3Client *s3.Client
	bucket   string
}

func NewScriptFetcher(s3Client *s3.Client, bucket string) *ScriptFetcher {
	return &ScriptFetcher{
		s3Client: s3Client,
		bucket:   bucket,
	}
}

func InitScriptFetcher() (*ScriptFetcher, error) {
	region, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		return nil, fmt.Errorf("AWS_REGION is not set")
	}
	bucket, ok := os.LookupEnv("AWS_BUCKET")
	if !ok {
		return nil, fmt.Errorf("AWS_BUCKET is not set")
	}
	endpoint, ok := os.LookupEnv("AWS_ENDPOINT")
	if !ok {
		return nil, fmt.Errorf("AWS_ENDPOINT is not set")
	}
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	return NewScriptFetcher(s3Client, bucket), nil
}

func (s *ScriptFetcher) FetchScript(ctx context.Context, hostname string) (string, error) {
	key := fmt.Sprintf("%s/index.js", hostname)
	slog.Debug("Fetching script", "key", key)
	res, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
