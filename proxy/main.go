package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	arg := os.Args[1]
	// Define the target URL for the reverse proxy
	target, err := url.Parse(arg) // Replace with your target URL
	if err != nil {
		log.Fatal(err)
	}

	// Create a new reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Modify the director function to add custom headers
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// req.Header.Add("X-Custom-Header", "CustomValue")
	}

	// Modify the response
	proxy.ModifyResponse = func(resp *http.Response) error {
		oldBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}

		newBody := append([]byte("Hello! rp!\n"), oldBody...)
		resp.Body = io.NopCloser(bytes.NewReader(newBody))
		resp.ContentLength = int64(len(newBody))
		resp.Header.Set("Content-Length", fmt.Sprint(len(newBody)))

		return nil
	}

	// Start the server
	log.Println("Starting reverse proxy server on :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
