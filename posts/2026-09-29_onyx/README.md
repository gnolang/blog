---
publication_date: 2026-09-29T12:00:00Z
slug: onyx-1
tags: [gnoland, ecosystem, updates, onyx, community, blog]
authors: [ryanlee19]
---

# Onyx Testnet Is Open: Test Your Code Before Deploying on Mainnet

Onyx, Gno.land's new testnet, is now live. Onyx replaces Pearl and is built to give developers a place to test their code under mainnet conditions before deploying it to mainnet.

## Chain Details

- Chain ID: `onyx-1`
- Launch: 28 September 2026, 00:00 UTC
- Launch version: `v1.5.0`, mainnet's binaries, unchanged
- 89 curated packages at genesis on the `/v0` layout, with a package list identical to mainnet's byte for byte
- Release tag: [chain/onyx](https://github.com/gnolang/gno/releases/tag/chain/onyx)
- Endpoints: `onyx.testnets.gno.land`, `rpc.onyx.testnets.gno.land`, `seed-1.onyx.testnets.gno.land`, and `seed-2.onyx.testnets.gno.land`

## Why Onyx

Pearl served Gno.land well through the road to mainnet. With mainnet now live, developers need a testnet that behaves the way mainnet does, so that code that works on the testnet works the same way in production. Onyx runs mainnet's code one release candidate ahead and is upgraded whenever mainnet is, which means every mainnet release is rehearsed on Onyx first.

## Same as Mainnet

Onyx's genesis follows mainnet's shape. It ships the same seven initial namespaces, with namespace enforcement (`r/sys/names`) from block 1. Governance starts from the same sole GovDAO T1 seed (aeddi), and code submission follows the same policy: packages added after genesis go through the gpao (package approvals oracle) before they are cleared, and `maketx run` is restricted to the seeded member. Testing on Onyx means testing against the rules your code will actually face on mainnet.

## Different From Mainnet

Onyx keeps mainnet's rules but uses testnet GNOT. In place of the independence-day allocation, genesis funds four accounts: the web faucet's and the faucet agent's dispensing accounts, aeddi, and the gpao oracle. Transfers are open from genesis, with no section 126 lock and no exemption list, and there is no vesting. The network launches with one founding validator, `gno-core-validator-1`.

## Pearl Sunsetting

With Onyx live, Pearl will be sunset within 24 hours. Onyx is a fresh chain, not a hardfork of Pearl, so packages and balances on Pearl do not carry over. Developers currently building on Pearl will need to redeploy on Onyx.

## Genesis and Binaries

The `chain/onyx` release carries `genesis.json`, `genesis.json.gz`, and `CHECKSUMS.txt`. The genesis sha256 is `4b006fd7ccdec052865accc84dd29b2b76f8b57b2560789a15eedaa88f0e26c5`. Binaries are attached to the version release (`v1.5.0` at launch), not to the chain release.

Validator operators should use the version listed in the Onyx upgrade ledger ([`UPGRADES.md`](https://github.com/gnolang/gno/blob/chain/mainnet/misc/deployments/onyx.gno.land/UPGRADES.md)) and must join with `--skip-genesis-sig-verification`. To regenerate genesis or join as a validator, see [misc/deployments/onyx.gno.land/](https://github.com/gnolang/gno/tree/chain/mainnet/misc/deployments/onyx.gno.land) in the repo.

---

### Links

- GitHub release: https://github.com/gnolang/gno/releases/tag/chain/onyx
- Web: https://onyx.testnets.gno.land
- RPC: https://rpc.onyx.testnets.gno.land
- Faucet: https://onyx.testnets.gno.land/faucet
- Status: https://status.onyx.testnets.gno.land
- Gnockpit: https://gnockpit.onyx.testnets.gno.land
- Docs: https://docs.gno.land
