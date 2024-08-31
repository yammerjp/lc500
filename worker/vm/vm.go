package vm

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	api "github.com/yammerjp/lc500/proto/api/v1"
	v8 "rogchap.com/v8go"
)

type VM struct {
	iso            *v8.Isolate
	ctx            *v8.Context
	script         *v8.UnboundScript
	responseWriter http.ResponseWriter
}

type VMContext struct {
	httpRequest  *api.HttpRequest
	httpResponse *api.HttpResponse
}

func NewVM() *VM {
	return &VM{
		iso: v8.NewIsolate(),
	}
}

func NewVMContext(req *http.Request, res *http.Response) (*VMContext, error) {
	httpRequestHeaders := make(map[string]*api.HeaderValue)
	for key, values := range req.Header {
		httpRequestHeaders[key] = &api.HeaderValue{
			Values: values,
		}
	}
	httpRequestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	httpRequest := api.HttpRequest{
		Method:  req.Method,
		Url:     req.URL.String(),
		Body:    string(httpRequestBody),
		Headers: httpRequestHeaders,
	}
	httpResponseHeaders := make(map[string]*api.HeaderValue)
	for key, values := range res.Header {
		httpResponseHeaders[key] = &api.HeaderValue{
			Values: values,
		}
	}
	httpResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	httpResponse := api.HttpResponse{
		StatusCode: int32(res.StatusCode),
		Body:       string(httpResponseBody),
		Headers:    httpResponseHeaders,
	}

	return &VMContext{
		httpRequest:  &httpRequest,
		httpResponse: &httpResponse,
	}, nil
}

func (vm *VM) CompileScript(scriptStr string) error {
	if vm.iso == nil {
		return errors.New("isolate not initialized")
	}
	if vm.script != nil {
		return errors.New("script already compiled")
	}
	script, err := vm.iso.CompileUnboundScript(string(scriptStr), "main.js", v8.CompileOptions{})
	if err != nil {
		return err
	}
	vm.script = script
	return nil
}

func (vm *VM) SetContext(setContextReq *api.SetContextRequest) error {
	if vm.iso == nil {
		return errors.New("isolate not initialized")
	}
	if vm.ctx != nil {
		return errors.New("context already set")
	}
	globalThis := v8.NewObjectTemplate(vm.iso)
	if err := globalThis.Set("readRequestUrl", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, setContextReq.HttpRequest.Url)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		return val
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readRequestMethod", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, setContextReq.HttpRequest.Method)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		return val
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readRequestHeader", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		values := setContextReq.HttpRequest.Headers[info.Args()[0].String()]
		if values == nil {
			return nil
		}
		val, err := v8.NewValue(vm.iso, values.Values[0])
		if err != nil {
			slog.Error("failed to create value", "error", err)
			return nil
		}
		return val
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readRequestHeaders", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		headersTmpl := v8.NewObjectTemplate(vm.iso)
		headers, err := headersTmpl.NewInstance(vm.ctx)
		if err != nil {
			slog.Error("failed to create headers", "error", err)
			return nil
		}
		for key, values := range setContextReq.HttpRequest.Headers {
			array := v8.NewObjectTemplate(vm.iso)
			o, err := array.NewInstance(vm.ctx)
			if err != nil {
				slog.Error("failed to create array", "error", err)
				return nil
			}
			for i, value := range values.Values {
				val, err := v8.NewValue(vm.iso, value)
				if err != nil {
					slog.Error("failed to create value", "error", err)
					return nil
				}
				err = o.SetIdx(uint32(i), val)
				if err != nil {
					slog.Error("failed to set value", "error", err)
					return nil
				}
			}
			headers.Set(key, o.Object().Value)
		}
		return headers.Object().Value
	})); err != nil {
		return err
	}

	if err := globalThis.Set("readRequestBody", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, setContextReq.HttpRequest.Body)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		return val
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readResponseStatusCode", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, setContextReq.HttpResponse.StatusCode)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to create value")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		return val
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readResponseHeader", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		values := setContextReq.HttpResponse.Headers[info.Args()[0].String()]
		array := v8.NewObjectTemplate(vm.iso)
		o, err := array.NewInstance(vm.ctx)
		if err != nil {
			slog.Error("failed to create array", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to create array")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		for i, value := range values.Values {
			val, err := v8.NewValue(vm.iso, value)
			if err != nil {
				slog.Error("failed to create value", "error", err)
				str, err := v8.NewValue(vm.iso, "failed to create value")
				if err != nil {
					slog.Error("failed to create error value, disposing isolate", "error", err)
					vm.iso.Dispose()
					return nil
				}
				vm.iso.ThrowException(str)
				return nil
			}
			err = o.SetIdx(uint32(i), val)
			if err != nil {
				slog.Error("failed to set value", "error", err)
				return nil
			}
		}
		return o.Object().Value
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readResponseHeaders", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		headersTmpl := v8.NewObjectTemplate(vm.iso)
		headers, err := headersTmpl.NewInstance(vm.ctx)
		for key, values := range setContextReq.HttpResponse.Headers {
			array := v8.NewObjectTemplate(vm.iso)
			o, err := array.NewInstance(vm.ctx)
			if err != nil {
				slog.Error("failed to create array", "error", err)
				return nil
			}
			for i, value := range values.Values {
				val, err := v8.NewValue(vm.iso, value)
				if err != nil {
					slog.Error("failed to create value", "error", err)
					return nil
				}
				err = o.SetIdx(uint32(i), val)
				if err != nil {
					slog.Error("failed to set value", "error", err)
					return nil
				}
			}
			headers.Set(key, o.Object().Value)
		}
		if err != nil {
			slog.Error("failed to create headers", "error", err)
			return nil
		}
		return headers.Object().Value
	})); err != nil {
		return err
	}
	if err := globalThis.Set("readResponseBody", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, setContextReq.HttpResponse.Body)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			return nil
		}
		return val
	})); err != nil {
		return err
	}

	if err := globalThis.Set("setStatus", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.WriteHeader(int(info.Args()[0].Int32()))
		return nil
	})); err != nil {
		return err
	}

	if err := globalThis.Set("setHeader", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.Header().Set(info.Args()[0].String(), info.Args()[1].String())
		return nil
	})); err != nil {
		return err
	}

	if err := globalThis.Set("setBody", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.Write([]byte(info.Args()[0].String()))
		return nil
	})); err != nil {
		return err
	}

	vm.ctx = v8.NewContext(vm.iso, globalThis)
	return nil
}

func (vm *VM) SetResponseWriter(w http.ResponseWriter) {
	vm.responseWriter = w
}

func (vm *VM) Dispose() {
	vm.iso.Dispose()
}

func (vm *VM) RunScript() error {
	if vm.ctx == nil {
		return errors.New("context not set")
	}
	if vm.script == nil {
		return errors.New("script not compiled")
	}
	if vm.responseWriter == nil {
		return errors.New("response writer not set")
	}
	val, err := vm.script.Run(vm.ctx)
	if err != nil {
		return err
	}
	slog.Info("script ran", "value", val)
	return nil
}
