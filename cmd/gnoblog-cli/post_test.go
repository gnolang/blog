package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestValidFlowSinglePost(t *testing.T) {
	var (
		cfg = &cliCfg{
			KeyName: "dev",
			Edit:    false,
		}
		io = mockIO{}
	)

	// Generate temporary dir
	sourceDir, err := os.MkdirTemp(".", "tmp")
	require.NoError(t, err)
	t.Cleanup(removeDir(t, sourceDir))

	// Write temp post
	postContent := generateTestPostContent(t)
	postPath := filepath.Join(sourceDir, "README.md")
	err = os.WriteFile(postPath, []byte(postContent), 0644)
	require.NoError(t, err)

	// Run execPost with mock data
	err = execPost(io, []string{postPath}, cfg)
	assert.NoError(t, err)
}

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
