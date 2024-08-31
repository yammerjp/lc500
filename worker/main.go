package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/yammerjp/lc500/worker/pool"
	"github.com/yammerjp/lc500/worker/response"
	lc500Vm "github.com/yammerjp/lc500/worker/vm"
)

func main() {
	pool := pool.NewPool()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	http.HandleFunc("/vm/init", func(w http.ResponseWriter, r *http.Request) {
		vmid := pool.Generate()
		w.Write([]byte(vmid))
	})

	http.HandleFunc("/vm/dispose", func(w http.ResponseWriter, r *http.Request) {
		pool.Dispose(r.URL.Query().Get("vmid"))
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/vm/compile", func(w http.ResponseWriter, r *http.Request) {
		script, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("failed to read script", "error", err)
			http.Error(w, "Failed to read script", http.StatusInternalServerError)
			return
		}
		err = pool.CompileScript(r.URL.Query().Get("vmid"), string(script))
		if err != nil {
			slog.Error("failed to compile script", "error", err)
			http.Error(w, "Failed to compile script", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/vm/setcontext", func(w http.ResponseWriter, r *http.Request) {
		vmid := r.URL.Query().Get("vmid")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("failed to read body", "error", err)
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		type VmCtxReq struct {
			HttpRequest struct {
				Method  string            `json:"method"`
				URL     string            `json:"url"`
				Headers map[string]string `json:"headers"`
				Body    string            `json:"body"`
			} `json:"httpRequest"`
			AdditionalContext map[string]string `json:"additionalContext"`
		}

		// json parse
		var vmCtxReq VmCtxReq
		err = json.Unmarshal(body, &vmCtxReq)
		if err != nil {
			slog.Error("failed to parse vm context", "error", err)
			http.Error(w, "Failed to parse vm context", http.StatusInternalServerError)
			return
		}

		req, err := http.NewRequest(vmCtxReq.HttpRequest.Method, vmCtxReq.HttpRequest.URL, strings.NewReader(vmCtxReq.HttpRequest.Body))
		if err != nil {
			slog.Error("failed to create request", "error", err)
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		for k, v := range vmCtxReq.HttpRequest.Headers {
			req.Header.Set(k, v)
		}

		vmCtx := lc500Vm.NewVMContext(req, vmCtxReq.AdditionalContext)
		err = pool.SetContext(vmid, vmCtx)
		if err != nil {
			slog.Error("failed to set context", "error", err)
			http.Error(w, "Failed to set context", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("ok"))
	})

	http.HandleFunc("/vm/run", func(w http.ResponseWriter, r *http.Request) {
		vmid := r.URL.Query().Get("vmid")

		receiver := response.NewReciever()
		err := pool.Run(vmid, receiver)
		if err != nil {
			slog.Error("failed to run script", "error", err)
			http.Error(w, "Failed to run script", http.StatusInternalServerError)
			return
		}
		pool.Dispose(vmid)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(receiver.ToJSON()))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request received", "method", r.Method, "url", r.URL.String())
		w.Write([]byte("Hello, World!"))
	})

	slog.Info("start HTTP server on port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("failed to start HTTP server", "error", err)
	}
}
