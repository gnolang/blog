---
publication_date: 2024-09-02T00:00:00Z
slug: monthly-dev-11
tags: [gnoland, ecosystem, updates, gnovm, gnobro]
authors: [Kouteki]
---

# The More You Gno 11: Introducing `gnobro`

As we're gearing up toward the mainnet launch, we haven't forgotten to give some love to minor features that makes everyone's lives easier. This time the spotlight's on **gnobro**.

# Gno Core Updates

## `gnobro` is Live

What's `gnobro`, you say? Simply put, it's a terminal based realm browser you can use to explore realms and improve the development experience on 'gnodev'. [See it in action](https://github.com/gnolang/gno/pull/2608).

## Changelog

- We've updated the release CI and [fixed issues](https://github.com/gnolang/gno/pull/2686) with `go-releaser`. Now all of the tools in the monorepo are released properly, with accompanying artifacts. 
- Added support for smooth `.md` file rendering in `gnoweb`, ahead of our plans to work on gnoweb 2.0. This allows packages that have READMEs and other documentation to render easily in gnoweb.
- Slew of GnoVM fixes, increasing stability
    - [Type comparison fix](https://github.com/gnolang/gno/pull/1890)
    - [Cyclic references in type declarations](https://github.com/gnolang/gno/pull/2081)
    - [Handling non-call expression valuedecl values](https://github.com/gnolang/gno/pull/2647)
    - And more pending reviews!
- We've added back [coverage support (CodeCov) for txtar tests](https://github.com/gnolang/gno/pull/2377), which make up a majority of our integration testing suite. The txtar tests for the `gnovm` package added an additional 5% coverage. We are currently assessing other packages that suffer from bad txtar coverage.
- We've added [support for more robust stack traces for Gno-code panics](https://github.com/gnolang/gno/pull/2145), providing a much better UX for the developer. You no longer need to dig through a 5k line log output to figure out what panicked in your Gno code; you'll see the exception stack trace instead.
- [Variable config command help output is drastically improved](https://github.com/gnolang/gno/pull/2399). In the past you'd need to know exactly what the configuration looks like before modifying or viewing the values. Now these values are conveniently present in the command help output.
- Last, but not least: let's welcome our new R&D Go Engineer, [Antoine](https://github.com/aeddi). He'll help us scale core components for gno and beyond!

# Events and Meetups

## Past events

### BUIDL With Cosmos / Web3 Summit, Germany

We've had a couple of cool talks in Berlin: [An Introduction to gno.land](https://www.youtube.com/watch?v=hTGeG0z09NU) by Leon Hudak and [Building the Interchain of Ecosystems](https://youtu.be/nhpqaQxcIUY) by Tobias Schwartz.

### Web3 Kamp

Web3 Kamp is a 9-day intensive camp held anually in Serbia, focusing on getting students involved in Web3. Check out this [X thread](https://x.com/_gnoland/status/1828443842221080778) for the highligh of our involvement.