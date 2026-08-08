---
publication_date: 2026-08-10T12:00:00Z
slug: sapphire
tags: [gnoland, ecosystem, updates, sapphire, community, blog]
authors: [ryanlee19]
---
# Sapphire is Live

We're pleased to announce that Sapphire, the latest test network for Gno.land, is live. Sapphire is the next iteration after Topaz, and it is a new chain, not a hardfork. Notice: Sapphire is an experimental test network.

As it was with Topaz, no state carries over from the previous testnet, so balances, realms, and registered names all start at zero. To be a part of Sapphire, you’ll have to receive GNOT from the faucet and re-register your names. Topaz will sunset within the next 24 hours.

### Chain Details
- Chain ID: sapphire-1
- 85 curated packages at genesis
- 2 founding validators, namespace enforcement (r/sys/names) live from block 1
- 3 pre-funded faucet accounts
- Protocol-level test transfers enabled
- Code baseline: master 1bf8b2826, 17 commits ahead of Topaz

The genesis builder also fixes the fee-payer balance-measurement bug found in the topaz PR review: node readiness now gates on committed state, zero balances are only trusted when the account provably exists, and the build asserts measurement/shipping params parity and tx-stream reconciliation. Every fee payer lands at exactly zero post-genesis; the only funded accounts are the three faucets.

## What's New in Sapphire?

### Breaking / Genesis-composition-relevant Updates
- commondao aligned with the Common DAO Spec — Council, default tally, per-DAO treasuries (#6012); with boards permissions migrated to the new p/nt/groups/v0 (#6009, #6010), p/nt/commondao/v0 and four transitive deps left the genesis set (85 packages vs Topaz's 90)
- boards2 public API made non-crossing again; hub moved into p/gnoland/boards/exts/hub (#5809)
- tm2: non-gas balances stored in their own keys (#6034)

### New Features
- grc20: NewToken event on token creation (#6042)
- grc20reg: transaction wrappers for registered tokens (#5962)
- tm2: configurable RPC keep-alive idle timeout (#6002)
- gnoweb: open render on package-name click, Browse link in explorer (#5966)

### Fixes
- tm2: snapshot-isolated query paths and write-proof bptree fast-index maintenance (#6018)
- gnodev: disable PeerExchange (#5984)

### Docs / Operations
- node-operator docs moved into gno.land/cmd/gnoland, gnops.io links dropped (#6039)
- "Running a node" routing page (#6038)

### Genesis + Binaries
- genesis.json (verify with sha256sum)
- Prebuilt gno, gnokey, gnoland, gnoweb binaries for darwin/linux × amd64/arm64

Want to regenerate genesis? See here.

### Validators
Please refer to this link to join Sapphire as a Validator. Further announcements will be made in the testnet-announcements channel on Discord.

**Disclaimer**
*Tokens on Sapphire testnet have no real-world value, and participation is for testing purposes only.*

