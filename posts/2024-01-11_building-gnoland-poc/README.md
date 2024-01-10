---
title: "Building Gno.land – Next Generation Smart Contract System"
publication_date: 2024-01-10T10:51:00Z
slug: bgl-poc1
tags: [building-gnoland,gnoland,proof-of-contribution]
author: [christina]
---

*Disclaimer: The Building Gno.land series aims to share the core concepts of our platform, mission, and tech stack, providing a snapshot of a current state on our builder’s journey. New episodes may deprecate previous ones, but no editions will be modified at any time. We encourage you to follow along as we create the rails for a more transparent and accountable society, and we welcome feedback and contributions to the repo.*

# Proof of Contribution (PoC) 

## I. What Is Proof of Contribution (PoC)?

Gno.land is secured by a novel consensus mechanism that makes our platform unique—Proof of Contribution (PoC). PoC prioritizes fairness and merit, rewarding the people most active on the platform and revolutionizing the concept of open-source rewards. By removing the voting power associated with being wealthy (holding tokens in Proof-of-Stake (PoS) networks or amassing mining hardware in Proof-of-Work (PoW) networks), PoC restructures the financial incentives that tend to corrupt blockchain projects in the long run and rewards contributors fairly for their work based on their expertise, commitment, and values. 

Gno.land contributors receive rewards and voting power according to their contribution level. These rewards increase as they make additional contributions, gain expertise, and are promoted up the Gno.land governing DAO’s (GovDAO) tier levels by higher-level contributors. So how does PoC work, what are its core features, and how does it lend security and decentralization to the platform? 

### Prioritizing Fairness and Alignment 

Proof of Stake (PoS) was a monumental leap forward for the blockchain industry, solving the energy-intensive requirements of Proof of Work (PoW) and enabling blockchains to scale for broader adoption (thanks to its minimal carbon footprint and faster throughput). However, like PoW, PoS has some disadvantages. For example, in PoS networks, participants receive rewards based on how many tokens they stake, which means their incentives for working on the chain are often purely financial. Validators accumulate vast net worths and don’t always hold values that align with the core development of the chain. 

Since validators are crucial in securing PoS networks, they should be paid fairly for their work and encouraged to contribute more. However, validators should not be purely financially (and certainly not politically) motivated, taking up competing positions and launching political campaigns to convince token holders to stake with them. This type of lobbying affects all aspects of the chain’s development—from governance to technical upgrades—and can lead to factionalism and misalignment. 

PoC makes the system fairer and more sustainable, ensuring participants are aligned and take actions that benefit the Gno.land community and the broader ecosystem. That’s why (unlike PoS) contributors receive rewards based on their contribution effort (tier level) rather than how many tokens they stake. They are thus incentivized and recognized for the quality of their work, ideas, and alignment, driving participation and active engagement. Governance is allocated to the people most likely to care for the ecosystem’s long-term success—the contributors who have spent the most time working toward it—from open-source developers to video creators and everyone in between.

### Rethinking Financial Incentives 

For long-term security and sustainability, PoC emphasizes project principles and values over monetary gains, replacing standard token incentives with a system that separates voting power from token ownership. Two reward systems are currently being considered (in addition to a hybrid system). For the first, contributors receive WORX units that weigh the amount of GNOT tokens (the native Gno.land gas token) earned each month. Each member of the same tier receives the same amount of WORX. At the end of the month, the total each member earned is divided by the total amount of WORX distributed that month to calculate a percentage. This percentage represents the percentage of Gno.land fees earmarked for contributors that each member will earn in GNOT. WORX will likely be cleared each month to prevent cumulative, exponential reward exploits over long periods of time. 

For the second, each tier level simply receives an amount of GNOT each month fixed to a USD value, similar to a salary. This would be combined with risk management and caps per tier level in order to promote long-term sustainability based on Gno.land fee generation. A hybrid of this system is also possible, either rewarding contributors of lower tiers one way and higher tiers the other or using both systems in tandem based on predefined conditions. This will be explored further in future tokenomics articles, models, and documentation.

Regardless, WORX units are not transferable, will not be listed on exchanges, and hold no monetary value. WORX units are more like shares that represent value provided by contributors and allow their work to be quantified compared to other contributors/tier levels. It’s important to stress that GNOT tokens do not influence governance on the platform in any way. Voting power is earned through contributions and distributed according to contribution effort, with each member of the same tier representing equal voting power that increases with their tier level. This creates a network of highly aligned contributors who care deeply about the platform they are building and strive to improve it.

GNOT, the native Gno.land gas token and the gas token of the Gno.land ecosystem, will be distributed via airdrop to qualifying ATOM stakers. It will also be available for purchase after that point (*more on Gno.land’s airdrop and tokenomics coming soon*). GNOT is used to pay all fees associated with the network and beyond, including transfers, IBC, ICS, and contract interactions, giving holders the chance to earn rewards from the economic activities of Gno.land.

### What Makes a Good Contribution?

WORX and/or GNOT can be earned through different types of contributions—not only coding and development expertise—but also through non-technical contributions, such as community building, governance involvement, constitutional proposals, teamwork, media creation, etc. The core focus is on alignment, not necessarily specific tasks. For example, an accepted proposal or merged code will raise or at least maintain the contributor’s tier level, allowing them to receive rewards during their time working between submissions. However, a proposal or code that has displayed a very high level of effort, detail, and aligned values (but is not merged) will also be considered in any proposals regarding contributor promotion.

This system allows the ecosystem to show appreciation for diverse forms of contributions and ‘useful failures’ that bring us closer to the solutions we adopt. It is designed to foster engagement, creativity, and collaboration while encouraging anyone aligned to contribute to growing the Gno.land chain and community.  

### How Are Contributions Assessed?

There is a strong human element to deciding what makes a good contribution, requiring knowledgeable human judges to exercise discretion. As such, contributions won’t be templated by default or rewarded automatically but assessed through Gno.land’s governing DAO, GovDAO. GovDAO is responsible for development and governance and is organized into tiers, as discussed above.

GovDAO members review, measure, and curate contributions, and the tokenomics of GovDAO incentivizes members to be effective and unbiased evaluators. They engage in discussions and assess contributions based on effort, time, and other relevant factors/metrics that contributors will have stored in their profiles. The decision-making rationale is transparent and visible through on-chain forums. Again, contributors are assigned a tier level and receive a corresponding reward each month according to their tier. As contributors join GovDAO, the DAO grows, giving Gno.land decentralization efficiency and a high Satoshi score. 

GovDAO is assisted by a network of knowledge-specific DAOs, such as an Engineering DAO, a Support DAO, an Operations DAO, and the EvaluationDAO, which comprises a trusted group of high-reputation contributors that help assess specific contributions. This enables secure collaboration and seamless integration (*more on Gno.land’s network of interconnected DAOs coming soon*.) 

### Sybil-Resistant and Secure

In addition to being fairer, more aligned, and sustainable, PoC is Sybil-resistant by design. In blockchains, a Sybil attack is where one or multiple attackers multiply their presence and influence by creating fake identities to sway major network decisions (for example, including malicious blocks). In terms of PoS, the Sybil resistance is purely monetary (people need to stake real money to get power), so an attacker that wants to carry out a Sybil attack on a PoS network needs to lock at least as much stake as that locked by honest validators.

PoC minimizes risks of Sybil attacks, takeovers, and alliances as the community vets every person who is given any power or sway in the network (including validator power) through the DAO, so at no point can anyone "spoof" identities and regain major sway. Moreover, Gno.land is built and secured by the merit and effort put into the project, as opposed to how many tokens someone can buy, rethinking financial incentives and making the platform Sybil-resistant and secure.

Through fairer rewards, restructured incentives, resistance to corruption and Sybil attacks, and a strong appreciation for all contributions, Gno.land is designed to be sustainable and fair. A censorship-resistant platform built, owned, and secured by a growing, aligned community for many generations to come.

*I. What Is Proof of Contribution? is the first in a series of articles to dive deeply into the philosophy, vision, mechanics, and work involved in developing a new consensus mechanism for the next generation of smart contract systems. Look out for subsequent editions and additional Building Gno.land series, and let us know what you think! Got questions? Join the Gno.land [Discord](https://discord.com/invite/S8nKUqwkPn) or follow us on [Twitter/X](https://x.com/_gnoland)*.
