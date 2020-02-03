package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintError(t *testing.T) {
	var tests = []struct {
		desc           string
		code           string
		msg            string
		extra          string
		origErr        error
		expectedOutput string
	}{
		{
			desc:           "all fields provided",
			code:           ErrCodeTimeout,
			msg:            "some failure message",
			extra:          "more stuff",
			origErr:        errors.New("Mayday mayday mayday"),
			expectedOutput: "timeout: some failure message\n\tmore stuff\n caused by: Mayday mayday mayday",
		},
		{
			desc:           "error not provided",
			code:           ErrCodeTimeout,
			msg:            "some failure message",
			extra:          "more stuff",
			expectedOutput: "timeout: some failure message\n\tmore stuff",
		},
		{
			desc:           "extra not provided",
			code:           ErrCodeTimeout,
			msg:            "some failure message",
			origErr:        errors.New("Mayday mayday mayday"),
			expectedOutput: "timeout: some failure message\n caused by: Mayday mayday mayday",
		},
		{
			desc:           "only the code and msg are provided",
			code:           ErrCodeTimeout,
			msg:            "some failure message",
			expectedOutput: "timeout: some failure message",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.desc, func(t *testing.T) {
			assert := assert.New(t)
			op := printError(testcase.code, testcase.msg, testcase.extra, testcase.origErr)
			assert.Equal(testcase.expectedOutput, op)
		})
	}

}
