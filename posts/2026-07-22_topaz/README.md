---
publication_date: 2026-07-22T12:00:00Z
slug: topaz
tags: [gnoland, ecosystem, updates, topaz, community, blog]
authors: [ryanlee19]
---

# Topaz is Live

We’re pleased to announce that Topaz, the latest test network for Gno.land, is live. Topaz is the next iteration after Test13, and it is a new chain, not a hardfork. **Notice: Topaz is an experimental test network.**

Because no state carries over from Test13, balances, realms, and registered names all start at zero. You’ll have to receive GNOT from the [faucet](https://topaz.testnets.gno.land/faucet) and re-register your names. Test13 will sunset within the next 24 hours.

### Chain Details

- Chain ID: `topaz-1`
- 90 curated packages at genesis
- 2 founding validators, namespace enforcement (`r/sys/names`) live from block 1
- 10 pre-funded faucets
- Unrestricted token transfers
- Code baseline: master [959cefd91](https://github.com/gnolang/gno/commit/959cefd916021d3a55e9b51f20d05ef618e7f357), 133 commits ahead of Test13

## What’s New in Topaz?

### Breaking / Consensus-relevant Updates
- `fix(grc20)!`: token id decoupled from symbol ([#5908](https://github.com/gnolang/gno/pull/5908))
- secp256k1 validator keys no longer supported ([#5949](https://github.com/gnolang/gno/pull/5949))
- package name must match the last path element at `addpkg` ([#5048](https://github.com/gnolang/gno/pull/5048))
- `cockroachdb/apd` replaced by `math/big.Rat` in the VM ([#5867](https://github.com/gnolang/gno/pull/5867))
- type-check + preprocess gas now metered at `AddPackage/Run` ([#5892](https://github.com/gnolang/gno/pull/5892))

### VM Correctness

Notable fixes from the large batch of gnovm fixes are:
- interface-bound method dispatch ([#5737](https://github.com/gnolang/gno/pull/5737))
- embedded field/method lookup ([#5721](https://github.com/gnolang/gno/pull/5721), [#5739](https://github.com/gnolang/gno/pull/5739))
- deterministic type-check verdicts ([#5893](https://github.com/gnolang/gno/pull/5893))
- tuple assignment order ([#5790](https://github.com/gnolang/gno/pull/5790))
- map-delete object reclaim ([#5882](https://github.com/gnolang/gno/pull/5882), [#5899](https://github.com/gnolang/gno/pull/5899))

### New Features
- sub-realm identities via `realm.Sub(subpath)` ([#5890](https://github.com/gnolang/gno/pull/5890))
- mempackage storage split into prod and test blobs ([#5891](https://github.com/gnolang/gno/pull/5891))
- vesting accounts PoC ([#5840](https://github.com/gnolang/gno/pull/5840))
- tm2 p2p address book ([#5861](https://github.com/gnolang/gno/pull/5861))
- concurrent ABCI queries via snapshot isolation ([#5431](https://github.com/gnolang/gno/pull/5431))
- gnoweb package overview source doc ([#5572](https://github.com/gnolang/gno/pull/5572))

### Genesis / Operator Tooling
- loud genesis-replay failures ([#5924](https://github.com/gnolang/gno/pull/5924))
- `GnoTxMetadata` provenance fields ([#5925](https://github.com/gnolang/gno/pull/5925))
- unified genesis signature-skip for replay txs ([#5926](https://github.com/gnolang/gno/pull/5926))
- `gnogenesis fork` provenance, tx patching + `inspect` ([#5927](https://github.com/gnolang/gno/pull/5927))
- valid, payable `fork valoper-seed` output ([#5928](https://github.com/gnolang/gno/pull/5928))
- `gnogenesis verify -skip-signature-check` ([#5944](https://github.com/gnolang/gno/pull/5944))

### Genesis + Binaries
- [genesis.json](https://github.com/gnolang/gno/releases/download/chain/topaz/genesis.json) (verify with `sha256sum`)
- Prebuilt `gno`, `gnokey`, `gnoland`, `gnoweb` binaries for darwin/linux × amd64/arm64

Want to regenerate genesis? See [here](https://github.com/gnolang/gno/blob/chain/topaz/misc/deployments/topaz.gno.land/README.md).

### Validators
Please refer to this [link](https://github.com/gnolang/gno/blob/chain/topaz/misc/deployments/topaz.gno.land/VALIDATOR.md) to join Topaz as a Validator. Once your validator is up and synced, you will have to register your valoper profile again with the same operator address you used in Test13. This enables us to add you directly to the Topaz valset. The team will soon announce on [Discord](https://discord.gg/gnoland) when Test13 will sunset in order to simplify the validator migration process.

---

#### Links
- Gnoweb: https://topaz.testnets.gno.land
- Faucet: https://topaz.testnets.gno.land/faucet
- RPC: https://rpc.topaz.testnets.gno.land
- Gnockpit: https://gnockpit.topaz.testnets.gno.land
- Status: https://status.topaz.testnets.gno.land
- Tx-indexer: https://indexer.topaz.testnets.gno.land/graphql

**Disclaimer**
*Tokens on Topaz testnet have no real-world value, and participation is for testing purposes only.*

