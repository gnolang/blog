package main

import (
	"context"
	"fmt"
	"math"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// Fallback chain limits, used when ConsensusParams can't be queried. These match
// the gnoland-1 mainnet values at the time of writing.
const (
	defaultMaxTxBytes int64 = 1_000_000
	defaultMaxGas     int64 = 3_000_000_000

	// txSizeSafetyRatio leaves headroom in each transaction for the signature
	// and amino framing that aren't counted when measuring individual messages.
	txSizeSafetyRatio = 0.8
)

// chainLimits holds the per-transaction constraints the batching logic respects.
type chainLimits struct {
	maxTxBytes int64
	maxGas     int64
}

// queryChainLimits reads the live block limits from the chain, falling back to
// sane mainnet defaults when the query fails.
func queryChainLimits(c gnoclient.Client) chainLimits {
	limits := chainLimits{maxTxBytes: defaultMaxTxBytes, maxGas: defaultMaxGas}

	res, err := c.RPCClient.ConsensusParams(context.Background(), nil)
	if err != nil || res == nil || res.ConsensusParams.Block == nil {
		return limits
	}

	if b := res.ConsensusParams.Block; b.MaxTxBytes > 0 {
		limits.maxTxBytes = b.MaxTxBytes
		if b.MaxGas > 0 {
			limits.maxGas = b.MaxGas
		}
	}

	return limits
}

// packBatches greedily groups messages so each batch's encoded size stays under
// sizeLimit. A single message larger than the limit is placed in its own batch
// (and is later rejected with a clear error when the signed tx exceeds the cap).
func packBatches(msgs []vm.MsgCall, sizeLimit int) [][]vm.MsgCall {
	var (
		batches [][]vm.MsgCall
		current []vm.MsgCall
		curSize int
	)

	for _, msg := range msgs {
		msgSize := len(amino.MustMarshal(msg))
		if len(current) > 0 && curSize+msgSize > sizeLimit {
			batches = append(batches, current)
			current = nil
			curSize = 0
		}
		current = append(current, msg)
		curSize += msgSize
	}

	if len(current) > 0 {
		batches = append(batches, current)
	}

	return batches
}

// deriveGasFee computes the total gas fee coin from a gas price and gas wanted,
// e.g. gasWanted=50000000 with gasPrice "1ugnot/1000gas" -> "50000ugnot".
func deriveGasFee(gasWanted int64, gasPrice string) (string, error) {
	gp, err := std.ParseGasPrice(gasPrice)
	if err != nil {
		return "", fmt.Errorf("invalid gas price %q: %w", gasPrice, err)
	}

	// fee = ceil(gasWanted * price.Amount / gas)
	fee := int64(math.Ceil(float64(gasWanted) * float64(gp.Price.Amount) / float64(gp.Gas)))
	if fee < 1 {
		fee = 1
	}

	return fmt.Sprintf("%d%s", fee, gp.Price.Denom), nil
}

// estimateBatchGas simulates a batch on the chain and returns the gas the node
// reports it needs. The simulated tx is signed with the given sequence number so
// it matches the account's on-chain state at estimation time.
func estimateBatchGas(c gnoclient.Client, batch []vm.MsgCall, accNumber, sequence uint64, limits chainLimits) (int64, error) {
	simCfg := gnoclient.BaseTxCfg{
		GasFee:         "1000000ugnot", // irrelevant to simulation, just needs to parse
		GasWanted:      limits.maxGas,
		AccountNumber:  accNumber,
		SequenceNumber: sequence,
		Memo:           txMemo,
	}

	tx, err := gnoclient.NewCallTx(simCfg, batch...)
	if err != nil {
		return 0, err
	}

	signedTx, err := c.SignTx(*tx, accNumber, sequence)
	if err != nil {
		return 0, err
	}

	return c.EstimateGas(signedTx)
}
