---
title: "The More You Gno: Gno.land Monthly Updates - 2"
publication_date: 2023-05-26T13:37:00Z
slug: monthly-dev-2
tags: [gnoland, gnovm, tm2]
authors: [christina]
---

**The More You Gno 2: Gno.land Developer Updates**

Over the past few weeks, our core devs and ecosystem contributors have been making massive strides on Gno.land. There’s a lot to cover in the second edition of *The More You Gno*, from updates on Tendermint2 and GnoVM to stack/frames management, Gno IDE, and plenty more. We’ll also see what some of the external teams contributing to the platform have been up to, including Gno.land’s first decentralized exchange, GnoSwap, and Adena compatibility with GRC20 tokens. Check it out.

**Tendermint2**

We’re making steady development progress on Tendermint2, which focuses on simplicity of design, minimal code, minimal dependencies, modular dependencies, and completeness. For the time being, Tendermint2 will stay in the main repo in a top-level folder named Tendermint2. This is the official location to develop and improve the consensus protocol until it is stable enough to be extracted from the Gno repo and become a standalone project. Currently, Tendermint2 depends on GnoVM, however, we are working to unlink this dependency and build a basic demo Tendermint2 chain and Client.

Tendermint2 JS/TS Client is a JavaScript/TypeScript client implementation for Tendermint2-based chains. The client will make it easier for developers to interact with Tendermint2 chains, with a simplified API for account and transaction management, removing a ton of manual work and allowing developers to focus on building their dApps. You can [read more about the client here](https://www.npmjs.com/package/@gnolang/tm2-js-client). In addition to the Tendermint2 JS/TS client, we also created a Gno JS/TS client that just extends the TM2 one to provide Gno-specific functionality. You can read more about this here.

**Game of Realms**

The incentivized competition to find the best contributors to Gno.land continues in phase one, with slow but steady progress being made. Nir1218 initiated an Evaluation DAO Kickoff discussion in [issue 792](https://github.com/gnolang/gno/pull/792) to initiate testing code for the key smart contract infrastructure that will power the Gno.land platform. We are also interviewing architects for the core team with experience in governance modules and creating new economies on-chain, and a new DevRel team member will be joining us soon to create awesome tutorials and documentation to advance Game of Realms further. Gno.land must be built by the community and we will not rush to push Game of Realms to the second phase until we have found quality contributors to complete the challenge tasks and become the platform’s first founding members.

**Gno IDE**

Our core development team is working on a web-based IDE for Gno.land that will greatly improve the developer experience, allowing builders to quickly spin up Gno realms and packages right on their browsers just by visiting a web app. Currently named Gno IDE but with a rebranding on the horizon, this intuitive product focuses on ease of use and improved UX, and will include all the features you’d expect from an IDE, such as auto compilation in the editor, debugging, extensive testing capability, and powerful APIs like IntelliJ to supercharge your programming.

Gno IDE currently has multiple modes to support different use cases, including a normal mode for everyday programming, similar to a standard code editor, a presentation mode for video calls or screen sharing, and an embedded mode to extend functionality, allowing you to embed the IDE directly into websites and blogs. You can also choose to edit your code in Emacs or Vim and easily switch between steps, from previous to next, making creating your tutorials and blog posts more intuitive. Watch out for more to come on Gno IDE soon, and if you want to contribute by creating a plugin for your favorite editor, open a PR to win contribution points.

**Stack/Frames Management**

The stack/frames is an integral part of the virtual machine (VM) and the language. Stack/frames provide context for smart contract developers, enabling them to access useful information, such as the original caller, or to determine if a contract is being called through another one. The current implementation is limited in scope and relies on fixed positions in the stack which can lead to inconsistencies.

There is an ongoing [issue 683 open here](https://github.com/gnolang/gno/issues/683) and we have continued to work on enhancing stack/frames development over the last month. We’re adding a new function in the standard library std.PrevRealm (previously GetRealmCaller). Currently, we only have GetOrigCaller, which returns the user calling the first realm. This is not secure and we need a way to call the previous caller. This will allow a realm to handle GRC20 treasuries. See [issue 667](https://github.com/gnolang/gno/pull/667) and [issue 634](https://github.com/gnolang/gno/issues/634) for further details.

**Dealing with Panics in Native Functions**

We have devised a solution for dealing with panics in native functions, [see pull request 732](https://github.com/gnolang/gno/pull/732). Previously, when there was a panic in a native function, we could not recover it in Gno code. An example of this was the assert origin call, which panicked if the call was not a direct call from a transaction. Based on discussions with contributors, we’ve agreed that native functions should never panic, but if they panic, they panic with machined Gno panic. This gives us the choice in a native function to code a Gno panic, or, if it's a very bad panic, use Go panic so that we know the Gno code is unable to recover it.

**Logic Upgrading**

Making it possible to upgrade your logic is definitely out of scope for the first version of Gno.land, however, it’s an important issue that we have begun to discuss so that we can place certain restrictions on it, such as allowing upgrades when we consider them safe enough to be compatible with imports. Another idea is to work on creating workflows where migrations become something official. This way, we could define ways to migrate a contract completely in a single transaction at the chain level. Once everything is working and approved as the previous contract is parsed or archived, the new one gets the data. We will revisit this topic after the first version of Gno.land reaches the mainnet.

**Garbage Collection**

In terms of garbage collection, we don’t have memory leaks as such but we do have defacto memory leaks. By the VM having references to all objects, they won’t be released by Go’s underlying GC. We have some form of reference counting but it is only done at the end of a transaction. We have implemented a mark-and-sweep garbage collector and are working on the VM runtime to manage the objects and signal to the garbage collector to release them when they are no longer needed. This is done by adding the notion of a heap, which is managed by the garbage collector.

**GnoVM**

Developing GnoVM is an ongoing task and we will likely need to fork the GnoVM to create different competing versions. GnoVM will be complete, limited in features, and serve as the only interpreter, an enduring reference point over time. Future versions of GnoVM will be designed to incorporate CosmWasm so that all Cosmos chains can have CosmWasm enabled and the VM can run directly on the browser and execute tasks on the browser without requiring to make an API call, making it faster. To do this, we can make a Gno compiler in WebAssembly without changing the code because Go supports WASM cross-compilation.

We plan on making a competing version of the original minimalist GnoVM, such as a Rust version with a JIT compiler using LLVM as a backend.

**Ecosystem Updates**

Since our last update, the Gno.land community continues to expand with awesome teams and contributors building cool infrastructure and projects on the platform. Below, we take a look at the largest developments of the past few weeks and extend a special thanks to everyone helping us build Gno.land.

**Teritori**

Teritori blockchain and multi-chain hub launched in November 2022, allowing IBC and non-IBC communities to connect, create groups, exchange tokens and NFTs, and launch new projects. Teritori’s idea for building on Gno.land is to create a multi-chain experience for users with a web portal, NFT marketplace, and social feed that will grow the community, and gradually integrate smart contracts and realms. This will promote Gno.land to more developers and showcase all the dApps being built through an easy-to-navigate dApp store. In the coming weeks, Teritori will work with the Onbloc team to integrate the Athena wallet into their portal as well as discuss ideas for promoting Game of Realms to new developers.

**Onbloc**

Onbloc is one of the Gno.land ecosystem’s most active contributors, responsible for building the Adena wallet and the block explorer Gnoscan. The team has also been working on creating an official Gno SDK that will allow developers to interact with Gno.land more easily, and remove some of the current friction. Onbloc opened [issue 701](https://github.com/gnolang/gno/issues/701) on GitHub primarily for developers who either have their own web app or are building a JavaScript app and want to work with Gno in some way. Currently, developers need to do a lot of manual work, which Gno SDK will abstract away, improving the workflow and developer experience. If you have any ideas or feedback, please contribute to the aforementioned issue.

In another cool development, Onbloc has rolled out a new feature in Adena and Gnoscan to provide support for GRC20 tokens. To store and send tokens, you can open your Adena wallet, click on "Manage Tokens”, navigate to the Custom Token page, and see which GRC20 tokens are available on Gno Testnet 3, searching by the symbol or path. To research on or discover tokens, head over to the Tokens page on Gnoscan for a full list of GRC20 tokens. You can click on any token on the list for detailed information, such as the total supply, owner, or other available functions built into the token. The Account Details page has also been updated to display all tokens owned by each address. You can help by checking out [issue 764](https://github.com/gnolang/gno/pull/764), which discusses adding bigint to support a wide range of numbers and encoding binary, and [issue 816](https://github.com/gnolang/gno/pull/816), which highlights a small bug the team runs into when coding.

Onbloc has also created a new [token resource page on GitHub](http://github/onbloc/gnotokenresources) for anyone to share or upload resources associated with their Gno.land project. This will serve as a shared knowledge pool about any dApp on the platform. If you wanted to create a decentralized exchange, for example, you would need all the information about the tokens available on Gno.land, such as their images, symbols, descriptions, links to websites, etc. Now you can find this in one handy GitHub repository. If you’re a developer or builder who wants your logo or any other static data posted, be sure to submit a PR.

And speaking of decentralized exchanges, Onbloc is also building Gnoswap, the first DEX to be powered by Gno.land, designed to simplify the concentrated liquidity experience and increase capital efficiency for traders. Its interface is built using TypeScript to be user-friendly, secure, and accessible for streamlining complex mechanisms such as price range configurations and staking as part of its core service. Contribute to its interface [here](https://github.com/gnoswap-labs/gnoswap-interface).

As for the contract side, Onbloc is actively working on its development with help from the core members of Gno.land. The code will be open-sourced for full transparency once the basic functions are ready.

**New Core Contributors**

We’re excited to welcome two new core team members, Antonio and Zack. Antonio joined us in April in the core team, bringing with him vast experience in IPFS, and writing Git servers in Go. Zack is our first “tinkerer in residence” and will try to bootstrap the ecosystem of small contracts and small libraries. He will also be writing apps and helping us design a system to better share and showcase our work with a super UX for team builders and open-source addicts.

Antonio is already hard at work researching a benchmarking dashboard that will show performance improvements or regressions when we change the code. He’s assessing whether to use GiHub to track actions or run our own machine to execute GitHub actions. Take a peek at his research so far on [issue 783 here](https://github.com/gnolang/gno/pull/783).

Zack is working on a microblog project. As an experienced web2 Go programmer, Zack is transitioning to web3. Since he’s interested in incentivized social networks, the microblog project will be his first realm, as a Twitter-style blog without titles, where each user has their own page based on their address. Check out [issue 391](https://github.com/gnolang/gno/pull/391) for more details.

**Developer Events**

Over the past few weeks, our core devs have been mainly focused on building but they’re preparing to speak at some exciting events in the coming months. Catch up with Manfred at BUIDL Asia, in Seoul, South Korea, from June 5 - 9. We’re co-hosting a side event with Onbloc, Code States, and Cosmostation on June 5, so be sure to register if you’re in town! We’ll also be at EthBelgrade in Serbia from June 2 - 4, and GopherCon in Berlin from June 26 - 29, so stop by and say hello.

*Do you want to contribute to Gno.land’s monthly updates? If you’re building on Gno.land and want to highlight your development, project, event, or idea, let us know and we’ll include your contribution.*
