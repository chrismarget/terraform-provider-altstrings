package customtype

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*StringWithAltsType)(nil)

type StringWithAltsType struct {
	basetypes.StringType
}

func (o StringWithAltsType) Equal(other attr.Type) bool {
	_, ok := other.(StringWithAltsType)
	return ok
}

func (o StringWithAltsType) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.IsKnown() {
		return NewStringWithAltsUnknown(), nil
	}

	if in.IsNull() {
		return NewStringWithAltsNull(), nil
	}

	var s string
	err := in.As(&s)
	if err != nil {
		return nil, err
	}

	return NewStringWithAlts(s), nil
}

func (o StringWithAltsType) ValueType(_ context.Context) attr.Value {
	return StringWithAlts{}
}

func (o StringWithAltsType) String() string {
	return "customtype.StringWithAltsType"
}

func (o StringWithAltsType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return NewStringWithAlts(in.String()), nil
}
