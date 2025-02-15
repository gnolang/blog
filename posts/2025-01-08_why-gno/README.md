---
publication_date: 2025-01-08T00:00:00Z
slug: why-gno
tags: [gnoland, gno]
authors: [thehowl, moul]
---

# Why do we need Gno and gno.land?

gno.land is a decentralized operating system to create decentralized platforms.
It is powered by Gno, a deterministic implementation of the Go programming
language, as a powerful tool to build composable packages and smart contracts on
a blockchain. Although gno.land can support a variety of applications and
implementations, among our core use-cases, we intend gno.land to be used as a
base for human-centric and transparent social platforms. In this blog post, we
outline the problems that gno.land addresses, and what solutions we are building
to solve them. While this summary is non-exhaustive and reflects the evolving
focus of the project, it offers an overview of our key objectives and ongoing
initiatives.

## The Social Problem

> Achieving Unstoppable Truth.

We need a social platform that can be a common ground for a plurality of contrasting voices, while being transparent in its moderation and governance, with built-in features to create schisms when that is not upheld.

- The current centralized social networking platforms are plagued with problems for users, content creators, and businesses:
	- Users are often constrained by the centralized ownership and control of platforms. Due to the network effect, it's very hard to move away from an established platform, or to create a "schism", granting users the possibility of true choice and requiring the platform to align with the interests of the users.
	- Content creators frequently need to adhere to unwritten "community guidelines" to preserve monetization or, worse, to avoid shadowbans that hide their content from their followers. Shadowbans are often attributed to opaque “algorithms”, making it difficult to challenge decisions or prove instances of discrimination, bias, or censorship.
	- Centralized, un-forkable platforms can have agendas, and their agendas are controlled by powerful entities. Users have little choice to "take their business elsewhere", as for many users the platforms themselves become an important medium for their business, or simply a way to connect to the people they care about. As it's been seen time and time again, control can be instrumentalized to track and profile people, sell them products they don't need or misinform them to vote for candidates that don't align with their best interests.
	- The data of the users is a commercial tool which can be used for advertising and improving massive AI models. In turn, the AI revolution promises to make the internet an unlivable hellscape of bots talking to each other who spread political propaganda and regurgitate incorrect information. Platforms must be resilient to bots and spam, fostering authentic communication and genuine content production among real users.
- What gno.land wants to achieve:
	- Forkability. Blockchains are distributed databases, which themselves can be constructed as Merkle trees - allowing to check out different points in time, and to cheaply allow forking of content, similarly to how git allows branching.
		- gno.land, like all blockchains, is forkable, but even within gno.land we can make all states, including those of social platforms, forkable, too.
		- We see forking as a feature that should not only be a possibility for the most expert of users, but a common solution for everyone; it should be allowed even on the contracts themselves.
		- It's not an option of last resort; forking in software is something that happens continuously. The tools we use, such as git, allow us to build forks that feed back work into the upstream. Similarly, constructing common knowledge and information databases can also happen in a model that accepts schisms and different ideas on how to achieve the objective.
	- Decentralization of hosting and governance, but centralization of the communication platform. Thanks to the nature of blockchains, we can create social platforms which are hosted by many different validators, all with the same exact state. Governance is decoupled from validation processes, allowing it to be guided by distinct principles and incentives.
		- With the Fediverse, we avoid sharing single points of failure and centralization of governance into a single entity (person or company).
		- As opposed to the Fediverse, gno.land platforms can be more reliable and be centralized in terms of discovering and interacting with other users.
	- Funding by the users. Prioritize user-aligned platforms over advertiser-centric models.
		- This happens practically by ensuring the transaction fees, or revenue models built on top of them, can adequately fund both the hosting and the governance/moderation.
	- Freedom of speech. Ensure a baseline of moderation to address harmful and illegal content, while enabling smaller communities to govern content management according to their preferences, also using algorithms and moderating systems.
		- The chain should incentivize thoughtful content creation, moderation, and active participation in the community; but not block out content, so that the platform can be truly neutral.
			- There can and should be algorithms that show users interesting content while blocking out what they don't want to see. But users need to be able to choose between different algorithms and understand what each one is targeting; rather than solely signing up for the one that maximises "engagement" or "screen-time".
- Why gno.land can achieve it:
	- By solving the incentive problem, we ensure that the chain can continue to work with a specific vision while being hard to subvert. The governance aims to expand the blockchain's potential and usages in real applications, rather than the capital of its investors.
	- By solving the technical problem, we can actually create a platform with the properties described above.
	- gno.land pays attention to creating scalable data structures that can host the content of the platform and enable free and universal access to it.


## The Incentive Problem

> Establishing Anti-Subversion Governance.

Most blockchain incentive schemes focus primarily on rewarding validators for securing consensus and maintaining network integrity. In order to build larger systems which are not just financially focused, we need to incentivize non-validators who nonetheless contribute to the advancement of the project, and make sure the chain governance is philosophically aligned, rather than just looking out for individual gains.

- Exclusively rewarding a chain's validators is a good way to secure the chain; but not a good basis to create larger ecosystems outside of financial transactions.
	- The Proof of Stake model has been successful in securing blockchains; but it has not been successful in creating incentives for using and advancing a blockchain, only for providing validation services.
	- Additionally, as governance is most often tied to the stake of each user, control of the chain can easily be subverted by powerful and rich actors who want to acquire sizable portions of the blockchain.
	- Similarly, Proof of Work puts the incentives in the hands of those who have the resources and the means to have cost-effective machines, incentivizing those who have the most resources or those who have good, insider access to primary resources (chips and electricity).
- Blockchains allow us to create effective economic models for building the internet which are fair to the participants involved in making it.
	- We should pay creators for making great content.
	- We should pay moderators who keep communities focused and true to their spirit.
	- We should pay those contributing to collective silos of knowledge like wikis, and verifying its information and accuracy.
	- We should pay software developers, as they create the applications and infrastructure.
	- We should pay auditors, who verify the security of smart contracts and packages.
	- Finally, we should pay validators, as they physically host the blockchain and make it possible.
	- The internet can no longer afford to only pay the companies who are making content hosting possible; we need to recognize the hours of work that go into collecting, creating, organizing and distributing information.
- The governance of the chain cannot be tied to how much capital each participant has.
	- It should instead work to make its government aligned on common goals and a general philosophy of advancing the state of the project.
	- Its government has to be directly responsible for hosting the chain, ensuring that all the way down to the validator control is never given up to a non-aligned "third party".
- The chain can and will have economic activity and present itself as a good place to build businesses; but the focus of the chain will go towards its applications rather than merely "decentralizing finance".
- The chain should still welcome "investors"; but not as a main focus. Investors can shine in creating interactions with other ecosystems, and creating ways to exchange tokens, making it easier to buy GNOT and increasing demand. But the true value of the platform is created by those who develop and create on top of it.
- Contribution-driven governance.
	- The core of the chain should consist of its most experienced members who have shaped it into what it is.
	- The governance also picks the validators of the chain, who will be chosen on alignment and technical capability rather than the capital they can stake.
	- Governance mechanisms are structured to reward contributors based on their verified contributions, with an emphasis on maintaining alignment with the network’s foundational principles.

## The Technical Problem

> Creating Timeless Code.

The technical innovation brought by Gno makes solving the above problems possible.

- We're at the Assembly age of web3.
	- Developing software is costly:
		- It's expensive to build basic things, both in terms of expertise in development, but also auditing and finding people that truly understand and can take advantage of the technology.
		- Many ecosystems employ domain-specific languages or restrictions on top of others (ie. WASM), which still requires specific training.
	- The development which is possible is very basic, similar to "interconnected spreadsheets".
	- Monolithic blockchains create limitations which de-facto limit the growth of the chain, only allowing for the entire system a very limited amount of computation in each block.
- What Solidity misses:
	- Solidity is a Domain Specific Language, further increasing the cost of adoption by new users, especially those who are still curious and uncommitted to working on a blockchain.
	- Financially focused. Solidity contains many native constructs pertaining specifically to smart contracts; like the global `ether`, the type `address`, the globals `msg`, `tx` and `block`. On the other hand, its design lacks simple ways to represent and handle floats, dates and times, and misses other general-purpose features programmers expect from most programming languages.
	- Hidden control flow: [fallback functions](https://docs.soliditylang.org/en/latest/contracts.html#fallback-function) create hidden control flow when sending tokens, causing problems like [re-entrancy attacks](https://docs.soliditylang.org/en/latest/security-considerations.html#reentrancy).
- How Gno solves it:
	- Composability:
		- Interfaces and closures allow to plug different code and data structures into each other.
		- Realms can be composed by re-using functionality of other realms or existing functionality published as packages.
		- IBC enables the blockchains using Gno to communicate with one another seamlessly, and build on the work happening elsewhere.
	- Low cost of adoption: Gno is a subset of Go with some different standard libraries. Learning Gno has a lower adoption barrier, as it is based on the widely-used Go programming language, familiar to approximately [two million developers worldwide](https://research.swtch.com/gophercount).
	- Low cost of building: quicker and simpler to write dApps (simple language, good already-audited p/ libraries); quicker and cheaper to audit them (all the source code is to be published on-chain).
	- App-focused: we are interested in social platforms, video game servers, and ways to share knowledge in a decentralized manner. We see dApps as systems that can make the web decentralized again and create a truly open internet for the time to come. DeFi is great - and an obvious use case for blockchains - Gno expands on what DeFi can do and enable many more apps, and ways to communicate with others.
	- Clear control flow: similarly to Go, there are no implicit getters, setters, and no method overloading; clear is better than clever.
	- Open source by default: all code on gno.land is published directly as source code, and published under the terms of the [Gno Network General Public License](https://gno.land/license) (on top of any license that the author may decide to use). The on-chain code can be perpetually inspected, modified and improved; becoming what we call "timeless code".
- What does Gno enable?
	- Re-usable "core" contracts, like user management, groups, permissions.
	- Open Web: all data is public and clonable by others; as well as the tools to organize it and present it.
	- Coordination (civil services without government backing), Wikipedia, platforms that cannot be subverted or sold.
- gno.land can be a GitHub for a wider ecosystem.
	- Our hope is to show with gno.land a much better way to develop decentralized applications.
	- Through IBC+ICS, other blockchains can look to connect with the gno.land chain and re-use its foundational code; as well as interact with its contracts.
	- gno.land, in the philosophy of [Cosmos](https://cosmos.network/), is not a monolithic, one-size-fits-all blockchain. The implementation of the programming language itself is not meant to be definitive, but rather to be a reference for creating other better, faster implementations that suit other use cases.

Learn more about the vision and mission driving gno.land's development by watching our [video](https://youtu.be/M920PO3yrHk).
