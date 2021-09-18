// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: registercenter.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/registercenter/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// RegisterCenterService registercenter.proto
	RegisterCenterService() pb.RegisterCenterServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		registerCenterService: pb.NewRegisterCenterServiceClient(cc),
	}
}

type serviceClients struct {
	registerCenterService pb.RegisterCenterServiceClient
}

func (c *serviceClients) RegisterCenterService() pb.RegisterCenterServiceClient {
	return c.registerCenterService
}

type registerCenterServiceWrapper struct {
	client pb.RegisterCenterServiceClient
	opts   []grpc1.CallOption
}

func (s *registerCenterServiceWrapper) ListInterface(ctx context.Context, req *pb.ListInterfaceRequest) (*pb.ListInterfaceResponse, error) {
	return s.client.ListInterface(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetHTTPServices(ctx context.Context, req *pb.GetHTTPServicesRequest) (*pb.GetHTTPServicesResponse, error) {
	return s.client.GetHTTPServices(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) EnableHTTPService(ctx context.Context, req *pb.EnableHTTPServiceRequest) (*pb.EnableHTTPServiceResponse, error) {
	return s.client.EnableHTTPService(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetServiceIpInfo(ctx context.Context, req *pb.ServiceIpRequest) (*pb.ServiceIpInfoResponse, error) {
	return s.client.GetServiceIpInfo(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetRouteRule(ctx context.Context, req *pb.GetRouteRuleRequest) (*pb.GetRouteRuleResponse, error) {
	return s.client.GetRouteRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) CreateRouteRule(ctx context.Context, req *pb.CreateRouteRuleRequest) (*pb.CreateRouteRuleResponse, error) {
	return s.client.CreateRouteRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) DeleteRouteRule(ctx context.Context, req *pb.DeleteRouteRuleRequest) (*pb.DeleteRouteRuleResponse, error) {
	return s.client.DeleteRouteRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetHostRule(ctx context.Context, req *pb.CetHostRuleRequest) (*pb.CetHostRuleResponse, error) {
	return s.client.GetHostRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) CreateHostRule(ctx context.Context, req *pb.CreateHostRuleRequest) (*pb.CreateHostRuleResponse, error) {
	return s.client.CreateHostRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) DeleteHostRule(ctx context.Context, req *pb.DeleteHostRuleRequest) (*pb.DeleteHostRuleResponse, error) {
	return s.client.DeleteHostRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetHostRuntimeRule(ctx context.Context, req *pb.GetHostRuntimeRuleRequest) (*pb.GetHostRuntimeRuleResponse, error) {
	return s.client.GetHostRuntimeRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) CreateHostRuntimeRule(ctx context.Context, req *pb.CreateHostRuntimeRuleRequest) (*pb.CreateHostRuntimeRuleResponse, error) {
	return s.client.CreateHostRuntimeRule(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetAllHostRuntimeRules(ctx context.Context, req *pb.GetAllHostRuntimeRulesRequest) (*pb.GetAllHostRuntimeRulesResponse, error) {
	return s.client.GetAllHostRuntimeRules(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetDubboInterfaceTime(ctx context.Context, req *pb.GetDubboInterfaceTimeRequest) (*pb.GetDubboInterfaceTimeResponse, error) {
	return s.client.GetDubboInterfaceTime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetDubboInterfaceQPS(ctx context.Context, req *pb.GetDubboInterfaceQPSRequest) (*pb.GetDubboInterfaceQPSResponse, error) {
	return s.client.GetDubboInterfaceQPS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetDubboInterfaceFailed(ctx context.Context, req *pb.GetDubboInterfaceFailedRequest) (*pb.GetDubboInterfaceFailedResponse, error) {
	return s.client.GetDubboInterfaceFailed(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *registerCenterServiceWrapper) GetDubboInterfaceAvgTime(ctx context.Context, req *pb.GetDubboInterfaceAvgTimeRequest) (*pb.GetDubboInterfaceAvgTimeResponse, error) {
	return s.client.GetDubboInterfaceAvgTime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
