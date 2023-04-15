---
title: "Gno.land Community Game of Realms AMA #1 - Recap"
publication_date: 2022-02-03T15:44:00Z
slug: gor-ama1
tags: [game-of-realms,gnoland,proof-of-contribution,dao,governance]
authors: [manfred]
---

With Game of Realms officially in phase one, core dev Manfred Touron jumped on Discord to answer Gno.land community questions about the ongoing high-stakes competition. From starting and end dates to participation requirements and a description of tasks, look for your answer below. If you have further questions or want to join our community, come and find us on the []Gno.land Discord](https://discord.com/channels/957002220384182312/1065646963825066044). The core team will be hosting regular “office hours” sessions soon so you can discuss your ideas with them directly.

## Q. How are the tasks in the issues assigned?

We received questions about how the tasks in the Game of Realms issues are assigned. Should submissions contain the whole implementation? Is the following task "available** when the previous one is completed? How is the “sync” happening?

**A.** TL;DR:

Everything should go smoothly and we will be leaving room for negotiation if any review looks invalid. Once it has been established, the evaluation DAO will enforce how to submit a contribution. In the meantime, there are official communication challenges that we encourage participants to use. People are also free to work in stealth mode, with the risk of finishing too late or losing points for being bad at collaborating.

----

We expect the current issues to be done by multiple people, in multiple steps. But anyone can try to make everything in stealth mode and open a PR with everything. Let's discuss the cases we believe will happen:

### Case 1

We're in phase 1, people want to contribute but can't manage to do everything, so they will try to participate as much as they can. They will participate on the issue or in Discord by indicating their desire to participate, by sharing ideas, reviewing others' work, giving feedback, clarifying, or whatever makes sense.

The only thing is that we're fully remote. We don't know each other, so everyone needs to be good at communication. At the end of a big task, i.e. the Evaluation DAO is finished, the core team will take all the small contributions and identify contributors, and then suggest how to split the task prize. We'll propose the split and allow room for public negotiations.

### Case 2

We're in phase 2, and a small contribution is done by an individual. We just review it, and that's done.

### Case 3

We're in phase 2, and a contribution is big and requires small steps. Probably, the Evaluation DAO will ask individual participants to submit their contributions so they can allocate points for the individual contributions. But maybe the Evaluation DAO prefers to review big tasks as a whole, and then split the prize, as we'll do in phase 1. We don’t have clarity on this at this stage, as it will be up to the implementers of the Evaluation DAO to design the best system for that case.

## Q. Will there be a leaderboard and place where we can submit evidence for tasks?

**A.** Not yet. The leaderboard will come in phase 2. One of the critical parts of the Evaluation DAO will be to allow contributors to submit evidence for tasks. Votes and point allocations will also be transparent. This will make sense for future Proof-of-Contributions, too. We'll also develop a leaderboard to make it easier to follow the competition, but this will probably come after the Evaluation DAO is running.

## Q. What will the overall tasks consist of?

**A.** Here is a non-exhaustive list:

* Onboard more contributors ([create tutorials and documentation](https://github.com/gnolang/gno/issues/408)
* Improve the project and implement more things
* Bootstrap our genesis of contributors for the future mainnet
* Experiment with Proof of Contribution by having a simpler system: Evaluation DAO
* Identify the best participants to propose jobs
* Identify the best organizations to propose partnerships

## Q. At what point in the Game of Realms timeline/phase are we?

**A.** We are at the beginning of phase 1. We plan to create a website soon so you can keep track of the status and, as I mentioned, a leaderboard will come in phase 2.

## Q. What will be the contributions, how will points be calculated, and are there tasks for non-programmers?

**A.** During phase 1, the tasks are relatively well defined, please read this:

https://github.com/gnolang/gno/issues/390

There are more tasks for programmers, but multiple parts are for non-programmers too.

During phase 2, it's hard to be sure about anything yet. Game of Realms is a competition to experiment with Proof-of-Contribution, which will replace Proof-of-Stake on Gno.land. If things go the way we imagine, then consider that the stakeholders (contributors** will allocate points to contributions that make sense for the project. The contributors won't lose points, but by allocating points, they will dilute their own point stack.

We expect the Evaluation DAO to attribute points to whatever makes sense to make the project better. We'll have some task ideas for phase 2, including for non-programmers. You can likely consider that even if the core team doesn’t control the DAO, its suggestions will be approved by the Evaluation DAO because we deeply want the project to be a success.

## Q. What are the requirements to start participating?

**A.** There is no requirement to start participating. You’ll need to do some KYC at the end of the competition to receive a prize. Feel free to fill out the form linked in the Register section of the following issue:

https://github.com/gnolang/gno/issues/390

This will allow us to contact you about the competition through our newsletter and set up prize payment later. Use the comment section of the issues or discuss them on Discord if you plan to work on specific tasks, so we can see that you’re actively working on a topic. It may be better to work with others and share a prize instead of taking the risk of implementing everything in stealth mode and not being the first.

## Q. Is there a fixed period of time for the end?

**A.** No. Phase 1 will be finished when we consider that enough materials have been implemented to switch to phase 2. This will probably take between 1-3 months. The end date for phase 2 will be announced during phase 2, which will probably last between 2-3 months. This is when we’ll send the prize rewards. After Game of Realms, people will continue to earn contribution points by contributing to the project, which will give them memberships on the future mainnet.

## Q. Is it possible to install a local testnet to get a proper local development environment?

**A.** You can find the answer in this GitHub issue. Subscribe to the issue to get updates:

https://github.com/gnolang/gno/issues/478

There are multiple ways to interact with Gno:

* Using gnodev allows you to use the GnoVM, without a blockchain. This method is super fast and allows you to use development patterns like TDD, where you test your implementation multiple times per minute.
* Running a localnet, by running the gnoland command and then configuring our tools like gnokey to use localhost:36657
* Using the staging network hosted on https://staging.gno.land reset regularly and you can use the hardcoded test key or use the faucet
* Using the official testnets

If you prefer to run a full blockchain node instead of just playing with GnoVM, you should play with the gnoland binary. This video shows how to do this in practice:
https://www.youtube.com/watch?v=-BlnEXCs0eI

Below is a further resource that may also help you:

https://test2.gno.land/r/boards:testboard/5

## Q. Will there be a list of what needs to be tested? When will the tests start?

**A.** The best place to look is on GitHub here:

https://github.com/gnolang/gno/issues/390

During phase 1, there are 3 official focuses:

- Evaluation DAO
- Tutorials
- Governance Module

The core team will actively review this and decide what contribution deserves to get prizes.

During phase 2, we’ll use the Evaluation DAO developed during phase 1 to review old contributions, even contributions made before the competition, as well as ongoing contributions. Right now, we have an issue gathering interesting topics for phase 2 here, but any contribution can be reviewed by the DAO, including things that are not listed:

https://github.com/gnolang/gno/issues/357

The competition was just announced, but we’ll review contributions made in the past, too, so it starts from the first commit, ~1-2y ago.

_Do you have more questions for Manfred? Would you like to know more about Gno.land, Gnolang, Game of Realms, or ways to contribute to our growing ecosystem? Drop us a question on Discord and watch out for our next AMA on Tuesday 7 Feb at 4 pm UTC._
