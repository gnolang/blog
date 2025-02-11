---
publication_date: 2025-02-12T00:00:00Z
slug: monthly-dev-14
tags: [gnoland, ecosystem, updates, gnovm]
authors: [Kouteki]
---

# The More You Gno 14: Mainnet Beta Checkpoint

Our work on the gno.land mainnet beta is in the final stages, so let's recap the remaining work. We've also launched a DevOps hub, and we're getting ready to launch Test6. It's gonna be a busy quarter!

## Mainnet Beta Progress

There are several key components that we need to complete before we launch the mainnet beta:
* [Constitution](https://github.com/gnolang/genesis) is being written, and we're looking for feedback. If you want to be part of the history, open an issue or a PR, leave comments and present your case.
* [Realm ownership spec](https://github.com/gnolang/gno/issues/2743) is a critical piece that will dictate how realms interact with one another. ADD MORE DETAILS AFTER MONDAY
* [GovDAO v3](https://github.com/gnolang/gno/issues/3078) is a unique realm tasked with governing the chain; it consists of contributors organized into multiple tiers.
* [Token lock & params keeper](https://github.com/gnolang/gno/pull/3176) allows us to lock token transfers, except for paying gas fees to add a package or call a contract. The restriction will eventually be unlocked through a GovDAO vote.
* [Garbage collector](https://github.com/gnolang/gno/issues/266) is capable of synchronous garbage collection, greatly reducing memory allocation and costs.
* Tokenomics work will finalize the token allocation, including the airdrop distribution. Once the work is done, you will be able to check whether you're eligible for the airdrop.

## Introducing the gno.land DevOps Hub

[gnops.io](https://gnops.io/) is a dedicated resource for developers to build secure, resilient workflows and designed to support newcomers & experts alike in managing DevOps for gno.land.

It is a collection of effective essentials, explainers & best practices for DevOps in gno.land. Security, resilience & automation are key. Over time, it’ll expand into a repository of tools for monitoring, automation & security. Its goal is to accelerate blockchain development by improving collaboration & security. gnops.io helps optimize workflows, mitigate risks & support high availability. Whether you run a chain or validate, it’s your go-to DevOps resource.

Want to contribute? gnops.io is community-driven! Share automation scripts, CI/CD guides, security audits & more. Future updates will reward top contributions. Check out the [Gnops GitHub repo](https://github.com/gnoverse/gnops) for more information.

## Onbloc Spotlight

Onbloc is a Seoul-based company that's an active contributor on the gno.land project. They're building GnoSwap, Adena Wallet and GnoScan, as well as contributing to the core. That's why we filmed our last visit to the Onbloc office, and showcased these amazing people.

[Inside Onbloc - YouTube](https://www.youtube.com/watch?v=7umxHYDJ4Bk)

## Events and Meetups

### FOSDEM 2025

We traveled to Brussels for FOSDEM 2025, one of the world’s largest and most influential open-source gatherings. This year marked FOSDEM’s 25th anniversary, drawing over 5,000 developers, engineers, and enthusiasts from across the globe.

Attending FOSDEM was inspiring, a powerful reminder of what can be achieved when people collaborate in a free, transparent, and open way. The depth of knowledge shared, the engaging discussions, and the sheer diversity of topics, from robotics and simulations to the Go ecosystem, highlight the incredible breadth of open-source innovation, and what’s possible when people build together. It went beyond code, FOSDEM was a first-hand experience to witness the global community driven by a shared ethos.
Check out the recap blog post [here](https://gno.land/r/gnoland/blog:p/fosdem-2025).

### Epitech Career Day

Last year, we had a strong Epitech University SCP cohort that we onboarded to gno.land, and as a follow up were invited to their annual student career day. This week, we hosted a talk titled "[Building dApps in gno.land](https://github.com/gnolang/workshops/blob/main/presentations/2025-02-05--epitech--leon/building-dapps.pdf)" to an audience of 40 students.