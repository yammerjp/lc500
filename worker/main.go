package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	v8 "rogchap.com/v8go"
)

// ロガーを整える

func workerHandler(w http.ResponseWriter, r *http.Request) {
	iso := v8.NewIsolate()
	globalThis := v8.NewObjectTemplate(iso)

	globalThis.Set("readHeader", v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(iso, r.Header.Get(info.Args()[0].String()))
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
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
			slog.Error("failed to read body", "error", err)
			str, err := v8.NewValue(iso, "failed to read body")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				iso.Dispose()
				return nil
			}
			iso.ThrowException(str)
			return nil
		}
		val, err := v8.NewValue(iso, string(body))
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				iso.Dispose()
				return nil
			}
			iso.ThrowException(str)
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
		slog.Error("failed to read script", "error", err)
		http.Error(w, "Failed to read script", http.StatusInternalServerError)
		return
	}

	script, err := iso.CompileUnboundScript(string(scriptStr), "main.js", v8.CompileOptions{})
	if err != nil {
		slog.Error("failed to compile script", "error", err)
		http.Error(w, "Failed to compile script", http.StatusInternalServerError)
		return
	}

	val, err := script.Run(ctx)
	if err != nil {
		slog.Error("failed to run script", "error", err)
		http.Error(w, "Failed to run script", http.StatusInternalServerError)
		return
	}

	slog.Info("script ran", "value", val)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	http.HandleFunc("/worker", workerHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request received", "method", r.Method, "url", r.URL.String())
		w.Write([]byte("Hello, World!"))
	})

	slog.Info("start HTTP server on port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("failed to start HTTP server", "error", err)
	}
}
