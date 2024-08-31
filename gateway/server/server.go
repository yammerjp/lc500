package server

import (
	"context"
	"net/http"

	"github.com/chebyrash/promise"
)

type Handler struct {
	workerClient     *WorkerClient
	scriptFetcher    *ScriptFetcher
	blueprintFetcher *BlueprintFetcher
}

func NewHandler(workerTarget string, workerInsecure bool, blueprintEndpoint string) (*Handler, error) {
	workerClientWrapper, err := NewWorkerClient(workerTarget, workerInsecure)
	if err != nil {
		return nil, err
	}
	scriptFetcher, err := InitScriptFetcher()
	if err != nil {
		return nil, err
	}
	blueprintFetcher := NewBlueprintFetcher(blueprintEndpoint)

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

func (h *Handler) HandleRequest(r *http.Request) (*WorkerResponse, error) {
	ctx := context.Background()

	promiseScript := promise.New(func(resolve func(string), reject func(error)) {
		script, err := h.scriptFetcher.FetchScript(ctx, r.Host)
		if err != nil {
			reject(err)
		}
		resolve(script)
	})

	promiseBlueprint := promise.New(func(resolve func(string), reject func(error)) {
		res, err := h.blueprintFetcher.FetchBlueprint(r)
		if err != nil {
			reject(err)
		}
		resolve(res)
	})

	workerRequest, err := h.workerClient.NewWorkerRequest(ctx, r)
	if err != nil {
		return nil, err
	}

	script, err := promiseScript.Await(ctx)
	if err != nil {
		return nil, err
	}

	if err := workerRequest.Compile(*script); err != nil {
		return nil, err
	}

	additionalContext, err := promiseBlueprint.Await(ctx)
	if err != nil {
		return nil, err
	}

	return workerRequest.Run(*additionalContext)
}
