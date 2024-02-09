---
title: "Building Gno.land - Proof of Contribution II"
publication_date: 2024-01-26T13:37:00Z
slug: bgl-poc2
tags: [gnoland, gnovm, tm2, PoC]
authors: [christina]
---

## II. Proof of Contribution vs Proof of Stake

Proof of Stake (PoS) is a robust consensus mechanism that provides a more environmentally friendly and scalable alternative to Proof of Work (PoW) and powers most of the web3 industry today. As PoS pioneers, Cosmos technology secures hundreds of blockchain projects and billions of dollars of digital assets, and Ethereum (launched as a PoW chain in 2015) made the historic switch to PoS in 2022. According to [ethereum.org](https://ethereum.org/en/developers/docs/consensus-mechanisms/pos), PoS is “more secure, less energy-intensive, and better for implementing new scaling solutions compared to the previous proof-of-work architecture.” However, as we briefly discussed in [*What Is Proof of Contribution?*](https://test3.gno.land/r/gnoland/blog:p/bgl-poc-1), PoS has vulnerabilities that can corrupt the network over time.

### The Limitations of Proof of Stake (PoS)

Beyond securing the network, the main goal of any consensus mechanism (PoW, PoS, DPoS, PoC, etc.) is to be as decentralized as possible and not reliant on any central actors. This can be measured by the Satoshi Score (or the Nakamoto coefficient), a quantitative measure that assesses a blockchain’s level of decentralization by calculating the minimum number of nodes needed to compromise a network or carry out a 51% attack. PoS systems can be bootstrapped within days (or even hours), starting off decentralized and achieving a high Satoshi Score.

The PoS chain Genesis allocates a default voting power to ~20-50 nodes, in general equally (or at least making sure that no single node has more than 5% of the voting power). This makes PoS chains decentralized enough (in theory) from block 0 with a near-perfect Satoshi score. However, in practice, PoS has two main issues. Because the system is dictated by money, PoS chains become imperfect over time. Anyone wealthy enough can stake their tokens progressively and use their accumulated power to sway decision-making on the chain—or take the network over completely.

The chain can limit the maximum voting power per validator node, but this is almost ineffective, as a malicious actor can carry out a Sybil attack on the network and create multiple validators to bypass the voting cap. Such an attack renders the max voting power per node useless and leaves the chain defenseless against a single organization or cartel gaining the majority of the voting power. PoS systems leave chains like Cosmos Hub and Ethereum at risk from such bad actors, cartels, and powerful protocols (such as Lido and Rocket Pool).

While Proof of Contribution (PoC) can’t prevent Sybil attacks on standard user accounts (when malicious actors create multiple accounts with a single computer and transfer tokens within a few hours), it does make it almost impossible for validator nodes to suffer Sybil attacks. Since the community vets every person who is given voting power or sway in the network (including validator power) through the DAO, at no point can anyone "spoof" identities and gain major sway. 

### Where Proof of Contribution (PoC) Excels

PoC is actually Proof of Authority (PoA) which, instead of offering up a resource like computing power or a financial stake, relies on validators staking their reputation. Anyone can join most public PoW and PoS networks without revealing their identity. However, by definition, PoA validators need to make themselves known and are selected based on their trustworthiness. This means PoA tends to work better when deployed in private or permissioned blockchains than in public platforms (because of this tendency toward centralization). 

PoC solves this problem, ensuring the network becomes increasingly decentralized over time by being governed by a decentralized entity, GovDAO. Like standard PoA chains, PoC chains launch with a handful of validators that must be identified and trusted by the network, meaning governance is centralized at the start, and the chain achieves a low Satoshi Score. The system is about contributing and earning contribution units, which are slow to gain and require human interaction. It takes months (or years) before there are enough actors in the DAO and sufficient voting power for the chain to be considered decentralized enough, according to the Nakamoto coefficient. 

PoC is thus slower to bootstrap than PoS and harder to achieve. You can think of PoC versus PoS as a marathon versus a sprint, whereby PoC starts slowly but then gains momentum over time, and PoS starts quickly but loses momentum over time (the graph below provides a visual representation of PoC versus PoS). 

[![Graph](https://gnolang.github.io/blog/2024-01-26_bgl-poc2/src/thumbs/graph-container.png)](https://gnolang.github.io/blog/2024-01-26_bgl-poc2/src/graph-container.png)

The GovDAO that owns the chain has a mandate to scale (to grow and decentralize) continuously as it adds more contributors. This means it becomes progressively larger over time, achieving high decentralization efficiency way beyond the initial fast sprint of PoS chains. Once established as a proven consensus mechanism and alternative to PoS, GovDAO can benefit from by any blockchain project (through an evolution of ICS) wanting to achieve decentralization and sustainability—PoC can secure Gno.land and the web3 industry at large.

### Security-Conscious by Design

Another advantage of PoC is that because it’s reliant on human interactions, it is more Sybil-resistant by design. As discussed, it’s almost impossible to split a validator node into two (or more) nodes, making conducting a Sybil attack infinitely difficult. Since contribution units are not transferrable or exchangeable, PoC cannot suffer from whales attempting to purchase voting power quickly. If someone wanted to take over the network, they would need to invest years of their time making meaningful contributions. Their attack would be so slow that it would easily be prevented by humans monitoring the decentralization and adjusting the parameters. 

Moreover, GovDAO will activate and deactivate new validators on request, establish a KYC system for validators, and manage promotions of contributors with votes. This removes the possibility of a takeover happening overnight since the only way to gain validator or voting power is by voting on governance requests, which is slow and managed by humans. This is in contrast to PoS systems which are powerful and fully automated yet defenseless against such coordinated attacks.

Gno.land is built on the very premise that such an attack on a PoC network would never happen as it would be entirely counter-intuitive. Since contributions are not only about expertise but also alignment, it is our hypothesis that longstanding contributors who have invested years of time and brainpower in developing the chain will do their best to protect it rather than destroy it. The DAO system will endure thanks to the mix of expertise and alignment and the amount and frequency of contributions. 

### Concluding Thoughts

Beyond separating voting power from net wealth, a core component of Proof of Contribution (PoC) is its focus on long-term sustainability. PoC makes the system fairer and more sustainable, ensuring participants are aligned and take actions that benefit the community and the broader ecosystem. PoC is slower to bootstrap and harder to achieve than PoS but focuses on long-term alignment and security. 

Unlike PoS, contributors receive rewards based on their contribution effort rather than how many tokens they stake. They are thus incentivized and recognized for the quality of their work, ideas, and alignment, driving participation and active engagement. Governance is allocated to the people most likely to care for the ecosystem’s long-term success—the contributors who have spent the most time working toward it.

*II. Proof of Contribution vs Proof of Stake is the second in a [series of articles](/r/gnoland/blog:p/bgl-poc1) to dive deeply into the philosophy, vision, mechanics, and work involved in developing a new consensus mechanism for the next generation of smart contract systems. Look out for subsequent editions and additional Building Gno.land series, and let us know what you think! Got questions? Join the Gno.land [Discord](https://discord.com/invite/S8nKUqwkPn) or follow us on [Twitter/X](https://x.com/_gnoland)*


