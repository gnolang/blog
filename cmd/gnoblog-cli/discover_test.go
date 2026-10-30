package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscoverTarget(t *testing.T) {
	t.Parallel()

	const page = `<!doctype html><html><head>
	<meta charset="UTF-8" />
	<meta name="gnoconnect:rpc" content="https://rpc.gno.land" />
	<meta name="gnoconnect:chainid" content="gnoland-1" />
	</head><body>hello</body></html>`

	t.Run("extracts rpc and chainid", func(t *testing.T) {
		t.Parallel()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte(page))
		}))
		defer srv.Close()

		rpc, chainID, err := discoverTarget(srv.URL)
		require.NoError(t, err)
		require.Equal(t, "https://rpc.gno.land", rpc)
		require.Equal(t, "gnoland-1", chainID)
	})

	t.Run("errors when meta tags missing", func(t *testing.T) {
		t.Parallel()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte(`<html><head></head><body>nope</body></html>`))
		}))
		defer srv.Close()

		_, _, err := discoverTarget(srv.URL)
		require.Error(t, err)
	})

	t.Run("errors on non-200", func(t *testing.T) {
		t.Parallel()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer srv.Close()

		_, _, err := discoverTarget(srv.URL)
		require.Error(t, err)
	})
}
