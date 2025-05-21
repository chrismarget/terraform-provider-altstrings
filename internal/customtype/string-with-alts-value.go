package customtype

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ basetypes.StringValuableWithSemanticEquals = (*StringWithAlts)(nil)

type StringWithAlts struct {
	basetypes.StringValue
	altValues []basetypes.StringValue
}

func (v StringWithAlts) Equal(o attr.Value) bool {
	other, ok := o.(StringWithAlts)
	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v StringWithAlts) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(StringWithAlts)
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

func (v StringWithAlts) Type(_ context.Context) attr.Type {
	return StringWithAltsType{}
}

func NewStringWithAltsNull() StringWithAlts {
	return StringWithAlts{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewStringWithAltsUnknown() StringWithAlts {
	return StringWithAlts{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewStringWithAlts(s string, alts ...string) StringWithAlts {
	altValues := make([]types.String, len(alts))
	for i, alt := range alts {
		altValues[i] = types.StringValue(alt)
	}

	return StringWithAlts{
		StringValue: basetypes.NewStringValue(s),
		altValues:   altValues,
	}
}
