// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: runtime.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterRuntimeServiceImp runtime.proto
func RegisterRuntimeServiceImp(regester transport.Register, srv RuntimeServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterRuntimeServiceHandler(regester, RuntimeServiceHandler(srv), _ops.HTTP...)
	RegisterRuntimeServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.orchestrator.runtime.RuntimeService",
	)
}

var (
	runtimeServiceClientType  = reflect.TypeOf((*RuntimeServiceClient)(nil)).Elem()
	runtimeServiceServerType  = reflect.TypeOf((*RuntimeServiceServer)(nil)).Elem()
	runtimeServiceHandlerType = reflect.TypeOf((*RuntimeServiceHandler)(nil)).Elem()
)

// RuntimeServiceClientType .
func RuntimeServiceClientType() reflect.Type { return runtimeServiceClientType }

// RuntimeServiceServerType .
func RuntimeServiceServerType() reflect.Type { return runtimeServiceServerType }

// RuntimeServiceHandlerType .
func RuntimeServiceHandlerType() reflect.Type { return runtimeServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		runtimeServiceClientType,
		// server types
		runtimeServiceServerType,
		// handler types
		runtimeServiceHandlerType,
	}
}
