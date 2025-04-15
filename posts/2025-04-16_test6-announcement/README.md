---
publication_date: 2025-04-16T00:00:00Z
slug: test6-announcement
tags: [gnoland, ecosystem, updates, testnet, test6]
authors: [michelleellen]
---

# Announcing Test6: The Latest gno.land Testnet


[![banner](https://gnolang.github.io/blog/2025-04-16_test6-announcement/src/thumbs/banner.jpg)](https://gnolang.github.io/blog/2025-04-16_test6-announcement/src/banner.jpg)

The newest gno.land testnet, [Test6](https://test6.testnets.gno.land/), is now live! This major milestone brings us closer to the anticipated beta mainnet launch. Test6 introduces essential features and infrastructure improvements designed to enhance network stability, and introduce critical pieces like governance, a token locking mechanism, and the new validator operator (valoper) realm ahead of the Beta Mainnet rollout.

Key Enhancements in Test6 include: 

- GovDAO V3: The latest iteration of gno.land’s primary decentralized governance model.
- Token Transfer Locking Mechanism: A controlled approach to token distribution before full transferability at the time of the feature complete mainnet launch.
- Valoper Registry: An on-chain, permissionless registry for valopers (validator operators) to register their public profile to be listed in the gno.land validator registry.
- GnoVM Enhancements: Performance improvements to the Gno Virtual Machine (GnoVM) for greater efficiency, security and scalability.

## GovDAO V3: Advancing gno.land’s Governance Model

The Test6 GovDAO will launch with all 3 Tiers with the intent to test them out for Beta Mainnet. These tier members are tasked with governance oversight and advancing the progressive decentralization of governance and will model what could occur in the Beta Mainnet.

GovDAO V3 overhauls gno.land’s previous governance iterations to support the finalized governance [specifications](https://gist.github.com/jaekwon/918ad325c4c8f7fb5d6e022e33cb7eb3) by implementing:

- Membership Tiers (T1, T2, T3) with distinct requirements.
- Voting power assignments based on tier membership.
- Modular governance logic, ensuring adaptability and structured decentralization.

[GovDAO V3](https://test6.testnets.gno.land/r/gov/dao/) consists of several key components that ensure a structured and decentralized governance system to manage gno.land. The GovDAO is fitted with an implemented upgrade mechanism, enabling governance-led upgrades through proposals. The GovDAO [member store](https://test6.testnets.gno.land/r/gov/dao/v3/memberstore) acts as a registry for GovDAO members to manage member data and tier assignments today and in future GovDAO versions. Lastly, the GovDAO V3 [implementation realm](https://test6.testnets.gno.land/r/gov/dao) holds the governance operational logic, and proposal statuses that can be migrated to new GovDAO instances.

## Token Transfer Locking Mechanism

During the Beta Mainnet phase, GNOT token distribution will occur, but direct token transfers will remain disabled. To facilitate this, a token locking mechanism has been implemented, ensuring that tokens can be distributed while disallowing transfers until the Beta Mainnet GovDAO puts up a proposal and votes positively to indicate that the network is ready.

Once tokens are distributed at beta mainnet, token holders will be able to use their GNOT on gno.land to interact with the chain, and make realm and package deployments.

For Test6, you can use the [Faucet Hub](https://faucet.gno.land/) to request tokens and start using the network.

## Valoper Registry Realm

The Valoper Registry Realm is a public directory of validator profiles, providing transparency and essential information for GovDAO members to evaluate new validator candidates to add to the gno.land valset.

If validators are interested in participating on the Test6 network, you can register your public validator profile to the [registry](https://test6.testnets.gno.land/r/gnoland/valopers). Two informational guides are available to help you started on your way to becoming a [testnet validator](https://gnops.io/articles/guides/become-testnet-validator/) and provide details on the [hardware specifications](https://gnops.io/articles/effective-gnops/validator-specs/) for gno.land.

You will need to include your Discord handle in the registry in order to receive an invite to the valoper Discord channel. This is where you will signal your interest in joining Test6, and a GovDAO member(s) will be responsible for making and voting on the proposals to add new validators to the Test6 network. This process is a step in testing validator onboarding and the governance system leading up to beta mainnet.

Please note: registering your validator profile and validating on Test6 does not guarantee a validator slot on the beta mainnet, or any future gno.land networks. However, active participation and contributions to testnets will help establish credibility and may improve chances for future validator acceptance. The initial validator amount and valset will ultimately be selected through GovDAO T1 governance proposals and acceptance.

The launch of Test6 signals the final groundwork needed to achieve the Beta Mainnet deliverables outlined in the Road to Mainnet and Beyond [post](https://gno.land/r/gnoland/blog:p/road-to-mainnet)
