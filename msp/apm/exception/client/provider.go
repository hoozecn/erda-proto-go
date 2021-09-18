// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: exception.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/exception/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.apm.exception",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                = reflect.TypeOf((*Client)(nil)).Elem()
	exceptionServiceClientType = reflect.TypeOf((*pb.ExceptionServiceClient)(nil)).Elem()
	exceptionServiceServerType = reflect.TypeOf((*pb.ExceptionServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.apm.exception-client":
		return p.client
	case "erda.msp.apm.exception.ExceptionService":
		return &exceptionServiceWrapper{client: p.client.ExceptionService(), opts: opts}
	case "erda.msp.apm.exception.ExceptionService.client":
		return p.client.ExceptionService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case exceptionServiceClientType:
		return p.client.ExceptionService()
	case exceptionServiceServerType:
		return &exceptionServiceWrapper{client: p.client.ExceptionService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.exception-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.apm.exception.ExceptionService",
			"erda.msp.apm.exception-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			exceptionServiceClientType,
			// server types
			exceptionServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
