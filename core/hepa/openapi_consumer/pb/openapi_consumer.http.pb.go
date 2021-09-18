// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: openapi_consumer.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// OpenapiConsumerServiceHandler is the server API for OpenapiConsumerService service.
type OpenapiConsumerServiceHandler interface {
	// +publish path: "/api/gateway/openapi/consumers"
	// GET /api/gateway/openapi/consumers
	GetConsumers(context.Context, *GetConsumersRequest) (*GetConsumersResponse, error)
	// +publish path: "/api/gateway/openapi/consumers"
	// POST /api/gateway/openapi/consumers
	CreateConsumer(context.Context, *CreateConsumerRequest) (*CreateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	// PATCH /api/gateway/openapi/consumers/{consumerId}
	UpdateConsumer(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}"
	// DELETE /api/gateway/openapi/consumers/{consumerId}
	DeleteConsumer(context.Context, *DeleteConsumerRequest) (*DeleteConsumerResponse, error)
	// +publish path: "/api/gateway/openapi/consumers-name"
	// GET /api/gateway/openapi/consumers-name
	GetConsumersName(context.Context, *GetConsumersNameRequest) (*GetConsumersNameResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	// GET /api/gateway/openapi/consumers/{consumerId}/packages
	GetConsumerAcl(context.Context, *GetConsumerAclRequest) (*GetConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/packages"
	// POST /api/gateway/openapi/consumers/{consumerId}/packages
	UpdateConsumerAcl(context.Context, *UpdateConsumerAclRequest) (*UpdateConsumerAclResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	// GET /api/gateway/openapi/consumers/{consumerId}/credentials
	GetConsumerAuth(context.Context, *GetConsumerAuthRequest) (*GetConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/consumers/{consumerId}/credentials"
	// POST /api/gateway/openapi/consumers/{consumerId}/credentials
	UpdateConsumerAuth(context.Context, *UpdateConsumerAuthRequest) (*UpdateConsumerAuthResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	// GET /api/gateway/openapi/packages/{packageId}/consumers
	GetEndpointAcl(context.Context, *GetEndpointAclRequest) (*GetEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/consumers"
	// POST /api/gateway/openapi/packages/{packageId}/consumers
	UpdateEndpointAcl(context.Context, *UpdateEndpointAclRequest) (*UpdateEndpointAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	// GET /api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz
	GetEndpointApiAcl(context.Context, *GetEndpointApiAclRequest) (*GetEndpointApiAclResponse, error)
	// +publish path: "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz"
	// POST /api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz
	UpdateEndpointApiAcl(context.Context, *UpdateEndpointApiAclRequest) (*UpdateEndpointApiAclResponse, error)
}

// RegisterOpenapiConsumerServiceHandler register OpenapiConsumerServiceHandler to http.Router.
func RegisterOpenapiConsumerServiceHandler(r http.Router, srv OpenapiConsumerServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_GetConsumers := func(method, path string, fn func(context.Context, *GetConsumersRequest) (*GetConsumersResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetConsumersRequest))
		}
		var GetConsumers_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetConsumers_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumers", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetConsumers_info)
				}
				r = r.WithContext(ctx)
				var in GetConsumersRequest
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
				if vals := params["env"]; len(vals) > 0 {
					in.Env = vals[0]
				}
				if vals := params["pageNo"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.PageNo = val
				}
				if vals := params["pageSize"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.PageSize = val
				}
				if vals := params["projectId"]; len(vals) > 0 {
					in.ProjectId = vals[0]
				}
				if vals := params["sortField"]; len(vals) > 0 {
					in.SortField = vals[0]
				}
				if vals := params["sortType"]; len(vals) > 0 {
					in.SortType = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateConsumer := func(method, path string, fn func(context.Context, *CreateConsumerRequest) (*CreateConsumerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateConsumerRequest))
		}
		var CreateConsumer_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "CreateConsumer", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateConsumer_info)
				}
				r = r.WithContext(ctx)
				var in CreateConsumerRequest
				if err := h.Decode(r, &in.Consumer); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["env"]; len(vals) > 0 {
					in.Env = vals[0]
				}
				if vals := params["projectId"]; len(vals) > 0 {
					in.ProjectId = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateConsumer := func(method, path string, fn func(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateConsumerRequest))
		}
		var UpdateConsumer_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumer", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateConsumer_info)
				}
				r = r.WithContext(ctx)
				var in UpdateConsumerRequest
				if err := h.Decode(r, &in.Consumer); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_DeleteConsumer := func(method, path string, fn func(context.Context, *DeleteConsumerRequest) (*DeleteConsumerResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteConsumerRequest))
		}
		var DeleteConsumer_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteConsumer_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "DeleteConsumer", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteConsumer_info)
				}
				r = r.WithContext(ctx)
				var in DeleteConsumerRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_GetConsumersName := func(method, path string, fn func(context.Context, *GetConsumersNameRequest) (*GetConsumersNameResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetConsumersNameRequest))
		}
		var GetConsumersName_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetConsumersName_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumersName", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetConsumersName_info)
				}
				r = r.WithContext(ctx)
				var in GetConsumersNameRequest
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
				if vals := params["env"]; len(vals) > 0 {
					in.Env = vals[0]
				}
				if vals := params["projectId"]; len(vals) > 0 {
					in.ProjectId = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetConsumerAcl := func(method, path string, fn func(context.Context, *GetConsumerAclRequest) (*GetConsumerAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetConsumerAclRequest))
		}
		var GetConsumerAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetConsumerAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumerAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetConsumerAcl_info)
				}
				r = r.WithContext(ctx)
				var in GetConsumerAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_UpdateConsumerAcl := func(method, path string, fn func(context.Context, *UpdateConsumerAclRequest) (*UpdateConsumerAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateConsumerAclRequest))
		}
		var UpdateConsumerAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateConsumerAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumerAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateConsumerAcl_info)
				}
				r = r.WithContext(ctx)
				var in UpdateConsumerAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_GetConsumerAuth := func(method, path string, fn func(context.Context, *GetConsumerAuthRequest) (*GetConsumerAuthResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetConsumerAuthRequest))
		}
		var GetConsumerAuth_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetConsumerAuth_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetConsumerAuth", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetConsumerAuth_info)
				}
				r = r.WithContext(ctx)
				var in GetConsumerAuthRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_UpdateConsumerAuth := func(method, path string, fn func(context.Context, *UpdateConsumerAuthRequest) (*UpdateConsumerAuthResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateConsumerAuthRequest))
		}
		var UpdateConsumerAuth_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateConsumerAuth_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateConsumerAuth", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateConsumerAuth_info)
				}
				r = r.WithContext(ctx)
				var in UpdateConsumerAuthRequest
				if err := h.Decode(r, &in.Credentials); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "consumerId":
							in.ConsumerId = val
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

	add_GetEndpointAcl := func(method, path string, fn func(context.Context, *GetEndpointAclRequest) (*GetEndpointAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetEndpointAclRequest))
		}
		var GetEndpointAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetEndpointAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetEndpointAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetEndpointAcl_info)
				}
				r = r.WithContext(ctx)
				var in GetEndpointAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "packageId":
							in.PackageId = val
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

	add_UpdateEndpointAcl := func(method, path string, fn func(context.Context, *UpdateEndpointAclRequest) (*UpdateEndpointAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateEndpointAclRequest))
		}
		var UpdateEndpointAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateEndpointAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateEndpointAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateEndpointAcl_info)
				}
				r = r.WithContext(ctx)
				var in UpdateEndpointAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "packageId":
							in.PackageId = val
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

	add_GetEndpointApiAcl := func(method, path string, fn func(context.Context, *GetEndpointApiAclRequest) (*GetEndpointApiAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetEndpointApiAclRequest))
		}
		var GetEndpointApiAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetEndpointApiAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "GetEndpointApiAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetEndpointApiAcl_info)
				}
				r = r.WithContext(ctx)
				var in GetEndpointApiAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "packageId":
							in.PackageId = val
						case "apiId":
							in.ApiId = val
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

	add_UpdateEndpointApiAcl := func(method, path string, fn func(context.Context, *UpdateEndpointApiAclRequest) (*UpdateEndpointApiAclResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateEndpointApiAclRequest))
		}
		var UpdateEndpointApiAcl_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateEndpointApiAcl_info = transport.NewServiceInfo("erda.core.hepa.openapi_consumer.OpenapiConsumerService", "UpdateEndpointApiAcl", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateEndpointApiAcl_info)
				}
				r = r.WithContext(ctx)
				var in UpdateEndpointApiAclRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
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
						case "packageId":
							in.PackageId = val
						case "apiId":
							in.ApiId = val
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

	add_GetConsumers("GET", "/api/gateway/openapi/consumers", srv.GetConsumers)
	add_CreateConsumer("POST", "/api/gateway/openapi/consumers", srv.CreateConsumer)
	add_UpdateConsumer("PATCH", "/api/gateway/openapi/consumers/{consumerId}", srv.UpdateConsumer)
	add_DeleteConsumer("DELETE", "/api/gateway/openapi/consumers/{consumerId}", srv.DeleteConsumer)
	add_GetConsumersName("GET", "/api/gateway/openapi/consumers-name", srv.GetConsumersName)
	add_GetConsumerAcl("GET", "/api/gateway/openapi/consumers/{consumerId}/packages", srv.GetConsumerAcl)
	add_UpdateConsumerAcl("POST", "/api/gateway/openapi/consumers/{consumerId}/packages", srv.UpdateConsumerAcl)
	add_GetConsumerAuth("GET", "/api/gateway/openapi/consumers/{consumerId}/credentials", srv.GetConsumerAuth)
	add_UpdateConsumerAuth("POST", "/api/gateway/openapi/consumers/{consumerId}/credentials", srv.UpdateConsumerAuth)
	add_GetEndpointAcl("GET", "/api/gateway/openapi/packages/{packageId}/consumers", srv.GetEndpointAcl)
	add_UpdateEndpointAcl("POST", "/api/gateway/openapi/packages/{packageId}/consumers", srv.UpdateEndpointAcl)
	add_GetEndpointApiAcl("GET", "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz", srv.GetEndpointApiAcl)
	add_UpdateEndpointApiAcl("POST", "/api/gateway/openapi/packages/{packageId}/apis/{apiId}/authz", srv.UpdateEndpointApiAcl)
}
