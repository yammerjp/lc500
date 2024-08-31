package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/chebyrash/promise"
	workerapi "github.com/yammerjp/lc500/proto/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HandlerBuilder struct {
	WorkerEndpoint    string
	BlueprintEndpoint string
}

func GenKey(r *http.Request) string {
	// TODO fix
	// return fmt.Sprintf("%s/%s", r.Host, r.URL.Path)
	slog.Info("GenKey", "host", r.Host, "path", r.URL.Path)
	return "script.js"
}

func fetchBlueprint(blueprintEndpoint string, r *http.Request) (string, error) {
	originalUrl, err := url.Parse(r.URL.String())
	if err != nil {
		return "", err
	}
	blueprintUrl, err := url.Parse(blueprintEndpoint)
	if err != nil {
		return "", err
	}
	originalUrl.Host = blueprintUrl.Host
	originalUrl.Scheme = blueprintUrl.Scheme

	req, err := http.NewRequest(r.Method, originalUrl.String(), r.Body)
	if err != nil {
		return "", err
	}
	req.Header = r.Header.Clone()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	headers := make(map[string]string)
	for k, v := range res.Header {
		// TODO support multiple values
		headers[k] = v[0]
	}

	type Response struct {
		Status  int               `json:"status"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
	}

	response := Response{
		Status:  res.StatusCode,
		Headers: headers,
		Body:    body,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func NewHandler(b *HandlerBuilder) (http.Handler, func(), error) {
	workerEndpointUrl, err := url.Parse(b.WorkerEndpoint)
	if err != nil {
		return nil, nil, err
	}
	if workerEndpointUrl.Scheme != "http" && workerEndpointUrl.Scheme != "https" {
		return nil, nil, fmt.Errorf("invalid scheme: %s", workerEndpointUrl.Scheme)
	}
	if workerEndpointUrl.Host == "" {
		return nil, nil, fmt.Errorf("host is empty")
	}
	if workerEndpointUrl.Path != "" {
		return nil, nil, fmt.Errorf("path is not empty: %s", workerEndpointUrl.Path)
	}
	dialOptions := []grpc.DialOption{}
	// if workerEndpointUrl.Scheme == "http" {
	dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// }
	slog.Info("workerEndpointUrl", "url", workerEndpointUrl)
	target := "worker:8080"
	// if workerEndpointUrl.Port() == "" {
	// 	target = workerEndpointUrl.Host
	// } else {
	// 	target = fmt.Sprintf("%s:%s", workerEndpointUrl.Host, workerEndpointUrl.Port())
	// }
	cc, err := grpc.NewClient(target, dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	close := func() {
		cc.Close()
	}
	workerClient := workerapi.NewWorkerClient(cc)

	keyGenerator := GenKey

	region, ok := os.LookupEnv("AWS_REGION")
	if !ok {
		return nil, nil, fmt.Errorf("AWS_REGION is not set")
	}
	bucket, ok := os.LookupEnv("AWS_BUCKET")
	if !ok {
		return nil, nil, fmt.Errorf("AWS_BUCKET is not set")
	}
	endpoint, ok := os.LookupEnv("AWS_ENDPOINT")
	if !ok {
		return nil, nil, fmt.Errorf("AWS_ENDPOINT is not set")
	}
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, nil, err
	}
	s3Client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	scriptFetcher := func(r *http.Request) (string, error) {
		obj, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(keyGenerator(r)),
		})
		if err != nil {
			return "", err
		}
		bodyBytes, err := io.ReadAll(obj.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}
	blueprintFetcher := fetchBlueprint

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		promiseScript := promise.New(func(resolve func(string), reject func(error)) {
			script, err := scriptFetcher(r)
			if err != nil {
				reject(err)
			}
			resolve(script)
		})
		promiseBlueprint := promise.New(func(resolve func(string), reject func(error)) {
			res, err := blueprintFetcher(b.BlueprintEndpoint, r)
			if err != nil {
				reject(err)
			}
			resolve(res)
		})

		resInitVM, err := workerClient.InitVM(ctx, &workerapi.InitVMRequest{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vmid := resInitVM.Vmid

		script, err := promiseScript.Await(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = workerClient.Compile(ctx, &workerapi.CompileRequest{
			Vmid:   vmid,
			Script: *script,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		headers := make(map[string]string)
		for k, v := range r.Header {
			// TODO support multiple values
			headers[k] = v[0]
		}
		additionalContext, err := promiseBlueprint.Await(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = workerClient.SetContext(ctx, &workerapi.SetContextRequest{
			Vmid:               vmid,
			HttpRequestMethod:  r.Method,
			HttpRequestUrl:     r.URL.String(),
			HttpRequestBody:    string(body),
			HttpRequestHeaders: headers,
			AdditionalContext:  *additionalContext,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resRun, err := workerClient.Run(ctx, &workerapi.RunRequest{
			Vmid:    vmid,
			Dispose: true,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(int(resRun.HttpResponseStatusCode))
		for k, v := range resRun.HttpResponseHeaders {
			w.Header().Set(k, v)
		}
		w.Write([]byte(resRun.HttpResponseBody))
	}), close, nil
}
