// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: runtime.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// RuntimeServiceHandler is the server API for RuntimeService service.
type RuntimeServiceHandler interface {
	// GET /api/runtimes/healthz
	Healthz(context.Context, *pb.VoidRequest) (*pb.VoidResponse, error)
	// GET /api/runtimes/{nameOrID}
	GetRuntime(context.Context, *GetRuntimeRequest) (*RuntimeInspect, error)
}

// RegisterRuntimeServiceHandler register RuntimeServiceHandler to http.Router.
func RegisterRuntimeServiceHandler(r http.Router, srv RuntimeServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_Healthz := func(method, path string, fn func(context.Context, *pb.VoidRequest) (*pb.VoidResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*pb.VoidRequest))
		}
		var Healthz_info transport.ServiceInfo
		if h.Interceptor != nil {
			Healthz_info = transport.NewServiceInfo("erda.orchestrator.runtime.RuntimeService", "Healthz", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Healthz_info)
				}
				r = r.WithContext(ctx)
				var in pb.VoidRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetRuntime := func(method, path string, fn func(context.Context, *GetRuntimeRequest) (*RuntimeInspect, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetRuntimeRequest))
		}
		var GetRuntime_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetRuntime_info = transport.NewServiceInfo("erda.orchestrator.runtime.RuntimeService", "GetRuntime", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetRuntime_info)
				}
				r = r.WithContext(ctx)
				var in GetRuntimeRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["applicationId"]; len(vals) > 0 {
					in.AppID = vals[0]
				}
				if vals := params["workspace"]; len(vals) > 0 {
					in.Workspace = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "nameOrID":
							in.NameOrID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Healthz("GET", "/api/runtimes/healthz", srv.Healthz)
	add_GetRuntime("GET", "/api/runtimes/{nameOrID}", srv.GetRuntime)
}