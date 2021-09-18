// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: tracing.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/oap/common/pb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*Span)(nil)

// Span implement urlenc.URLValuesUnmarshaler.
func (m *Span) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "traceID":
				m.TraceID = vals[0]
			case "spanID":
				m.SpanID = vals[0]
			case "parentSpanID":
				m.ParentSpanID = vals[0]
			case "startTimeUnixNano":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartTimeUnixNano = val
			case "endTimeUnixNano":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.EndTimeUnixNano = val
			case "name":
				m.Name = vals[0]
			case "relations":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
			case "relations.traceID":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.TraceID = vals[0]
			case "relations.resID":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResID = vals[0]
			case "relations.resType":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResType = vals[0]
			case "relations.resourceKeys":
				if m.Relations == nil {
					m.Relations = &pb.Relation{}
				}
				m.Relations.ResourceKeys = vals
			}
		}
	}
	return nil
}
