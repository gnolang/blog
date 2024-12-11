---
publication_date: 2024-12-10T00:00:00Z
slug: road-to-mainnet
tags: [gnoland]
authors: [adrianakalpa, michelleellen, thehowl, salmad3, Kouteki, leohhhn]
---

# The Road to Mainnet & Beyond

### Gno the Difference

gno.land is evolving as a cutting-edge decentralized operating system, setting a
new standard for building and governing diverse applications. At its core, it 
provides an open-source smart contract platform designed to address systemic 
challenges in centralized systems—censorship, lack of user autonomy, misaligned
incentives, and non-transparent operations.

The true strength of gno.land lies in its broad applicability across domains. 
It provides builders with the means to create secure, scalable, and trustless 
applications that redefine the possibilities of decentralized systems.

gno.land has been advancing through a phased development process aimed at 
delivering critical functionality at each stage. The network architecture 
ensures that every milestone reinforces the foundation of a reliable and scalable
decentralized ecosystem that prioritizes transparency, forkability, and governance.

Let’s take a closer look at the progress, key milestones, and what’s in store 
for the future of gno.land.

## Milestones Overview

The priority is to establish a secure and stable network. Once operational and 
thoroughly tested, the infrastructure for rewards distribution and development 
will activate, accelerating the growth of Gno's ecosystem for both builders and 
users.

Here’s a look back at what we’ve accomplished so far.

### 2021: The Beginning

gno.land’s journey began in 2021, carried out by Jae Kwon, who laid the 
groundwork to develop Gno's most revolutionary component, the Gno Virtual Machine (GnoVM), and bootstrap a simple
Tendermint node capable of executing Gno code.

The GnoVM is a virtual machine capable 
of deterministically running a subset of Go code, where smart contracts are
composed together by simply “importing” them, much like working with any standard
Go package. The initial work was focused on creating the first iteration of the 
VM, aiming to implement a complete feature set of Gno while accounting for
various edge cases.

[By the end of 2021](https://github.com/gnolang/gno/tree/fa412d4bb38f6d6485d9b61784b34e5d5ca64554),
rudimentary tooling was available for experimenting with the VM, an extensive
suite of tests was in place, and an initial version of Gno’s state persistence 
was implemented. The first commits for the 
[`boards` realm](https://github.com/gnolang/gno/tree/master/examples/gno.land/r/demo/boards)
were introduced, and a functional chain executable was developed, enabling the 
addition of new packages and the ability to call functions within them.

### 2022: The First Testnets

The full scope of the project only became evident in 2022, following the 
implementation of the initial functionality. During this period,
the Gno project experienced a significant increase in contributors, 
resulting in a surge of new features, bug fixes, and developments
across multiple areas beyond just the VM. The Gno core team 
grew to three full-time engineers, with hiring underway to add more.

- Three testnets were launched (test1, test2 and test3), each with significant 
improvements in usability and ease of access for newcomers, and more example realms.
- The `boards` realm was improved, with the introduction of new features including
editing posts, moderation, and links to easily construct replies to posts using
`gnokey`. Additionally, thanks to the `users` realm, Gno users were able to register
a username which then showed up on the `boards` realm (among others).
- The Onbloc team began the development of the first gno.land wallet,
[Adena](https://www.adena.app/), marking the first steps to creating web dApps
using gno.land.
- The RPC layer got many improvements, including `qeval` and other features to 
inspect Gno code.
- Additional examples were added on creating GRC20 tokens, as well as string 
formatting utilities.
- The first [workshops](https://github.com/gnolang/workshops) and presentations
were posted on how to program in Gno.
- On the Gno side, the VM was consistently improved to fix more bugs and 
increase its safety for use on-chain.

[End-of-year status](https://github.com/gnolang/gno/tree/cabc1ad9a36319f68fed2bc888018bfcbff228a7)

### 2023: The Growth

In 2023, the Gno team grew to twelve engineers, many of whom continue to form the 
Gno core development team today. Around the same time, gno.land began seeing its
first real-world applications, putting its capabilities to the test.

Additionally, new developer tools were introduced to enhance the Gno development
experience. These include utilities such as `gno doc`, `gno repl`, `gno lint`, 
alongside `gnodev`, which is the foundational tool for building realms on gno.land
today.

- [GnoChess](https://github.com/gnoverse/gnochess) was a first experiment to 
program complex Gno smart contracts, capable of running the backend of a chess 
server. The Gopher community had the opportunity to experience GnoChess for the 
first time at GopherCon 2023, where gno.land highlighted Gno’s unique features 
and its potential to transform the decentralized application landscape.
- [Test3](https://github.com/gnolang/gno/releases/tag/chain%2Ftest3.0), launched
at the end of 2022, was continued in 2023, allowing builders to experiment with
running full nodes and building indexers. It gave everyone the possibility of 
having a persisting history of state.
- Many API decisions, like `std.PrevRealm()`, were discussed and implemented.
- The [License](https://github.com/gnolang/gno?tab=License-1-ov-file) for gno.land
was established as the *Gno Network Public License*, a fork of the GNU Afferro 
General Public License.
- [Berty](https://berty.tech/) and [Teritori](https://teritori.com/) became core 
ecosystem builders, creating integrated applications, communication tools, and 
social platforms.
- The first official documentation was released, allowing more users to be 
onboarded onto gno.land without having to learn directly from the source.
- The Beta version of [Gno Playground](https://play.gno.land/) launched, providing
a space to experiment Gno programs and realms directly from the web.

[End-of-year status](https://github.com/gnolang/gno/tree/c3a4993335f8881989eb49e4efea6a3c1310061d)

### 2024: The Refinement

2024 saw the Gno core team coming together, prioritizing the creation of a stable
chain to launch as soon as possible. Throughout the year, the team focused on 
delivering great stability on gno.land and being able to fix the most critical 
issues which stand in the way of launching the first stable network.

The launch of Test4 on July 15, 2024, was a major milestone in gno.land’s journey.
One of the main objectives of this testnet was to collect valuable knowledge and 
expertise from operations like running a full node, testing validator coordination,
deploying stable Gno dApps, creating and testing tools that require persisted 
state and transaction history.

During Test4, thousands of bug fixes and significant advancements were implemented,
including improvements to multi-node testnets and essential updates to `gnodev`.

**Here’s a breakdown of the most important key features that make Test4 unique:**

- **Multi-node Testnet:** This was the first testnet that was capable of running
on multiple nodes, making important steps towards decentralization
- **Permanent and stable Testnet:** Unlike previous versions, Test4 was a 
permanent testnet, giving early adopters and builders a consistent environment 
to build and test dApps while also providing them with the most up-to-date stable 
environment
- **GovDAO Proof-of-Concept:** Together with the launch of Test4, the first 
framework for decentralized governance was introduced, allowing selective 
contributors to begin participating in decision-making processes.

Test4 provided critical insights into running a multinode network and facilitating
on-chain governance. The result was a stable blockchain network with 7 validator 
nodes, as well as [5 governance proposals](https://test4.gno.land/r/gov/dao/v2),
3 of them executed successfully based on the two-thirds majority.

Later in 2024, [Test5](https://test5.gno.land/) was released. Main key elements 
that Test5 enabled include: expanding the current validator set, which now consists
of 17 validators, pushing on more GnoVM fixes, releasing GovDAO v2 and overall 
improving the quality-of-life for builders. View testnet’s current activity and 
deployed realms on [Gnoscan](https://gnoscan.io).

The Gno Core Team also greatly focused on creating a stable experience on Gno:

- Fixing many different bugs on the VM and making its behaviour more predictable 
and better documented.
- Improving “quality of life” tools that show the power of developing on gno.land,
like the [transaction indexer](https://github.com/gnolang/tx-indexer/) and `gnodev`
- Improving performance in critical paths and making behaviors of different APIs
consistent, and clearing out the standard libraries.
- Improving the behavior of integration tools, like `gnoclient` and `gnokey`, as 
well as kick-starting development of a mobile `gnokey` application.

**The road ahead starts with the preparation of the gno.land’s Beta Launch. 
Here is how the path forward looks like.**

## Beta Launch (UPCOMING)

We’re excited to announce that as 2024 comes to a close, we are gearing up for 
**Beta Launch**, anticipated in early 2025. Here’s what you can expect 
at genesis:

1. Contributors who have built gno.land will form the initial Governance DAO ("GovDAO"),
distributing membership into three tiers—T1 (highest), T2, and T3 (lowest)—based 
on the size and impact of their contributions as outlined in Jae Kwon's
[GovDAO draft](https://gist.github.com/jaekwon/918ad325c4c8f7fb5d6e022e33cb7eb3).
Membership is governed by age limits, contribution criteria, and a structured 
payment system tied to the GovDAOPayTreasury to ensure compensation directly aligns
with each contributor's level of contribution. **The GovDAO proposal remains an 
initial draft subject to revisions before final implementation.**
2. gno.land will intentionally maintain a modest transaction processing speed at
genesis; however, the network will enable early adopters to experience its
features and begin accruing rewards based on their contributions to the project
**While not immediately claimable, these rewards will become redeemable in future
phases to recognize early participants for their contributions**.
3. gno.land will operate under a Constitution that establishes clear guidelines
on the community's rights, liberties, obligations, governance, economics, and 
both current and future implementations of the gno.land chain. **The gno.land 
contributor community will participate and engage in this effort.**
4. The genesis token distribution of GNOT, gno.land’s native token, will primarily
allocate the majority of the supply to the Cosmos Hub (ATOM holders) community, 
rewarding their alignment with gno.land’s mission. The remaining GNOT allocations
will be used to support gno.land’s community pool, various Builders DAOs or 
builders communities, and the core team builders (through NewTendermint), which 
will play a crucial role in shaping the platform’s future growth and development.
**The GNOT token will be distributed at launch but will initially be non-transferable.**
5. While the GNOT token will not be transferable, users of the network will be 
able to use it to pay for gas fees. Token holders will be able to view their
balances, with expected support from Ledger, Adena, and,potentially, other 
community wallets. This will allow users to deploy and interact with realms on
gno.land. **Please note that the current chain design does not involve any staking.
Instead, the voting power is determined by a member’s reputation.**
6. Builders will have the opportunity to build and test dApps, thereby actively
contributing to the platform’s growth. **Although the GNOT token is not going to
be transferable, it will still be possible to mint, burn and transfer both 
*coins*** (a native kind of token, accessible through smart contracts using the 
banker module) **and tokens implemented in code**, i.e. *GRC20 tokens*, similar to 
their ERC counterparts.

Builders should be aware that during gno.land’s Beta network state, the core team
may, in exceptional cases, need to revert the chain state. Such instances might
include:

- **Identification of a GnoVM bug** that requires a rollback and replay of transactions.
- **Detection of a functional issue within a realm,** where the fix is well-defined,
and replaying transactions is required to maintain consistency.

The Beta Launch will equip contributors with essential tools to build and 
experiment on gno.land. We have published a 
[Launch Milestone](https://github.com/gnolang/gno/milestone/7) to track progress
and explore the specific initiatives we plan to address up until launch.

**Key components of the network at Genesis include:**

- **GnoVM** : For the Beta Launch, the GnoVM is anticipated to be in a relatively
stable and nearly feature-complete state. Bugs are an expected part of this 
phase; caution is advised to navigate potential issues.
- **Tendermint2**: Tendermint2 will mark a significant leap in consensus 
performance and scalability. Key improvements include faster block finality, 
more robust validator management, and greater resilience to node failures.
- **GovDAO Formation:** GovDAO, gno.land’s decentralized governance system, will 
be live at Genesis. Composed of the most active contributors, GovDAO will take 
on responsibility for governing gno.land on-chain, handling decisions like 
protocol upgrades, parameter changes, and other critical decisions to ensure 
the network's long-term sustainability and community alignment.
- **Boards:** Boards will operate as a single realm where users can create and 
participate in forums and message boards, subject to board-specific permissions.
Each board will be managed through its own administration DAO.


**Key tools for builders available at the Beta Launch:**

- **Core dev tools:** A fully-featured set of tools for local development of Gno 
dApps and libraries, including `gnokey`, a powerful CLI key management tool and 
client, `gnoweb`, a universal web frontend for gno.land, and `gnodev`, a local 
development solution stack offering a built-in gno.land node with a hot-reload 
feature for smart contracts.
- [**Gno Playground**](https://play.gno.land/): A browser-based tool to write,
test, run, deploy, and share Gno code, inspired by the Go playground. It offers
a simple and interactive way to experiment with Gno without the need for a local 
development environment.
- [**Gno Studio**](https://gno.studio/): A powerful unified development platform 
designed to drive dApp development on gno.land. This product suite will include 
a full-featured IDE, workspace management, production-grade capabilities, and 
integrated tools to support efficient development workflows. It will eventually 
feature an extensive marketplace for realm templates and ready-to-use solutions.
Today, the beta version of Studio features a minimal version of the IDE via Gno
Playground, along with Connect, an app designed for realm interaction.
- **Gno.me:** Envisioned to be a dynamic ecosystem of realms serving as an 
educational and community hub to empower builders through learning, collaboration,
and engagement. The initial release, which is estimated to be ready by the Beta 
Launch, will feature tutorials, a tooling section, and ecosystem updates, with
plans to evolve into a vibrant space for community interaction and a comprehensive
educational platform.
- [**docs.gno.land**](https://docs.gno.land/): The official gno.land documentation,
covering basic concepts, getting started tutorials, developer guides, references, and more.
- [**gno/examples**](https://github.com/gnolang/gno/tree/master/examples): 
Ready-to-use example smart contracts and libraries written in Gno.
- **Community Spaces:** Builders will have access to expanded tooling to build
and deploy applications in [gnoverse](https://github.com/gnoverse), a community-led
GitHub space, ensuring **gno.land** becomes a productive ground to encourage
future innovation.

## Mainnet Launch (UPCOMING)

While our Beta Launch will introduce the core features of the gno.land network,
the **Mainnet Launch** will focus on expanding **gno.land**’s interoperability 
with other chains. By providing a fully-featured, interoperable blockchain, 
**gno.land** will allow builders to integrate their products and services with 
the gno.land network, therefore scaling the ecosystem.

**Here are some key features of the Mainnet Launch:**

- **Interoperability with IBC and ICS:** **gno.land** will integrate Inter-Blockchain
Communication (IBC) and Interchain security (ICS), allowing seamless interaction 
between **gno.land** and other IBC and ICS compatible chains. During this 
milestone, modified versions of IBC and ICS will be integrated, versions that 
are fully compatible with existing Cosmos chains, called IBCx and ICSx.
- **Gno SDK:** Builders will gain access to the Gno SDK, making the entire process
of building and deploying dApps on the network much easier.
- **Launchpad for dApps:** The release of a specific Launchpad for dApps will 
allow builders to quickly deploy and scale their applications on gno.land.
- **Reward Distribution:** Contributors will be able to collect the rewards they
have accumulated from different contributions until this phase of the network.
- **Performance Enhancements:** The Mainnet Launch will also introduce significant
performance upgrades, including sharding and increased TPS, to support more 
complex applications.
- **Transfers enabled:** During this phase of the launch, GNOT tokens become 
transferable. This transition will mark a significant milestone, allowing users 
to engage in token transactions, interact with the network and its applications, 
and contribute to the broader economic activities within the gno.land ecosystem.
- **Boards**: A full social interaction dApp will be available. Users will continue
to create and participate in forums and message boards, with the added ability
to build their own decentralized social media platforms using a customizable 
framework. Forkability and DAO-driven governance will enable scalable, user-driven
communities.

## Call to Contribute

gno.land puts at the center of its identity the contributors that are helping to
create and shape the project. Contributing to gno.land’s development today means 
you have the opportunity to make a notable impact at such a critical stage in the
project’s roadmap.

For now, we are looking for the earliest adopters; those curious to explore a 
new way to build smart contracts and eager to be a part of the next generation of
blockchain programming. Join us and [become a contributor today](https://gno.land/contribute)!

## Conclusion

With more than three years in the making, gno.land is now closer than ever to
launching. The **Beta Launch,** while representing an important step in gno.land’s
journey, should be seen as only the beginning.

Future milestones are envisioned to bring improved performance, 
interoperability with other blockchain networks through IBCx and ICSx, the 
ability for builders to integrate their products and services with gno.land, 
onboarding tools for different user segments, enhanced support for builders and 
much more.

We will host a Q&A on [Discord](https://discord.com/invite/gnoland), Friday, 
December 13th at 5:30 pm CET, for the community to join and learn more about 
gno.land and upcoming milestones. Please add questions ahead of time to this 
[form](https://docs.google.com/forms/d/1DN8YneJ2-qvzsdnGsk_ZVA4CnBthUop3tgwXCodlF3M/edit) 
so we can make sure to address them.
