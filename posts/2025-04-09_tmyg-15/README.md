---
publication_date: 2025-04-09T00:00:00Z
slug: monthly-dev-15
tags: [gnoland, ecosystem, updates, gnovm]
authors: [Kouteki]
---

# The More You Gno 15: Getting Ready for the Mainnet Beta

As we gear up for the launch of the Gno mainnet beta, this edition of _The More You Gno_ takes a look at the progress we’ve made, what’s currently in the works, and what’s coming next. From critical governance updates to improvements in storage and gas mechanisms, we’re laying the groundwork for a stable and feature-rich mainnet. Let’s dive in.

## Mainnet Beta Progress

There are several key components that we need to complete before we launch the mainnet beta.

Done:
* [Token lock & params keeper](https://github.com/gnolang/gno/pull/3176) allows us to lock token transfers, except for paying gas fees to add a package or call a contract. The restriction will eventually be unlocked through a GovDAO vote.
* [GovDAO v3](https://github.com/gnolang/gno/issues/3078): new iteration of the governance realms that introduces the multi-tiered membership.

In review:
* [Garbage collector](https://github.com/gnolang/gno/pull/3789): capable of synchronous garbage collection, greatly reducing memory allocation and costs.
* [Realm ownership spec](https://github.com/gnolang/gno/pull/4028): a critical piece that will dictate how realms interact with one another.
* [Gas storage fee mechanism](https://github.com/gnolang/gno/issues/3418): a system to manage storage fees for transactions by locking GNOT proportional to additional storage usage. This ensures fair accounting for storage costs and provides incentives for efficient storage use.
* [Balance checker and genesis distribution](https://github.com/gnolang/gno/pull/3899): a realm and interface enabling you to quickly check any account balance.
* [Drip faucet](https://github.com/gnolang/gno/pull/3808): since tokens initially can't be purchased or traded, the faucet will enable you to cover gas fees while using gno.land.
* [Gas fee distribution](https://github.com/gnolang/gno/pull/3956): the cornerstone of the transaction fee logic, routing fees to the collector addresses.

In development:
* [std refactor](https://github.com/gnolang/gno/issues/3874): the division into packages should help to reduce the scope of each package, helping us create modular packages.

## Test6 Preview

Test6 has been delayed because we wanted to launch it with both the token lock and GovDAO v3 components, so we could test them together and figure out how they could break. The new testnet will also introduce some new components, like the username registry and the valoper (validator operator) registry; prospective validators will be able to register a public profile before getting nominated to become a network validator.

We aim to launch Test6 next week, so stay tuned!

## Events and Meetups

### Introduction to gno.land @ Belgrade

If you're in the Belgrade area, join us for an [in-person workshop](https://lu.ma/gyesps56) on April 10, where we’ll explore Gno, and give you a solid understanding of gno.land and how it differs from other smart contract platforms.