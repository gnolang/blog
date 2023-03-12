---
title: Intro to Gnoland - The Smart Contract Platform to Improve Our Understanding of the World
publication_date: 2022-11-21T17:13:00Z
slug: intro
tags: [gnoland,gnosh,gnot,permissionless,consensus,proof-of-contribution,dao,governance,ibc,democracy,freedom]
authors: [christina,jae,manfred]
---

_Welcome to Gno.land. This is the official site to learn about the Gnolang (Gno) programming language and the Gno.land smart contract platform, as well as understand the motivations behind Gno and our core values and mission. We’re starting a series of blog posts and holding regular community calls and AMAs so that you can stay up to date with upcoming developments and dive deeper into the Gno World Order. Stay tuned._

## What Is Gno.land?

Gno.land (pronounced no-land) is a layer 1 smart contract platform invented by Jae Kwon, co-founder of Cosmos and Tendermint, to address multiple issues in the blockchain space — in particular, the ease of use and intuitiveness of smart contract programming platforms. Beyond offering succinctness, composability, expressivity, and completeness not found in any other smart contract platform, we aim to challenge the regime of information censorship that we find ourselves living in today.

By using the programming language Gnolang (Gno), an interpreted version of the widely-used Golang (Go) language, using a state-of-the-art VM written in Go, we want to lower the barrier to entry to web3 and make it simple for developers (particularly existing web2 developers) to write smart contracts and other blockchain applications without having to learn a programming language that is limited by design or exclusive to a single blockchain ecosystem.

### Gnolang (Gno) Is Essential to Broader Adoption of Web3

For web3 to grow in a sustainable way, we need technological solutions that are designed for the blockchain with programming languages that are universally adopted, secure, composable, and complete. The main programming language currently used for creating smart contracts, Solidity, is designed for one purpose only (writing smart contracts) and lacks the completeness of a general-purpose language.

Solidity removes many of the complexities that blockchain programming requires (such as memory management, ensuring that the code is deterministic, and understanding how the entire tech stack is implemented) allowing developers to quickly build succinct smart contracts. However, Solidity is only used for smart contracts on EVM-compatible blockchains (like Ethereum, Polygon, or EVMOS) and its design is limited by the limitations of the EVM. In addition, developers have to learn several languages if they want to understand the whole stack or work across different ecosystems.

Go, on the other hand, is a well-designed complete programming language with its foundation based on composable structures, designed by the creators of Plan 9. This allows developers to rapidly accelerate application development and adopt a modular structure by reusing and reassembling existing modules without building from scratch. They can embed one structure inside another in an intuitive way while preserving localism, and the language specification is simple, successfully balancing practicality and minimalism.

Go is widely used, especially among existing web2 developers. It’s easier to learn and can be used to program almost anything, such as GoEthereum or Tendermint. Every part of the Gno.land stack is written in Go so that one person can understand the entire system just by studying a relatively small code base. The Go language is so well designed that the Gnolang smart contract system will become the new gold standard for smart contract development and other blockchain (and even non-blockchain) applications.

### Security Is a Built-in Feature of Go (Golang)

Beyond object embedding, closures, importing of modules, composability of programs, and interfaces that allow you to implement a specific set of functions, Go supports secure programming through exported/non-exported fields, enabling “least-authority” design. It is easy to create objects and APIs that expose only what should be accessible to callers while hiding what should not be simply by the capitalization of letters, thus allowing a succinct representation of secure logic that can be called by multiple users.

Another major advantage of Go is that the language comes with an ecosystem of great tooling, like the compiler and third-party tools that statically analyze code. Gno inherits these advantages from Go directly to create a smart contract programming language that is safe and helps developers to write secure code relying on the compiler, parser, and interpreter to give warning alerts for common mistakes.

### How Gnolang (Gno) Differs from Golang (Go)

[![Go and Gno](https://gnolang.github.io/blog-assets/intro/thumbs/go-and-gno.png)](https://gnolang.github.io/blog-assets/intro/go-and-gno.png)

_Image 1: Gnolang - Like Go but specific to the blockchain_

Gno is around 99% identical to Go and most people can code in Gno from day one, even minute one. The Gno.land programming environment comes with blockchain-specific standard libraries, but any code that doesn’t use the blockchain-specific logic can run in Go with minimal processing. On the other hand, some libraries that don’t make sense in the blockchain context are not available in the Gno.land programming environment, such as network or operating-system access.

Otherwise, Gno loads and uses many standard libraries that power Go, so most of the parsing of the source code is the same. Under the hood, the Gno code is parsed into an abstract syntax tree (AST) and the AST itself is used in the interpreter, rather than byte code as in many virtual machines such as Java, Python, or WASM. This makes even the Gno VM accessible to any Go programmer. The novel design of the Gno VM interpreter allows  Gno to freeze and resume the program by persisting and loading the entire memory state. This allows (smart contract) programs to be succinct, as the programmer doesn’t have to serialize and deserialize objects to persist them into a database (unlike programming applications with the Cosmos SDK).

The composable nature of Go/Gno allows for type-checked interactions between contracts, making Gno.land safer and more powerful, as well as operationally cheaper and faster. Smart contracts on Gno.land will be light, simple, more focused, and easily interoperable — a network of interconnected contracts rather than siloed monoliths that limit interactions with other contracts.

[![Gnolang code example](https://gnolang.github.io/blog-assets/intro/thumbs/code-example.jpg)](https://gnolang.github.io/blog-assets/intro/code-example.jpg)

_Image 2: Code snippet from the Gno programming language_

Today, Gno.land is the only blockchain instance in the world that supports Gno but tomorrow, there will be many chains with different names such as mydapp.zone, or mydao.xyz. Gno.land is the name of ONE chain and is not a name that will be used by other Gnolang-powered chains. Gno.land will remain a minimal hub with three main utilities:

* Managing cross-Gnolang-chain fees/licenses
* To be the (or an) official home for the best smart contracts
* To provide new models of governance (w/ DAO modules)

### Earning Rewards Through Proof-of-Contribution (PoC)

There are four main ways to earn rewards through PoC on the Gno.land chain:

* Pre-defined tasks (technical or otherwise)
* Pre-defined bounties
* Retroactive bounties
* Vesting-style rewards for core members

Bounties rewards (both pre-defined and retroactive) will be decided with “local rules,” through the agreement of the DAO with everything on-chain and transparent. If one human were to abuse the system, it would trigger and the bad actor would be slashed. We’ll go into depth on how you can earn rewards in an upcoming post.

### Durable Solutions to Improve Our Understanding of the World

One of our inspirations for the Gno.land project is the gospels, which built a system of moral code that lasted for thousands of years. Part of Gno.land’s endurance will be having a minimal production implementation that becomes a reference for other implementations and a basis for education to elevate people's understanding of blockchains.

Gno.land aims to appeal to web developers, dApp developers, and blockchain builders to create solutions that help people improve their understanding of the world. With the barrage of misinformation delivered today from various factions, it’s impossible to separate the real from the fake. This causes a state of gridlock. We are living in a regime of information censorship spanning all important topics from climate change to global pandemics — a vast coordinated effort to prevent people from understanding the truth.

By just browsing Reddit, searching with Google, and scrolling through Facebook, Twitter, or Instagram, people are deliberately being [misled](https://twitter.com/lhfang/status/1587095890983936000) about key global issues that we all deserve clarity on. This is as malevolent as any type of censorship regime in the world — and we need to come together to challenge it and break the wall of censorship to achieve a functional democracy at last.

### Gno.land’s Current Phase of Development

Gno.land is currently running in its third testnet and there will be several more testnets before the platform is production ready. Modern civilization wasn’t built in a day, and neither will Gno.land rush into committing to an exact launch date. However, the next development, an incentivized testnet called ‘Game of Realms’, is scheduled for Q1 2023.

Game of Realms will be similar to ‘Game of Stakes’ on the Cosmos Hub and will reward the earliest and best contributors. If you would like to find out more about Game of Realms, Gno.land, Gnolang, or anything else,  join us for our first community call with Gno.land Founder, Jae Kwon on November 22nd, at 4pm UTC on our [Discord channel](https://discord.gg/YFtMjWwUN7). We look forward to seeing you.
