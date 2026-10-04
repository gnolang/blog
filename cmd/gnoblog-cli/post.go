package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto"
)

const txMemo = "Posted from gnoblog-cli"

type cliCfg struct {
	Publish       bool
	Edit          bool
	GasWanted     int64
	GasAdjustment float64
	GasPrice      string
	GasFee        string
	ChainId       string
	BlogRealmPath string
	Target        string
	DryRun        bool

	KeyName               string
	GnoHome               string
	Remote                string
	Quiet                 bool
	InsecurePasswordStdIn bool
	ConfirmInput          bool
}

func newPostCommand(io commands.IO) *commands.Command {
	var (
		fs  = flag.NewFlagSet("Post", flag.ExitOnError)
		cfg = &cliCfg{}
	)

	// Register flagset
	cfg.RegisterFlags(fs)

	// Make Post command
	return commands.NewCommand(
		commands.Metadata{
			Name:       "post",
			ShortUsage: "post <FILE OR FILES_DIR> [flags]",
			LongHelp:   `Post one or more files. Passing in a file will post that single file, while passing in a directory will search for all README.md files and batch post them.`,
		},
		cfg,
		func(_ context.Context, args []string) error {
			return execPost(io, args, cfg)
		},
	)
}

// RegisterFlags registers flags for cliCfg
func (cfg *cliCfg) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.KeyName,
		"key",
		"",
		"name of keypair to use for deployment",
	)
	fs.BoolVar(&cfg.Publish,
		"publish",
		false,
		"publish blogpost",
	)
	fs.BoolVar(&cfg.Edit,
		"edit",
		false,
		"edit mode",
	)
	fs.Int64Var(&cfg.GasWanted,
		"gas-wanted",
		0,
		"gas requested per tx (0 = auto-estimate via simulation)",
	)
	fs.Float64Var(&cfg.GasAdjustment,
		"gas-adjustment",
		1.3,
		"multiplier applied to the simulated gas estimate",
	)
	fs.StringVar(&cfg.GasPrice,
		"gas-price",
		"1ugnot/1000gas",
		"gas price used to derive the fee when --gas-fee is unset",
	)
	fs.StringVar(&cfg.GasFee,
		"gas-fee",
		"",
		"flat gas payment fee (overrides --gas-price when set)",
	)
	fs.StringVar(&cfg.ChainId,
		"chainid",
		"dev",
		"chain ID (auto-set by --target)",
	)
	fs.StringVar(&cfg.Target,
		"target",
		"",
		"gno.land web endpoint to auto-discover --remote and --chainid (e.g. gno.land)",
	)
	fs.BoolVar(&cfg.DryRun,
		"dry-run",
		false,
		"simulate and print the tx plan without broadcasting",
	)
	fs.StringVar(&cfg.BlogRealmPath,
		"pkgpath",
		"gno.land/r/gnoland/blog",
		"blog realm path",
	)
	fs.StringVar(&cfg.GnoHome,
		"home",
		gnoenv.HomeDir(),
		"home directory",
	)
	fs.StringVar(&cfg.Remote,
		"remote",
		"localhost:26657",
		"remote node URL",
	)
	fs.BoolVar(&cfg.InsecurePasswordStdIn,
		"insecure-password-stdin",
		false,
		"WARNING! take password from stdin",
	)
	fs.BoolVar(&cfg.ConfirmInput,
		"confirm-input",
		false,
		"ask user to confirm input",
	)
}

func execPost(io commands.IO, args []string, cfg *cliCfg) error {
	if len(args) != 1 {
		return ErrInvalidNumberOfArgs
	}

	if cfg.KeyName == "" {
		return ErrEmptyKeyName
	}

	// Auto-discover remote & chainid from a gno.land web endpoint if requested
	if cfg.Target != "" {
		rpc, chainID, err := discoverTarget(cfg.Target)
		if err != nil {
			return fmt.Errorf("discovering target %q: %w", cfg.Target, err)
		}
		cfg.Remote = rpc
		cfg.ChainId = chainID
		fmt.Printf("discovered target %q -> remote=%s chainid=%s\n", cfg.Target, rpc, chainID)
	}

	// Ask user for confirming the ChainID
	if cfg.ConfirmInput && !askForConfirmation(cfg.ChainId, cfg.KeyName) {
		fmt.Println("operation canceled by the user, exiting")
		return nil
	}

	// Stat passed in arg
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		return fmt.Errorf("unable to stat %q: %w", args[0], err)
	}

	// Gather candidate post paths (a directory is searched for README.md files)
	var paths []string
	if fileInfo.IsDir() {
		paths, err = findFilePaths(args[0])
		if err != nil {
			return err
		}
	} else {
		paths = []string{args[0]}
	}

	// Initialize the RPC client first (no password needed) so we can resolve
	// which posts are new and show them before the user unlocks their key.
	rpc, err := initRPCClient(cfg)
	if err != nil {
		return err
	}
	queryClient := gnoclient.Client{RPCClient: rpc}

	// Plan which posts will be added/edited, and show them before the password
	plans, err := planPosts(queryClient, cfg, paths)
	if err != nil {
		return err
	}
	if len(plans) == 0 {
		return fmt.Errorf("%w, exiting", ErrNoNewOrChangedPosts)
	}
	printPlan(cfg, plans)

	// Now unlock the key and sign
	var pass string
	if cfg.Quiet {
		pass, err = io.GetPassword("", cfg.InsecurePasswordStdIn)
	} else {
		pass, err = io.GetPassword("Enter password:", cfg.InsecurePasswordStdIn)
	}
	if err != nil {
		return err
	}

	signer, err := initSigner(cfg, pass)
	if err != nil {
		return err
	}

	client := gnoclient.Client{
		Signer:    signer,
		RPCClient: rpc,
	}

	return post(client, cfg, plans)
}

func post(c gnoclient.Client, cfg *cliCfg, plans []postPlan) error {
	// Get account info
	account, err := c.Signer.Info()
	if err != nil {
		return fmt.Errorf("getting signer info failed: %w", err)
	}

	address := account.GetAddress()

	signingAcc, _, err := c.QueryAccount(address)
	if err != nil {
		return fmt.Errorf("query account %q failed: %w", account, err)
	}

	baseNonce := signingAcc.GetSequence()
	accNumber := signingAcc.GetAccountNumber()

	// Build a message for every planned post (the caller is only known now that
	// the key is unlocked)
	msgs := make([]vm.MsgCall, len(plans))
	for i, pl := range plans {
		msgs[i] = pl.toMsgCall(address, cfg.BlogRealmPath)
	}

	// Split the messages into transactions that respect the chain's per-tx size
	// limit, so a growing blog never bounces off the MaxTxBytes wall.
	limits := queryChainLimits(c)
	sizeLimit := int(float64(limits.maxTxBytes) * txSizeSafetyRatio)
	batches := packBatches(msgs, sizeLimit)

	action := "posted"
	if cfg.Edit {
		action = "edited"
	}

	fmt.Printf("Prepared %d post(s) across %d transaction(s) (chain limits: maxTxBytes=%d, maxGas=%d).\n",
		len(msgs), len(batches), limits.maxTxBytes, limits.maxGas)

	total := 0
	for i, batch := range batches {
		nonce := baseNonce + uint64(i)

		// In a real run, each broadcast commits before the next batch, so the
		// on-chain sequence advances with the batch index. In a dry run nothing
		// is broadcast, so simulation must keep signing against the base nonce.
		estNonce := nonce
		if cfg.DryRun {
			estNonce = baseNonce
		}

		// Resolve gas: honor an explicit --gas-wanted, otherwise simulate.
		gasWanted := cfg.GasWanted
		if gasWanted <= 0 {
			estimated, err := estimateBatchGas(c, batch, accNumber, estNonce, limits)
			if err != nil {
				return fmt.Errorf("estimating gas for batch %d/%d: %w", i+1, len(batches), err)
			}
			gasWanted = int64(float64(estimated) * cfg.GasAdjustment)
		}
		if gasWanted > limits.maxGas {
			return fmt.Errorf("batch %d/%d needs %d gas, above chain max %d; reduce batch size",
				i+1, len(batches), gasWanted, limits.maxGas)
		}

		// Resolve fee: honor an explicit --gas-fee, otherwise derive from price.
		gasFee := cfg.GasFee
		if gasFee == "" {
			gasFee, err = deriveGasFee(gasWanted, cfg.GasPrice)
			if err != nil {
				return err
			}
		}

		txCfg := gnoclient.BaseTxCfg{
			GasFee:         gasFee,
			GasWanted:      gasWanted,
			AccountNumber:  accNumber,
			SequenceNumber: nonce,
			Memo:           txMemo,
		}

		tx, err := gnoclient.NewCallTx(txCfg, batch...)
		if err != nil {
			return fmt.Errorf("building tx for batch %d/%d: %w", i+1, len(batches), err)
		}

		signedTx, err := c.SignTx(*tx, accNumber, nonce)
		if err != nil {
			return fmt.Errorf("signing batch %d/%d: %w", i+1, len(batches), err)
		}

		// Guard against a batch that would exceed the tx-size limit.
		if encoded := int64(len(amino.MustMarshal(signedTx))); encoded > limits.maxTxBytes {
			return fmt.Errorf("batch %d/%d encodes to %d bytes, above chain max %d; reduce batch size",
				i+1, len(batches), encoded, limits.maxTxBytes)
		}

		if cfg.DryRun {
			fmt.Printf("[dry-run] batch %d/%d: %d post(s), gas-wanted=%d, gas-fee=%s\n",
				i+1, len(batches), len(batch), gasWanted, gasFee)
			total += len(batch)
			continue
		}

		if _, err := c.BroadcastTxCommit(signedTx); err != nil {
			return fmt.Errorf("broadcasting batch %d/%d: %w", i+1, len(batches), err)
		}

		total += len(batch)
		fmt.Printf("batch %d/%d: %s %d post(s) (gas-wanted=%d, gas-fee=%s)\n",
			i+1, len(batches), action, len(batch), gasWanted, gasFee)
	}

	if cfg.DryRun {
		fmt.Printf("[dry-run] would have %s %d post(s) across %d transaction(s).\n", action, total, len(batches))
	} else {
		fmt.Printf("Successfully %s %d post(s) across %d transaction(s)!\n", action, total, len(batches))
	}

	return nil
}

// postPlan is a parsed post together with the realm function that will be called
// for it: ModAddPost for a new post, ModEditPost in edit mode.
type postPlan struct {
	post *Post
	verb string
}

func (p postPlan) toMsgCall(caller crypto.Address, pkgPath string) vm.MsgCall {
	return vm.MsgCall{
		Caller:  caller,
		Send:    nil,
		PkgPath: pkgPath,
		Func:    p.verb,
		Args: []string{
			p.post.Slug,
			p.post.Title,
			p.post.Body,
			p.post.PublicationDate.Format(time.RFC3339),
			strings.Join(p.post.Authors, ","),
			strings.Join(p.post.Tags, ","),
		},
	}
}

// planPosts parses each candidate file and, using read-only chain queries,
// decides which posts need to be sent: new posts in normal mode, or existing
// posts in edit mode. It needs no signer, so the plan can be shown before the
// user unlocks their key. Posts already on chain are skipped in normal mode, so
// a batch never fails because a single message references a published slug.
func planPosts(c gnoclient.Client, cfg *cliCfg, paths []string) ([]postPlan, error) {
	plans := make([]postPlan, 0, len(paths))

	for _, postPath := range paths {
		postFile, err := os.Open(postPath)
		if err != nil {
			return nil, fmt.Errorf("cannot open file %q: %w", postPath, err)
		}

		p, err := parsePost(postFile)
		postFile.Close()
		if err != nil {
			fmt.Printf("skipping post at %q, cannot parse: %v\n", postPath, err)
			continue
		}

		// Define function to call on the blog realm
		verb := "ModAddPost"

		// Check if Post already exists on chain
		existsExpr := "PostExists(\"" + p.Slug + "\")"
		exists, _, err := c.QEval(cfg.BlogRealmPath, existsExpr)
		if err != nil {
			slog.Error("error while checking if Post exists", "error", err, "slug", p.Slug)
		}

		bExists := strings.Contains(exists, "true")
		if cfg.Edit {
			if !bExists {
				return nil, fmt.Errorf("%s is not on chain yet - disable the edit flag", p.Title)
			}
			// If Post exists, and user wants to edit it, use ModEditPost
			verb = "ModEditPost"
		} else if bExists {
			// if a post is already on chain, and we are not editing it, just skip it
			continue
		}

		plans = append(plans, postPlan{post: p, verb: verb})
	}

	return plans, nil
}

// printPlan lists the posts that will be added or edited, before the password
// prompt, so the user can review them before unlocking their key.
func printPlan(cfg *cliCfg, plans []postPlan) {
	action := "add"
	if cfg.Edit {
		action = "edit"
	}

	fmt.Printf("%d post(s) to %s on %s (chainid=%s):\n", len(plans), action, cfg.BlogRealmPath, cfg.ChainId)
	for _, pl := range plans {
		fmt.Printf("  - %s [%s]\n", pl.post.Title, pl.post.Slug)
	}
}

func parsePost(reader io.Reader) (*Post, error) {
	var p Post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid Post frontmatter: %w", err)
	}

	body := string(rest)
	p.Title, err = extractTitle(body)
	if err != nil {
		return nil, err
	}

	p.Body = removeTitle(body, p.Title)

	if len(p.Tags) != 0 {
		p.Tags = removeWhitespace(p.Tags)
	}

	if p.PublicationDate == nil {
		now := time.Now()
		p.PublicationDate = &now
	}

	return &p, nil
}
