package main

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

func initSigner(cfg *cliCfg, password string) gnoclient.SignerFromKeybase {
	kb, _ := keys.NewKeyBaseFromDir(cfg.GnoHome)
	return gnoclient.SignerFromKeybase{
		Keybase:  kb,
		Account:  cfg.KeyName,
		Password: password,
		ChainID:  cfg.ChainId,
	}
}

func initRPCClient(cfg *cliCfg) rpcclient.Client {
	remote := cfg.Remote // todo: validate if chainID & remote match?
	return rpcclient.NewHTTP(remote, "/websocket")
}
