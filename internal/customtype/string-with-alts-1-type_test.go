package customtype_test

import (
	"context"
	"testing"

	"github.com/chrismarget/terraform-provider-altstrings/customtype"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/require"
)

func TestStringWithAltValuesType_ValueFromTerraform(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in          tftypes.Value
		expectation attr.Value
		expectedErr string
	}{
		"true": {
			in:          tftypes.NewValue(tftypes.String, "foo"),
			expectation: customtype.NewStringWithAlts1("foo"),
		},
		"unknown": {
			in:          tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			expectation: customtype.NewStringWithAlts1Unknown(),
		},
		"null": {
			in:          tftypes.NewValue(tftypes.String, nil),
			expectation: customtype.NewStringWithAlts1Null(),
		},
		"wrongType": {
			in:          tftypes.NewValue(tftypes.Number, 123),
			expectedErr: "can't unmarshal tftypes.Number into *string, expected string",
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			got, err := customtype.StringWithAlts1Type{}.ValueFromTerraform(ctx, tCase.in)
			if tCase.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, tCase.expectedErr, err.Error())
				return
			}

			require.Truef(t, got.Equal(tCase.expectation), "values not equal %s, %s", tCase.expectation, got)
		})
	}
}
