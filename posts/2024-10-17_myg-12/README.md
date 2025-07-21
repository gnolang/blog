---
publication_date: 2024-10-17T00:00:00Z
slug: monthly-dev-12
tags: [gnoland, ecosystem, updates, gnovm]
authors: [Kouteki]
---

# The More You Gno 12: Gno Bounties

Welcome to the 12th edition of The More You Gno! This month's highlight is our bounty program; solve an interesting challenge in Gno, get paid.

## Gno Bounties

Here's how our bounty program works. We've [labelled issues in our repo](https://github.com/gnolang/gno/labels/bounty) that are eligible for the bounty program. 

1. Find a bounty you want to work on
2. Submit a draft PR
3. Ask questions and get feedback early

Bounties are sorted by t-shirt sizes, ranging from $500 to $32,000 USD.

Find out more [here](https://gno.land/contribute).

## Introducing Gnoverse

One of the gno.land tenets is censorship resistance. For that, we need greater decentralization. That's why we're opening [Gnoverse](https://github.com/gnoverse), a new GitHub organization dedicated to community-driven projects, as well as other minot projects not critical to the gno.land infrastructure.

Contributors will be able to experiment with Gno related ideas and projects, collaborating without interference and oversight from the core engineering team.

We've migrated a dozen repos from [Gnolang](https://github.com/gnolang) to [Gnoverse](https://github.com/gnoverse). Whether you're a developer, designer, or just passionate about Gno, your input is invaluable. Feel free to submit pull requests, suggest ideas, or simply join the conversation.

## Gno Core Updates

The engineering team got together in Turin, Italy, to work through the details of the upcoming main.gno.land launch. We are almost ready to talk about this publicly, but you can hear some spoilers on the [video snippets](https://x.com/_gnoland/status/1844779439160213861) we caught.

### Changelog

- [Fixed repo-level benchmark workflows](https://github.com/gnolang/gno/pull/2716), for easier performance regression monitoring. Viewable [here](https://gnolang.github.io/benchmarks/). This allows us to identify how individual PRs affect performance.
- vel benchmark workflows, for easier performance regression monitoring. Viewable here.
Impact: We can now track on individual PRs if there is a performance regression, and monitor long-term performance.
- [Valid type comparisons for bools](https://github.com/gnolang/gno/pull/2725).
- [GnoVM slice memory allocation order fix](https://github.com/gnolang/gno/pull/2781).
- [Value declaration loop fix](https://github.com/gnolang/gno/pull/2074).
- [Support `len` and `cap` on an array pointer](https://github.com/gnolang/gno/pull/2709) .

## Events and Meetups

### Past events

#### gno.land Contributor Tech Discussions

We've revamped the contributor calls to showcase the cool stuff being built on our platform, and have technical discussions on the challenges we face. The [1st video is out](https://www.youtube.com/watch?v=4YUOTt5bDJc), and new ones will be published every two weeks.

#### Go Meetup - Turin, Italy

During our engineering retreat in Turin, we took the opportunity to connect with the local Go community and hold a gno.land workshop. Our very own Morgan Bazalgette walked the attendees through the gno.land project, followed by a live coding session where Morgan built a simple messaging board. See the video [here](https://youtu.be/b3zRbVcJxyE).
