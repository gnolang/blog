---
publication_date: 2026-08-10T12:00:00Z
slug: sapphire
tags: [gnoland, ecosystem, updates, sapphire, community, blog]
authors: [ryanlee19]
---
# Sapphire is Live

We're pleased to announce that Sapphire, the latest test network for [Gno.land](https://gno.land/), is live. Sapphire is the next iteration after Topaz, and it is a new chain, not a hardfork. ***Notice: Sapphire is an experimental test network.***

As it was with Topaz, no state carries over from the previous testnet, so balances, realms, and registered names all start at zero. To participate, you’ll have to receive GNOT from the [faucet](https://faucet.gno.land/) and re-register your names. Topaz will sunset within the next 24 hours.

### Chain Details
- Chain ID: `sapphire-1`
- 85 curated packages at genesis
- 2 founding validators, namespace enforcement (`r/sys/names`) live from block 1
- 3 pre-funded faucet accounts
- Protocol-level test transfers enabled
- Code baseline: master [1bf8b2826](https://github.com/gnolang/gno/commit/1bf8b2826b7deec7648c6c17851b089cb078d04e), 17 commits ahead of Topaz

A balance-measurement bug caught during [Topaz's PR review](https://github.com/gnolang/gno/pull/5945#pullrequestreview-4733742520) is also fixed in the genesis builder for this release, so fee payer balances are now verified more strictly before a node is considered ready. As a result, the only accounts funded at genesis are the three faucets; every other fee payer starts at zero.

## What's New in Sapphire?

### Breaking / Genesis-composition-relevant Updates
- commondao aligned with the Common DAO Spec — Council, default tally, per-DAO treasuries ([#6012](https://github.com/gnolang/gno/pull/6012)); with boards permissions migrated to the new `p/nt/groups/v0` ([#6009](https://github.com/gnolang/gno/pull/6009)), ([#6010](https://github.com/gnolang/gno/pull/6010)), `p/nt/commondao/v0` and four transitive deps left the genesis set (85 packages vs Topaz's 90)
- boards2 public API made non-crossing again; hub moved into `p/gnoland/boards/exts/hub` ([#5809](https://github.com/gnolang/gno/pull/5809))
- tm2: non-gas balances stored in their own keys ([#6034](https://github.com/gnolang/gno/pull/6034))

### New Features
- grc20: `NewToken` event on token creation ([#6042](https://github.com/gnolang/gno/pull/6042))
- grc20reg: transaction wrappers for registered tokens ([#5962](https://github.com/gnolang/gno/pull/5962))
- tm2: configurable RPC keep-alive idle timeout ([#6002](https://github.com/gnolang/gno/pull/6002))
- gnoweb: open render on package-name click, Browse link in explorer ([#5966](https://github.com/gnolang/gno/pull/5966))

### Fixes
- tm2: snapshot-isolated query paths and write-proof bptree fast-index maintenance ([#6018](https://github.com/gnolang/gno/pull/6018))
- gnodev: disable `PeerExchange` ([#5984](https://github.com/gnolang/gno/pull/5984))

### Docs / Operations
- node-operator docs moved into `gno.land/cmd/gnoland`, gnops.io links dropped ([#6039](https://github.com/gnolang/gno/pull/6039))
- "Running a node" routing page ([#6038](https://github.com/gnolang/gno/pull/6038))

## Genesis + Binaries
- [genesis.json](https://github.com/gnolang/gno/releases/download/chain/sapphire/genesis.json) (verify with sha256sum)
- Prebuilt `gno`, `gnokey`, `gnoland`, `gnoweb` binaries for `darwin`/`linux` × `amd64`/`arm64`

Want to regenerate genesis? See [here](https://github.com/gnolang/gno/blob/chain/sapphire/misc/deployments/sapphire.gno.land/README.md).

### Validators
Please refer to this [link](https://github.com/gnolang/gno/blob/chain/sapphire/misc/deployments/sapphire.gno.land/VALIDATOR.md) to join Sapphire as a Validator. Further announcements will be made in the [testnet-announcements channel](https://discord.gg/4dAT8KbtU) on Discord.

---

#### Links

- Gnoweb: https://sapphire.testnets.gno.land
- Faucet: https://faucet.gno.land/
- RPC: https://rpc.sapphire.testnets.gno.land
- Gnockpit: https://gnockpit.sapphire.testnets.gno.land
- Status: https://status.sapphire.testnets.gno.land
- Tx-indexer: https://indexer.sapphire.testnets.gno.land/graphql 

**Disclaimer**
*Tokens on Sapphire testnet have no real-world value, and participation is for testing purposes only.*
