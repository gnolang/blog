---
title: "Gno.land Community Technical AMA #1 - Recap"
publication_date: 2022-12-05T16:15:00Z
slug: tech-ama1
tags: [gnoland,gnosh,gnot,permissionless,consensus,proof-of-contribution,dao,governance,ibc,democracy,freedom]
authors: [manfred, jae]
---

Your questions, observations, and feedback are vital to our core development team. Not only do they give us an understanding of the types of applications and features the community would like to see but they help us formulate better ideas for developing Gno.land as we go. Before we dive into our second **Discord AMA on November 22nd @4pm UTC**, check out the community questions from our first technical AMA below answered by core Gno.land devs Jae Kwon and Manfred Touron.

### Why did you choose Golang over Rust?

**Jae**: “With parallelism offered by ICS1 [Interchain Security 1], the bottleneck becomes speed of innovation with safe code, rather than bare metal performance. So here, garbage collection, concurrency, embeddable structures, and clear spec are good primitives for the next-generation smart contract language.

Rust (or components of Rust** may be used to implement faster clients for gno.land in the future, but in terms of mindshare, I don't think Rust can flip Go due to its design choices. That's not to say that Rust is any worse than Go; they are different.”

### Will Gno be its own hub? Will Gno provide ICS-like security to its own community?

**Jae**: “Gno.land can be a "hub," like "git hub" is a "hub," but that doesn't mean it will offer ICS. If other chains solve ICS1 better, it makes sense for gno.land to be IBC-connected to zones that are not ICS1 replicated/secured with gno.land validators.

If we consider that validators of gno.land are better as contributors to the gno.land ecosystem (rather than general validator service providers** we may be more comfortable contributing to an awesome ecosystem but not entering the validator-as-a-service business.

It makes more sense to me that Cosmos Hub validators should own that business, which will eventually require validators to run their own server stacks and have data center infrastructure.”

### How can one become a validator?

**Jae**: “First, one has to become a member. We have not yet defined the full member system, but we will figure that out along the way. For now, we can say that we want first and foremost members who also validate, rather than impartial validators that only validate.”

### How does Gno validate work? PoS? Proof of Contribution?

**Manfred**: “The contributors DAO will elect validators and validators will have the same amount of power. They'll be focused on validating and will receive rewards for that job.”

### What is Proof of Contribution? What kind of contribution will be credited?

**Manfred**: “Proof-of-Contribution is a way to replace Proof-of-Stake with a metric based on the contributions. It's a variation of Proof-of-Authority where the authority is a DAO of contributors. After the 'Game of Realms** competition, we'll reward the best contributors with a tiered membership in the first version of Proof-of-Contributions DAO. The voting power and everything related to staking will be distributed across the contributors.

Later, we'll add more flexibility to the membership with $GNOSH, allowing more accurate and fair rewards. Validators won't receive voting power with staking. The DAO will elect them, and they will all receive the same amount of power. Validators will receive rewards for their technical work, not for the amount of staked tokens they are bound to.”

### Is there a document or resource that describes the key concepts in a Gno smart contract?

**Manfred**: “We have yet to get a single top-level documentation, sorry. You can find documentation in the code, README files, issues, etc. We need to improve this. The community will be able to work on this during Game of Realms.”

### Is there a big-picture diagram of the ecosystem?

**Jae**: cosmos hub <-- "ec2+DTCC"
gno.land <-- "github for gno"
(cosmos hub etc) ICS zones <-- "holy grail" scalable smart contracts
your chain <-- "gno inside"
your app <-- "import gno.land/..."
blockchain-based communications/coordination/discourse platform <-- us
// DTCC: "https://www.investopedia.com/terms/d/dtcc.asp" // my point is, be a good reliable token hub with good governance.”

### I'm a developer (PHP, Python**. How can I become a Gno developer? Please advise me on where to start.

**Manfred**: “Start learning Go! One of the long-term goals of Gno is to make writing contracts as easy as writing web2 apps. The language is already strong in that direction, but we still need to catch tooling, documentation, tutorials, and language improvements. You need to have a good level with Golang and be autonomous to start building on Gno.

One of the Game of Realms tracks will be to work on everything related to onboarding more people. This will be the best place to write specific tutorials to onboard people from other ecosystems or languages.”

What are Realms, and what is r board?

**Jae**: “A realm is a Gno package with state, that represents a smart contract with storage and coins. The other Gno packages don't have state, and so are "pure" packages that can be imported from other realm or non-realm packages. Like land-tax, realms must be whitelisted or pay storage upkeep for their state. You can create new realms by uploading a new package with the package directory starting with /r/REALM/NAME.

/r/demo/boards is a Gno package that renders a message board. It is a proof of concept message board written in Gno. Since we need to preserve messages, it is a stateful (realm** package. You can see the files of the demo boards, like:

https://test3.gno.land/r/demo/boards/board.gno

### How do external packages get imported?

**Manfred**: “Example: when you call your smart contract from Go during testing, how can/should that smart contract load external packages?

A gnolang can only import other gnolang contracts/libraries that were published on-chain. If you want to import an external Golang library, you need to port it to Gno, and publish it as a library, then you can import it from a top-level contract.

gnodev test is an exception, it basically creates an in-memory Gnolang VM, publishes the dependencies (automatically detected**, and executes the test. The tool can act differently from the real on-chain experience. Note that we'll improve the gnodev so it can automatically download on-chain contracts or use custom local paths, to support advanced development workflows.”

### What is a Gnode?

**Jae**: “I don't like the name "Gnode" because it's too generic, but the idea is to build Gno-based building blocks for GnoDAOs, as MyGnode embeds components (of owners, treasury, board, etc.** here:

https://github.com/gnolang/gno/commit/b9128b1d69f02dbb49be883e0c70fe9d3fc40dcc

**Manfred**: “We can change the name 🙂. A Gnode is a DAO implementation that implements an interface allowing them to interact. A Gnode can have a parent and have children. Top-down interactions may be funding, grants, and approvals. Bottom-up interactions may be reporting or voting. The implementation is flexible. You can have DAOs managing a Gnode, its treasury, and voting the cross-Gnode interactions. You can have Gnodes with an elected leader or one driven by a bot or another blockchain. One of the goals of Game of Realms will be to propose various implementations of Gnodes.

At the level of Gnoland, we will probably have a top-level Gnoland Gnode managing a global treasury and vision. Then various technical and non-technical child Gnodes manage subsets of the treasury and their tasks. They may also have children. With IBC2, Gnodes could be distributed across different chains.”

### What is the timeline for IBC2?

**Jae**: “After the launch of gno.land, IBC2 is permissionless innovation anyone can try for, so I imagine not long after that. After initial implementations, I bet we will want to tweak/optimize the Merkle tree further, but this can come after IBC2 demos.”

### Can you tell us more about Game of Realms?

**Manfred**: “Game of Realms is a competition to build the first contracts of Gnoland and experiment with proof of contributions. The first step of the competition will be to build the missing tools for the second step. So people will compete to write the DAO that will review the other contributions and allocate points.

The rest of the competition will be about competing to write the best contracts for well-known categories or make non-technical contributions. At the end, we'll have strong foundations (libraries, rules, tutorials, dApps** to help upcoming builders to start in better conditions. The best contributors will earn rewards and membership in the future DAO of contributors that will co-own the chain.

We'll have the first version of a Proof-of-Contributions-based DAO of contributors. Focus on one of the official tracks: build a contract suite to compete with Cosmos' governance module to eventually complete Cosmos Hub governance. Realm boards are basic discussion contracts that can be used for discussions, and be extended for governance, launchpad, or other things mixing discussions and DAO actions.”

### Is it possible to build code with gno.land directly online?

**Jae**: “We will make the sandbox staging.gno.land environment easy to access, and that will be preferable to testing on gno.land directly. The gno codebase tries to remain minimal so it shouldn't be difficult to run it locally.”

**Manfred**: “I've seen people writing contracts from VSCode on an online VSCode instance. Someone could create a VSCode template configured to communicate with staging by default with a dummy wallet containing tokens.”

### Is there a plan to be able to use the Gno VM with a Cosmos SDK-based chain?

**Manfred**: “This is one of the plans, yes. And not only on Cosmos SDK. But we don't have a clear plan about how it will happen yet.”

### How about interoperability?

**Jae**: “Regarding interoperability, will it be between Gno chains, with Cosmos, or with more chains outside of Cosmos? If it is with chains outside of Cosmos, which ones, in the short and long term? I think if the latter were to come to pass, the world of web3 and NFT could be awesome. Short run, Cosmos SDK-based chains with IBC1 for code import and cross-chain smart contract calls; but with IBC2/Gno it's really up to the smart contract logic.”

### Are Gno.land tokenomics deflationary?

**Jae**: “There will be $GNOT, and this token will be used for spam prevention fee payment, and it will be deflationary. Previously, we discussed $GNOSH as a secondary token, but we have moved away from the $GNOT/$GNOSH model and will keep $GNOT while making gno.land more about membership among levels of peers.

I think we need an alternative to the Cosmos Hub that is more people-centric than stake-centric, and where alignment is not bought or sold but depends on contributions and value alignment proven over time. The hope is that by moving away from a pure tokenomics perspective and moving into the realm of politics and ethics along with general economics we can curate a different kind of culture.”

### Are there any collaborations with other projects to build on Gno?

**Jae**: “Yes, why don't we make this truly open, in the style of free software, so that we can build upon a common VM design? The only thing I want to retain control over for a temporary duration of time is the regulation of trademarks, like "gno", "gno**", "*gno" (but you can use the license to fork this project however you want); and we want proper attribution, but the AGPL fork license suggests how we can work together collaboratively.

The GNO VM can be used on any chain if it follows the AGPL style license, which we are calling the "Gno GPL". Blockchains can still be composed of components licensed with compatible open source software. We can collaborate indirectly by working and contributing to the same codebase, and know that the code we are building together will always be available for you to use for your chain, as long as it remains and is offered as GNO free software.

So anyone can build GNO smart contracts into their chain for free, according to the license we are deriving from the GNU (not GNO** AGPL license. You don't have to pay gno.land or anyone if the license is followed. Example: we will collaborate with the Cosmos Hub and Cosmos/ATOM community to offer gno DAOs to be hosted by ICS1, and help bring collaboration tools for Cosmos. So this is how gno works with Cosmos Hub assuming ICS1 is solved. As for gno.land, we can start off with an independent gno instance for the Cosmos Hub's gno shards, and later allow the IBC importing of vetted code from gno.land/*.”

### Apart from Adena, are there any plans for another wallet?

Jae: “I think what we need are a few competing base implementations that best leverage the framework they build upon, rather react or minimal vue; and to create common core libraries along the way if reasonable. But there ought to be more than one approach for such a key component, with special care taken into consideration for security. Like, I don't agree with Keplr asking so easily for a 12/24-word mnemonic, even if the implementation is secure, it is going to become a problem. PSA btw.”

### Wen mainnet?

**Jae**: “Some time by Q2 next year would be good. But as policy, we can't commit to a date, because everything has to be ready first before the official launch. Our thesis is that having the DAO with sub-DAOs will allow us to reach the end result in a faster way via some form of parallelism. First, we need DAOs to assess new code, and better UX for managing something like upgrades to the Cosmos Hub. Once we have the DAO running on testx.gno.land, for some x > 4, and we have checked all vital TODOs, we will know that we are ready for "mainnet."

_Do you have more questions for Manfred or Jae? Would you like to know more about Gno.land, Gnolang, Game of Realms, or ways to contribute to our growing ecosystem? Drop us a question on Discord and be sure to join us for our second **AMA on December 6th @4pm UTC.**_
