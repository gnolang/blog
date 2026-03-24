---
publication_date: 2026-03-26T00:00:00Z
slug: beta-mainnet
tags: [gnoland, ecosystem, community]
authors: [michelleellen, Ticojohnny]
---

# Gno.land Beta Mainnet is Live

Gno.land’s Beta Mainnet marks the beginning of a new era in application development. Engineered from the ground up, Gno.land is the world’s first language-based, multi-user operating system. It pushes the limits of existing blockchain infrastructure and programming languages by making shared computation, persistent state, and composability native to the system itself.

Similar to the way Cosmos changed how blockchains intercommunicate, and the Tendermint Byzantine Fault Tolerant Consensus Protocol changed how blockchain developers approach consensus algorithms, Gno.land will change the way developers think about programming.

## Gno.land Genesis

Started in 2022 by Jae Kwon, founder of Tendermint and Cosmos, Gno.land set out to bring next-generation smart contract programming to the Cosmos ecosystem. What began as a passion project evolved into a dedicated engineering effort: a novel Layer1 platform, the Gno language (derived from Go), the Gno virtual machine, an application layer, and a universal browser and rendering system to make deployed code human-readable and interactive.

Gno.land is a paradigm shift in both programming languages and operating systems. The goal is to become the world’s open knowledge base for the new millennium, starting with a powerful development environment for building next-level decentralized applications.

## Launch Phases

### **Phase I: The Beta Mainnet (Present)**

The launch of the Beta Mainnet includes the first distribution of the GNOT token, an operational GnoVM-powered smart contract network, and the inaugural release of GovDAO, Gno.land’s on-chain governing system. An initial treasury has been set up, managed by GovDAO members to fund future network development and growth. There will be a period of time where the virtual machine will be hardened and optimized with user provided contracts and automated fuzz tooling (which is currently being developed). Finally, the [Gno.land Constitution](https://github.com/gnolang/gno/blob/master/docs/CONSTITUTION.md) is in a preliminary state, and is expected to be ratified by a Supermajority Decision of GovDAO now that the Beta Mainnet is live.

Note: For now, token transfers are disabled during Beta; GNOT is usable within the network for supported on-chain functions.

### **Phase II: Mainnet**

This phase will enable GNOT transfers. The activation of GNOT transfers will be proposed and voted on by GovDAO. Once transfers are enabled, Gno.land’s economy will begin forming with the community’s help. The priority then becomes growth: expand awareness, attract a core group of new Go developers, prototype dApps at speed, and equip builders with better tools.

We will also see numerous optimizations of the Gno virtual machine both during the Beta Mainnet and Mainnet phases. More to come later!

### **Phase III: Expanding the Network**

The next engineering task following Beta Mainnet, focuses on secure interoperability with a priority on bridging AtomOne and Gno.land. This includes integrating and testing key Cosmos technologies: Inter-Blockchain Communication (IBC) and [Interchain Security (ICS)](https://allinbits.com/blog/rethinking-ics/).

*Timelines for each phase will be communicated as they solidify.*

<img width="1051" height="567" alt="Screenshot 2026-03-24 at 22 00 04" src="https://github.com/user-attachments/assets/f18ba51a-035e-488a-bebc-d15938f490f5" />


## GNOT Distribution Details

Of the 1 billion total GNOT token supply, 70% is being distributed to the community in three categories.

- Cosmos recipients: 35% has been allocated to a partial governance snapshot taken on [May 20th, 2022 (block 10,562,840)](https://www.mintscan.io/cosmos/block/10562840)
- AtomOne recipients: 23.1% has been allocated to ATONE and PHOTON holders from a snapshot dated at [December 25th, 2025 (block 6439117](https://www.mintscan.io/atomone/block/6439117))
- Gno contributors: 11.9% has been allocated for distribution to prior and new Gno.land ecosystem contributors. This category will be managed in the future by a GovDAO subDAO.

*Allocation figures and eligibility are based on snapshots and current records and may be adjusted to correct errors, prevent abuse, or comply with legal requirements.*

From the remaining 30% GNOT supply, 7% has been reserved for current and future private token purchase rounds (if any), subject to applicable law, while the remaining 23% will be managed by NewTendermint, LLC. GovDAO and NewTendermint,LLC together will develop, support and maintain the Gno.land ecosystem.

## Beta Means Beta

Beta Mainnet is a milestone, not the finish line. While the network is stable enough for builders and community members to begin experimenting, occasional chain instability, including network interruptions or restarts, as well as breaking changes may alter existing features or contracts. Data resets *will* occur as development progresses. This is intentional; learning how to recover from these issues and developing tooling to make the process smoother is part of the objective of the beta phase.

The key objectives are to deploy and build realms, packages, and contracts to validate core functionality, stress-test the network to explore its performance limits and validator behavior, and report and document any findings, bugs, or improvement opportunities to the community.

Every bug found, every realm deployed, and every part of Gno.land tested brings the network one step closer to a fully production-ready mainnet.

## Join the Beta Mainnet Experience

There are many ways to contribute, whether you are a developer, a validator, or a network enthusiast, there’s a few ways to push the ecosystem forward.

**Gnomies**

For Cosmos and AtomOne recipients, you can check your GNOT balance with your Cosmos Hub or AtomOne address using the [Coin Balance Checker](https://gno.land/r/gnoland/coins). You have two options for managing your GNOT tokens and interacting with the network:

- Creating or importing your existing wallet through the [Adena Wallet](https://www.adena.app/)
- Using [Gnokey CLI](https://docs.gno.land/users/interact-with-gnokey), the command-line wallet for Gno.land

*WARNING: Never enter your mnemonic (secret 12, 18, or 24 words) on any device that is or will ever be connected to the internet, not even on the Adena Wallet or Gnokey CLI. While you can check your GNOT balance, and use the Gno.land network, token transfers are currently not available.*

If you didn’t receive GNOT, you can still explore Beta Mainnet via the [Faucet Hub](https://faucet.gno.land/) by requesting tokens from a drip faucet. The [Staging network](https://staging.gno.land/) is also available for experimentation using a separate test token.

### **Developers**

If you have the skills to go further at this early stage, you can make meaningful code contributions, tooling, smart contract examples, and killer applications. Your contributions are integral to shaping and bootstrapping the network.

Developers can start by deploying a simple realm, exploring the example contracts, or submitting a pull request either directly in the [Gno repository](https://github.com/gnolang), or in [Gnoverse](https://github.com/gnoverse), a collaborative space for experimental and innovative projects inspired by the Gno.land ecosystem. The best way to jump right in is visiting the [Gno.land Documentation](https://docs.gno.land/)**.**

Contributors to the ecosystem may eventually deploy GRC-20 smart contracts enabling future on-chain swap functionality with [Gnoswap](https://beta.gnoswap.io/), the first decentralized exchange (DEX) on Gno.land.

**Validators**

Validators can join the set, help test scaling, and prepare for future upgrades. Beta Mainnet validators are added to the valset only through a GovDAO proposal. Unlike other networks where validators stake their funds and earn rewards, Gno.land’s validator slots are given to contributors with proven track records. You can register your validator for consideration on the Gno.land [Valoper Registry](https://gno.land/r/gnops/valopers) and learn more [here](https://gnops.io/articles/guides/become-testnet-validator/) on how to get started.

**Everyone**

The Gno.land block explorer, [GnoScan](https://gnoscan.io/), gives you an overview of Gno.land networks like Beta Mainnet as well as Staging, or ongoing testnets. It includes important data for an overview of the network’s status as well as a myriad of transactions made including total storage deposits, realm deployments, and fungible tokens (GRC20 token standard).

## Recognition and the Road Ahead

Phase II is expected to enable token transfers and expand development and user tooling. Phase III will focus on interoperability and building bridges to the Cosmos ecosystem. Each phase builds toward establishing Gno.land as the go-to operating system and platform for a censorship-free and economically viable alternative to today’s broken internet.

Early adopters and contributors will always be remembered within the network. Whether your impact comes through core code, or creative realms, we'll spotlight your work and share it with the wider community. Now is the time. The Beta Mainnet is only the beginning.

Join our community:

- Webpage: https://gno.land/
- GitHub: https://github.com/gnolang/
- Docs: https://docs.gno.land/
- X: https://x.com/_gnoland
- Discord: [Discord.gg/gnoland](http://discord.gg/gnoland)
- Telegram: https://t.me/join_gnoland
- Gno.land Linker: https://gno.land/links

*Disclaimer:*

*This announcement is for informational purposes only and is not an offer to sell or a solicitation to buy any token or any other instrument. It includes forward-looking statements and targets that may change; actual outcomes may differ materially. Participation may be restricted or unavailable in certain jurisdictions. Beta features and network behavior may change, including through upgrades, restarts, or data resets.*
