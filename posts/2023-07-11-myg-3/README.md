---
title: "The More You Gno: Gno.land Monthly Updates - 3"
publication_date: 2023-07-11T13:37:00Z
slug: monthly-dev-3
tags: [gnoland, gnovm, tm2]
authors: [christina]
---

## The More You Gno - Gno.land Monthly Updates 3

We’ve been busy since the last edition of *The More You Gno,* with the Gno.land core team and ecosystem partners present at various global developer events. We’ve visited many gnomes (and gnomes-in-the-making) around the world from Berlin to Belgrade, spreading the word about Gno.land and growing our expanding community. Aside from all the networking, Gno.land is taking shape with a new iteration of our website, the Gno.land Funding and Grants Program, and a host of developer updates as always. Let’s dive in.

## Gno by Example

We recently launched [Gno by Example](https://gno-by-example.com/), our equivalent to both [Solidity by Example](https://solidity-by-example.org/) and [Go by Example](https://gobyexample.com/), where you can see tutorials and code snippets to help you learn and get more easily onboarded to Gno.land. Gno by Example is designed to be community-run with a front-end app and tutorials in markdown. There’s also a specific markdown syntax where you can embed certain file fragments to make your tutorials more structured. We’d love to build this into the ultimate resource center for Gno.land, so feel free to [contribute](https://github.com/gnolang/gno-by-example) with new tutorials and sections. Contributions here are eligible for rewards from the Game of Realms competition.

## GnoVM

We continue developing GnoVM and invite you to provide feedback on what can be improved. This month, there have been a lot of discussions about how to improve native bindings and use the Gno machine in native function calls. Native function calls are well-defined in Go code generation and Go templates but need some modifications for GnoVM. For example, since new native functions already exist in the Gno code, when we try to define a native function, calling the function doesn’t yield the desired result. We’ve created a bunch of panics and tried writing out native functions to see what goes on for them, in an investigation that will go on for the next few weeks. Got any ideas? Please contribute. ([PR 859](https://github.com/gnolang/gno/pull/859)).

## Testnets

Talk about testnets has come up a lot in recent weeks and how to best proceed. Some gnomes are asking for a multi-node testnet to allow for great experimentation, whereas others prefer to keep the testnet single-node. There are advantages and disadvantages to both approaches and we are still listening to feedback and ideas. However, we will likely keep testnet 3 single-node and focus on the language while having a second dedicated multi-node testnet where devs can get creative, think outside of the box, test performance, consensus, and everything they need to push the chain to its limits. We’ve created a new [Hackerspace](https://github.com/gnolang/hackerspace) Repository for the multi-node testnet to prevent spam on the main repo, so please use it to share your scripts, posts, snippets, etc.

## Native Coins and GRC-20 Tokens

We uncovered some significant issues with the banker module ([PR 393](https://github.com/gnolang/gno/pull/393)) regarding minting and burning tokens with the package minter. It was not scoping, filtering, or minting tokens correctly, making it possible to mint and burn unlimited tokens, including GNOT. We want to allow any realm to create its own token and run multiple tokens on their chains, but we need a prefix for security to resolve the issue and allow anyone to create GRC20 smart-contract-based coins but not native coins. We continue to work with small fixes on this issue and will reopen the PR soon.

## Gno.land Funding and Grants Program

Last month we released our Funding and Grants Program to encourage more developers, researchers, educators, and tinkerers to interact with Gno.land. If you’re interested in experimenting with Gnolang (Gno) and building innovative dApps, tooling, products, or infrastructure, check out our GitHub [Funding and Grants](https://github.com/gnolang/ecosystem-fund-grants) page for further information on how you can apply. Start contributing to Gno.land or Game of Realms as this is a prerequisite of the funding and grant application process.

## Developer Relations

The Gno core team is growing! We hired a new DevRel last month and are looking to take on another dev for this open position, so if you’re interested, head over to our [careers page](https://jobs.lever.co/allinbits) and apply! You can expect to see a lot more documentation, FAQs, tutorials, and onboarding materials in the coming weeks and months.

## Ecosystem Updates

Our community of gnomes continues to expand, making tons of activity and progress over the past few weeks. Let’s see what they’ve been up to below.

## Onbloc

Onbloc has been super active this month attending and co-hosting IRL events and networking to find new gnomes about town. Among other updates, Onbloc has completed the first integration of Tendermint2 JS with the Adena wallet and will continue to swap out their existing libraries with TM2JS wherever applicable to ensure that they are as tightly integrated as possible. The team has also open-sourced the Gnoscan block explorer, so if you’re interested in contributing, hop on over to [Gnoscan](https://gnoscan.io/) or the [GitHub repo](https://github.com/onbloc/gnoscan).

## Teritori

Another of our first cohorts from the Grants program, Teritori continues to churn out awesome work and expand its growing team. This month, Teritori has been busy integrating Adena with the Teritori app and working on the DAO contract to build a DAO deployer and various DAO standards and templates for DAO creation. Teritori’s target is to focus on a moderation DAO that can be used for content moderation in social feeds and boards. In the coming weeks, the team plans to integrate the DAO contract into the UI to allow the community to launch a DAO and experiment on the testnet. They have also made an effort to really integrate Gno users by adding .gno at the end of nicknames for people to use. All our grant recipients are documenting their journeys in the hackerspace repo, check out [Teritori’s](https://github.com/gnolang/hackerspace/issues/7) journey.

## Resident Tinkerer, Zack

Another grant receiver, Zack, has been making significant progress on his microblogging project. You can check out the specs on GitHub ([PR 791](https://github.com/gnolang/gno/pull/791)) or watch the informative tutorial video, [Go to Gno: How to Build a Microblog](https://www.youtube.com/watch?v=F-_dadxcRJM). You’ll find this especially useful if you have a background in Go and need some additional insights to turn your hand to blockchain coding. Zack has also been working on an implementation of a smart contract for creating and transferring text-based NFTs that conform to haiku poetry standards (find out more on GitHub ([PR 860](https://github.com/gnolang/gno/pull/860)). Other than that, Zack continues his Gnolang journey, “learning and having a lot of fun.”

## EthSeoul, BUIDL Asia, and Getting to Gno

June saw members of our core team heading over to Seoul, South Korea, for a week of networking, talks, and events. Our VP of Engineering Manfred Touron gave a keynote on the evolution of smart contracts and an introduction to Gno.land for participants of EthSeoul, followed by a fascinating dive into Proof of Contribution at BUIDL Asia, where we also had a booth. It was an honor to meet so many talented and motivated Korean developers and contributors from around the globe. Seoul is a hotbed of up-and-coming talent and we’ll definitely be back soon.

We also had the chance to meet with our most active ecosystem contributors Onbloc and co-hosted an event together, Getting to Gno, at the Code States developer academy along with long-time Cosmos builders, Cosmostation. Attendees had the chance to hear about what the core team is building and see some of the great work of our community. A massive thanks to everyone involved, it’s awesome to be BUIDLing together! Read more about our Korean adventures in this [fab write-up by Onbloc](https://medium.com/onbloc/2023-buidl-asia-recap-894c60a1c0f).

EthSeoul - [Watch the talk here](https://www.youtube.com/watch?v=_iSsStlmxoU)

BUIDL Asia - [Watch the talk here](https://www.youtube.com/watch?v=v6k3NHm5vcE)

## EthBelgrade

Core contributor Milos Zivkovic rocked the Gno.land presence at EthBelgrade in Serbia, giving an introductory workshop about Gno.land, called 'Alice in Gno.land'. Being the first Ethereum conference organized in Serbia, there were lots of attendees from all over the Balkans. Participants joined in a journey through the enchanting realm of Gnolang and the Gno.land platform. Most of the participants were not aware of Goland before but were avid Gophers eager to learn more about the application of the Gno language in blockchains.

## GopherCon Berlin

The Gno.land team also had a blast last month at the European edition of GopherCon in Berlin. We had a booth at the event for two days, where we networked, talked about all things Gno, made some amazing connections, and even shared some live code! We’re looking to build an active, open-source Gopher contributor group in Gno.land, so stay tuned for more on that soon.

Coming up later this month, Gno.land is an official sponsor of EthCC, Paris, July 17-20. Stop by our booth to pick up some swag, say hey, and ask your questions about Gno.land. You can also catch us at the Nebular Summit for a keynote and workshop by our VP of Engineering, Manfred Touron.

*Do you want to contribute to Gno.land’s monthly updates? If you’re building on Gno.land and want to highlight your development, project, event, or idea, let us know and we’ll include your contribution.*
