// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: meta.proto, metric.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/metric/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.monitor.metric",
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
	clientsType                 = reflect.TypeOf((*Client)(nil)).Elem()
	metricMetaServiceClientType = reflect.TypeOf((*pb.MetricMetaServiceClient)(nil)).Elem()
	metricMetaServiceServerType = reflect.TypeOf((*pb.MetricMetaServiceServer)(nil)).Elem()
	metricServiceClientType     = reflect.TypeOf((*pb.MetricServiceClient)(nil)).Elem()
	metricServiceServerType     = reflect.TypeOf((*pb.MetricServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.monitor.metric-client":
		return p.client
	case "erda.core.monitor.metric.MetricMetaService":
		return &metricMetaServiceWrapper{client: p.client.MetricMetaService(), opts: opts}
	case "erda.core.monitor.metric.MetricMetaService.client":
		return p.client.MetricMetaService()
	case "erda.core.monitor.metric.MetricService":
		return &metricServiceWrapper{client: p.client.MetricService(), opts: opts}
	case "erda.core.monitor.metric.MetricService.client":
		return p.client.MetricService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case metricMetaServiceClientType:
		return p.client.MetricMetaService()
	case metricMetaServiceServerType:
		return &metricMetaServiceWrapper{client: p.client.MetricMetaService(), opts: opts}
	case metricServiceClientType:
		return p.client.MetricService()
	case metricServiceServerType:
		return &metricServiceWrapper{client: p.client.MetricService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.metric-client", &servicehub.Spec{
		Services: []string{
			"erda.core.monitor.metric.MetricMetaService",
			"erda.core.monitor.metric.MetricService",
			"erda.core.monitor.metric-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			metricMetaServiceClientType,
			metricServiceClientType,
			// server types
			metricMetaServiceServerType,
			metricServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
