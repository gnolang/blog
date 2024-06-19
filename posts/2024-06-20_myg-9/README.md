---
publication_date: 2024-06-20T00:00:00Z
slug: monthly-dev-9
tags: [gnoland, ecosystem, updates, gnovm, tm2, test4, gnostudio, connect]
authors: [Kouteki]
---

# The More You Gno 9: Realm to Realm Interaction, Faucet Hub, New Test4 Launch Date and More

This edition brings several major pieces of news; Gno Studio Connect beta, and Faucet Hub. We're moving up the Test4 launch by 2 weeks, so we could get to mainnet faster. We'll also be seeing each other at GopherCons EU and US, so make sure to stop by and say hi! Oh, and our [new gnome logo is live](https://gno.land/r/gnoland/blog:p/the-gnome)!

We're also covering all the major and minor code improvements; if you want to dive deeper, we have the weekly engineering [updates](https://github.com/gnolang/meetings/issues/37) and video [recordings](https://www.youtube.com/watch?v=9ch7MhKNBmw&list=PL7nP7r1QiDktMCdw1ydQo2crM3y6Zk7E4) available. 

# Gno Core Updates

## Introducing Gno Studio Connect

Last month [we talked](https://medium.com/@gnostudio/introducing-gno-studio-the-premier-builder-suite-for-gno-land-d1b4d82b46de) about the need for a premier builder suite for Gno.land. As Gno.land expands into a universe of realms, we saw the need for a set of tools; tools empowing the community to create and use succint and composable realms on Gno.land. Last year we launched [Gno Playground](https://play.gno.land/). Now we're adding another tool to the toolbox - [Gno Studio Connect](https://gno.studio/connect). Currently in beta, Connect allows seamless access to realms, making it simple to explore, interact, and engage with Gno.land's smart contracts through function calls. Try out your first realm interaction via Connect by taking our [gnoyourdate](https://gno.studio/connect/view/gno.land/r/gnostudio/gnoyourdate?network=test3#Vote) poll. 

## Faucet Hub is Live

With the [Faucet Hub](https://faucet.gno.land) you can now easily and effortlessly use all Gno.land ecosystem faucets, including Portal Loop, staging, future testnets and implementation partner chains. It's easily extensible, with the goal of having a single stop for every possible Gno.land test token you might need.

## Test4 Launch Scheduled for July 15

After announcing the tentative [Test4](https://gno.land/r/gnoland/blog:p/test4-explained) launch for the end of July/early August, we realized we're working faster than anticipated. That's why **we're moving the official Test4 launch to July 15, 2024**!

## 550 TPS

Recent [supernova](https://github.com/gnolang/supernova) tests showed that we have **~550 TPS** on a single node machine (M3 Mac; 100k txs, 1s block time). This is a huge step up from last year, when the TPS performance we had varied from 7-20 TPS with the same setup.

## New DevOps team member

Please welcome Sergio Matone! A DevOps with a strong Go background, he'll be a key player in improving and managing the Gno.land infrastructure.

## Belgrade Retreat

The core team travelled from across the world to Belgrade, Serbia, where we hunkered down and solved a number of issues blocking Test4 as well as future releases all the way to the mainnet. Full recap [here](https://gno.land/r/gnoland/blog:p/gnomes-in-serbia).

## Changelog

- Dropped support for [all unused DB implementations](https://github.com/gnolang/gno/pull/1714), apart from leveldb and boltdb. Without excessive DB implementations that need maintaining, we can focus on optimizing the ones we actually use, and squeeze out performance from systems we already have. 

- Resolved a Long Standing CI Issue With BFT Tests. [The CI is finally green across the board](https://github.com/gnolang/gno/pull/1515). We still have some flaky tests, but the CI will no longer constantly fail on the BFT CI due to irregular channel usage. This, couple with our CI rework over the past month, significantly improved our GitHub development process.

- Lots of spring cleaning efforts with [gnokey](https://github.com/gnolang/gno/pull/1212), [gnoland](https://github.com/gnolang/gno/pull/1985), and our other tools. We took this chance to also add extensive coverage and regression tests, in case they were missing before.

- [Finally merged in Event / Emit support](https://github.com/gnolang/gno/pull/1653) in Gno! Users and Gno clients can finally fetch on-chain events as soon as they happen, and have a rich context for them as well. Support for these has already been propagated to our [tx-indexer](https://github.com/gnolang/tx-indexer/pull/43).

- [Merged in the VM gas consumption fixes](https://github.com/gnolang/gno/pull/1430), which standardize VM gas usage across the board for users. Having predictable gas costs, and gas costs that take into account VM operations is only fair for both the node operators and the chain users.

- [Reworked the node directory structure](https://github.com/gnolang/gno/pull/1944), and made [genesis.json usage explicit](https://github.com/gnolang/gno/pull/1972), paving the way to an easier multinode future. This effort enables us to easily orchestrate and configure Gno blockchain nodes, especially with the upcoming devnet / testnet launch.

- [Overhauled our monorepo GH actions](https://github.com/gnolang/gno/pull/2040), with them now avoiding double-work, and being much snappier. CI for PRs now runs much quicker and more stable, due to these optimizations.

- [Finished migration to Goreleaser](https://github.com/gnolang/gno/pull/2101). All of our important tools and binaries have a clear build and release schedule (for Docker, and beyond!), with us implementing nightly, dev and `master` releases on the monorepo.

- [Bumped the gnovm test coverage](https://github.com/gnolang/gno/pull/2143) from ~34% to ~67%. With upcoming changes to the GnoVM, we needed a good safety net in the form of a testing suite that will alert us if funny business is going on after a change.

- [Many, many bug fixes, small UX improvements and QoL changes](https://github.com/gnolang/gno/commits/master/?since=2024-05-01&until=2024-06-03&before=edb321f85beff98536a3a1a7111febaea0771a5e+35) that keep the lights on and systems running smoothly.

# Ecosystem Updates

## Onbloc

 Multinode/Validator Docs:  https://github.com/gnolang/gno/pull/2285
2. `test4` Validator Initiative Discussion: https://github.com/gnolang/hackerspace/issues/69
3. GnoSwap Portal Loop Migration Blocker: https://github.com/gnolang/gno/issues/2283
    - Liquidity Pool realm deployment failing on portal-loop
    - Attempts tried:
        - Changing package path / reducing realm size / setting gax-wanted to max / changing client env
        - Cloning gno repo every single recent commit to re-produce on local but never failed.
    - Update: Potential cause → Certain failing txs are causing corrupt cache files `cacheNodes`

- Gno Core
    - You can check our Merged, Awaiting Review, and TODO PRs in [our hackerspace. ](https://github.com/gnolang/hackerspace/issues/29)
- [Onbloc API Docs](https://onbloc-info.gitbook.io/onbloc-api-docs): Permissionless JSON-RPC methods to Gno.land's official networks (portal-loop, test3) for any individuals or teams building on Gno.land (feedback is welcome!)
- GnoSwap
    - Contract migration to portal-loop (Tartget: May 24th)
    - Applying `std.Emit` to contracts and APIs ([apply emit event in contract](https://github.com/gnoswap-labs/gnoswap/pull/217))
    - VWAP (Volume-Weighted Average Price) implementation for improved pricing ([gnoswap-labs/vwap](https://github.com/gnoswap-labs/vwap))
- Misc
    - Add links to [GnoStudio Connect in GnoScan](https://gnoscan.io/realms/details?path=gno.land/r/michelle/mymood)
    - ![image](https://hackmd.io/_uploads/HJrWausQC.png)
- Full update: [Onbloc Hackerspace](https://github.com/gnolang/hackerspace/issues/29#issuecomment-2124697977), [Onbloc Kanban](https://github.com/orgs/gnolang/projects/23)


## Teritori

- GnoVM

    - **Add stacktrace functionality and replace some uses of `Machine.String`**: This pull request is currently in review discussions with Morgan ([PR #2145](https://github.com/gnolang/gno/pull/2145)).
    - **Go2Gno loses type info**: This issue is Merged ([PR #2016](https://github.com/gnolang/gno/pull/2016)).
    - **Avoid instantiation of banker on pure package**: This change is awaiting review and merge ([PR #2248](https://github.com/gnolang/gno/pull/2248)).
    - **Missing length check in value declaration**: This pull request is also awaiting review and merge ([PR #2206](https://github.com/gnolang/gno/pull/2206)).
    - **Issue: File line not set on `ValueDeclr`, `ImportDecl`, `TypeDecl`**: The issue has been resolved and closed ([Issue #2220](https://github.com/gnolang/gno/issues/2220)).
    - **File line not set on `ValueDeclr`, `ImportDecl`, `TypeDecl`**: The fix has been successfully merged ([PR #2221](https://github.com/gnolang/gno/pull/2221)).

- Gno lint

    - **Printing all the errors from goparser**: This pull request has been successfully merged ([PR #2011](https://github.com/gnolang/gno/pull/2011)).
    - **Lint all files in folder before panicking**: This pull request is awaiting review and merge ([PR #2202](https://github.com/gnolang/gno/pull/2202)).

- DAO SDK (still waiting for review)
  PR: [#1925](https://github.com/gnolang/gno/pull/1925)
  
- **GnoVM**

    - **Cannot use struct as key of a map**: We resolved the issue where structs couldn't be used as keys in maps. This PR has been merged ([PR #2044](https://github.com/gnolang/gno/pull/2044)).
    - **Go2Gno loses type info**: This issue is still awaiting review and merge ([PR #2016](https://github.com/gnolang/gno/pull/2016)).
    - **Gno Issue with pointer**: We  proposed a solution ([Issue #2060](https://github.com/gnolang/gno/issues/2060)).
    - **Stacktrace functionality**: We added stacktrace functionality and replaced some uses of `Machine.String` ([PR #2145](https://github.com/gnolang/gno/pull/2145)).
    - **Recover not working correctly with runtime panics**: We created an issue to address this problem ([Issue #2146](https://github.com/gnolang/gno/issues/2146)).
    - **Panic when adding a package with subpaths**: We worked on this issue and waiting for review and merge  [PR #2155](https://github.com/gnolang/gno/pull/2155)).

- **Gno lint**

    - **Printing all the errors from goparser**: This improvement is waiting for review and merge ([PR #2011](https://github.com/gnolang/gno/pull/2011)).

- **DAO SDK**

    - **DAO SDK**: Waiting Review and merge: [PR #1925](https://github.com/gnolang/gno/pull/1925).

- **Project Manager**
Since we have already a lot in review, before opening a PR on the Gno repo, we're taking time to:
    - Polish the "private" [atomic PR](https://github.com/TERITORI/gno/pull/20)
    - Polish the UI
    - Set-up e2e testing with gnodev and a gno-js-client wallet, you can see a demo recording [here](https://github.com/TERITORI/gno/pull/20), the end-goal is to run e2e tests in CI

## Dragos

### ZenTasktic 
- Zentasktic User  (3rd grant milestone) implemented: https://github.com/irreverentsimplicity/zentasktic-user
    - updated docs for all 3 projects:
            - https://github.com/irreverentsimplicity/zentasktic-core/blob/main/README.md
            - https://github.com/irreverentsimplicity/zentasktic-project/blob/main/README.md
            - https://github.com/irreverentsimplicity/zentasktic-user/blob/main/README.md

- zentasktic (the package)
    - big overhaul, allowing for keeping the Assess-Decide-Do logic on the `zentasktic` package, but save the data locally, in the realm importing the package. PR here with some more explanation: https://github.com/irreverentsimplicity/zentasktic-core/pull/1

- zentasktic-project (the realm)
    - backend finished, repo here: https://github.com/irreverentsimplicity/zentasktic-project
    - question: test failing on RemoveWorkDuration with `panic: reflect: reflect.Value.SetString using value obtained using unexported field` on WorkDuration? AddWorkDuration and EditWorkDuration tests are passing...

### Flippando
* hackerville.xyz website updated for the upcoming airdrop: https://hackerville.xyz
* airdrop script in testing
* flippando NFT airdrop
    - copy added to the main flippando website, with airdrop mechanics and due date: https://gno.flippando.xyz/airdrop
    - minor updates to the website in preparation for this (airdrop mode: https://gno.flippando.xyz/playground)
    - spoiler, it will be on June 8th (2024, for conformity)

## Berty

- tx-indexer genesis [PR34](https://github.com/gnolang/tx-indexer/pull/34) (related to [PR1941](https://github.com/gnolang/gno/pull/1941)? ) / Jeff
  - blank screen bug fixed
- dSocial latest features / Iuri
  - reply to a post
  - view other user's posts
  - others
- UI conversation with Alexis - to plan soon

- dSocial demo app
  - Released on Test Flight and Google Play
  - To get an invitation, send your email to Iuri on Signal. Please say if you have an iPhone or Android phone.
  - Using custom indexer on the Berty production server
- Stress testing
  - Finalizing report
  - What is the PR to watch for resolving https://github.com/gnolang/gno/issues/1577 ?
    - Referenced PR https://github.com/gnolang/gno/issues/1576
  - Comment on CodeMagic as a possibility

## Var Meta

- issue: https://github.com/gnolang/gno/issues/2053
PR: https://github.com/gnolang/gno/pull/2108
Des: limit import path length
Status: Merged
- issue: https://github.com/gnolang/gno/issues/2192
PR: https://github.com/gnolang/gno/pull/2242
Des: Restric the maketx call function can only call to a realm
Status: Merged
- issue: https://github.com/gnolang/gno/issues/1998
PR: https://github.com/gnolang/gno/pull/2149
Des: This PR defines a GasUsed() func and a defaultInvokeCost in gas within std package. Simple feature let realm developer know the gas used at the time function is called.
Status: Wait for reviewing
- ISSUE: New issue about GnoPlayGround RUN and TEST func in different browsers (safari, chrome)
Link: https://github.com/gnolang/gno/issues/2270
Status: NO PR is made, waiting core team
- WIP: Sponsor TX
PR: https://github.com/gnolang/gno/pull/2209
Des: This PR aims to facilitate a transaction that should have been from A(signer) to B(Address/Realm) (A would pay the gas fee). Instead, A will delegate C(signer) to sign the transaction from A to B (C will pay the gas fee).
Status: Under reviewing
- PR: https://github.com/gnolang/gno/pull/2249
Issue: https://github.com/gnolang/gno/issues/2232
Des: To consolidate our returns on gnokey queries, I propose that we make them return JSON strings
Status: Waiting for #1776 to be merged
- PR: https://github.com/gnolang/gno/pull/2198
Issue: https://github.com/gnolang/gno/issues/2193
Des: Propose refactoring p/ownable from a struct with a specific implementation to a more minimal interface, allowing for custom ownership logic while retaining the current struct as the default implementation.
Status: Waiting for reviewing
- PR: https://github.com/gnolang/gno/pull/2225
Issue: https://github.com/gnolang/gno/issues/2193
Des: I propose integrating certain utility smart contracts from Ethereum into Gnoland. Now i'm working on defining the Bitmap, NonceManager and Queue packages, which can provide essential functionality for the Gnoland ecosystem.
Status: Waiting for reviewing
- PR: https://github.com/gnolang/gno/pull/2234
Issue: https://github.com/gnolang/gno/issues/2231
Des: We should either implement this, or remove the flag until the functionality is implemented. Same goes for the –prove flag.
Status: Merged

- PR: 
    - Support extensions like Metadata, RoyaltyInfo for GRC721  https://github.com/gnolang/gno/pull/1962 (Merged)
    - Deprecate & remove std.CurrentRealmPath() https://github.com/gnolang/gno/pull/2087 (Reviewing)
    - limit package path length https://github.com/gnolang/gno/pull/2108 (Aprroved)
    - Implement Bitmap package https://github.com/gnolang/gno/pull/2115 (Need review)
    - Implement Nonces package https://github.com/gnolang/gno/pull/2123 (Need review)
    - GasUsed() for GnoVM std https://github.com/gnolang/gno/pull/2149
- Issues:
    - Panic when getting keypair information https://github.com/gnolang/gno/issues/2133
    - Proposal: Integrate Sponsor Mechanism for Transaction Fees https://github.com/gnolang/gno/issues/2152

## Student Contributor Program

**Mustapha**
* Made a V0 Auction dapp ([PR#2265](https://github.com/gnolang/gno/pull/2265)) 

**Antonio**
- GNOWLEDGE
A [realm](https://github.com/iam-agf/Gnowledge) to simulate a Stackoverflow styled. Sharing to get some feedback. You can interact with it via https://github.com/iam-agf/Gnowledge-website .

**Antonio**
* https://github.com/iam-agf/Shiken 

# New Content

- [Key/Value Stores: How We Improved the Performance of Our tx-indexer by 10x](https://gno.land/r/gnoland/blog:p/kv-stores-indexer)
- [Introducing Gno Studio, the Premier Builder Suite for Gno.land](https://gno.land/r/gnoland/blog:p/gno-studio-intro)
- [Introducing our new Gno.land logo: the gnome](https://gno.land/r/gnoland/blog:p/the-gnome)
- [Gnomes Spotted in Belgrade, Serbia: Recap from the Engineering Retreat and Golang Serbia Meetup](https://gno.land/r/gnoland/blog:p/gnomes-in-serbia)
- [Test4 Explained](https://gno.land/r/gnoland/blog:p/test4-explained)

# Events and Meetups

## Past events

- GoLang Serbia Meetup / Belgrade / May 23. We used the Gno core team's retreat in Belgrade to connect with the local Go developers and possible contributors over the next few months to build the ecosystem. [Full recap](https://gno.land/r/gnoland/blog:p/gnomes-in-serbia).
- We're currently wrapping up **GopherCon EU** in Berlin, expect an update soon!

## Upcoming events

- **GopherCon US** / Chicago / July 7th - 10th
- **Nebular Summit** / Brussels / July 12th/13th

## Discord Developer Office Hours

Every week on Thursday at 2:30 pm CEST, we host office hours on [Discord](https://discord.com/invite/d24CT5b9cd?event=1252310282450112595) to answer questions, discuss updates, and catch up with the community. We'd love to see you there!