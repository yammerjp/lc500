package main

import (
	"fmt"
	"io"
	"net/http"

	v8 "rogchap.com/v8go"
)

func workerHandler(w http.ResponseWriter, r *http.Request) {
	iso := v8.NewIsolate()
	globalThis := v8.NewObjectTemplate(iso)

	globalThis.Set("readHeader", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(iso, r.Header.Get(info.Args()[0].String()))
		if err != nil {
			fmt.Println(err)
			str, err := v8.NewValue(iso, err.Error())
			if err != nil {
				fmt.Println(err)
				iso.Dispose()
				return nil
			}
			iso.ThrowException(str)
			return nil
		}
		return val
	}))

	globalThis.Set("readBody", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			iso.Dispose()
			return nil
		}
		val, err := v8.NewValue(iso, string(body))
		if err != nil {
			fmt.Println(err)
			iso.Dispose()
			return nil
		}
		return val
	}))

	globalThis.Set("setStatus", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		w.WriteHeader(int(info.Args()[0].Int32()))
		return nil
	}))

	globalThis.Set("setHeader", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		w.Header().Set(info.Args()[0].String(), info.Args()[1].String())
		return nil
	}))

	globalThis.Set("renderBody", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		w.Write([]byte(info.Args()[0].String()))
		return nil
	}))
	ctx := v8.NewContext(iso, globalThis)

	scriptStr, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to read script", http.StatusInternalServerError)
		return
	}

	script, err := iso.CompileUnboundScript(string(scriptStr), "main.js", v8.CompileOptions{})
	if err != nil {
		http.Error(w, "Failed to compile script", http.StatusInternalServerError)
		return
	}

	val, err := script.Run(ctx)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to run script", http.StatusInternalServerError)
		return
	}
	fmt.Println(val)
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
