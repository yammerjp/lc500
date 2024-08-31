package vm

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	v8 "rogchap.com/v8go"
)

type VM struct {
	iso            *v8.Isolate
	ctx            *v8.Context
	script         *v8.UnboundScript
	responseWriter http.ResponseWriter
}

type VMContext struct {
	httpRequest    *http.Request
	injectedParams map[string]string
}

func NewVM() *VM {
	return &VM{
		iso: v8.NewIsolate(),
	}
}

func NewVMContext(r *http.Request, injectedParams map[string]string) *VMContext {
	return &VMContext{
		httpRequest:    r,
		injectedParams: injectedParams,
	}
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

func (vm *VM) SetContext(vmCtx *VMContext) error {
	if vm.iso == nil {
		return errors.New("isolate not initialized")
	}
	if vm.ctx != nil {
		return errors.New("context already set")
	}
	globalThis := v8.NewObjectTemplate(vm.iso)
	err := globalThis.Set("readHeader", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		val, err := v8.NewValue(vm.iso, vmCtx.httpRequest.Header.Get(info.Args()[0].String()))
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
	}))
	if err != nil {
		return err
	}

	err = globalThis.Set("readBody", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		body, err := io.ReadAll(vmCtx.httpRequest.Body)
		if err != nil {
			slog.Error("failed to read body", "error", err)
			str, err := v8.NewValue(vm.iso, "failed to read body")
			if err != nil {
				slog.Error("failed to create error value, disposing isolate", "error", err)
				vm.iso.Dispose()
				return nil
			}
			vm.iso.ThrowException(str)
			return nil
		}
		val, err := v8.NewValue(vm.iso, string(body))
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
	}))
	if err != nil {
		return err
	}

	err = globalThis.Set("readInjectedParam", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		key := info.Args()[0].String()
		slog.Info("reading injected param", "key", key)
		val, ok := vmCtx.injectedParams[key]
		if !ok {
			slog.Error("injected param not found", "key", key)
			return nil
		}
		v8val, err := v8.NewValue(vm.iso, val)
		if err != nil {
			slog.Error("failed to create value", "error", err)
			return nil
		}
		return v8val
	}))
	if err != nil {
		return err
	}

	err = globalThis.Set("setStatus", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.WriteHeader(int(info.Args()[0].Int32()))
		return nil
	}))
	if err != nil {
		return err
	}

	err = globalThis.Set("setHeader", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.Header().Set(info.Args()[0].String(), info.Args()[1].String())
		return nil
	}))
	if err != nil {
		return err
	}

	err = globalThis.Set("renderBody", v8.NewFunctionTemplate(vm.iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		vm.responseWriter.Write([]byte(info.Args()[0].String()))
		return nil
	}))
	if err != nil {
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
