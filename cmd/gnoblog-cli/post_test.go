package main

import (
	"context"
	"testing"
	"time"
)

func TestInputs(t *testing.T) {
	type cliCfg struct {
		Publish       bool
		Edit          bool
		GasWanted     int64
		GasFee        string
		ChainId       string
		BlogRealmPath string

		KeyName               string
		GnoHome               string
		Remote                string
		Quiet                 bool
		InsecurePasswordStdIn bool
	}

	testTable := []struct {
		name        string
		cfg         *cliCfg
		expectedErr error
	}{
		{
			name: "No key provided",
			cfg: &cliCfg{
				KeyName: "",
			},
			expectedErr: ErrEmptyKeyName,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFn()

			testCase.cfg

		})
	}

}
