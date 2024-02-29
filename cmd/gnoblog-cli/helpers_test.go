package main

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/integration"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	"github.com/gnolang/gno/tm2/pkg/bft/node"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/log"
	"github.com/stretchr/testify/require"

	"testing"
)

func newInMemoryNode(t *testing.T) (*node.Node, string) {
	t.Helper()
	// Set up in-memory node
	config, _ := integration.TestingNodeConfig(t, gnoenv.RootDir())

	//// Init Signer & RPCClient
	//signer := newInMemorySigner(t, "tendermint_test")
	//rpcClient := rpcclient.NewHTTP(remoteAddr, "/websocket")

	// defer stop node after the call to this function!
	return integration.TestingInMemoryNode(t, log.NewNoopLogger(), config)
}

func newInMemorySigner(t *testing.T, chainid string) *gnoclient.SignerFromKeybase {
	t.Helper()

	mnemonic := integration.DefaultAccount_Seed
	name := integration.DefaultAccount_Name

	kb := keys.NewInMemory()
	_, err := kb.CreateAccount(name, mnemonic, "", "", uint32(0), uint32(0))
	require.NoError(t, err)

	return &gnoclient.SignerFromKeybase{
		Keybase:  kb,      // Stores keys in memory or on disk
		Account:  name,    // Account name or bech32 format
		Password: "",      // Password for encryption
		ChainID:  chainid, // Chain ID for transaction signing
	}
}
