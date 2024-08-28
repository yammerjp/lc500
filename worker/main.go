package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	v8 "rogchap.com/v8go"
)

func main() {
	// http serverを起動する

	http.HandleFunc("/worker", func(w http.ResponseWriter, r *http.Request) {
		iso := v8.NewIsolate()
		ctx := v8.NewContext(iso)

		requestBodyJson, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		requestBody := map[string]interface{}{}
		if err = json.Unmarshal(requestBodyJson, &requestBody); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		}

		scriptStr, ok := requestBody["script"].(string)
		if !ok {
			http.Error(w, "Failed to get script", http.StatusInternalServerError)
		}
		script, err := iso.CompileUnboundScript(scriptStr, "main.js", v8.CompileOptions{})
		if err != nil {
			http.Error(w, "Failed to compile script", http.StatusInternalServerError)
		}

		val, err := script.Run(ctx)
		if err != nil {
			http.Error(w, "Failed to run script", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val.String()))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("HTTPサーバーを起動しています...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTPサーバーの起動に失敗しました: %v\n", err)
	}
}
