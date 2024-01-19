package main

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

func initSigner(cfg *postCfg, password string) gnoclient.SignerFromKeybase {
	kb, _ := keys.NewKeyBaseFromDir(cfg.gnoHome)
	return gnoclient.SignerFromKeybase{
		Keybase:  kb,
		Account:  cfg.keyName,
		Password: password,
		ChainID:  cfg.chainId,
	}
}

func initRPCClient(cfg *postCfg) rpcclient.Client {
	remote := cfg.remote // todo: validate if chainID & remote match?
	return rpcclient.NewHTTP(remote, "/websocket")
}
