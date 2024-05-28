---
publication_date: 2024-05-30T00:00:00Z
slug: test4-explained
tags: [gnoland, test4, testnet, poc, proof-of-contribution, tm2]
authors: [Kouteki]
---

# Test4 Explained

We recently [announced](https://gno.land/r/gnoland/blog:p/monthly-dev-8) the Test4 milestone on our road to mainnet. It's a big step in terms of complexity, so let's delve into the nuts and bolts, and how Test4 fits into the bigger picture.

## Why Test4 is Significant

Our choice to go with Go stemmed from our desire to lower the barrier to entry for new developers. That meant that we needed an intuitive programming language with widespread usage; developers would be able to write smart contracts without being forced to learn another language. Go fits the role perfectly. 

The flip side is that we can't just port a VM and call it a day. We have to build most tools and services from scratch.

In that sense, Test4 is the technological leap from a single validator POC to a first iteration of an actual functioning decentralized blockchain network. 

It's also the first stable testnet since the project's inception over 2 years ago.

## Multinode capability

Like most blockchains, we're looking to boost resilience and security through a number of validator nodes. But since we're not forking a VM, we're building the whole thing from scratch; starting from the ability to emit events, to being able to subscribe to smart contracts, to making a change based on those events. All this is needed e.g. to add or remove a validator node in the Gno.land network, since the validator management is done completely on-chain.

The initial validator set will not be incentivized; it will be a small group of the most active contributors, as well as internal AiB teams, who will thoroughly test and break the network, and document all the things!

## The First Permanent Testnet

Test4 is the first serious iteration of an official Gno.land testnet. Our aim is not to have any breaking changes, meaning you can use it to develop your dApp, identify bugs, test features in an environment that's as close to the upcoming mainnet as possible. Being a testnet, you can use the [Faucet Hub](https://faucet.gno.land/) to obtain test tokens and use dApps without any real consequences.

Now, that isn't to say that Test4 will be the ultimate testnet. If we uncover major breaking issues, we'll deploy Test5 and so on.

We promise, however, to come up with a better naming convention going forward.

## Indexer and Transaction Explorer Included

A staple of a functional blockchain network is being able to search, track and verify transactions. This enables the developers to confirm their smart contracts are working as intended, as well as troubleshoot issues. Test4 comes with a transaction explorer, allowing developers to move off custom Test3 forks over to Test4 for convenience.

## Regular Release Schedule for Binaries

With breaking changes being deployed before the Test4 launch, we can now switch to the rolling release model. Updates will go live on a regular, most likely biweekly basis, enabling developers to test their dApps on the latest available version.

[Portal Loop](https://docs.gno.land/concepts/portal-loop/) will keep up with the binaries, and run on the latest version.

## Proof-of-Contribution Scaffolding

A key Gno.land concept is Proof of Contribution; the idea that the reward goes not to the one who pays to win either via mining rig or big stake, but to the one that invests time to contribute to the community.

Contributions will be managed on-chain via DAOs; E.g. a contributor with core contributions will be a part of a DAO that recognizes those contributions, and the rewards will be tied to the level the contributor has in that DAO.

Test4 will lay down the foundation for PoC. We'll be adding proposal support, as well as voting support that executes those proposals on-chain. The result will be a modified version of PoA that will be supported on a protocol level, and managed by the admin. This will make the future switch to PoC a much more seamless affair.

## Release Date

Our goal is to complete development by the end of July 2024, and launch in early August 2024. We've opened up the [Test4 milestone](https://github.com/gnolang/gno/milestone/4) and the Gno.core team's [project board](https://github.com/orgs/gnolang/projects/38/views/5), so you could track the progress. 

We're also getting ready to open up some of the issues to the community, so stay tuned!