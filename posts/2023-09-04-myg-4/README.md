---
title: "The More You Gno: Gno.land Monthly Updates - 4"
publication_date: 2023-09-04T13:37:00Z
slug: monthly-dev-4
tags: [gnoland, gnovm, tm2]
authors: [christina]
---

We’ve had more on our plates than ever over the last few weeks, with a huge team presence in Paris at EthCC and Nebular Summit in July, an opening talk at Stanford Blockchain Club in August by Gno.land’s founder Jae Kwon, and some awesome contributions from Gno.land grantees and ecosystem partners, including the first demos of Gnoswap and Teritori’s social platform and DAO deployer. We continue to make solid progress on GnoVM, an alternative VM in Rust, Tendermint2, native bindings, and much more. Check out our latest developer updates below.

## Upgrade Strategy for AVL Between GitHub and test3.gno.land

One ongoing discussion is about an incompatibility bug that affects many things we do on Gno.land. The current AVL implementation on the testnet is outdated and does not match the AVL implementation users get when they pull in the latest master branch. Therefore, building and deploying contracts on a local Gno chain (with the latest master changes) and deploying those same contracts on the testnet may fail due to this incompatibility. We need to find a way to seamlessly integrate these two approaches. Ideally, when you write code on the master branch on GitHub, it should work on the testnet as well.

In [issue 970](https://github.com/gnolang/gno/issues/970), you can find details of five different proposed solutions to implement this upgrade strategy, from resetting the whole blockchain (which would mean losing on-chain content and debugging information) to implementing a migration feature specifically for testnets that allows developers to rename packages and patch their contracts before publishing them. There are pros and cons to each proposal, and we continue to work together to find the best way forward.

## Encoding JSON and the Discussion Around Reflection

Some contributors have highlighted the need for native JSON encoding, and we are discussing how best to approach it. See [issue 808](https://github.com/gnolang/gno/issues/808) for further details. One idea is to copy the code from encoding JSON in the standard library Go and take it over to Gno, but we would need to have reflection to do that. So, the important question here is whether we want to have reflection and, if so, what it should look like. We could emulate Go’s reflection package with some added elements, like being able to inspect the realm state, but we would need to be extremely careful about how we do this.

For example, should users be able to read private fields of external packages through reflection or even *ufmt*, or could that introduce a problem? It would be simpler, and the language capability security would be tighter and easier to understand if we made accessing private fields impossible, but that would also make it limited. We could consider supporting reflection as an internal user package and whitelisting and encoding JSON. This way, new encoding packages would have to be whitelisted because they’re using the reflection package. We could also mark reflection as unsafe so developers know they must carefully audit their work.

Another solution is the partial implementation of reflection. In [issue 971](https://github.com/gnolang/gno/issues/971), Gno.land core engineer Petar discusses introspection, which involves implementing reflection as Go has it now but enabling only one of its two main capabilities: the ability to inspect types, but not the ability to modify code. The main difference between introspection and reflection is that, since it is done at compile time, it is completely type-safe. This discussion is ongoing.

## Alternative GnoVM Implementations

To deliver the best possible virtual machine, we’re working on two different implementations of GnoVM. Petar has spent the last three weeks developing a new GnoVM implementation written in Rust. His work is still private as the machine is not yet ready for public use, but he will soon make the code public for your inspection. Rust gives the ability to write more performant code and, in some scenarios, the Rust GnoVM can run up to 20 times faster than the GnoVM at roughly 87 milliseconds compared to 2,000 milliseconds on a Fibonacci benchmark, which is a considerable improvement in speed.

Since one of Gno.land’s core features is that the entire tech stack is written in Go, we’re unsure if everyone will appreciate a Rust GnoVM or whether it aligns with our vision. However, it’s always good to provide alternatives, and, Petar argues, as long as the VM carries out the same functions (and does so more cheaply), most developers won’t mind what language the VM is written in.

Rust has a few other features that some developers may favor over Go, such as more tools for creating languages, advanced garbage collector libraries that allow you to change the algorithm without changing the runtime (by swapping out a tricolor algorithm for a generational one, for example), and built-in data structures that solve many issues. For example, we needed a deterministic map that is fairly fast. With Rust’s Btree in the standard library, this was simple, Petar only had to implement the native Go map type with a Btree map from the standard library. This took just a few minutes.

Core team dev Marc has also started an initiative to improve the Go GnoVM so that it is faster and offers a clean and user-friendly interface. He believes the debate over the VM is more about whether to have a VM that is bytecode-defined or AST-defined (rather than speed). Marc has been comparing the fundamental differences between the two and noted that the bytecode version is 15 times faster than the AST. This means that changing to Rust would give an increased performance of 2-3 times.

The VM must be fast, secure, and performant in many ways. In either version, the AST will be stored on the blockchain, whereas the bytecode is only an internal representation that doesn't affect the users. We must still consider any potential architecture consequences between bytecode and AST before deciding whether to change. Marc’s WIP code is still in a private repo, but you’ll be able to inspect it soon and make a comparison of the VM implementations in the coming weeks. The decision about the direction of GnoVM is still very much TBD; however, the Rust GnoVM will not replace the Go GnoVM but will complement it, eventually giving validators the choice of which to run.

## Defining Wording for People/Documentation and Consistency

[Issue 1024](https://github.com/gnolang/gno/issues/1024) discusses the need to define the wording we use throughout our documentation, for example, how we name a module, package, sub-module, etc. Once we have the wording defined, we will set the GnoVM to only accept elements with the correct naming. The importance of wording affects the design choice of the whole project and how we go about versioning for the best possible user experience.

For example, is mt/board/admin part of the same realm of mt boards, or is it its own realm? Can we work with both by adding patterns to have some realms responsible for hosting data and others responsible for having more privileged actions? How do we split a complex realm into sub-libraries and sub-realms? We want to define the documentation and the logic for this and have begun to touch on this issue. We will discuss this in greater depth in the upcoming developer calls.

## Improving the GRC20/Foo20 APIs

When working on the specs for a Merkle airdrop contract, Albert came against some issues with users initiating airdrop reward claims (see [PR 906](https://github.com/gnolang/gno/pull/906) for more details). Currently, when the Merkle airdrop contract tries to execute the reward claim for the user, an instance of the GRC20 contract is used for transferring. Within the GRC20 implementation Transfer() method, the caller (token sender) is fetched using the standard library method std.PrevRealm().

However, calling this method in the Merkle airdrop context returns the user as the caller, not the Merkle airdrop contract, which is an unexpected functionality. We are discussing different ways to tackle this issue efficiently. However, each solution would require possible changes to the GRC20 API and subsequent token implementations. Additionally, as part of [PR 952](https://github.com/gnolang/gno/pull/952), we are looking into improving the standard GRC20 API and possibly resolving the ambiguity with standard library calls that are causing the mentioned issues.

## Client Optimized for CLI, Not Mobile

Our newest contributor to Gno.land, Berty, is developing the mobile version of Gno, which means writing a mobile app to interact directly with the blockchain. The team is facing some issues as they need a client library with utility functions like sign and broadcast, which are used by the command line. This code (tm2/pkg/crypto/keys/client) is not ready for external users yet, and the Gno client is designed for CLI. However, Berty needs a way to interact with the Gno chain from their application and to call the logic without adding the full CLI.

From the existing TypeScript/JavaScript client library (gno-js-client and tm2-js-client), Berty should be able to build out a Go client library by exclusively using the RPC endpoints of the node itself (just like gno-js and tm2-js work), and not having to worry about importing private logic like transaction broadcasting. The team is writing its own framework to call Go code for Gno from Java, Swift, and React Native mobile apps that creates a transaction and sends it (see [PR 1047](https://github.com/gnolang/gno/pull/1047)).

They are working on an API that interacts with the blockchain and lets them export the code without having to write their own utilities. The API will be minimal, and update the Tendermint2 build script by moving tm2txsync from tm2/cmd to gno.land/cmd (see more details in [PR 1080](https://github.com/gnolang/gno/pull/1080) here). For the time being, Berty will copy the code and use the objects directly until a more convenient API is complete.

## Tendermint2 Development

In [PR 546](https://github.com/gnolang/gno/pull/546), we introduce file-based transaction indexing. Transaction index parsing should be done as a separate process from the main node, meaning other services can be instantiated to index transactions as readers. The current problem is that there is no way to figure out whether a transaction has failed after it’s been sent out with a broadcast sync, or fetch any kind of receipt information or error reason in the delivered transaction.

So, we’ve started working on an event indexer to index Gno node events, which include transactions. Soon, developers and users will be able to ask the event indexer what happened to the transaction or in which state in its execution it's currently at, and also to retrieve information on other events like block commits as they happen.

## Extending the Functionality of Go

In [issue 919](https://github.com/gnolang/gno/issues/919), Petar proposes extending the functionality of Go by adding constant data structures, arrays, slices, etc. He believes this would benefit users, as they wouldn’t need to create special functions as in Go to simulate this behavior, and it would also catch bugs when there is mutation. There has been a discussion, and Jae has similar ideas with the notion of “invar” expressions, where the resulting value can only be read, not mutated or stored. This would fix the bug where if you pass a pointer (that represents part of your contract state) to another contract, the other party can “steal” it by assigning it to their state, and your contract would fail to execute.

Morgan believes that we should take a different approach as slices have the semantic in Go, where the underlying array is always heap-allocated and modifiable. Introducing constant slices would thus necessarily have to introduce concepts regarding im/mutability of values without the matching constructs that a language like Rust has. To make a compromise and keep compatibility with the Go spec, we are likely to implement this in a transpiler (gnoffeescript) that would implement this feature and be able to transpile to valid Go.

## Grantee and Ecosystem Updates

As you can see, we’ve made a ton of development progress over the last few weeks. We’re also steadily adding more gnomes to our community of builders, and they’re working on the core infrastructure of Gno.land, as well as the permissionless dApps the platform will house. Let’s see what they’ve been up to since the last update.

## Onbloc

Onbloc has been busy, as always, with a slew of updates for us over the last few weeks. The team has been developing Gnoswap, the first Gno.land automated market maker with concentrated liquidity, and they gave us a live demo. On the front end, which is still a work in progress, you can find a one-stop venue for traders to view all the information about tokens on gno.land, so you don’t have to move between Gnoswap and a token aggregator like CoinGecko. You can also see incentivized pools sorted by liquidity, volume, APR, liquidity mining rewards, etc., and a wallet page to check your balances. You will also be able to deposit or withdraw assets from the Interchain when IBC is enabled.

Check out the work they’ve done so far on the Onbloc [hackerspace](https://github.com/gnolang/hackerspace/issues/29). The team has also released [the documentation](https://docs.gnoswap.io/) about what you can expect from Gnoswap, the rationale behind their design choices, some information about tokenomics, a preview of the UI, and more. Their main focus is on delivering a smooth and welcoming user experience and abstracting away the difficult mechanisms of concentrated liquidity so that the interface is as minimal and simple as possible.

The team will be ready to launch Gnoswap as soon as gno.land reaches mainnet. Feature updates and enhancements will be aligned with the development of the core Gno Stack.  The code for Gnoswap has now been [open-sourced](https://github.com/gnoswap-labs), so you can take a look at everything they’ve done and even make suggestions. In the coming weeks, Onbloc will also work on building core Gno.land infrastructure to support an earlier launch. Find details of this in Onbloc’s [grant submission](https://github.com/gnolang/ecosystem-fund-grants/pull/4). And be sure to check out Onbloc’s informative 6-episode [blog series](https://medium.com/@gnoswaplabs/why-gno-introducing-gnoswap-dd6acc22e6a1) that features the history of blockchain and exchanges, a deep dive into the Gno Stack, and an introduction to Gnoswap, where they share details of their journey and insights.

## Teritori

We also saw an awesome demo from the Teritori team, which you can check out at app.teritori.com. Simply connect your Adena wallet to create a user name, start interacting with the social feed, create your own DAO, and add members. The team is working on more extensive documentation to explain how it works in more detail. While still a work in progress, Teritori has developed a cool flagging system that allows you to unfollow content you don’t like or flag content as inappropriate. If posts receive many flags, users can vote on whether to ban them, creating a healthy and supportive social environment free from derogatory content monitored by a like-minded community through a moderation DAO.

The team continues its work on DAO interfaces and has built a useful tool for speeding up the deployment of packages as a workaround until we’ve decided how to best tackle realm versioning. They are also working on the escrow system, which will be useful for the freelance marketplace, and presenting DAO standards documentation.

## Berty

We have a new contributing team to Gno.land from the Berty private messaging app. This team is working on a mobile version of Gno.land, implementing the WESH protocol, which is available by Bluetooth, local WIFI, or other means, and provides secure censorship-resistant communication between devices. The plan is to be able to provide an alternative transport for Gno applications when the internet is not available and build the skeleton/foundations that enable developers to create Gno-centric mobile apps more easily in the future. Berty brings a ton of experience in off-grid communication and getting apps to run on mobile devices, both Android and iOS.

The team has created its own [testnet](http://testnet.gno.berty.io/), which you are welcome to test out and play around with, although they will be restarting and rebooting without prior notice, so be aware that your work could be wiped. In the few short weeks they’ve been working with us, Berty has already finished their first Proof of Concept, a simple app running on iOS and Android. They copied code from the gnokey command line, and now it’s installing and running on mobile and interacting with the blockchain.

Now, Berty is working on a nicer UI for the app and will propose a project to create a formal framework called GnoMobile, which will allow anyone to create their own app and run it on mobile. We look forward to seeing their demo soon.

## Golang Working Group

In other news, we've started a bi-weekly [Gnome Golang Working Group](https://github.com/gnolang/hackerspace/issues/15) where we get together and discuss various topics, such as the language-related and theory elements of Go and Gno. We also aim to identify meaningful and reasonable ways to contribute to Golang, Gophers, and the general open-source community and improve our visibility there. We hope to attract more Go devs to the project and provide a “blockchain-less” experience for web2 Go devs.

We've had two meetings so far, and some recent hackerspace issues have already emerged from the discussions. One in particular that we’re actively evaluating is Gnoffee, a transpiler tool inspired by the likes of [CoffeeScript](https://coffeescript.org/) for Go and Gno integration. Gnoffee would be a powerful standalone tool to enhance Go and Gno (blockchain) projects by generating code and seamlessly integrating new features without manual coding. Find out more at the link above.

## EthCC and Nebular Summit

The Gno.land team was in full force in Paris at the end of July for EthCC, where we met many passionate developers and spread the word about Gno.land and, specifically, how Gnolang compares and contrasts to Solidity. We had a booth during the conference manned by the Gno.land team complete with awesome swag and a continuous presentation in the background playing on a full-screen television.

At Nebular Summit, our VP of Engineering, Manfred Touron, [gave a talk](https://www.youtube.com/watch?v=CtxBajCcTYQ) called ‘Gnolang for Developers: Examining the Core Stack,’ where he broke down the major components of Gno, demonstrated how the upcoming Gno SDK compares with the existing Cosmos SDK, and explained why Gno.land is an excellent choice for accessible and sustainable blockchain development.

## Blockchain Application Stanford Summit (BASS)

Jae opened the [Blockchain Application Stanford Summit (BASS)](https://bass.sites.stanford.edu/) event, attended by thousands of students and future blockchain developers. He gave an overview of Gno.land, GnoVM, and Gnolang, and explained the features that make our platform paradigm-shifting and timeless. He also dove into the core of why we’re building Gno.land – to provide a censorship-resistant platform for truth discovery that helps people improve their understanding of the world in an era of information censorship and control.

Coming up later this month, you can catch up with the Gno.land team at [DappCon Berlin](https://www.dappcon.io/) from September 11-13, where we’ll be delivering an informative keynote and hosting a side event to get to gno you better. If you find yourself in Barcelona for [Web3 Family](https://web3fc.xyz/) on September 23, you can join in a Gno coding workshop. You’ll also be able to meet the team at [GopherCon US](https://www.gophercon.com/) in San Diego. We’re hosting an action-packed workshop, ‘Chess: The Gnolang Way,’ on Gopher Community Day, where you can learn to build a web3 chess server on Gno.land and compete for cool prizes in an ongoing chess tournament throughout the event. More details coming soon. That’s all for now! Be sure to check back again with us for the next edition of *The More You Gno* to keep up with all our progress.

*Do you want to contribute to Gno.land’s monthly updates? If you’re building on Gno.land and want to highlight your development, project, event, or idea, let us know, and we’ll include your contribution.*
