---
title: "The More You Gno: Gno.land Monthly Updates"
publication_date: 2023-04-15T13:37:00Z
slug: monthly-dev-1
tags: [gnoland, gnovm, tm2]
authors: [christina]
---

We made progress across the board at Gno.land last month, from onboarding more devs to receiving an influx of contributions to the Game of Realms contest. To encourage development and discourse, we set up a biweekly public developer call in addition to our biweekly Office Hours sessions. Anyone can join, ask questions, and give their suggestions on how to shape the Gno.land platform and become a contributor. Last month, we covered several pressing topics from Gno IDE and Gno.land website language, to GnoVM, IBC, and ICS. Jae also came back to the circuit in March with two IRL workshops for devs at side events during EthDenver and Game Developer Conference (GDC) in San Francisco.

## Developer Updates

You can find the live streams of the new biweekly public developer calls on [Gno.land YouTube](https://www.youtube.com/@_gnoland/videos) as well as access the agendas on [GitHub](https://github.com/gnolang/meetings/blob/main/notes/2023_03_15_dev_call_notes.md). The main talking points this month were Gno IDE, Gno.land website language and UX, garbage collection, bug fixes, and how to bring IBC and ICS to the platform. We are working on all these issues concurrently but the order of release will be Gno.land mainnet, IBC, and then ICS (this is reflected in the DAG below).



[![Gno.land mini DAG](https://gnolang.github.io/blog/2023-04-15_myg-march/thumbs/mini-dag.png)](https://gnolang.github.io/blog/2023-04-15_myg-march/mini-dag.png)

## Gno.land Website Language

We want to add more features for developers, such as libraries to make writing interfaces better and more consistent. There is an open topic for frontend developers with typography skills and library developers to create a UI framework for markdown or a custom rendering system.

Internally, our core team is working on improvements to Gno.land’s website, making it easier to navigate with shorter columns while ensuring the text is markdown centric and readable in plain text and the GitHub rendering machine. We hope to achieve this using CSS and having classes for vertical columns, without having to make an extension to the markdown parser.

## Gno IDE

Gno.land developer experience team is working on a web-based Gno IDE for quickly building Gno realms and packages right on your browser by just visiting a web app. Gno IDE will provide much improved UX for everything around building a realm (including making the testing easier), and additional features like autocompletion in the editor. Gno IDE will contain all the features you would expect from an IDE as well as valuable APIs for devs building tools around Gno.land with the public Gno Infrastructure.

[![Gno IDE](https://gnolang.github.io/blog/2023-04-15_myg-march/thumbs/gno-ide.png)](https://gnolang.github.io/blog/2023-04-15_myg-march/gno-ide.png)

Gno IDE will have multiple modes to support different use cases. The normal mode will be used during everyday developments (as you’re familiar with from other code editors). The presentation mode is for high accessibility and readability. You can use it during video calls or physical workshops while projecting your screen to an audience. The third and perhaps most interesting mode is the embedded mode. Use this mode to embed the IDE into websites and blogs. This feature is especially useful for tutorials to test out sample code, run it on the real testnets, and play with it.

## IBC and ICS

As depicted in the DAG above, Gno.land mainnet will launch first, followed by IBC and then ICS. We will focus on implementing IBC1, as we strongly believe in the ICS model and want to be a consumer of an existing Cosmos chain. We want a common ICS implementation that works across many hubs because Gno.land is a type of hub that will need its own ICS to scale while providing GnoVM on consumer chains on the Cosmos Hub. Our next step now is to find the best way to configure ICS for Gno.land and make GnoVM available as a consumer chain in the Cosmos Hub system.

Regarding IBC, we will use the current implementation that was written for the Cosmos SDK and port that over to Tendermint2. We anticipate some issues along the way including security patches that need to be applied to our code base. There are multiple ongoing directions and discussions about how to bridge Gno.land’s smart contracts to IBC, which are essentially Interchain smart contract interactions.

One possibility is to have an API that submits events to a queue of outgoing events, and another queue to receive and consume events asynchronously. This mechanism could work for IBC2 to have rich inter-contract Interchain features, and the same API could work for Interchain plus smart contract interactions that require advanced options. We discussed a proposal to create a standard for Interchain contracts so that IBC2 could eventually be standardized eliminating limitations by applying it with an EVM, other languages, and CosmWasm.

This protocol could be based on Protobuf or a similar well-known syntax definition protocol so that we can push the Interchain to the next level. IBC2 will be safe and fast and replace vulnerable atomic bridges between multiple technologies. This is a major update that we are committed to developing and we need help identifying all the challenges involved. Working on IBC integration, separate from the Gno.land mainnet launch, will require significant time to understand how the light client system works. If you’re interested in taking on this task, let us know and we’ll set up a group. IBC will likely be the most important challenge of Game of Realms phase 2.

## Garbage Collection

Currently, our work on garbage collection does not address the problem in the traditional Golang sense of dealing with memory efficiency. Instead, we are progressively optimizing and improving the main state tree by automating the clean-up of orphan nodes. The next phase will be targeting the official garbage collector component to begin work on memory management as we have some common Golang garbage collection challenges, but are identifying some uncommon ones too.

We need to consider elements like where to hold our objects because this is tied to releasing them in a concurrent lock-free way. We also need a good data structure. This is ongoing research as of now to implement a dedicated routine to synchronously clean stuff in a non-blocking way.

## Game of Realms

This month, we have seen a massive uptick in contributions to Game of Realms phase one with a tidal wave of issues, general discussions, and PRs. One of the biggest things we worked on was adding support for MOD, which is a version of Go mod with an easier interface to manage your dependencies and version your dependencies. You can track the ongoing issue on GitHub [here](https://github.com/gnolang/gno/issues/390).

There have been some really strong contributions to the Evaluation DAO and governance module, as well as a big CLI refactor that went into our code base. We've also seen people contribute contracts like GRC 1155 or general improvements to existing realms, with many suggestions for fixing bugs. Finding bugs and reporting what people want is a good indication that the Gno.land platform is being picked up and gaining adoption.

You can find the Office Hours recordings that cover Game of Realms on YouTube [here](https://www.youtube.com/watch?v=JTmNg-b6Lcs).

## Developer Events Stateside

Gno.land hosted a lively meetup during EthDenver last month where Gno.land founder and core dev Jae Kwon gave a talk for Solidity developers called “Gno.land, the Inevitable Next Generation Smart Contract Platform." He compared and contrasted Gno.land and Gnolang to Solidity, and showed Ethereum developers how the GnoVM shifts the smart contract paradigm. You can watch the [recording here](https://www.youtube.com/watch?v=IJ0xel8lr4c).

Also in March, Jae hosted a gaming workshop at a side event during the infamous Gaming Developer Conference (GDC) in San Francisco. “Gno.land for Game Developers, Building Your App in Web3," showed participants a sample gaming app built on the Gno.land platform and offered them the chance to try their hand at writing a smart contract for their app with Gno.

## Virtual Events - How to Build a Forum

Core tech lead at Gno.land Miloš Živković held a virtual workshop for Go devs called “How to Build a Forum.” He showed how Gnolang is a fast and simple way to build and launch smart contracts using the Gnolang interpreter virtual machine that interprets Gno and eliminates the need for any servers or ORNs.

The VM allows for the memory state of your Gno.land application to persist automatically after every transactional function call, which is a completely new way to handle transaction volume and memory recall. You can watch the [full tutorial here](https://github.com/gnolang/workshops).

*We’d like the community to get involved in Gno.land’s monthly updates, so if you’re building on Gno.land and want to highlight your development, project, event, or idea, let us know and we’ll include your contribution.*
