---
title: "The More You Gno: Gno.land Monthly Updates - 7"
publication_date: 2024-01-22T00:00:00Z
slug: monthly-dev-7
tags: [gnoland, ecosystem, updates, gnovm, tm2]
authors: [christina]
---
Welcome to the latest edition of *The More You Gno*, your regular source of updates from the Gno.land core team and contributor ecosystem. After a well-deserved rest during the holiday break, we’re kicking off 2024 with renewed energy and plenty of exciting initiatives, including a new staging testnet (the Portal Loop), the official Gno.land documentation page, several merged PRs (including native bindings!), and many updates across the board. Dive in to find out what we’re working on and what our ecosystem partners and grantees have been up to.

## Gno Core Team Updates TL;DR

Short on time? Skim the highlights from the core team in the list below. You’ll find additional details in the next section if you want to explore any topic in greater detail.
- **Native Bindings** - If you’ve been following our journey or experimenting with the platform, you’ll hear virtual champagne pops as Morgan’s ongoing work with native bindings is finally merged [PR 859](https://github.com/gnolang/gno/pull/859).
- **Gnodev** - Thanks to Guilhem’s `gnodev` initiative [PR 1386](https://github.com/gnolang/gno/pull/1386), you can now create and develop contracts with a single command.
- **Gno.land Offical Docs** - Check out [docs.gno.land](https://docs.gno.land) for how-to guides, getting started, and an overview of key concepts of the platform.
- **Effective Gno** - Taking inspiration from *Effective Go*, Manfred’s begun listing common patterns and examples of the differences between Gno and Go.
- **Assignment in GnoVM** - Jae is working on approaches to fixing assignment in the GnoVM and issues that deal with persistence [(issue 1326)](https://github.com/gnolang/gno/issues/1326). 
- **Portal Loop** - The [Portal Loop](https://portal.gnoteam.com) has been released on a staging domain and is being tested.
- **Roadmap** - We’re working on a fully-fledged Gno.land roadmap and will share a detailed DAG and important goals and milestones with you soon.
- **Tendermint2 Update** - There are several PRs aimed at removing the dependencies between Tendermint2 and GnoVM.
- **Gno.land Tokenomics** - We continue to make progress in defining the structure of Gno.land’s DAOs and the design of reward schemes for contributors.
### Native Bindings (PR859) Has Been Merged
[PR 859](https://github.com/gnolang/gno/pull/859) (native bindings) was submitted by Morgan in May 2023 to improve calling Go code from Gno standard libraries, all while improving `gno doc` documentation for standard library functions. Native functions are _declared_ in Gno code, but their definition (the underlying code) only exists in Go: this is similar to how Go and many other systems languages implement assembly functions. Overall, the addition will now allow us to better support precompilation (transpiling Gno code to Go) for all Gno-specific standard libraries, like [`std`](https://docs.gno.land/reference/standard-library/std/address/), and have a system for defining such functions that is transparent to code analysis tools like `gno doc` and `gnols`.
### Gnodev Has Been Merged
[PR 1386](https://github.com/gnolang/gno/pull/1386) (`gnodev`) has been merged. Gnodev is a tool to locally develop Gno realms which automatically re-deploys your contracts when you change the files, similar to JavaScript frameworks `npm run dev`. There are some additional features being worked on to improve the experience, including browser hot-reload (for the full front-end JavaScript experience!)—and Gno core developers who have worked on realms all agree that thanks to `gnodev`, they can finally stop visiting their therapist every week. Play around with it, and let us know how you get on. There may be a few bugs still and Guilhem is happily accepting feedback.
###  The Gno.land Official Documentation Page Is Live
We’re excited to have the Gno.land Official Documentation page live on the [https://docs.gno.land](https://docs.gno.land) domain. This will always be a work in progress as we expand the docs, make iterations to existing issues, and refine some of the core concepts, but it’s an excellent resource for anyone wanting to find out more about Gno and for onboarding new developers to the platform. A big thanks to the Onbloc team, whose developer portal was a huge inspiration for this. We’re looking for feedback, so leave your reviews and let us know where the docs can be improved and what else you would like to see.
### Effective Gno
Manfred has been working on a document called [Effective Gno (PR 1000)](https://github.com/gnolang/gno/pull/1000), which takes inspiration from *[Effective Go](https://go.dev/doc/effective_go)* and will become an important reference document for Gno devs to explore common patterns and crucial differences in how we program compared to Go. We’ll be iterating on this as we progress, but you can already find plenty of examples. If you’re just getting into Gno and coming from a Go background, this is a great resource. Read this document and provide some comments if you have any. 
### The Portal Loop Beta Is Live
The Portal Loop Beta has been released on a staging domain, and you can check it out now at [https://portal.gnoteam.com](https://portal.gnoteam.com). The Portal Loop will replace the Gno.land website once we’ve finished squashing bugs and adding features. We’re still testing it and have identified several issues. For example, from the last three merged PRs, only one triggered a redeploy when we expected two or three deploys. We will also add a faucet.

As we continue to evolve the Portal Loop out of its early development stages, transaction volume and general activity will increase. However, currently, there are insufficient transit testing transactions. One of the tasks we want to do to prove that the Portal Loop is working well enough is to write a kind of monitoring-oriented oracle that will try to make transactions, perhaps incrementing a counter every minute. We’re looking for help writing a script or a daemon for this oracle, so let us know if you want to contribute to [issue 1443](https://github.com/gnolang/gno/issues/1443). Once the Portal Loop is finished, we will focus on testnet 4.
###  Assignment Issues in the GnoVM
Morgan came across a bug [issue 1326](https://github.com/gnolang/gno/issues/1326), which returned an error about an [“unexpected unreal object”](https://tenor.com/es/view/cranizox-gif-8576622211330078986) when assigning a local variable to a dereferenced global variable in the GnoVM. Jae has been spending some time working on approaches to solving this and fixing assignment that will also work for saving escaped objects that don't have a parent (like variables whose pointers are referenced on a persisted object). This is a tough one to figure out, so if there are any other VM issues that deal with persistence and detached parentless objects, now is the time to add them to Jae’s plate.  
### An Update on Tendermint2
[PR 1483](https://github.com/gnolang/gno/pull/1483) has the same goal as [PR 1438](https://github.com/gnolang/gno/pull/1438): to make Tendermint2 completely independent of GnoVM and Gno.land. This continues a project started many months ago to separate Gno into three separate components: the Tendermint2 consensus engine, the Gno programming language and VM, and Gno.land, the blockchain combining both together. This way, we’re working towards making it possible to build other blockchains that use Tendermint2 (like AtomOne!), the GnoVM, or both!
### Gno.land Engineering Retreat
In the last *The More You Gno*, we covered the Gno.land and AIB company-wide retreat, an invaluable opportunity to work together, code together, and get to know our peers outside of work. It was such a success that the Gno core dev team held another retreat in December in Rouen, France, where many of the above issues and PRs were tackled and merged. We look forward to more productive and frequent face-to-face meetings in the year ahead.
### Gno.land DAOs and Tokenomics
With the input of Manfred, Jae, and the rest of the team, Michael continues to make advancements on Gno.land’s system of DAOs and tokenomics. One key change since the last edition is that the WorxDAO (responsible for governance and all issues related to development in Gno.land) will now be known as the GovDAO. The DAO will likely have seven tiers but initially launch with three or four. The main benefits of moving up tiers are increased voting power, increased monthly rewards, and the authority to promote members from lower tiers. GovDAO will be assisted by WorxDAO, which will encompass several different sub-DAOs, such as engineering, funding, and projects. 

We’re currently exploring different reward systems for contributors, whereby each member of the same tier level will receive the same amount of rewards, either directly or indirectly, in the GNOT native gas token or USD, in a type of salary-based scheme. We may also elect to distribute rewards based on a contribution/work “hash difficulty” (total number and tier split of active contributors that month). We may also adopt a hybrid of these two models. 

Michael is also working on a bounty system to make Game of Realms (GoR) more accessible and evaluating contributions easier for judges. High ranking GoR competitors will likely receive Gno.land tier levels based on their leaderboard placing in addition to ATOM rewards. It’s important to note that these discussions are ongoing, and the information here may be deprecated. 
### Making Testing Faster

Thanks to Petar, [PR 1417](https://github.com/gnolang/gno/pull/1417), we have improved the entire VM testing suite runtime by around four minutes, which is an incredible achievement. We just need to refactor some test scenarios that are not very concurrent-friendly, but this PR makes interacting with the platform so much easier.

### Bug Fixes and Miscellaneous Items

Thanks to Joon from Onbloc, we were able to add support for octals without 'o' (check out [PR 1331](https://github.com/gnolang/gno/pull/1331) for more details), and thanks to Dragos [PR 1309](https://github.com/gnolang/gno/pull/1309), we extended the GRC721 interface so that it now supports setting a token URI. These are both extremely welcomed contributions, and we appreciate our ecosystem partners.

From the core team, a special shout out to Dylan for killing it fixing bugs, and getting many PRs ([PR 1451](https://github.com/gnolang/gno/pull/1451), [PR 1315](https://github.com/gnolang/gno/pull/1315), and [PR 1305](https://github.com/gnolang/gno/pull/1305), to name a few) merged over the last few weeks. Props also go to Marc for [PR 1177](https://github.com/gnolang/gno/pull/1177), which has just been merged, which fixes append in certain key situations. We’ve also welcomed a new security engineer, Kristov, to the team.

## Grantee and Ecosystem Updates

### Onbloc

Onbloc has been on a roll, giving us an internal demo of Gnoswap beta just before the Christmas break and a public demo of its awesome Pool Incentivization feature during the last contributor sync call. With Pool Incentivization, anyone can add extra rewards on top of swap fees for LP stakers. This will help bootstrap initial liquidity for new-coming projects by attracting liquidity providers until sufficient organic trading volume is secured. Onbloc is also actively developing Adena’s Airgap feature and has improved the sign-in flow for security enhancement along with some refactoring. There will be a demo coming up in the next few weeks. Onbloc will also be researching airdrop trends and aiming to identify some of the most coveted DEX features users want to see for Gnoswap to streamline the onboarding process.

Regarding Gno core, Onbloc core dev Byeongjoon Lee has developed a JSON parser for Gno, giving us a live demo during the last contributor sync. This allows the conversion or accessing of data from contracts in the JSON format, which will improve the Gno developer experience. His code is currently under review by the core team in [PR 1415](https://github.com/gnolang/gno/pull/1415). Dive deeper into Onbloc’s Builder Journey in the [hackerspace issue 29](https://github.com/gnolang/hackerspace/issues/29).

### Teritori

Teritori continues the challenging work of developing Gno Project Manager, a web app that allows anyone to create, fund, review, or manage projects fully on-chain. During the last contributors' call, the team gave a demo of the work achieved so far, in particular regarding the escrow system and completing project milestones so contributors can be paid once each one is completed rather than having to wait until the project finalization. 

Gno Project Manager is a complex goal, and the team has run into some issues with edge cases they hadn’t bargained for in the relationships between grantees and funders. The team is looking for feedback and help identifying edge cases, so if you have any in mind, let them know. Teritori is also working on the conflict solver module and improving the social feed on [https://app.teritori.com/feed?network=gno-teritori](https://app.teritori.com/feed?network=gno-teritori), as well as providing more detailed documentation on their work, which they’ll be releasing in the coming weeks.

### Berty

The Berty team has been busy working on GnoSocial backend implementation. The initial feature set has been implemented [here](https://github.com/gnolang/gnosocial/blob/main/realm/public.gno), including posting and replying to messages and reposting threads. You can keep up with Berty’s journey on GnoSocial in [hackerspace issue 51](https://github.com/gnolang/hackerspace/issues/51), which contains many issues and PRs, such as implementing calls, running tests, and fixing bugs. We’re super excited about pushing the limits of scalability with Berty’s decentralized social platform, and we’ll be looking forward to more demos in the coming weeks.
### Dragos
Dragos has successfully launched the Flippando game, and you can try it out on the [testnet here](https://gno.flippando.xyz/flip). If you haven’t been following the progress, Flippando is an on-chain memory game that you can play with your choice of styles, such as dice, colors, and hexagrams. Once you successfully complete a matrix, you can mint the end result as an NFT, which can later be assembled into larger, more complex NFTs to create digital artwork. You can find out more about the game, its creator, and the official roadmap on the site. We’ll also release a blog post soon from Dragos sharing his experience porting Flippando from Solidity to Gno, so stay tuned!
### Varmeta  
Varmeta’s update was brief this week since the contributor sync call ran over. We look forward to hearing more about the team’s progress in developing the Unity SDK for Gno next time. You can read more about it on Varmeta’s [hackerspace issue 43](https://github.com/gnolang/hackerspace/issues/43).

*Do you want to contribute to Gno.land's monthly updates? If you're building on Gno.land and want to highlight your development, project, event, or idea, let us know, and we'll include your contribution. That's all for now! Keep track of our progress by following our socials [Twitter/X](https://twitter.com/_gnoland) and [Discord](https://discord.com/invite/tF2X8M6cVj) and watch out for the next edition of The More You Gno in a few weeks.* 
