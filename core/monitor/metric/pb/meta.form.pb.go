// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: meta.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ListMetricNamesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetricNamesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetricMetaRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetricMetaResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetricGroupsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetricGroupsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetMetricGroupRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetMetricGroupResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MetricGroup)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GroupMetricMeta)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MetricMeta)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NameDefine)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TagDefine)(nil)
var _ urlenc.URLValuesUnmarshaler = (*FieldDefine)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ValueDefine)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Group)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MetaMode)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TypeDefine)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Aggregation)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Operation)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TagFilter)(nil)

// ListMetricNamesRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricNamesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// ListMetricNamesResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricNamesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ListMetricMetaRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricMetaRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "metrics":
				m.Metrics = vals
			}
		}
	}
	return nil
}

// ListMetricMetaResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricMetaResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ListMetricGroupsRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricGroupsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "mode":
				m.Mode = vals[0]
			}
		}
	}
	return nil
}

// ListMetricGroupsResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListMetricGroupsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetMetricGroupRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetMetricGroupRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "id":
				m.Id = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "version":
				m.Version = vals[0]
			case "format":
				m.Format = vals[0]
			case "appendTags":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.AppendTags = val
			}
		}
	}
	return nil
}

// GetMetricGroupResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetMetricGroupResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &MetricGroup{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &MetricGroup{}
				}
				m.Data.Id = vals[0]
			case "data.meta":
				if m.Data == nil {
					m.Data = &MetricGroup{}
				}
				if m.Data.Meta == nil {
					m.Data.Meta = &MetaMode{}
				}
			}
		}
	}
	return nil
}

// MetricGroup implement urlenc.URLValuesUnmarshaler.
func (m *MetricGroup) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "meta":
				if m.Meta == nil {
					m.Meta = &MetaMode{}
				}
			}
		}
	}
	return nil
}

// GroupMetricMeta implement urlenc.URLValuesUnmarshaler.
func (m *GroupMetricMeta) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "metric":
				m.Metric = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// MetricMeta implement urlenc.URLValuesUnmarshaler.
func (m *MetricMeta) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				if m.Name == nil {
					m.Name = &NameDefine{}
				}
			case "name.key":
				if m.Name == nil {
					m.Name = &NameDefine{}
				}
				m.Name.Key = vals[0]
			case "name.name":
				if m.Name == nil {
					m.Name = &NameDefine{}
				}
				m.Name.Name = vals[0]
			}
		}
	}
	return nil
}

// NameDefine implement urlenc.URLValuesUnmarshaler.
func (m *NameDefine) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// TagDefine implement urlenc.URLValuesUnmarshaler.
func (m *TagDefine) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// FieldDefine implement urlenc.URLValuesUnmarshaler.
func (m *FieldDefine) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "unit":
				m.Unit = vals[0]
			}
		}
	}
	return nil
}

// ValueDefine implement urlenc.URLValuesUnmarshaler.
func (m *ValueDefine) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "value":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Value = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value = val
					} else {
						m.Value = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// Group implement urlenc.URLValuesUnmarshaler.
func (m *Group) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			case "order":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Order = int32(val)
			}
		}
	}
	return nil
}

// MetaMode implement urlenc.URLValuesUnmarshaler.
func (m *MetaMode) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TypeDefine implement urlenc.URLValuesUnmarshaler.
func (m *TypeDefine) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Aggregation implement urlenc.URLValuesUnmarshaler.
func (m *Aggregation) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "aggregation":
				m.Aggregation = vals[0]
			case "name":
				m.Name = vals[0]
			case "result_type":
				m.ResultType = vals[0]
			}
		}
	}
	return nil
}

// Operation implement urlenc.URLValuesUnmarshaler.
func (m *Operation) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "operation":
				m.Operation = vals[0]
			case "name":
				m.Name = vals[0]
			case "multi":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Multi = val
			}
		}
	}
	return nil
}

// TagFilter implement urlenc.URLValuesUnmarshaler.
func (m *TagFilter) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tag":
				m.Tag = vals[0]
			case "op":
				m.Op = vals[0]
			case "value":
				m.Value = vals[0]
			}
		}
	}
	return nil
}
