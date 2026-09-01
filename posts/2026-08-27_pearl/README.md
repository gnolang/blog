---
publication_date: 2026-08-27T12:00:00Z
slug: pearl
tags: [gnoland, ecosystem, updates, pearl, community, blog]
authors: [ryanlee19]
---

# Pearl is Live

As August comes to a close, we're pleased to announce that Pearl, the latest test network for [Gno.land](https://gno.land/), is live. Pearl, the next iteration after [Sapphire](https://github.com/gnolang/gno/releases/tag/chain/sapphire), is a new chain, not a hardfork. Notice: Pearl is an experimental test network.

As it was with Sapphire, no state carries over from the previous testnet, so balances, realms, and registered names all start at zero. To participate, you'll have to receive GNOT from the [faucet](https://faucet.gno.land/) and [re-register your names](https://docs.gno.land/resources/users-and-teams/).

## Chain Details

- Chain ID: `pearl-1`
- Release tag: [`chain/pearl`](https://github.com/gnolang/gno/releases/tag/chain/pearl)
- 85 curated packages at genesis
- 3 founding validators with [operator keyed valoper profiles](https://gnops.io/articles/guides/become-testnet-validator/), [namespace enforcement (`r/sys/names`)](https://gno.land/r/sys/names) live from block 1
- 3 pre-funded faucet accounts
- Protocol-level test transfers enabled
- Genesis built from the repo's [`examples/`](https://github.com/gnolang/gno/tree/master/examples) tree in about a minute, booting in seconds
- Code baseline: master [`c04f8793d`](https://github.com/gnolang/gno/commit/c04f8793dacd3f340a2dda3776340cb62b0eeaa1), 25 commits ahead of Sapphire's baseline
- Target commit: [`c4c72fdd2`](https://github.com/gnolang/gno/commit/c4c72fdd288c757e8da0d93aae867fa479b1b15c)
- Full release comparison: [`sapphire...pearl`](https://github.com/gnolang/gno/compare/tags/chain/sapphire...tags/chain/pearl)
- Release candidate PR: [#6091](https://github.com/gnolang/gno/pull/6091)

At 3 validators weighted 60 power each, one validator going offline sits exactly at the one-third halt boundary. This has been accepted for launch and will be addressed as additional validator partners join.

## What's New in Pearl?

The headline addition in this release is **genesis vesting accounts**. The genesis builder now supports a configurable list of vested accounts, where continuous schedules unlock linearly and delayed schedules unlock as a cliff. Malformed entries, duplicate accounts, and invalid schedules are all caught and rejected before any node runs, and vesting account creation is now exercised on every build. Pearl ships with ten vested test accounts covering the full schedule matrix, including accounts already fully vested at genesis, several linear unlock windows, cliff schedules, a future start date, and one schedule straddling genesis itself.

## Breaking / Consensus Relevant Updates

- Inert-package flow completed, including a `MsgRun` allowlist and delegated params ([#6088](https://github.com/gnolang/gno/pull/6088), [#6093](https://github.com/gnolang/gno/pull/6093))
- `AddPackage` now type checks production files only, while test files are stored and syntax parsed ([#6025](https://github.com/gnolang/gno/pull/6025))
- Three separate cases where coins could be lost or over-authorized around the send envelope are fixed ([#6062](https://github.com/gnolang/gno/pull/6062))
- GovDAO allowlist lockdown, proposal page escaping, and executor disclosure ([#6068](https://github.com/gnolang/gno/pull/6068))
- Per-file Go version now pinned in the consensus type-check ([#5978](https://github.com/gnolang/gno/pull/5978))

## VM Correctness

- Type-switch with a sole `case nil:` fixed ([#5766](https://github.com/gnolang/gno/pull/5766))
- Clause-local names on switch fallthrough ([#6056](https://github.com/gnolang/gno/pull/6056))
- `goto` out of nested loops ([#5963](https://github.com/gnolang/gno/pull/5963))
- Map composite-literal keys inside for loops ([#6037](https://github.com/gnolang/gno/pull/6037))
- `ComputeMapKey` now metered on the realm-restore path ([#5710](https://github.com/gnolang/gno/pull/5710))
- Unsupported tilde operator now rejected at parse time ([#6057](https://github.com/gnolang/gno/pull/6057))

## Performance

- Size-aware encode, decode, and render paths ([commit `acd01fa29`](https://github.com/gnolang/gno/commit/acd01fa29))
- Runtime blocks recycled through a per machine pool ([#5813](https://github.com/gnolang/gno/pull/5813))
- Vestigial `Defer.Parent` dropped ([#5856](https://github.com/gnolang/gno/pull/5856))

## New Features

- `grc721` core rewritten on the `Token`/`PrivateLedger`/`Teller` axis, with stackable metadata, royalty, and enumerable extensions ([#6072](https://github.com/gnolang/gno/pull/6072), [#6073](https://github.com/gnolang/gno/pull/6073))
- GnoVM chain params reader API ([#6094](https://github.com/gnolang/gno/pull/6094))
- p2p: `config.P2P.Seeds` wired into the switch ([#6023](https://github.com/gnolang/gno/pull/6023)), plus a dial loop backoff fix ([#6054](https://github.com/gnolang/gno/pull/6054))
- New audit pattern harness with executable security pattern fixtures ([#5835](https://github.com/gnolang/gno/pull/5835))

## Genesis + Binaries

- [`genesis.json`](https://github.com/gnolang/gno/releases/download/chain/pearl/genesis.json) — verify with `sha256sum` before use:

  ```
  c45fe60c8c8a1f859d9e4d5aad7ce4d100ff0eb78302e71318ba0de481a8dc91  genesis.json
  ```

- Prebuilt `gno`, `gnokey`, `gnoland`, and `gnoweb` binaries for `darwin`/`linux` × `amd64`/`arm64`, with a `CHECKSUMS.txt` covering every asset
- Do you want to regenerate genesis or join as a validator? See the [Pearl deployment docs](https://github.com/gnolang/gno/tree/chain/pearl/misc/deployments/pearl.gno.land)

## Validators

Please refer to the [Pearl validator guide](https://github.com/gnolang/gno/blob/chain/pearl/misc/deployments/pearl.gno.land/VALIDATOR.md) to join as a validator. Further announcements will be made in the [testnet announcements channel](https://discord.gg/4dAT8KbtU) on Discord.

---

### Links
- Gnoweb: https://pearl.testnets.gno.land
- Faucet: https://faucet.gno.land/
- RPC: https://rpc.pearl.testnets.gno.land
- Gnockpit: https://gnockpit.pearl.testnets.gno.land
- Status: https://status.pearl.testnets.gno.land
- Tx-indexer: https://indexer.pearl.testnets.gno.land/graphql

---

### Disclaimer

Tokens on Pearl testnet have no real world value, and participation is for testing purposes only.
