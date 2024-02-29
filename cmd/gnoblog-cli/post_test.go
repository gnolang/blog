package main

import (
	"context"
	"testing"
	"time"
)

func TestInputs(t *testing.T) {
	testTable := []struct {
		name        string
		io          mockIO
		cfg         *cliCfg
		args        []string
		expectedErr error
	}{
		{
			name: "No key provided",
			io: mockIO{
				getPassword: func() (string, error) {
					return "pass", nil
				},
			},
			cfg: &cliCfg{
				KeyName: "",
			},
			args: []string{
				"path1",
			},
			expectedErr: ErrEmptyKeyName,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFn()

			err := execPost(mockIO{}, []string{"lol"}, testCase.cfg)
			if err != nil {
				return
			}
		})
	}

}
