// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: base.proto

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

// BaseServiceHandler is the server API for BaseService service.
type BaseServiceHandler interface {
	////////// pipelines
	// POST /api/v2/pipelines
	PipelineCreate(context.Context, *PipelineCreateRequest) (*PipelineCreateResponse, error)
	// GET /api/pipelines
	PipelinePaging(context.Context, *PipelinePagingRequest) (*PipelinePagingResponse, error)
	// GET /api/pipelines/{pipelineID}
	PipelineDetail(context.Context, *PipelineDetailRequest) (*PipelineDetailResponse, error)
	// PUT /api/pipelines/{pipelineID}
	PipelineOperate(context.Context, *PipelineOperateRequest) (*PipelineOperateResponse, error)
	// DELETE /api/pipelines/{pipelineID}
	PipelineDelete(context.Context, *PipelineDeleteRequest) (*PipelineDeleteResponse, error)
	// POST /api/pipelines/{pipelineID}/actions/run
	PipelineRun(context.Context, *PipelineRunRequest) (*PipelineRunResponse, error)
	// POST /api/pipelines/{pipelineID}/actions/cancel
	PipelineCancel(context.Context, *PipelineCancelRequest) (*PipelineCancelResponse, error)
	// POST /api/pipelines/{pipelineID}/actions/rerun
	PipelineRerun(context.Context, *PipelineRerunRequest) (*PipelineRerunResponse, error)
	// POST /api/pipelines/{pipelineID}/actions/rerun-failed
	PipelineRerunFailed(context.Context, *PipelineRerunFailedRequest) (*PipelineRerunFailedResponse, error)
	// POST /api/pipelines/actions/batch-create
	DeprecatedPipelineBatchCreate(context.Context, *DeprecatedPipelineBatchCreateRequest) (*DeprecatedPipelineBatchCreateResponse, error)
	// POST /api/pipelines/actions/pipeline-yml-graph
	PipelineYmlGraph(context.Context, *PipelineYmlGraphRequest) (*PipelineYmlGraphResponse, error)
	// task
	// GET /api/pipelines/{pipelineID}/tasks/{taskID}
	PipelineTaskDetail(context.Context, *PipelineTaskDetailRequest) (*PipelineTaskDetailResponse, error)
	// GET /api/pipelines/{pipelineID}/tasks/{taskID}/actions/get-bootstrap-info
	PipelineTaskGetBootstrapInfo(context.Context, *PipelineTaskGetBootstrapInfoRequest) (*PipelineTaskGetBootstrapInfoResponse, error)
	// GET /api/pipelines/actions/callback
	PipelineCallback(context.Context, *PipelineCallbackRequest) (*PipelineCallbackResponse, error)
	// GET /api/pipeline-snippets/actions/query-details
	QueryPipelineSnippet(context.Context, *PipelineSnippetQueryRequest) (*PipelineSnippetQueryRequest, error)
}

// RegisterBaseServiceHandler register BaseServiceHandler to http.Router.
func RegisterBaseServiceHandler(r http.Router, srv BaseServiceHandler, opts ...http.HandleOption) {
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

	add_PipelineCreate := func(method, path string, fn func(context.Context, *PipelineCreateRequest) (*PipelineCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineCreateRequest))
		}
		var PipelineCreate_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineCreate_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineCreate", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineCreate_info)
				}
				r = r.WithContext(ctx)
				var in PipelineCreateRequest
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

	add_PipelinePaging := func(method, path string, fn func(context.Context, *PipelinePagingRequest) (*PipelinePagingResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelinePagingRequest))
		}
		var PipelinePaging_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelinePaging_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelinePaging", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelinePaging_info)
				}
				r = r.WithContext(ctx)
				var in PipelinePagingRequest
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
				if vals := params["anyMatchLabel"]; len(vals) > 0 {
					in.AnyMatchLabelsQueryParams = vals
				}
				if vals := params["anyMatchLabels"]; len(vals) > 0 {
					in.AnyMatchLabelsJSON = vals[0]
				}
				if vals := params["branch"]; len(vals) > 0 {
					in.Branches = vals
				}
				if vals := params["branches"]; len(vals) > 0 {
					in.CommaBranches = vals[0]
				}
				if vals := params["clusterName"]; len(vals) > 0 {
					in.ClusterNames = vals
				}
				if vals := params["mustMatchLabel"]; len(vals) > 0 {
					in.MustMatchLabelsQueryParams = vals
				}
				if vals := params["mustMatchLabels"]; len(vals) > 0 {
					in.MustMatchLabelsJSON = vals[0]
				}
				if vals := params["notStatus"]; len(vals) > 0 {
					in.NotStatuses = vals
				}
				if vals := params["source"]; len(vals) > 0 {
					in.Sources = vals
				}
				if vals := params["sources"]; len(vals) > 0 {
					in.CommaSources = vals[0]
				}
				if vals := params["status"]; len(vals) > 0 {
					in.Statuses = vals
				}
				if vals := params["statuses"]; len(vals) > 0 {
					in.CommaStatuses = vals[0]
				}
				if vals := params["triggerMode"]; len(vals) > 0 {
					in.TriggerModes = vals
				}
				if vals := params["ymlName"]; len(vals) > 0 {
					in.YmlNames = vals
				}
				if vals := params["ymlNames"]; len(vals) > 0 {
					in.CommaYmlNames = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_PipelineDetail := func(method, path string, fn func(context.Context, *PipelineDetailRequest) (*PipelineDetailResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineDetailRequest))
		}
		var PipelineDetail_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineDetail_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineDetail", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineDetail_info)
				}
				r = r.WithContext(ctx)
				var in PipelineDetailRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineOperate := func(method, path string, fn func(context.Context, *PipelineOperateRequest) (*PipelineOperateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineOperateRequest))
		}
		var PipelineOperate_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineOperate_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineOperate", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineOperate_info)
				}
				r = r.WithContext(ctx)
				var in PipelineOperateRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineDelete := func(method, path string, fn func(context.Context, *PipelineDeleteRequest) (*PipelineDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineDeleteRequest))
		}
		var PipelineDelete_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineDelete_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineDelete", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineDelete_info)
				}
				r = r.WithContext(ctx)
				var in PipelineDeleteRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineRun := func(method, path string, fn func(context.Context, *PipelineRunRequest) (*PipelineRunResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineRunRequest))
		}
		var PipelineRun_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineRun_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineRun", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineRun_info)
				}
				r = r.WithContext(ctx)
				var in PipelineRunRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineCancel := func(method, path string, fn func(context.Context, *PipelineCancelRequest) (*PipelineCancelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineCancelRequest))
		}
		var PipelineCancel_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineCancel_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineCancel", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineCancel_info)
				}
				r = r.WithContext(ctx)
				var in PipelineCancelRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineRerun := func(method, path string, fn func(context.Context, *PipelineRerunRequest) (*PipelineRerunResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineRerunRequest))
		}
		var PipelineRerun_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineRerun_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineRerun", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineRerun_info)
				}
				r = r.WithContext(ctx)
				var in PipelineRerunRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_PipelineRerunFailed := func(method, path string, fn func(context.Context, *PipelineRerunFailedRequest) (*PipelineRerunFailedResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineRerunFailedRequest))
		}
		var PipelineRerunFailed_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineRerunFailed_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineRerunFailed", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineRerunFailed_info)
				}
				r = r.WithContext(ctx)
				var in PipelineRerunFailedRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
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

	add_DeprecatedPipelineBatchCreate := func(method, path string, fn func(context.Context, *DeprecatedPipelineBatchCreateRequest) (*DeprecatedPipelineBatchCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeprecatedPipelineBatchCreateRequest))
		}
		var DeprecatedPipelineBatchCreate_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeprecatedPipelineBatchCreate_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "DeprecatedPipelineBatchCreate", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeprecatedPipelineBatchCreate_info)
				}
				r = r.WithContext(ctx)
				var in DeprecatedPipelineBatchCreateRequest
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

	add_PipelineYmlGraph := func(method, path string, fn func(context.Context, *PipelineYmlGraphRequest) (*PipelineYmlGraphResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineYmlGraphRequest))
		}
		var PipelineYmlGraph_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineYmlGraph_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineYmlGraph", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineYmlGraph_info)
				}
				r = r.WithContext(ctx)
				var in PipelineYmlGraphRequest
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

	add_PipelineTaskDetail := func(method, path string, fn func(context.Context, *PipelineTaskDetailRequest) (*PipelineTaskDetailResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineTaskDetailRequest))
		}
		var PipelineTaskDetail_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineTaskDetail_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineTaskDetail", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineTaskDetail_info)
				}
				r = r.WithContext(ctx)
				var in PipelineTaskDetailRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
						case "taskID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.TaskID = val
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

	add_PipelineTaskGetBootstrapInfo := func(method, path string, fn func(context.Context, *PipelineTaskGetBootstrapInfoRequest) (*PipelineTaskGetBootstrapInfoResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineTaskGetBootstrapInfoRequest))
		}
		var PipelineTaskGetBootstrapInfo_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineTaskGetBootstrapInfo_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineTaskGetBootstrapInfo", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineTaskGetBootstrapInfo_info)
				}
				r = r.WithContext(ctx)
				var in PipelineTaskGetBootstrapInfoRequest
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
						case "pipelineID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.PipelineID = val
						case "taskID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.TaskID = val
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

	add_PipelineCallback := func(method, path string, fn func(context.Context, *PipelineCallbackRequest) (*PipelineCallbackResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineCallbackRequest))
		}
		var PipelineCallback_info transport.ServiceInfo
		if h.Interceptor != nil {
			PipelineCallback_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "PipelineCallback", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PipelineCallback_info)
				}
				r = r.WithContext(ctx)
				var in PipelineCallbackRequest
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

	add_QueryPipelineSnippet := func(method, path string, fn func(context.Context, *PipelineSnippetQueryRequest) (*PipelineSnippetQueryRequest, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineSnippetQueryRequest))
		}
		var QueryPipelineSnippet_info transport.ServiceInfo
		if h.Interceptor != nil {
			QueryPipelineSnippet_info = transport.NewServiceInfo("erda.core.pipeline.base.BaseService", "QueryPipelineSnippet", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, QueryPipelineSnippet_info)
				}
				r = r.WithContext(ctx)
				var in PipelineSnippetQueryRequest
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

	add_PipelineCreate("POST", "/api/v2/pipelines", srv.PipelineCreate)
	add_PipelinePaging("GET", "/api/pipelines", srv.PipelinePaging)
	add_PipelineDetail("GET", "/api/pipelines/{pipelineID}", srv.PipelineDetail)
	add_PipelineOperate("PUT", "/api/pipelines/{pipelineID}", srv.PipelineOperate)
	add_PipelineDelete("DELETE", "/api/pipelines/{pipelineID}", srv.PipelineDelete)
	add_PipelineRun("POST", "/api/pipelines/{pipelineID}/actions/run", srv.PipelineRun)
	add_PipelineCancel("POST", "/api/pipelines/{pipelineID}/actions/cancel", srv.PipelineCancel)
	add_PipelineRerun("POST", "/api/pipelines/{pipelineID}/actions/rerun", srv.PipelineRerun)
	add_PipelineRerunFailed("POST", "/api/pipelines/{pipelineID}/actions/rerun-failed", srv.PipelineRerunFailed)
	add_DeprecatedPipelineBatchCreate("POST", "/api/pipelines/actions/batch-create", srv.DeprecatedPipelineBatchCreate)
	add_PipelineYmlGraph("POST", "/api/pipelines/actions/pipeline-yml-graph", srv.PipelineYmlGraph)
	add_PipelineTaskDetail("GET", "/api/pipelines/{pipelineID}/tasks/{taskID}", srv.PipelineTaskDetail)
	add_PipelineTaskGetBootstrapInfo("GET", "/api/pipelines/{pipelineID}/tasks/{taskID}/actions/get-bootstrap-info", srv.PipelineTaskGetBootstrapInfo)
	add_PipelineCallback("GET", "/api/pipelines/actions/callback", srv.PipelineCallback)
	add_QueryPipelineSnippet("GET", "/api/pipeline-snippets/actions/query-details", srv.QueryPipelineSnippet)
}
