package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	lc500Vm "github.com/yammerjp/lc500/worker/vm"
)

func workerHandler(w http.ResponseWriter, r *http.Request) {
	vm := lc500Vm.NewVM()

	scriptStr, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("failed to read script", "error", err)
		http.Error(w, "Failed to read script", http.StatusInternalServerError)
		return
	}
	vm.CompileScript(string(scriptStr))
	vm.SetContext(lc500Vm.NewVMContext(r))
	vm.SetResponseWriter(w)
	vm.RunScript()
	vm.Dispose()
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
