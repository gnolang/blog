package main

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

func initSigner(cfg *cliCfg, password string) (gnoclient.SignerFromKeybase, error) {
	kb, err := keys.NewKeyBaseFromDir(cfg.GnoHome)
	if err != nil {
		return gnoclient.SignerFromKeybase{}, err
	}

	return gnoclient.SignerFromKeybase{
		Keybase:  kb,
		Account:  cfg.KeyName,
		Password: password,
		ChainID:  cfg.ChainId,
	}, nil
}

func initRPCClient(cfg *cliCfg) (rpcclient.Client, error) {
	remote := cfg.Remote // todo: validate if chainID & remote match?
	client, err := rpcclient.NewHTTPClient(remote)
	if err != nil {
		return nil, err
	}

	return client, nil
}
