---
publication_date: 2024-07-22T00:00:00Z
slug: monthly-dev-10
tags: [gnoland, ecosystem, updates, gnovm, tm2, test4, gnostudio, connect]
authors: [Kouteki]
---

# The More You Gno 10: Test4 is Live!

This edition focuses on Test4, a major milestone towards our mainnet. Test4 is the first true multinode testnet featuring DAOs and on-chain governance, offering a preview of what’s to come. Much excitement, such wow.

We document everything in our weekly engineering updates and video recordings.

# Gno Core Updates

## Test4 is Live

With 7 validator nodes running and 3 more about to be added via on-chain governance, Test4 is the first multinode testnet that has the complexity and the feature set we want for the gno.land mainnet. Find out more about it in [this article](https://www.gno.land/r/gnoland/blog:p/test4-live).

## Changelog

- Gnoweb live
- RPC live
- TX indexer available
- Test4 faucet added to the [Faucet Hub](https://faucet.gno.land/)
- Merged in [Gno type check](https://github.com/gnolang/gno/pull/1426) support, resolving the long-standing issue with Gnolang's type checking on a VM level, making it more stable for development.
- Added support for [transpiling gno standard libraries](https://github.com/gnolang/gno/pull/1695), as part of a bigger effort to stabilize the GnoVM with native binding support (which was added a while back). The Gno transpiler now uses Gno's standard libraries instead of Go's. This also eliminates the need for things like `stdshim` and an std whitelist.
- [Continued to improve upon v1 of the GOVDAO implementation](https://github.com/gnolang/gno/pull/2379), with additional improvements coming later this week ahead of test4. We want to launch with a minimal govdao implementation for test4, which will be centralized in the beginning. We will use the govdao mechanism for managing on-chain validator sets.
- [Embraced JSON output](https://github.com/gnolang/gno/pull/2393) as standard for configuration and secrets fetching. DevOps engineers can rejoice; it's now super easy to read and parse node values.
- Published [v1 of the validator documentation](https://github.com/gnolang/gno/pull/2285) ahead of the test4 launch. Having easy to understand orchestration docs is critical to easily onboarding node operators and validators. We will continue to improve upon the documentation, and have more use-cases and examples for orchestration.
- [Improved the performance](https://github.com/gnolang/gno/pull/2140) of `for` loops and `if` statements. The performance almost doubled for these super-common Gno statements.
- [Migrated](https://github.com/gnolang/gno/pull/2424) the `libtm` (Tendermint consensus engine) implementation to the monorepo. You can check it out [here](https://github.com/gnolang/gno/tree/master/tm2/pkg/libtm). We plan to adopt this engine implementation in TM2, shortly after the test4 launch. The blog post is coming soon on the official Gno blog.

# Events and Meetups

## Past events

### GopherCon US 

We sponsored and attended GopherCon US - full recap [here](https://gno.land/r/gnoland/blog:p/discover-gno-gc24). We participated in the [Challenge series](https://www.gophercon.com/agenda/session/1281366), held a [workshop on building a decentralized app](https://www.youtube.com/watch?v=lwL2VyjaV-A), and had a lot of great conversations on the hallway track. We also set up a [raffle realm](https://gno.land/r/gc24/raffle) - Gophers we able to join the raffle using the Adena wallet, Gno Playground and Connect. You'll see the snippets of the atmosphere in [the promo video](https://x.com/_gnoland/status/1811438404800057560) we put together.

### Nebular Summit

We had a great time in Brussels at the Nebular Summit. Manfred was on the agenda with a lightning talk, and the core team held a workshop. Catch a part of the event atmosphere [in this video](https://x.com/_gnoland/status/1812867888501477470).

## Upcoming events

## Discord Developer Office Hours

Every Thursday at 2:30 pm CEST, we host office hours on [Discord](https://discord.com/invite/d24CT5b9cd?event=1252310282450112595). Join us to get your questions answered, discuss updates, and catch up with the community. We'd love to see you there!