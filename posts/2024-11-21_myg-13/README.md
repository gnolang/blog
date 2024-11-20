---
publication_date: 2024-11-21T00:00:00Z
slug: monthly-dev-13
tags: [gnoland, ecosystem, updates, gnovm]
authors: [Kouteki]
---

# The More You Gno 13: Test5, gnoverse and more

This edition of "The More You Gno" is packed with updates and highlights. Test5 testnet was launched; a new community space called gnoverse is now available. The mainnet launch milestone is up and the development is speeding up. It's an exciting time to be a part of the gno.land project.

## Test5 is live

On November 13 we launched Test5, a new iteration of the multinode testnet. The validator set has been expanded from 7 to 17 nodes, plus a number of non-validator nodes. It also boasts GovDAO v2 that expands the on-chain voting capabilities, and a number of fixes and quality of live improvements for the developers. 

Test5 is already available on Adena and Faucet Hub, so feel free to give it a try.

## Introducing gnoverse

Launched two weeks ago, [gnoverse](https://github.com/gnoverse) serves as a collaborative hub for experimental and innovative projects inspired by the gno.land ecosystem. It’s a space for incubating ideas and exploring the full potential of Gno.

This GitHub space maintains a strong connection with the [gnolang](https://github.com/gnolang) community, fostering collaboration and resource sharing among developers and enthusiasts. Our goal is to support and highlight projects that enhance the gno.land experience, and to be much more hands-off in the process.

We value community contributions! Whether you’re a developer, designer, or simply enthusiastic about Gno, your input matters. Join us by submitting pull requests, sharing ideas, or engaging in discussions. If you need a new repo for your project, we'd be happy to oblige.

## Mainnet launch

While we're getting an official announcement ready, take a sneak peek at the [gno.land mainnet launch milestone](https://github.com/gnolang/gno/milestone/7) on GitHub. The scope has been finalized for the most part, and our engineering teams are covering a lot of ground. We'll most likely launch at least one more testnet between now and the mainnet launch, so stay tuned.

### Changelog

All the cool stuff we've done, outlined as bullet points:
- [for loops maintain the same block on iteration, which is referenced in any closures generated within](https://github.com/gnolang/gno/issues/1135)
    - [feat(gnovm): handle loop variables](https://github.com/gnolang/gno/pull/2429)
- [Running gno test causes a panic when struct variables are redeclared in loops](https://github.com/gnolang/gno/issues/3013)
    - [fix(gnovm): fix issue with locally re-definition](https://github.com/gnolang/gno/pull/3014)
- [feedback needed: how are you using gnoland genesis ...?](https://github.com/gnolang/gno/issues/2824)
    - [chore: move gnoland genesis to contribs/gnogenesis](https://github.com/gnolang/gno/pull/3041)
- [feat: r/gov/dao v2](https://github.com/gnolang/gno/pull/2581)
- [Test4 re-release](https://github.com/gnolang/gno/issues/3060)
- [Test5 release](https://github.com/gnolang/gno/issues/3061)
    - [feat: add initial test5.gno.land deployment](https://github.com/gnolang/gno/pull/3092)
- [Proposal: Implementing Versioning Across Application Serialization Structures](https://github.com/gnolang/gno/issues/1838)
- [Running go vet on the project fails](https://github.com/gnolang/gno/issues/2954)
    - [fix(tm2): rename methods to avoid conflicts with (un)marshaler interfaces](https://github.com/gnolang/gno/pull/3000)
- [Fix println DoS vector](https://github.com/gnolang/gno/issues/3075)
    - [fix(gnovm): don't print to stdout by default](https://github.com/gnolang/gno/pull/3076)
- [Add r/sys/params + params genesis support](https://github.com/gnolang/gno/pull/3003)
- [Support metadata for genesis txs](https://github.com/gnolang/gno/pull/2941)
- [Setup "nice" stale PRs' bot](https://github.com/gnolang/gno/issues/1445)
    - [ci: add a stale bot for PRs](https://github.com/gnolang/gno/pull/2804)
- [feat: add contribs/gnomigrate](https://github.com/gnolang/gno/pull/3063)
- [chore: update tx-archive in portal loop](https://github.com/gnolang/gno/pull/3064)

## Events and Meetups

[4th episode](https://youtu.be/hRZ7iU4bovM?si=2tOUeuIyiSWxELGv) of the Contributor Technical Discussions is out! Watch the rest [here](https://www.youtube.com/playlist?list=PL7nP7r1QiDktMCdw1ydQo2crM3y6Zk7E4).