package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
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
		{
			name: "Too many args provided",
			io: mockIO{
				getPassword: func() (string, error) {
					return "pass", nil
				},
			},
			args: []string{
				"path1",
				"path2",
				"path3",
			},
			expectedErr: ErrInvalidNumberOfArgs,
		},
		{
			name: "0 args provided",
			io: mockIO{
				getPassword: func() (string, error) {
					return "pass", nil
				},
			},
			args:        []string{},
			expectedErr: ErrInvalidNumberOfArgs,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFn()

			err := execPost(testCase.io, testCase.args, testCase.cfg)
			if err != nil {
				return
			}
		})
	}
}

// Helpers
func generateTestPostContent(t *testing.T) string {
	t.Helper()

	return `---
title: "Example Post"
publication_date: 2022-05-02T13:17:22Z
slug: example-post
tags: [post, tag1, "gno.land"]
authors: [test, author, leon]
---

# Test Header

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint 
occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
}

func removeDir(t *testing.T, dirPath string) func() {
	return func() {
		err := os.RemoveAll(dirPath)
		require.NoError(t, err)
	}
}

// todo add integration tests once a good way for it is available.
