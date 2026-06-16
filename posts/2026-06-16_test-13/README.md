---
publication_date: 2026-06-16T12:00:00Z
slug: test-13
tags: [gnoland, ecosystem, updates, test13, community, blog]
authors: [ryanlee19]
---

# Test13 is Live

We’re pleased to announce that Test13 is live. As we draw nearer to mainnet launch, Test13 marks a significant milestone in the development of the Gno.land blockchain.

## What’s New in Test13?

Test13 is a consensus-breaking release for the Gno testnet line with the following changes.

#### Interrealm Semantics Overhaul
- Packages under `/p/` and the standard library are now persistent frozen Realms
- Cross-realm panics from `/p/` abort instead of being catchable by `recover()`
- The old `N_Readonly` provenance taint has been removed in favor of PkgID ownership checks at pointer dereference boundaries

#### Additional Updates
- Hardfork v3 support
- GovDAO-driven chain halt handling
- Coordinated `halt_height` upgrades
- Validator-set limits
- Session accounts
- Improved genesis/replay tooling
- Calibrated gas and storage changes
- Safer markdown handling
- Gnoweb search and UX improvements
- Gnokey behavior updates
- Broad VM/RPC/web/security hardening

Read the full [change log here](https://github.com/gnolang/gno/releases/tag/chain/test13).

## What Test13 Means for Our Community

In the practical sense, Test13 opens up a range of possibilities for developers:

- Deploy and run applications using faucet tokens with no risk of monetary loss
- Iterate on Gno code, realms, and packages in a production-like setting
- Onboard early users and gather feedback before deploying on mainnet
- Turn ideas into decentralized applications and test product-market fit
- Estimate costs of building certain applications to determine the financial feasibility of deploying them

For the broader community, Test13 is a chance to explore Gno.land firsthand. Keep an eye out for DeFi and gaming experiences from ecosystem projects like [GnoSwap](https://x.com/gnoswaplabs) and [Akkadia](https://x.com/akkadialand) as they come online. As the ecosystem grows, so will the ways to engage with Gno.land. For our validators, we’ll give you an update on what needs to be done next.

To get started, visit [Test13](https://test13.testnets.gno.land/), receive test GNOT through our [faucet](https://test13.testnets.gno.land/faucet), and explore available [RPC endpoints](https://rpc.test13.testnets.gno.land/) if needed.

## Why Test13 Matters

Test13 is a significant step forward in Gno.land's journey toward mainnet. The release delivers foundational improvements across security, gas modeling, VM correctness, and chain upgrade infrastructure, each addressing core engineering challenges and reducing risk for the network in the long term.

Mainnet is when the network becomes permanent, real value starts to flow, and Gno.land’s vision of a transparent and composable blockchain platform becomes reality. Test13 brings that vision meaningfully closer.

## What’s Next?

We've put Test13 through rigorous internal testing, and now it's time to open it up to the community. By deploying contracts, transacting, and exploring the network, you play a direct role in validating that the chain is ready for what comes next.

And what comes next is important. The core team is actively preparing for mainnet, which includes the Token Generation Event (TGE), live GNOT transfers, on-chain governance through GovDAO, Inter-Blockchain Communication (IBC) and Interchain Security (ICS) integrations with AtomOne, continued optimization of the GnoVM, improved developer tooling, and the broader growth of the Gno.land ecosystem.

Welcome to Test13. Join the conversation on [Discord](https://discord.gg/gnoland) and [Telegram](https://t.me/join_gnoland).
