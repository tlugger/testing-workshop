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
			description:  "empty password",
			password:     "",
			expectErr:    true,
			expectErrMsg: "MUST_PROVIDE_PASSWORD",
		},
		{
			description:  "short password",
			password:     "a",
			expectErr:    true,
			expectErrMsg: "PWD_TOO_SHOT",
		},
		{
			description:  "long password",
			password:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			expectErr:    true,
			expectErrMsg: "PWD_TOO_LONG",
		},
		{
			description:  "common password",
			password:     "password",
			expectErr:    true,
			expectErrMsg: "PWD_NOT_SECURE",
		},
		{
			description:  "no uppercase password",
			password:     "asecurepassword",
			expectErr:    true,
			expectErrMsg: "PWD_MUST_INCLUDE_UPPERCASE_LETTER",
		},
		{
			description:  "no lowercase password",
			password:     "ASECUREPASSWORD",
			expectErr:    true,
			expectErrMsg: "PWD_MUST_INCLUDE_LOWERCASE_LETTER",
		},
		{
			description:  "no special char password",
			password:     "ASecurePassword",
			expectErr:    true,
			expectErrMsg: "PWD_MUST_INCLUDE_SPECIAL_CHARACTER",
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
