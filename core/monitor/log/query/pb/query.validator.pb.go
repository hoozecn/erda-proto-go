// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *LogItem) Validate() error {
	return nil
}
func (this *GetLogRequest) Validate() error {
	return nil
}
func (this *GetLogByRuntimeRequest) Validate() error {
	return nil
}
func (this *GetLogByOrganizationRequest) Validate() error {
	return nil
}
func (this *GetLogResponse) Validate() error {
	for _, item := range this.Lines {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lines", err)
			}
		}
	}
	return nil
}
func (this *GetLogByRuntimeResponse) Validate() error {
	for _, item := range this.Lines {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lines", err)
			}
		}
	}
	return nil
}
func (this *GetLogByOrganizationResponse) Validate() error {
	for _, item := range this.Lines {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lines", err)
			}
		}
	}
	return nil
}
