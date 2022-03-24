package pwdvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	var testcases = []struct {
		description  string
		password     string
		expectErr    bool
		expectErrMsg string
	}{
		{
			description:  "valid password",
			password:     "2KerDQ89FMCnwgR!LAf69zf",
			expectErr:    false,
			expectErrMsg: "",
		},
		{
			description:  "common password",
			password:     "password",
			expectErr:    true,
			expectErrMsg: "PWD_NOT_SECURE",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			actualErr := ValidatePassword(tc.password)
			if tc.expectErr {
				assert.Error(t, actualErr)
				assert.Equal(t, tc.expectErrMsg, actualErr.Error())
				return
			}
			assert.NoError(t, actualErr)
		})
	}
}
