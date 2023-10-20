// Disclaimer: This version is a rough draft intended for experimentation with low-level libraries outside the monorepo.
// The plan is to work on a clean SDK and refactor this file accordingly.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gnolang/gno/tm2/pkg/amino"
	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	keysclient "github.com/gnolang/gno/tm2/pkg/crypto/keys/client"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/gnolang/gno/tm2/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/std"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		// if !errors.Is(err, flag.ErrHelp) {
		if err != flag.ErrHelp {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}
		os.Exit(1)
	}
}

type publishOpts struct {
	Debug    bool
	Publish  bool
	PkgPath  string
	EditMode bool

	// gnokey
	keysclient.BaseOptions
	GasWanted       int64
	GasFee          string
	ChainID         string
	KeyNameOrBech32 string

	// internal
	accountNumber  uint64
	sequenceNumber uint64
	pass           string
	kb             keys.Keybase
}

func (opts *publishOpts) flagSet() *flag.FlagSet {
	defaultHome := keysclient.DefaultBaseOptions.Home

	fs := flag.NewFlagSet("blog publish", flag.ExitOnError)
	fs.BoolVar(&opts.Debug, "debug", false, "verbose output")
	fs.BoolVar(&opts.Publish, "publish", false, "publish blogpost")
	fs.BoolVar(&opts.EditMode, "edit", false, "edit instead of add")
	fs.Int64Var(&opts.GasWanted, "gas-wanted", 2000000, "gas requested for tx")
	fs.StringVar(&opts.GasFee, "gas-fee", "1000000ugnot", "gas payment fee")
	fs.StringVar(&opts.ChainID, "chainid", "staging", "")
	fs.StringVar(&opts.PkgPath, "pkgpath", "gno.land/r/gnoland/blog", "blog realm path")

	// keysclient.BaseOptions
	fs.StringVar(&opts.Home, "home", defaultHome, "home directory")
	fs.StringVar(&opts.Remote, "remote", "staging.gno.land:36657", "remote node URL")
	fs.BoolVar(&opts.Quiet, "quiet", false, "for parsing output")
	fs.BoolVar(&opts.InsecurePasswordStdin, "insecure-password-stdin", false, "WARNING! take password from stdin")
	return fs
}

func run(args []string) error {
	var opts publishOpts

	root := &ffcli.Command{
		ShortUsage: "gnoblog-cli KEY <POST...>",
		ShortHelp:  "(re)publish a blogpost",
		FlagSet:    opts.flagSet(),
		Exec: func(ctx context.Context, args []string) error {
			if len(args) < 2 {
				return flag.ErrHelp
			}
			opts.KeyNameOrBech32 = args[0]
			posts := args[1:]
			return doPublish(ctx, posts, opts)
		},
	}

	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	if !opts.Debug {
		log.SetOutput(io.Discard)
	}

	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}

	return nil
}

func doPublish(ctx context.Context, posts []string, opts publishOpts) error {
	log.Println("opts:", string(amino.MustMarshalJSON(opts)))

	kb, err := keys.NewKeyBaseFromDir(opts.Home)
	if err != nil {
		return err
	}
	opts.kb = kb
	info, err := kb.GetByNameOrAddress(opts.KeyNameOrBech32)
	if err != nil {
		return err
	}
	caller := info.GetAddress()
	log.Println("key:", caller)

	gasFee, err := std.ParseCoin(opts.GasFee)
	if err != nil {
		return fmt.Errorf("parse gas fee: %w", err)
	}

	res, err := makeRequest("auth/accounts/"+caller.String(), []byte(opts.PkgPath), opts)
	if err != nil {
		return fmt.Errorf("make request: %w", err)
	}
	var qret struct{ BaseAccount std.BaseAccount }
	amino.MustUnmarshalJSON(res.Data, &qret)
	log.Println("qret", qret)
	opts.accountNumber = qret.BaseAccount.AccountNumber
	opts.sequenceNumber = qret.BaseAccount.Sequence

	io := commands.NewDefaultIO()
	opts.pass, err = io.GetPassword("Enter password.", opts.InsecurePasswordStdin)
	if err != nil {
		return err
	}

	// FIXME: detect if post exists and automatically use ModEditPost instead of ModAddPost
	// FIXME: detect if post needs an update before makings TXs (same content)
	// FIXME: group Msgs in a single TX
	// FIXME: auto-increment sequence number

	for _, postPath := range posts {
		postFile, err := os.Open(postPath)
		if err != nil {
			return fmt.Errorf("open post %q: %w", postPath, err)
		}

		p, err := parsePost(postFile)
		if err != nil {
			return fmt.Errorf("parse post %q: %w", postPath, err)
		}

		// fmt.Printf("%+v\n", p)

		res, err := makeRequest("vm/qfuncs", []byte(opts.PkgPath), opts)
		if err != nil {
			return fmt.Errorf("make request: %w", err)
		}
		var fsigs vm.FunctionSignatures
		amino.MustUnmarshalJSON(res.Data, &fsigs)
		log.Println("fsigs", fsigs)

		/*
			   msg := vm.MsgCall{
				Caller:  caller,
				PkgPath: opts.PkgPath,
				Func:    "Render",
				Args:    []string{"p/" + p.Slug},
			  }
		*/
		verb := "ModAddPost"
		if opts.EditMode {
			verb = "ModEditPost"
		}

		msg := vm.MsgCall{
			Caller:  caller,
			PkgPath: opts.PkgPath,
			Func:    verb,
			Args: []string{
				p.Slug,
				p.Title,
				p.Body,
				strings.Join(p.Tags, ","),
			},
		}
		tx := std.Tx{
			Msgs:       []std.Msg{msg},
			Fee:        std.NewFee(opts.GasWanted, gasFee),
			Signatures: nil,
			Memo:       "from gnoblog-cli",
		}

		log.Println("tx", string(amino.MustMarshalJSON(tx)))
		bres, err := broadcastTx(tx, opts)
		if err != nil {
			return err
		}
		data := string(bres.DeliverTx.Data)
		println("DATA", data)
	}

	return nil
}

func broadcastTx(tx std.Tx, opts publishOpts) (res *ctypes.ResultBroadcastTxCommit, err error) {
	cli := rpcclient.NewHTTP(opts.Remote, "/websocket")

	signers := tx.GetSigners()
	if tx.Signatures == nil {
		log.Println("signers", signers)
		for range signers {
			tx.Signatures = append(tx.Signatures, std.Signature{
				PubKey:    nil,
				Signature: nil,
			})
		}
	}

	err = tx.ValidateBasic()
	if err != nil {
		return nil, fmt.Errorf("validateBasic: %w", err)
	}

	signbz := tx.GetSignBytes(opts.ChainID, opts.accountNumber, opts.sequenceNumber)
	log.Printf("sign bytes: %X", signbz)

	sig, pub, err := opts.kb.Sign(opts.KeyNameOrBech32, opts.pass, signbz)
	if err != nil {
		return nil, err
	}

	addr := pub.Address()
	found := false
	for i := range tx.Signatures {
		if signers[i] == addr {
			found = true
			tx.Signatures[i] = std.Signature{
				PubKey:    pub,
				Signature: sig,
			}
		}
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("addr %v (%s) not in signer set", addr, opts.KeyNameOrBech32))
	}

	log.Println("tx", string(amino.MustMarshalJSON(tx)))

	bz := amino.MustMarshal(tx)

	bres, err := cli.BroadcastTxCommit(bz)
	if err != nil {
		return nil, err
	}

	// opts.sequenceNumber++

	if bres.CheckTx.IsErr() {
		return nil, errors.New("check transaction failed %+v\nlog %s", bres, bres.CheckTx.Log)
	}
	if bres.DeliverTx.IsErr() {
		return nil, errors.New("deliver transaction failed %+v\nlog %s", bres, bres.DeliverTx.Log)
	}
	log.Println("OK, data", string(bres.DeliverTx.Data))
	return bres, nil
}

func makeRequest(qpath string, data []byte, opts publishOpts) (res *abci.ResponseQuery, err error) {
	opts2 := rpcclient.ABCIQueryOptions{
		// Height:
		// Prove:
	}
	cli := rpcclient.NewHTTP(opts.Remote, "/websocket")
	qres, err := cli.ABCIQueryWithOptions(qpath, data, opts2)
	if err != nil {
		return nil, err
	}
	if qres.Response.Error != nil {
		log.Printf("Log: %s\n", qres.Response.Log)
		return nil, qres.Response.Error
	}
	return &qres.Response, nil
}
