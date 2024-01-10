---
title: "The More You Gno: Gno.land Monthly Updates - 5"
publication_date: 2023-10-10T13:37:00Z
slug: monthly-dev-5
tags: [gnoland, gnovm, tm2]
authors: [christina]
---

It's been another productive month, packed with developer calls, live events, new contributors, a large team presence at the Go community's biggest event of the year, GopherCon 2023, and the launch of a PoC gaming dApp on Gno.land, GnoChess. We uncovered a bunch of bugs in the code and some issues with the GnoVM, and made further progress on the Go and Rust VMs, the banker module bug, Gnofee, and much more. Check out the updates below.

## Building a Web3 Chess Server on Gno.land - GnoChess

Most of our work over the last few weeks has been dedicated to [GnoChess](https://gnochess.com/), a [PoC gaming dApp](https://test3.gno.land/r/gnoland/blog:p/chess-gc23) unveiled at GopherCon 2023. As gold event sponsors, we wanted to provide gopher attendees with a memorable experience – and a little friendly competition – while battle-testing the Gno.land platform. As our first gaming dApp, developing GnoChess was extremely useful for our team in many ways. We managed to attract 61 players to the game during the event, including some die-hard web2 gophers who wanted to show off their moves and discover more about Gno.

Several PRs were opened as a result of our endeavors, and, beyond the conference, GnoChess taught us a lot about where we're at with Gno, how to successfully build complex dApps on top of the platform, and how well we work as a team. We uncovered some key issues and breaking behavior in the GnoVM, made our JavaScript clients much more reliable in their communications with the Gno.land node, and unearthed further issues that lead to complex errors and potential security flaws that must be addressed before mainnet.

For example, appending nil to a slice of errors resulted in a panic, or conditional statements like if not supporting custom boolean types. The GnoVM doesn't currently perform terminating statement analysis, which results in a cryptic panic message ([issue 1086](https://github.com/gnolang/gno/issues/1086)), and mixing untyped (negative) floats and integers in arithmetic sometimes drops the sign ([issue 1152](https://github.com/gnolang/gno/issues/1152)). The issues uncovered while developing GnoChess were discussed extensively in the public developer calls of [Sept 6](https://www.youtube.com/watch?v=BBBqgycMjqU) and [Sept 20](https://www.youtube.com/watch?v=WrxFVPR55G0), and referenced in the [GitHub meeting agenda](https://github.com/gnolang/meetings/issues/31). Most of the issues are common in software development and fairly simple to fix by making some implementation changes or adjustments to design choices.

While developing GnoChess, our engineers took on the role of expert platform users rather than core team members. This approach was very useful as it pushed the platform to new limits, and allowed us to dive deep into many aspects of the project, creating a culture of sharing by opening up issues for each bug and asking for feedback and support. We'll definitely take a similar approach for future app development and onboarding new devs to Gno. We'll be releasing a retrospective of our experiences in the coming weeks. In the meantime, if you want to build a dApp on Gno.land, check out the GnoChess repo, where you can find a useful [tutorial](https://github.com/gnolang/gnochess/blob/main/tutorial/01_getting_started/README.md) or watch the recording of the GopherCon workshop, '[Chess: The Gnolang Way](https://www.youtube.com/watch?v=JQh7LhqW7ns).'

## The Battle of the Virtual Machines

Core engineers Marc and Petar continue their excellent work developing two different VMs for Gno, one in Go and one in Rust. In the coming weeks, we'll have a face-off, comparing and contrasting their features, efficiency, speed, and performance, so watch this space! For now, the definition of the virtual machine is stable for both, and they are no longer working on the virtual machine definition. They are mainly focusing on code generation; everything from parsing to scanning to parsing and compiling. Let's see how they are shaping up.

### Rust VM

Petar has developed a Rust implementation not only of the virtual machine but of the whole chain, including the compiler. He has written a Go compiler entirely in Rust and has even started experimenting with changing the compiler to implement the Invar proposal from Jae. Further progress includes porting a part of the parser and scanner from the Go compiler to Rust (almost a direct translation from Go to Rust) and making it stable. 

In addition, Petar has completed work on typed nil values and improving the recursive closures of Go, which were not working with Gno code and needed additional pointers. He has also implemented Iota and hooked up the garbage collector. In the coming weeks, Petar will be working to smooth out bugs and implement type aliases, as well as implementing function analysis for the dependency graph. The dependency graph is necessary for compiling global types in the correct order, so, for example, when type A refers to type B, you need to compile type B first so that when type A refers to it, type B exists.

### Go VM

Marc is currently rewriting a parser and a scanner from scratch. His work is not as far along as Petar's, but he's getting closer, and the code generation works well. He is currently refactoring and building a single-pass compiler that can perform a **syntax-directed translation**, which means there are no intermediate data structures between the source code and the byte code. This is a much simpler design that should compile faster and be easier to maintain, but it requires a complete redesign. 

Marc believes his Go parser will be easier to maintain and understand than the one in Rust and benefit the user since the entire stack is written in Go. However, to assess the best implementation of the VMs, Marc has started a Go **test shoot project, which is a script** that will run many samples to verify that the compiler (in Go, Rust, or any other implementation) conforms to Go's specifications. Marc and Petar will open their repos soon, and the next edition of The More You Gno will highlight how the GnoVM works. 

## Gnoffee: Coffeescript for Go and Gno

Gnoffee (hackerspace [issue 22](https://github.com/gnolang/hackerspace/issues/22)) will be a powerful standalone tool to elevate the development process of Go and Gno by generating code and integrating new features, eliminating manual coding. We aim to create a custom variation of Golang that preserves similar readability, maintains compatibility, and enables being able to code in Gno very quickly when you know how to code in Go. How do we go about this? 

Regarding compatibility, one possibility is to propose all our changes to Golang and wait for approval before we start developing. However, this is likely to take some time. Another approach is to use a way to transpile TypeScript for JavaScript or Coffeescript for JavaScript, so it's another language passing through a program that creates standard valid Golang and will generate valid Gnolang. With this simple method, we can experiment with missing features like new native types, and new keywords, and when we have new features in mind, we can develop what we lack. 

For instance, it does not make sense to have extra security for your exported variables when you write a library in Go. However, in Gno, it is very important to ensure that everything you expose cannot be modified by other contracts. This means finding a way to expose constants and other readable elements without risking their values being overwritten.

Besides allowing us to carry out all types of experimentation more easily, Gnofee could eventually be a way for the Go team to measure the potential adoption of Gno. Gnofee is not a priority for the mainnet, but we're excited to work on this important initiative.

## META Multinode Testnet

The discussions about single and multinode testnets have been ongoing, so we opened an issue to establish a multinode testnet focused on multi-validator experimentation, including stability, benchmarking, and lifecycle management. This multinode testnet aims to provide a platform for in-depth explorations and evaluations of multi-validator setups, while we maintain the single-node test3+.gno.land set up, primarily dedicated to showcasing the VM and providing examples. Visit hackerspace [issue 9](https://github.com/gnolang/hackerspace/issues/9) if you want to participate in this initiative or share your insights.

## Banker Module Bug

The banker module bug is a known issue that needs to be fixed before the mainnet because, currently, it's still possible to mint new GNOT tokens from any contract. Several fixes have been suggested, and our goal is to merge [PR 875](https://github.com/gnolang/gno/pull/875) put forward by Onbloc to change the denomination of the coins minted by the banker. Merging this PR is currently blocked by 2 small failing checks, but we are close to resolving this issue.

## Preserving Go Comments in Protobuf

In [issue 1157](https://github.com/gnolang/gno/issues/1157), Jeff from Berty raises the question about preserving Go comments in the Receiver field. Currently, Amino converts the code, but the proto message Receiver field doesn't have the comment. Manfred agrees that informative comments are helpful. However, he doesn't want to create a complex Protobuf configuration. We will continue to discuss this issue to look for solutions, but for now, Berty will parse the original Go source code and get the comments this way.

## Multi-Sig and Security Features

Several contributors, including Teritori, are working on built-in multi-sig support in Gno.land, where Gnokey supports a multi-sig setup. We also want to introduce additional ways to improve the UX and security of Gno.land (and web3 in general). An idea we currently have is to add a new layer in authentication, creating something similar to browser cookies that we can name sessions. The chain will have two tables, one with the public key for an account and one with a public key for sessions linked to an account. From your main account, you can create a session with self-destructing features, such as destructing after one hour without usage or after 24 hours. The goal would be to allow more complex and secure flows when starting your operations. We may not want this for multi-sig, but it comes under the same family of security and privacy features.

For example, imagine a wallet like Adena uses your key, a passphrase, or a ledger. It will sign a new public key that you just created in memory. Each time you close your browser, the memory is cleared. You can also have a logout button to call on the blockchain to delete all your sessions or simply wait for the session to be self-destructed, especially if the session was just in memory on your side. We will continue to develop this idea.

## New Team Member

We're excited to welcome a new DevRel team member to Gno.land, Leon, who's been in blockchain development for two years and is passionate about engineering and teaching. Leon has taught languages, development, math, and music privately, as well as an OS fundamentals class at his previous faculty. Welcome on board!

## Grantee and Ecosystem Updates

As Gno.land core continues to advance, so does our blossoming ecosystem, with new contributors and community members turning their eyes to Gno. The overriding theme of this last month has been collaboration, and we're pleased to see gnomes working together to overcome their obstacles and push their projects forward. Let's see what they've worked on over the last few weeks.

### Onbloc

Onbloc is powering ahead, contributing to Gno.land core, making upgrades and improvements to Adena and Gnoscan, and developing the Gnoswap DEX. Last month, Onbloc released the patched version 1.8.0 of Adena, which includes some UI and UX enhancements, such as more intuitive account management settings, a copy icon next to the names of the accounts, and some bug fixes. This release also comes with new injection methods to enable dApps to request users to add a custom gno.land network or switch to an existing one. Check out the [release note](https://github.com/onbloc/adena-wallet/releases/tag/v1.8.0) for more details.

Onbloc has open-sourced the code for Gnoswap on this GitHub [repo here](https://github.com/gnoswap-labs/gnoswap). You can also find a guide to running unit tests. The team continues to improve the Gnoswap interface, focusing on the earn and staking pages, the graphs for positions, and some components for adding and removing liquidity and providing pool incentives. They're working on the next iteration of the interface, with the governance and airdrop pages, and developing the front-end logic to integrate with Gnoswap realms and APIs. Onbloc also contributed to Gno core, adding PRs for fixes to testing and the banker module. Keep up with Onbloc through their [hackerspace journey](https://github.com/gnolang/hackerspace/issues/29) and check out their latest initiative [Gnodesk](https://medium.com/onbloc/gnodesk-week-2-of-sept-2023-5edbc451bba7), which delivers weekly highlights and updates from Gno.land.

### Teritori

Teritori has been working on improvements since the last update and open-sourcing all their work, including the DAO deployer and the Moderation module. You can visit the Teritori DAO tooling repo to find the complete documentation and new realms to easily deploy your DAO. There is also a tutorial on creating your own DAO using the framework. 

The team has made extensive progress on the Justice DAO deployer, a module that can be used for third-party arbitration when there is a problem with the escrow system in a decentralized freelance marketplace. The Justice DAO can resolve potential conflicts between the seller and the buyer and implements randomness to choose the judges to solve problems without conflicts of interest. The content flagging system, which highlights the content that users deem to be inappropriate, has been tweaked and improved. Keep up with Teritori's [hackerspace journey here](https://github.com/gnolang/hackerspace/issues/7).

### Berty

Berty has already completed the first phase of the project and published the [technical proposal](https://github.com/gnolang/gnomobile/issues/15) to develop the Gno mobile framework. The team is now busy with the second phase of implementing the proposal and the gRPC interface, which is working with the local socket on Android and iOS. Jeff has been trying to use Amino, and, now that Iuri is back from vacation, the team will work on improving other parts of the interface. Check out their latest [demo](https://www.loom.com/share/c0f68f707d3e47089c2fdbd2698fc92f), which shows an example user interface with wallet functions and blockchain communication. 

Onbloc has laid the foundations for Gno mobile apps with the Adena mobile wallet, so Berty will use some of this code in the mobile framework and work with Onbloc to ensure a similar user experience across all Gno apps.

### Flippando

Dragos, the developer behind new grantee Flippando, is an experienced mobile app developer. Flippando is a simple on-chain memory game, which is currently written in Solidity and deployed on several testnets, including Goerli, Polygon, Near, Aurora, Evmos, and a Saga chainlet. Fippando started as a project for Dragos to learn Solidity but has already been the winner of two hackathons in Korea. It can be deployed relatively easily on any machine and is currently being ported to Gno.land. Dragos is exploring which user intersection can be more beneficial for this and will show us a demo in the coming weeks. Soon, we'll have two gaming dApps on Gno.land – Flippando and GnoChess! Read about Flippando in the [hackerspace journey](https://github.com/gnolang/hackerspace/issues/33).

### New Contributor Joseph Kato 

We have a new contributor to Gno.land who showed a demo last month of what he's been working on, a language server to run tests and scripts. Joseph is a major Go fan looking to get into web3 and was super excited to come across Gno. While interacting with Gno.land, he found many IDE-like features that he missed when working on files, so he decided to work with an LSP implementation—gnols—with the goal of making these features available to all contributors regardless of editor preference, starting with Sublime Text and Neovim and moving on to IntelliJ, Golang, and Emacs. This is a welcome addition for anyone who has ever developed a realm in Gno. Check out his [hackerspace](https://github.com/gnolang/hackerspace/issues/34) page for more details. 

## DappCon, Berlin

Manfred was back in Berlin in September at the Radial System presenting 'Gno.land: The Key To Perpetual Transparency,' where he discussed how Gno.land offers a familiar, seamless experience for code sharing and a sustainable and transparent path for blockchain development. 

## Web3 Family

Core dev Miloš Živković gave a talk at Web3 Family in Barcelona last month, 'Gno.land and Gnolang: The Dynamic Duo of Blockchain Development.' He presented a brief history of smart contract development and the issues associated with existing platforms, such as limitations in design and security. He introduced Gno and showed how we make web3 accessible and blockchain development more intuitive and secure. Catch the [talk here](https://www.youtube.com/watch?v=0K-jr_Ad3bI).

## GopherCon 2023

Gno.land was out in force at GopherCon 2023 with a well-stocked booth at the conference and an awesome workshop building a web3 chess server on Gno.land. Both Manfred and Jae were at the booth championing Gnolang to Gophers, and we received a lot of positive feedback, some new contributions, fresh PRs, and exposure for Gno.land in web2 circles. It was also a fabulous chance for the team to meet for valuable face-to-face time.

That's all for now! Be sure to check back again with us for the next edition of The More You Gno to keep up with all our progress.
Do you want to contribute to Gno.land's monthly updates? If you're building on Gno.land and want to highlight your development, project, event, or idea, let us know, and we'll include your contribution.
