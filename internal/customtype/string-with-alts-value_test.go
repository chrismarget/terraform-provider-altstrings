package customtype_test

import (
	"context"
	"testing"

	"github.com/chrismarget/terraform-provider-altstrings/internal/customtype"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/require"
)

func TestStringWithAltValues_StringSemanticEquals(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		currentValue  customtype.StringWithAlts
		givenValue    basetypes.StringValuable
		expectedMatch bool
	}{
		"equal - no alt values": {
			currentValue:  customtype.NewStringWithAlts("foo"),
			givenValue:    customtype.NewStringWithAlts("foo"),
			expectedMatch: true,
		},
		"equal - with alt values": {
			currentValue:  customtype.NewStringWithAlts("foo", "bar", "baz"),
			givenValue:    customtype.NewStringWithAlts("foo"),
			expectedMatch: true,
		},
		"semantically equal - given matches an alt value": {
			currentValue:  customtype.NewStringWithAlts("foo", "bar", "baz", "qux"),
			givenValue:    customtype.NewStringWithAlts("baz"),
			expectedMatch: true,
		},
		"semantically equal - current matches an alt value": {
			currentValue:  customtype.NewStringWithAlts("baz"),
			givenValue:    customtype.NewStringWithAlts("foo", "bar", "baz", "qux"),
			expectedMatch: true,
		},
		"not equal": {
			currentValue:  customtype.NewStringWithAlts("foo", "bar", "baz", "qux"),
			givenValue:    customtype.NewStringWithAlts("FOO"),
			expectedMatch: false,
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			t.Parallel()

			match, diags := tCase.currentValue.StringSemanticEquals(context.Background(), tCase.givenValue)
			require.Equalf(t, tCase.expectedMatch, match, "Expected StringSemanticEquals to return: %t, but got: %t", tCase.expectedMatch, match)
			require.Nil(t, diags)
		})
	}
}
