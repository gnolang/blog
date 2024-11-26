
---
publication_date: 2024-12-10T00:00:00Z
slug: why-gno
tags: [gnoland]
authors: [thehowl, moul]
---

# Why Gno?

gno.land is a decentralized operating system that enables human-centric social platforms, powered by Gno, a composable programming language to build them. In this blog post, we outline the problems that gno.land addresses, and what solutions we are building to solve them. This is non-exhaustive and, in many ways, evolving; but it can give you a good idea of what the team is focusing on today, and why.

## The Social Problem

> Achieving Unstoppable Truth.

We need a social platform that can be a common ground for a plurality of contrasting voices, while being transparent in its moderation and governance, with built-in features to create schisms when that is not upheld.

- The current centralised social networking platforms are plagued with problems for users, content creators and businesses.
    - Users are hostages to the owners of the platform. Due to the network effect, it's very hard to move away from an established platform, or to create a "schism", granting users the possibility of true choice and requiring the platform to align with the interests of the users.
    - Content creators also often have to follow the often unwritten "community guidelines", in order to maintain monetization or worse, avoid shadowbans which obscure their content to all of their followers, which are just easily hidden under the guise of "the algorithm decided". Because the behaviour of the algorithms is so opaque, it is hard to establish discrimination, bias and censorship when the above happens.
    - Centralized, unforkable platforms can have agendas, and their agendas be controlled by powerful entities. Users have little choice to "take their business elsewhere", as for many users the platforms themselves become an important medium for their business, or simply a way to connect to the people they care about. As it's been seen time and time again, control can be instrumentalized to track and profile people, sell them products they don't need or misinform them to vote for candidates that don't align with their best interests.
    - The data of the users is a commercial tool which can be used for advertising and improving massive AI models. In turn, the AI revolution promises to make the internet an unlivable hellscape of bots talking to each other who spread political propaganda and regurgitate incorrect information. We need platforms that are resilient to bots and spam and where real people can communicate and promote genuine content.
- What gno.land wants to achieve.
	- Forkability. Blockchains are distributed databases, which themselves can be constructed as merkle trees - allowing to check out different points in time, and to cheaply allow forking of content, similarly to how git allows branching.
	    - gno.land, like all blockchains, is forkable, but even within gno.land we can make all states, including those of social platforms, forkable, too.
	    - We see forking as a feature that should not only be a possibility for the most expert of users, but a common solution for everyone, ie. it should be allowed even on the contracts themselves.
	    - It's not an option of last resort; forking in software is something that happens continuously and the tools we use (like git) allow us to build forks that feed back work into the upstream. Similarly, constructing common knowledge and information databases can also happen in a model that accepts schisms and different ideas on how to achieve the objective.
	- Decentralization of the hosting and the governance, but centralization of the communication platform. Thanks to the nature of blockchains, we can create social platforms which are hosted by many different validators, all with the same exact state. The governance can be apart from the validators, and be guided by different principles and incentives.
	    - With the Fediverse, we share avoiding single points of failure and centralization of governance into a single entity (person or company).
	    - As opposed to the Fediverse, gno.land platforms can be more reliable and be centralized in terms of discovering and interacting with other users.
	- Funding by the users. Align the platforms we use with the users' interests, rather than the advertisers'.
	    - This happens practically by ensuring the transaction fees, or revenue models build on top of them, can adequately fund both the hosting and the governance/moderation.
	- Freedom of speech. Guarantee a baseline of moderation to remove outright harmful and illegal content, and hand off moderation to the choices of smaller communities, algorithms and moderating systems built on top.
	  - The chain should incentivize thoughtful content creation, moderation, and active participation in the community; but not block out content, so that the platform can be truly "neutral".
      - There can and should be algorithms that show users interesting content while blocking out what they don't want to see. But users need to be able to choose between different algorithms and understand what each one is targeting; rather than solely signing up for the one that maximises "engagement" or "screen-time".
- Why gno.land can achieve it:
  - By solving the incentive problem, we ensure that the chain can continue to work with a specific vision while being hard to subvert. The governance aims to expand the blockchain's potential and usages in real applications, rather than the capital of its investors.
  - By solving the technical problem, we can actually create a platform with the properties described above.
  - gno.land pays attention to creating scalable data structures, that can host the content of the platform and enable free and universal access to it.

## The Incentive Problem

> Establishing Anti-Subversion Governance.

Most blockchain incentive schemes focus primarily on rewarding validators for securing consensus and maintaining network integrity.  In order to build larger systems which are not just financially focused, we need to incentivize non-validators who nonetheless contribute to the advancement of the project, and make sure the chain governance is philosophically aligned, rather than just looking out for individual gains.

- Exclusively rewarding a chain's validators is a good way to secure the chain; but not a good basis to create larger ecosystems outside of financial transactions.
  - The Proof of Stake model has been successful in securing blockchains; but it has not been successful in creating incentives for using and advancing a blockchain, only for providing validation services.
  - Additionally, as governance is most often tied to the stake of each user, control of the chain can easily be subverted by powerful and rich actors who want to acquire sizable portions of the blockchain.
  - Similarly, Proof of Work puts the incentives in the hands of those who have the resources and the means to have cost-effective machines, incentivizing those who have the most resources or those who have good, insider access to primary resources (chips and electricity).
- Blockchains allow us to create effective economic models for building the internet which are fair to the participants involved in making it.
  - We should pay creators making great content.
  - We should pay moderators who keep communities focused and true to their spirit.
  - We should pay those contributing to collecting and creative colletive silos of knowledge like wikis, and verifying its information and accuracy.
  - We should pay software developers, as they create the applications and infrastructure
  - We should pay auditors, who verify the security of smart contracts and packages.
  - Finally, we should pay validators, as they phisically host the blockchain and make it possible.
  - The internet can no longer afford to only pay the companies who are making content hosting possible; we need to recognize the hours of work that go into collecting, creating, organizing and distributing information.
- The governance of the chain cannot be tied to how much capital each participant has.
  - It should instead work to make its government aligned on common goals and a general philosophy of advancing the state of the project.
  - Its government has to be directly responsible for hosting the chain, ensuring that all the way down to the validators control is never given up to a non-aligned "third party".
- The chain can and will have economic activity and present itself as a good place to build businesses; but the focus of the chain will go towards its applications rather than merely "decentralizing finance".
- The chain should still welcome "investors"; but not as a main focus. Investors can shine in creating interactions with other ecosystems, and creating ways to exchange tokens, making it easier to buy GNOT and increasing demand. But the true value of the platform is created by those who develop and create on top of it.
- Contribution-driven governance.
    - The core of the chain should be composed of the most experienced members of the chain who have helped it shape into what it is.
    - The governance also picks the validators of the chain, who will be chosen on alignment and technical capability rather than the capital they can stake.

## The Technical Problem

> Creating Timeless Code.

The technical innovation brought by Gno makes solving the above problems possible.

- We're at the Assembly age of web3.
    - Developing software is costly:
	    - It's expensive to build basic things, both in terms of expertise in development, but also auditing and finding people that truly understand and can take advantage of the technology.
	    - Many ecosystems employ domain-specific languages or restrictions on top of others (ie. WASM), which still require specific training.
	- The development which is possible is very basic, similar to "interconnected spreadsheets".
	- Monolithic blockchains create limitations which de-facto limit the growth of the chain, only allowing for the entire system a very limited amount of computation in each block.
- What Solidity misses:
	- Solidity is a Domain Specific Language, further increasing the cost of adoption by new users, especially those who are still curious and uncommmitted to working on a blockchain.
	- Financially focused. Solidity contains many native constructs pertaining specifically to smart contracts; like the global `ether`, the type `address`, the globals `msg`, `tx` and `block`. On the other hand, its design lacks simple ways to represent and handle floats, dates and times, and misses other general-purpose features programmers expect from most programming languages.
	- Hidden control flow: [fallback functions](https://docs.soliditylang.org/en/latest/contracts.html#fallback-function) create hidden control flow when sending tokens, causing problems like [re-entrancy attacks](https://docs.soliditylang.org/en/latest/security-considerations.html#reentrancy).
- Why Gno solves it:
	- Composability:
	    - Interfaces and closures allow to plug different code and data structures into each other.
	    - Realms can be composed re-using functionality of other realms or existing functionality published as packages.
	    - IBC enables the blockchains using Gno to communicate with one another seamlessly, and build on the work happening elsewhere.
	- Low cost of adoption: Gno is a subset of Go with some different standard libraries. Learning Gno represents a smaller cost of adoption, as the skill can be widely used outside of gno.land; and it also a language that is used by [around two million developers worldwide](https://research.swtch.com/gophercount).
	- Low cost of building: quicker and simpler to write dapps (simple language, good already-audited p/ libraries); quicker and cheaper to audit dapps (all the source code is to be published on-chain).
	- App-focused: we are interested in social platforms, video game servers, and ways to share knowledge in a decentralised manner. We see dApps as systems that can make the web decentralised again and create a truly open internet for the time to come. De-fi is great - and an obvious usecase for blockchains - Gno expands on what de-fi can do and enable many more apps, and ways to communicate with others.
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
