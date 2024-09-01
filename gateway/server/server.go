package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/chebyrash/promise"
)

type Handler struct {
	workerClient     *WorkerClient
	scriptFetcher    *ScriptFetcher
	blueprintFetcher *BlueprintFetcher
}

func NewHandler(workerTarget string, workerInsecure bool, blueprintTarget string) (*Handler, error) {
	workerClientWrapper, err := NewWorkerClient(workerTarget, workerInsecure)
	if err != nil {
		return nil, err
	}
	scriptFetcher, err := InitScriptFetcher()
	if err != nil {
		return nil, err
	}
	blueprintFetcher := NewBlueprintFetcher(blueprintTarget)

	return &Handler{
		workerClient:     workerClientWrapper,
		scriptFetcher:    scriptFetcher,
		blueprintFetcher: blueprintFetcher,
	}, nil
}

func (h *Handler) Close() {
	h.workerClient.Close()
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.HandleRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteTo(w)
}

func (h *Handler) HandleRequest(req *http.Request) (*WorkerResponse, error) {
	slog.Info("Handling request", "url", req.URL.String())
	ctx := context.Background()

	promiseScript := promise.New(func(resolve func(string), reject func(error)) {
		hostname, _, err := net.SplitHostPort(req.Host)
		if err != nil {
			reject(err)
		}
		script, err := h.scriptFetcher.FetchScript(ctx, hostname)
		if err != nil {
			reject(err)
		}
		resolve(script)
	})
	slog.Info("async Fetched script", "script", promiseScript)

	promiseBlueprint := promise.New(func(resolve func(*http.Response), reject func(error)) {
		res, err := h.blueprintFetcher.Fetch(req)
		if err != nil {
			reject(err)
		}
		resolve(res)
	})
	slog.Info("async Fetched blueprint", "blueprint", promiseBlueprint)

	workerRequest, err := h.workerClient.NewWorkerRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	slog.Info("created worker request", "workerRequest", workerRequest)

	script, err := promiseScript.Await(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("await fetching script", "script", &script)

	if err := workerRequest.Compile(*script); err != nil {
		return nil, err
	}
	slog.Info("compiled worker request", "workerRequest", workerRequest)

	res, err := promiseBlueprint.Await(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("await fetching blueprint", "blueprint", res)

	workerResponse, err := workerRequest.Run(req, *res)
	if err != nil {
		return nil, err
	}
	slog.Info("worker response", "workerResponse", workerResponse)

	return workerResponse, nil
}
