---
publication_date: 2026-09-16T12:00:00Z
slug: gnoland-1
tags: [gnoland, ecosystem, updates, mainnet, community, blog]
authors: [ryanlee19]
---

# The Beginning of a Decentralized Future: Gno.land Mainnet is Here

We are pleased to announce that Gno.land's mainnet is live. This marks the beginning of a new, decentralized future, and we are excited to have you join us.

## Chain Details

- Chain ID: `gnoland-1`
- Launch: 2026-09-12T15:00:00Z
- 89 curated packages at genesis, on the `/v0` layout
- Release tag: [chain/mainnet](https://github.com/gnolang/gno/releases/tag/chain/mainnet)
- Endpoints: `gno.land`, `rpc.gno.land`, `seed-1.gno.land`, and `seed-2.gno.land`

## Balances and Vesting

GNOT balances come from the independence-day allocation: 3,262,481 accounts totalling 1.333 billion GNOT, covering airdrops, investor buckets, treasuries, the public sale, and settled partner distributions. The allocation is sha256-pinned and reconciled against genesis at build time.

Transfers were locked at genesis under Constitution section 126, with a 91-address exemption list. GovDAO proposal #4 has since lifted the lock, and ugnot is transferable chain-wide. Nearly every account still carries a section 132 vesting schedule: 3,262,417 of the 3,262,481 rows vest 96% continuously from 2026-09-11 to 2028-09-11.

## Governance

Governance started with a sole GovDAO T1 member (aeddi); the remaining six confirmed members join post-genesis through regular proposals — two have already been seated, bringing T1 to three. Code submission is inert from block 1: any package added after genesis is parked until the gpao (package-approvals oracle) clears it.

## Validators and Namespaces

Four founding validators secure the network at launch: Gnocore, OnBloc, Samourai Crew, and Berty, each running its organization's real ceremony consensus key. Namespace enforcement (`r/sys/names`) is on from block 1.

## What Changed Since Pearl Testnet

Mainnet release focused on hardening the chain for a real, value-bearing launch. Gas metering and package authority APIs were tightened for consensus safety, packages moved to a versioned `/v0` layout, and vesting became a native account field rather than a separate account type.

On the genesis side, the transfer lock, the inert code-submission policy, and the fee-payer and vested-allocation merge were all built specifically for mainnet's launch conditions. Node and VM work focused on performance and correctness: faster genesis loading, new bank transfer events, bounded-parallel queries, and a handful of gas and range-handling bug fixes. The gpao package-approvals oracle also saw meaningful hardening ahead of launch, and a new txtar harness now exercises multi-validator clusters in tests.

The full changelog is in the [GitHub release notes](https://github.com/gnolang/gno/releases/tag/chain/mainnet).

## Genesis and Binaries

`genesis.json` and prebuilt `gno`, `gnokey`, `gnoland`, and `gnoweb` binaries for darwin and linux (amd64/arm64) are attached to the release, along with a `CHECKSUMS.txt` covering every asset. Container images are tagged `chain-mainnet`. To regenerate genesis or join as a validator, see [misc/deployments/mainnet.gno.land/](https://github.com/gnolang/gno/tree/chain/mainnet/misc/deployments/mainnet.gno.land) in the repo.

---

### Links

- GitHub release: [https://github.com/gnolang/gno/releases/tag/chain/mainnet](https://github.com/gnolang/gno/releases/tag/chain/mainnet)
- Gnoweb: https://gno.land
- RPC: https://rpc.gno.land
- Gnockpit: https://gnockpit.gno.land
- Status: https://status.gno.land
- Tx-indexer: https://indexer.gno.land/graphql
