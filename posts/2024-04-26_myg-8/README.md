---
publication_date: 2024-04-26T00:00:00Z
slug: monthly-dev-8
tags: [gnoland, ecosystem, updates, gnovm, tm2]
authors: [Kouteki]
---

# The More You Gno 8: Portal Loop, Test4 and More

Welcome to the 8th installment of The More You Gno. Because we have a lot of great updates, we're experimenting with a different article format that's cleaner and more to the point.

If you want to dive deeper, we have the weekly engineering [updates](https://github.com/gnolang/meetings/issues/37) and video [recordings](https://www.youtube.com/watch?v=9ch7MhKNBmw&list=PL7nP7r1QiDktMCdw1ydQo2crM3y6Zk7E4) available. 

## Gno Core Updates

This edition is packed with great updates, but the most important ones are the launch of the Portal Loop testnet, and the announcement of the next milestone on the road to mainnet - Test4.

### Test4 Milestone is Live

The next major milestone on our road to mainnet is Test4. This testnet will have several important improvements:

- Multinode capability
- Rolling release of testnet binaries
- No resets

The initial validator set will be comprised of the most prolific contributors and All-in-Bits teams. We're aiming to launch Test4 in July, and with 32% of the milestone complete, our confidence in the launch date is high.

The entire milestone can be seen [here](https://github.com/gnolang/gno/milestone/4).

### Portal Loop is Live

Portal Loop is an always-up-to-date staging testnet that allows for using the latest version of Gno, Gno.land, and TM2. It can run the latest code from the master branch on the Gno monorepo, while preserving most/all the previous the transaction data. 

Find out more about Portal Loop in the [documentation](https://docs.gno.land/concepts/portal-loop/).

We've also enabled the [gno.land faucet](https://gno.land/faucet) to receive testnet funds for Portal Loop.

### Gno Playground now supports multiple networks

With Portal Loop live and Test4 coming, it made more and more sense to be able to use Gno Playground on other networks. That's why we've added support for not just Test3 and Portal Loop, but also any custom network you might want to spin up.

The network selection dropdown is in [Gno Playground](https://play.gno.land/) at the top right. You'll also find there an option to add your own custom network. 

### Changelog

We've also done a ton of other work to make everyone's lives easier.

- Docs are now searchable on [docs.gno.land](http://docs.gno.land/). No more manual searching!
- With the help of Onbloc, our tx-indexer [just went supersonic with GraphQL](https://github.com/gnolang/tx-indexer/pull/20). Users can now have fine grained filtered GraphQL queries, which are curiously quick. This is a quantum leap in data serving for the indexer.
- [Added support for `gnoland secrets`](https://github.com/gnolang/gno/pull/1593), part of our [larger effort](https://github.com/gnolang/gno/issues/1836) to overhaul the way Gno chains / nodes are initialized. Now users can directly manage, display and verify their node secrets, without having to rely on manual CLI magic.
- Finally merged in [support for a non-reflect JSON parser package](https://github.com/gnolang/gno/pull/1415). Rejoice, JSON support is now provided out of the box with an easy to use Gno package, thanks to the Onbloc team.
- Merged in support for [int256](https://github.com/gnolang/gno/pull/1848) and [uint256](https://github.com/gnolang/gno/pull/1778) packages. Since Gno does not support `big.Int` yet, these packages are a nice and much-needed replacement for large number logic.
- [Implemented shadowing rules](https://github.com/gnolang/gno/pull/1793) in the GnoVM. This VM fix prevents packages shadowing global-level identifiers.
- `gnodev` now [automatically reloads the gnoweb page](https://github.com/gnolang/gno/pull/1457) whenever you update the realm's code. The goal of gnodev is to make realm developers' life easier. This change further reaches that goal, creating a feedback loop for development similar to that experienced by front-end developers (ie `nodemon`).
- Plethora of quality of life fixes and improvements, including [docs updates](https://github.com/gnolang/gno/pull/1741), [command fixes](https://github.com/gnolang/gno/pull/1716) and [conformity to the standard go tooling suite](https://github.com/gnolang/gno/pull/1407). Updates like these keep the lights running and the UX smooth for end-users.
- [Merged in the `cford32` package](https://github.com/gnolang/gno/pull/1572) to the Gno package examples. Packages like these help the community and implementation partners in utilizing Gno effectively.
- Cleaned up and exposed an [endpoint for querying transaction results](https://github.com/gnolang/gno/pull/1546) directly from the node. Even though our transaction indexer (standalone tool) offers many bells and whistles, having the ability to fetch transaction results on the node was the final missing piece in keeping the TM2 RPC minimal, but functional.
- Gnodev now [supports account premines](https://github.com/gnolang/gno/pull/1938), soon to support transaction predeploys. It is also now installed out of the box with make install. This means that you can now specify an account balance to be premined when using gnodev, no clunky balance file changes required!
- Added support for [WS clients, updated JSON-RPC batch processing](https://github.com/gnolang/gno/pull/1498). Switching all of our tools, and implementation partner tools from the HTTP client to the WS client will be a massive boost in terms of network overhead.
- We've updated the `gnoland` (node) command [to include](https://github.com/gnolang/gno/pull/1954) the `genesis` management suite. As part of our bigger effort to overhaul the chain initialization and management flow, this was a necessary step in achieving a 2 step setup (`gnoland init` and `gnoland start`).
- Updated the deployment flows (releases) for the vast majority of our tools (faucet, indexer, supernova, main gno binaries) using go-releaser. Now the core team can tag specific tool / repo versions and have a prepared release, with all release binaries ready to go.
- The gnoland node [now supports telemetry](https://github.com/gnolang/gno/pull/1762), and exposes metrics out of the box. This allows the node to be tracked in time-series databases like Prometheus, and better record the health and performance of the node.

## Ecosystem Updates

### Onbloc

* Adena wallet [v1.10.0](https://github.com/onbloc/adena-wallet/releases/tag/v1.10.0) released 
* Adena [Security Evaluation Report](https://github.com/adr-sk/adena-extension/blob/main/audits/Adena%20Security%20Evaluation%20Feb%202024.pdf) published by the AiB security team 
* Published a [short video guide](https://www.youtube.com/watch?v=TfSzp1_MaOI) on how to broadcast and sign a transaction using the AirGap Account.

### Teritori

- Linked the GitHub account with on-chain address. The realm allows a bot to add a public key. The bot can sign the address and GitHub ID account. This functionality permits everyone to link their account on the realm seamlessly.
- Social feed, DAO SDK, and Moderation DAO deployed on portal loop: https://app.teritori.com/feed?network=gno-portal

### Dragos

- Flippando
    - Public faucet set up at https://faucet.flippando.xyz
    - Fixed and slightly improved the multi-chain dashboard at https://flippando.xyz
- Zentasktic, an on-chain project management application 
    - There is the Zentasktic core first complete implementation. A package containing an unopinionated implementation of the Assess-Decide-Do workflow. 
    - Docs and code: https://github.com/irreverentsimplicity/zentasktic-core

### Berty

- In order to avoid the misconception that it's an official social dapp, GnoSocial was renamed to [dSocial](https://github.com/gnolang/dsocial)
- Demo video of the customer journey, showcasing the updated social app UI https://www.loom.com/share/621f151459b040b0a2e6a22adf96b371
- Building out a series of Gno Native Kit tutorials ([Episodes one and two](https://www.youtube.com/watch?v=N1HLyQDHGQ0&list=PL7nP7r1QiDktseW7786wrp23ipBZFLPb6&index=2))

### Student Contributor Program

We've started a [program](https://github.com/gnolang/student-contributors-program) to work with student cohorts directly through their universities. The first university cohort is based in Roeun with four students currently particpating in this initiative. You can see some of the work in our hackerspace [repo](https://github.com/gnolang/hackerspace/issues/59). 

- Malek was cited for the 2 PRs ([ufmt multibyte](https://github.com/gnolang/gno/pull/1889) & and [todolist example V0](https://github.com/gnolang/gno/pull/1811)).
- Malek also finished V0 of GnoFundMe, a crowdfunding example.
- Theo made a [pull request](https://github.com/gnolang/memeland/pull/21) on issue#11 of the memeland repo. 

### New Content

We're regularly building out written and video tutorials and content. We recently released the '[How to build a realm using Gnodev](https://www.youtube.com/watch?v=Hp4aeRsPt3g)' video tutorial that starts with the basics of initializing a gno.mod file, creating a Gno realm, and implementing simple functions for effective realm visualization. 

## Events and Meetups

### Past events

#### Go to Gno Korea (March 23)

Together with Onbloc we hosted the Go to Gno workshop. The 8-hour session included a lecture, a builder challenge, and a networking event. 25+ developers from dev communities, Web3 startups, and the Cosmosnauts in Korea have attended the event for a hands-on experience of writing a simple app in Gno. Check out the [Onbloc's recap](https://medium.com/onbloc/go-to-gno-recap-intro-to-the-gno-stack-with-memeland-284a43d7f620) for more details and pictures.

#### Gno.land Tokyo meetup (April 11)

Intro to Gno session in Tokyo was held in front of 20+ developers and web3 enthusiasts. They got to hear about Gno.land, how it's developing as an ecosystem, our plans for the future, and how they could get involved. That's the TL;DR, full recap is on [our blog](https://gno.land/r/gnoland/blog:p/gno-tokyo).

### Upcoming events

We've got three major events coming up
- **GopherCon EU** / Berlin / June 17th - 20th
- **GopherCon US** / Chicago / July 7th - 10th
- **Nebular Summit** / Brussels / July 12th/13th

If you plan to attend any of these events, find us at our booth or on the hallway track, and let's talk!