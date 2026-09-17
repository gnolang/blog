package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindFilePaths(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		files    []string
		expected []string
	}{
		{
			name:     "exact name is picked up",
			files:    []string{"2026-01-01_post/README.md"},
			expected: []string{"2026-01-01_post/README.md"},
		},
		{
			name:     "miscased name is picked up too",
			files:    []string{"2026-01-01_post/ReadME.md"},
			expected: []string{"2026-01-01_post/ReadME.md"},
		},
		{
			name:     "unrelated files are ignored",
			files:    []string{"2026-01-01_post/README.md", "2026-01-01_post/src/image.png", "notes.md"},
			expected: []string{"2026-01-01_post/README.md"},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			root := t.TempDir()

			for _, file := range testCase.files {
				path := filepath.Join(root, file)

				require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
				require.NoError(t, os.WriteFile(path, []byte("# title"), 0o644))
			}

			paths, err := findFilePaths(root)
			require.NoError(t, err)

			expected := make([]string, 0, len(testCase.expected))
			for _, file := range testCase.expected {
				expected = append(expected, filepath.Join(root, file))
			}

			require.ElementsMatch(t, expected, paths)
		})
	}
}
