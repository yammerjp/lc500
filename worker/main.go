package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	v8 "rogchap.com/v8go"
)

type WorkerRequest struct {
	Script      string `json:"script"`
	HttpRequest struct {
		Method  string            `json:"method"`
		Url     string            `json:"url"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
	} `json:"httpRequest"`
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
	iso := v8.NewIsolate()

	workerRequestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	workerRequest := WorkerRequest{}
	if err := json.Unmarshal(workerRequestBytes, &workerRequest); err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	method, err := v8.NewValue(iso, workerRequest.HttpRequest.Method)
	if err != nil {
		http.Error(w, "Failed to create method", http.StatusInternalServerError)
		return
	}
	url, err := v8.NewValue(iso, workerRequest.HttpRequest.Url)
	if err != nil {
		http.Error(w, "Failed to create url", http.StatusInternalServerError)
		return
	}
	headers := v8.NewObjectTemplate(iso)
	for key, value := range workerRequest.HttpRequest.Headers {
		headerValue, err := v8.NewValue(iso, value)
		if err != nil {
			http.Error(w, "Failed to create header value", http.StatusInternalServerError)
			return
		}
		if err := headers.Set(key, headerValue); err != nil {
			http.Error(w, "Failed to set header", http.StatusInternalServerError)
			return
		}
	}
	body, err := v8.NewValue(iso, workerRequest.HttpRequest.Body)
	if err != nil {
		http.Error(w, "Failed to create body", http.StatusInternalServerError)
		return
	}

	v8Lc500 := v8.NewObjectTemplate(iso)
	v8Lc500.Set("method", method)
	v8Lc500.Set("url", url)
	v8Lc500.Set("headers", headers)
	v8Lc500.Set("body", body)
	globalThis := v8.NewObjectTemplate(iso)
	globalThis.Set("lc500", v8Lc500)
	ctx := v8.NewContext(iso, globalThis)
	script, err := iso.CompileUnboundScript(workerRequest.Script, "main.js", v8.CompileOptions{})
	if err != nil {
		http.Error(w, "Failed to compile script", http.StatusInternalServerError)
		return
	}

	val, err := script.Run(ctx)
	if err != nil {
		http.Error(w, "Failed to run script", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val.String()))
}

func main() {
	http.HandleFunc("/worker", workerHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("HTTPサーバーを起動しています...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTPサーバーの起動に失敗しました: %v\n", err)
	}
}
