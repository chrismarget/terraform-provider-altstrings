package customtype

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*StringWithAlts1Type)(nil)

type StringWithAlts1Type struct {
	basetypes.StringType
}

func (o StringWithAlts1Type) Equal(other attr.Type) bool {
	_, ok := other.(StringWithAlts1Type)

	return ok
}

func (o StringWithAlts1Type) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.IsKnown() {
		return NewStringWithAlts1Unknown(), nil
	}

	if in.IsNull() {
		return NewStringWithAlts1Null(), nil
	}

	var s string
	err := in.As(&s)
	if err != nil {
		return nil, err
	}

	return NewStringWithAlts1(s), nil
}

func (o StringWithAlts1Type) ValueType(_ context.Context) attr.Value {
	return StringWithAlts1{}
}

func (o StringWithAlts1Type) String() string {
	return "customtype.StringWithAlts1Type"
}

func (o StringWithAlts1Type) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return NewStringWithAlts1(in.String()), nil
}
