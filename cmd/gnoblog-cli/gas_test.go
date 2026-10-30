package main

import (
	"strings"
	"testing"

	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/stretchr/testify/require"
)

func makeMsg(bodySize int) vm.MsgCall {
	return vm.MsgCall{
		PkgPath: "gno.land/r/gnoland/blog",
		Func:    "ModAddPost",
		Args:    []string{"slug", "title", strings.Repeat("x", bodySize)},
	}
}

func TestPackBatches(t *testing.T) {
	t.Parallel()

	msgs := []vm.MsgCall{makeMsg(100), makeMsg(100), makeMsg(100)}
	msgSize := len(amino.MustMarshal(msgs[0]))

	t.Run("all fit in one batch", func(t *testing.T) {
		t.Parallel()
		batches := packBatches(msgs, msgSize*3+100)
		require.Len(t, batches, 1)
		require.Len(t, batches[0], 3)
	})

	t.Run("split across batches by size", func(t *testing.T) {
		t.Parallel()
		// Room for ~2 messages per batch.
		batches := packBatches(msgs, msgSize*2+msgSize/2)
		require.Len(t, batches, 2)
		require.Len(t, batches[0], 2)
		require.Len(t, batches[1], 1)
	})

	t.Run("oversized message gets its own batch", func(t *testing.T) {
		t.Parallel()
		batches := packBatches(msgs, 1) // limit below any single message
		require.Len(t, batches, 3)
		for _, b := range batches {
			require.Len(t, b, 1)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		t.Parallel()
		require.Empty(t, packBatches(nil, 1000))
	})
}

func TestDeriveGasFee(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name      string
		gasWanted int64
		gasPrice  string
		expected  string
		expectErr bool
	}{
		{"mainnet floor price", 50_000_000, "1ugnot/1000gas", "50000ugnot", false},
		{"rounds up", 1500, "1ugnot/1000gas", "2ugnot", false},
		{"minimum one", 1, "1ugnot/1000gas", "1ugnot", false},
		{"invalid price", 1000, "not-a-price", "", true},
	}

	for _, tc := range testTable {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fee, err := deriveGasFee(tc.gasWanted, tc.gasPrice)
			if tc.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, fee)
		})
	}
}
