package customtype

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ basetypes.StringValuableWithSemanticEquals = (*StringWithAlts1)(nil)

type StringWithAlts1 struct {
	basetypes.StringValue
	altValues []basetypes.StringValue
}

func (v StringWithAlts1) Equal(o attr.Value) bool {
	other, ok := o.(StringWithAlts1)
	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v StringWithAlts1) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(StringWithAlts1)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	// check new value against our "main" value
	if v.Equal(newValue) {
		return true, diags
	}

	// check new value against our "alt" values
	for _, altValue := range v.altValues {
		if altValue.Equal(newValue.StringValue) {
			return true, diags
		}
	}

	// check our value against new "alt" values
	for _, altValue := range newValue.altValues {
		if altValue.Equal(v.StringValue) {
			return true, diags
		}
	}

	return false, diags
}

func (v StringWithAlts1) Type(_ context.Context) attr.Type {
	return StringWithAlts1Type{}
}

func NewStringWithAlts1Null() StringWithAlts1 {
	return StringWithAlts1{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewStringWithAlts1Unknown() StringWithAlts1 {
	return StringWithAlts1{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewStringWithAlts1(s string, alts ...string) StringWithAlts1 {
	altValues := make([]types.String, len(alts))
	for i, alt := range alts {
		altValues[i] = types.StringValue(alt)
	}

	return StringWithAlts1{
		StringValue: basetypes.NewStringValue(s),
		altValues:   altValues,
	}
}
