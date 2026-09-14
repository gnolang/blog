package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// gno.land web pages advertise the chain they front through `gnoconnect` <meta>
// tags in their <head>, e.g:
//
//	<meta name="gnoconnect:rpc" content="https://rpc.gno.land" />
//	<meta name="gnoconnect:chainid" content="gnoland-1" />
var (
	reGnoconnectRPC     = regexp.MustCompile(`(?i)<meta[^>]+name=["']gnoconnect:rpc["'][^>]+content=["']([^"']+)["']`)
	reGnoconnectChainID = regexp.MustCompile(`(?i)<meta[^>]+name=["']gnoconnect:chainid["'][^>]+content=["']([^"']+)["']`)
)

// discoverTarget fetches the target gno.land web endpoint and extracts the RPC
// endpoint and chain ID from its `gnoconnect` <meta> tags, so the user only has
// to point the CLI at a website (e.g. `gno.land`) instead of manually wiring up
// --remote and --chainid.
func discoverTarget(target string) (rpc, chainID string, err error) {
	url := target
	if !strings.Contains(url, "://") {
		url = "https://" + url
	}

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("fetching %q: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("fetching %q: unexpected status %s", url, resp.Status)
	}

	// The <head> holds the meta tags; 1MB is plenty and bounds untrusted input.
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", "", fmt.Errorf("reading %q: %w", url, err)
	}

	rpcMatch := reGnoconnectRPC.FindSubmatch(body)
	chainMatch := reGnoconnectChainID.FindSubmatch(body)
	if rpcMatch == nil || chainMatch == nil {
		return "", "", fmt.Errorf("%q does not expose gnoconnect meta tags (rpc/chainid)", url)
	}

	return string(rpcMatch[1]), string(chainMatch[1]), nil
}
