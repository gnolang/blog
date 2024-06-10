---
publication_date: 2024-06-13T00:00:00Z
slug: reaching-consensus
tags: [gnoland, test4, consensus, tm2, libtm]
authors: [zivkovicmilos, petar-dambovaliev]
---

# Reaching Consensus: Developing Fault-tolerant SMTs using Golang

Every once in a while, an effort comes along for the Gno core engineering team that demands immense focus, adaptability and coordination, in order to successfully improve a critical part of the [gno.land](http://gno.land/) blockchain system. This article is meant to cover one such critical module - the Tendermint consensus engine.

Over the course of the next following sections, you’ll discover on a practical level what the Tendermint consensus engine is all about, how it works, and how we ended up implementing it for the purposes of [Gno.land](http://gno.land/), and the wider blockchain and open-source community.

If you want to jump straight to the source code, you can find the GitHub repository [here](https://github.com/gnolang/libtm).

It is worth noting that we will  primarily be discussing the use of  the Tendermint consensus engine from an angle of a distributed system, that doesn’t necessarily have anything to do with blockchains. 

## What are consensus engines, really?

Most blockchain systems that are distributed use some *consensus mechanism*. In essence, this is a process in which remote machines (blockchain nodes) agree that a given value is valid. The nature of their agreement (the process) and the value on which they agree on are completely arbitrary, and are up to the specific blockchain system to decide on and use.

Some of these *consensus algorithms* have specific properties that make them more resilient than others — consensus algorithms can be crash tolerant, meaning the process of reaching an agreement on a value tolerates nodes being *faulty* by crashing, in which case they stop while the rest of the nodes continue operating. An example of this consensus algorithm is [Raft](https://raft.github.io/). On the other hand, a more resilient consensus algorithm is also *byzantine-tolerant*, in which case the nodes participating in the consensus process can not only crash, but be malicious — meaning an arbitrary range of faulty behaviors is tolerated. Examples of byzantine-tolerant consensus algorithms include [PBFT](https://pmg.csail.mit.edu/papers/osdi99.pdf), [IBFT](https://arxiv.org/abs/2002.03613), [HotStuff](https://arxiv.org/abs/1803.05069), and the one we will focus on in this article — [Tendermint](https://arxiv.org/abs/1807.04938).

For the sake of clarity, it is worth noting an important distinction very early on, before we dive into the nitty-gritty. In this article we make a clear distinction between consensus protocols, the validator set modification algorithms, and consensus engines, the actual state machine changes required for reaching consensus. 

In a distributed consensus-run system, the participants are usually called *nodes*. In the blockchain context, they are called *validators*, since they validate the proposed chain state changes through the consensus process. The process for which these actors *enter and exit* the validator set (node consensus set) is called a *consensus protocol*. Most popular examples of consensus protocols include:
* **Proof of Authority (PoA)** - a permissioned system in which consensus participants (validators) are selectively chosen / voted in, by the existing participant set.
* **Proof of Stake (PoS)** - a permission-less system in which anyone can stake something (usually the network’s native currency) to gain entry into the validator set (become consensus participants).

There are many different flavors for the system in which the validator set is chosen — for the sake of simplicity we group this process under the *consensus protocol* term.

[![valset](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/valset.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/valset.png)

We mentioned earlier that consensus is a process in which nodes agree upon a specific value. We denote this actual protocol in which they agree upon a value as the *consensus engine*. 

The consensus engine represents the internal, most critical process of the consensus algorithm - the state machine. It gives answers to questions like - who proposes the consensus value (in leader-based systems), how is validation done, but most importantly — what are the steps taken by validators to reach consensus?

[![machine](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/machine.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/machine.png)

### Terminology
Consensus algorithms like Tendermint work on a round-based model, meaning consensus is attempted to be reached on a given value within a round, for a given height. If consensus is not reached on `round X` of `height Y`, then a new consensus round is started for `round X+1`, `height Y`. This round changing occurs until consensus is reached for the given height (in the context of blockchains, the height is the block number). We will discuss later on what the conditions are for increasing the consensus round number.

We call this tuple denoting the current consensus height and round `(h, r)` a ***view***.

For leader-based consensus algorithms like Tendermint, one node (validator) called the ***proposer*** (who is deterministically selected by all participating nodes) needs to propose a value for the current consensus view. This ***proposal*** value can in theory be anything, but in the context of blockchain systems, this is usually a block containing transactions (state transitions).

Specific to Tendermint, in order to reach consensus, there needs to be an agreement between at least `2F+1` validator nodes (where `F` denotes the maximum number of faulty nodes that are tolerated in the system). In the broader Tendermint ecosystem, validators are not equal, since the consensus protocol mostly used with Tendermint is Proof of Stake, and the validator’s power is proportional to their stake. In this case, `F` denotes the maximum voting power of faulty nodes that are tolerated in the system. In practice, `2F+1` is 2/3+ (supermajority) of the validator set’s voting power.

## How does Tendermint work?

The term “[Tendermint](https://docs.tendermint.com/v0.34/introduction/what-is-tendermint.html)” refers to a much larger scope of functionality than what we discuss in this article.  By “Tendermint” here, we solely refer to the Tendermint consensus engine (state machine).

A good graphic that shows an overview of the Tendermint state machine can be found on the official [Tendermint Core documentation](https://docs.tendermint.com/v0.34/introduction/what-is-tendermint.html#consensus-overview):

[![polka](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/polka.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/polka.png)

However, we want to tackle this question of Tendermint’s inner-workings by weaving in the theory from the paper titled *The latest gossip on BFT consensus*, and applying it in practice.

On a surface-level, the Tendermint consensus engine has 3 states:
- **Propose** - the *proposer* for the current view (height, round) sends out a *proposal* to all other validator nodes. Non-proposers wait for the proposal and verify it
- **Prevote** - the validator nodes wait for 2F+1 `PREVOTE` messages, potentially locking on a proposal
- **Precommit** - the validator nodes wait for 2F+1 `PRECOMMIT` messages, potentially saving the block to the block store, and starting a new consensus round (with a view `(h+1, 0)`)

These states are not as straightforward as they’re made out to be — there is an added complication of *state timers* which guarantee protocol liveness, and other asynchronous situations like consensus round jumps within the same height. We will discuss these mechanisms later on, as they are critical to making sure Tendermint is ticking correctly.

### I. Propose

The *propose* state is the initial consensus state for the Tendermint consensus engine. This is the only state within the state machine that has different logic flows for proposers and non-proposers.

In the *propose* state, only one node is the proposer for the given view (height, round), and the rest are simply supposed to wait for the proposal.

The proposer for the view has 2 choices when proposing a value:
1. craft an entirely new proposal (block). The paper denotes this process with `getValue()`
2. utilize an old (locked) proposal, that was proposed in an earlier round (same height, earlier round), but whose consensus was never reached (for whatever reason)

We will discuss “proposal locking” in the `Prevote` section, for now consider that this is a special case where the proposer does not actually construct anything new, but simply relays a proposal that was sent in the past (earlier round), but not agreed upon fully (committed).

[![paper-1](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-1.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-1.png)
*Source: The latest gossip on BFT consensus, page 6*

As outlined in *Algorithm 1*, the initialization step takes place before the beginning of the actual consensus process (state machine). `StartRound` is meant to kick off the Tendermint state machine.

Upon figuring out the proposal value, the proposer broadcasts it to the rest of the nodes. The specifics of network gossip implementations are not discussed in the actual paper / consensus protocol.

For non-proposers of the view, they simply initiate the `propose` state timer, that’s meant to advance the state machine further if no valid proposal comes in from the proposer of the view. 

[![paper-2](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-2.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-2.png)
*Source: The latest gossip on BFT consensus, page 6*

Once non-proposers receive a proposal from the proposer of the current view, they validate it using an external function, denoted here as valid(...). The proposal is accepted in 2 situations:
* if it is a fresh proposal (not a locked one), and it is valid according to the standards of a verifying method (`valid(...)`) .
* if it is an older proposal (locked value),  it is valid according to the standards of a verifying method (`valid(...)`) *and* the proposal matches the validator’s in-memory locked proposal.

In case of a successful proposal verification, the node broadcasts a `PREVOTE` message with a valid ID of the proposal. The identification method of the proposal is not defined by the Tendermint protocol itself, but commonly in practice this value is the hash of the actual block being proposed (in blockchain systems).

In case the proposal verification fails (proposal is invalid), the node broadcasts a `PREVOTE` message, with a `nil` value for the proposal ID, indicating that this step in the consensus process failed.

The purpose of broadcasting a message with a `nil` value is to ensure the consensus engine’s liveness in case of a faulty event. Note that Tendermint does not have a specific round-change state like PBFT, or IBFT, so the only way to move the protocol along is *to go through all consensus states, and in order to succeed or fail in the end*.

In either case, after the node broadcasts a `PREVOTE` message, it transitions into the *prevote* state.

#### Propose timer

Non-proposers for the view start a `propose` state timer.

[![paper-3](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-3.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-3.png)
*Source: The latest gossip on BFT consensus, page 6*

This timer is meant to go off after a fixed interval of time, in order to prevent the node from being stuck in a state where it’s still waiting on a proposal.

When the timer asynchronously ticks off, the node broadcasts the `PREVOTE` message with a `nil` ID value, indicating failure to receive a valid proposal. After the broadcast, it moves to the *propose* state.

### II. Prevote

The *prevote* state is the second state of the Tendermint consensus engine, and it’s the common state for all consensus participants (proposers and non-proposers).

Nodes end up in this state in the following scenarios:
- the proposer for the view sent a valid proposal
- the proposer for the view didn’t send a valid proposal
- the proposer for the view didn’t send a valid proposal in time (timeout triggered)

The *prevote* state can be seen as an “alignment” state between nodes because this is the first state in which consensus participants need to reach a supermajority on whether the current consensus round is proceeding as it should, or if there is some funny business happening.

To successfully move over to the *precommit* state, validators need to receive `2F+1` `PREVOTE` messages with a valid proposal ID. In other words, the supermajority of the validator set’s voting power needs to send out `PREVOTE` messages with an attached proposal ID (of the proposal accepted in the *propose* state!).

[![paper-4](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-4.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-4.png)
*Source: The latest gossip on BFT consensus, page 6*

As soon as these conditions are met, the participants broadcast a `PRECOMMIT` message, with the attached proposal ID (same one as with the `PREVOTE` messages, and derived from the proposal accepted in the *propose* state). After broadcasting the `PRECOMMIT` message, participants move over to the *precommit* state.

#### Block locking

An important characteristic for the *prevote* state is the block locking mechanism that is triggered as soon as conditions for a valid transition to *precommit* are met.

Before broadcasting a valid `PRECOMMIT` message, the validators “save” (in-memory) the proposal value in which the supermajority of the validator set’s voting power agreed upon, and the round in which it happened (denoted as `lockedValue` and `lockedRound`, respectively).

The purpose of this block locking dance is to be able to propose / verify the same proposal if the current consensus round fails in the following steps (for example, not enough valid `PRECOMMIT` messages received). The idea behind this mechanism is that a proposal agreed upon by a supermajority of the validator set is most probably a valid one, and should be considered for future rounds in case the original consensus round in which the proposal was brought up did not complete successfully.

If we go back through the conditions for the initial Tendermint engine state transitions, we will see that there is a strong requirement for “locked” values.

[![paper-5](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-5.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-5.png)
*Source: The latest gossip on BFT consensus, page 6*

These checks essentially enforce that if there was a proposal in some previous round, that reached a supermajority, it needs to be *proposed again* in a future round.

#### Jump to Precommit

There is a situation in which the validators move over to the *precommit* state, without receiving a supermajority of valid `PREVOTE` messages (with a valid proposal ID):

[![paper-6](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-6.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-6.png)
*Source: The latest gossip on BFT consensus, page 6*

In case a validator receives `2F+1` `PREVOTE` messages with a `nil` proposal ID, they simply broadcast a `PRECOMMIT` message with a `nil` ID value (indicating failure), and move over to the *precommit* state.

This `nil` propagation mechanism is part of the liveness property discussed in the *propose* section, which requires a validator to go through all Tendermint states before starting a new consensus round.

#### Prevote timer

All validators start a *prevote* state timer the moment they receive a supermajority (`2F+1`) of *any* `PREVOTE` message (either with a valid or with a `nil` proposal ID).

[![paper-7](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-7.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-7.png)
*Source: The latest gossip on BFT consensus, page 6*

When the timer asynchronously ticks off, the node broadcasts the `PRECOMMIT` message with a `nil` ID value, indicating failure to reach a supermajority on a proposal. After the broadcast, it moves to the *precommit* state.

[![paper-8](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-8.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-8.png)
*Source: The latest gossip on BFT consensus, page 6*

### III. Precommit

The *precommit* state is the final state in the Tendermint consensus engine, common for all validators.

Much like the *prevote* state, *precommit* consists of the validators waiting for a supermajority of messages of a specific type — in this case the `PRECOMMIT` message, with a valid proposal ID.

[![paper-9](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-9.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-9.png)
*Source: The latest gossip on BFT consensus, page 6*

After the validators exchange `2F+1` valid `PRECOMMIT` messages (with a valid proposal ID), they commit the proposal to their local storage, increase the height and move over to begin the consensus process anew from `(height + 1, round = 0)`.
Since the consensus process is finalized successfully for the current view, the in-memory values for locked values are cleared, along with the complete message log.

Of course, the proposal ID contained within the `PRECOMMIT` messages needs to match the proposal ID in the previously received supermajority of `PREVOTE` messages, and be derived from the proposal in the *propose* state.

#### Precommit timer

Much like with the *prevote* state, the moment validators receive a supermajority (`2F+1`) of *any* `PRECOMMIT` message (either with a valid or with a `nil` proposal ID) they initiate a *precommit* state timer.

[![paper-10](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-10.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-10.png)
*Source: The latest gossip on BFT consensus, page 6*

When the timer asynchronously ticks off, the node begins the consensus process anew with an increased round number (goes back to the *propose* state). Since the *precommit* state is the last one in the chain, the only next step left for the engine is to go back to the beginning.

[![paper-11](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-11.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-11.png)
*Source: The latest gossip on BFT consensus, page 6*

### IV. Round jumps

There is a special scenario within the Tendermint consensus engine that bypasses the notion of going through all consensus states in order to start a new consensus round.

[![paper-12](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/paper-12.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/paper-12.png)
*Source: The latest gossip on BFT consensus, page 6*

In this scenario, if at any point in time the validator receives `F+1` (called a *faulty majority*) of *any* message type, for a **future** round (round higher than in the current view), the validator performs a round jump to the given round by starting the consensus process anew.

### V. Timers

In the Tendermint consensus process, timers play an important role in making sure the protocol stays “alive”. They move things along if, for whatever reason, the flow stops.

By their nature, they are completely asynchronous, and can go off at point in time once they are set by the protocol.

Tendermint does not enforce specific initial values for each state timer, but provides suggestions on making sure the timers serve their purpose:
`timeoutX(r) = initTimeoutX + r * timeoutDelta`

By following this formula, it is ensured that each round state timer is:
- influenced by the current view round (higher rounds == longer timers)
- influenced by a constant `delta` increment (can be different for each state)

In the Gno blockchain, these values are defined as the following:
```toml
timeout_propose = "3s"
timeout_propose_delta = "500ms"  
timeout_prevote = "1s"
timeout_prevote_delta = "500ms"
timeout_precommit = "1s"
timeout_precommit_delta = "500ms"
```

### VI. Tying it all together

Understanding the Tendermint consensus engine involves grasping its three core states: *propose*, *prevote*, and *precommit*, and the transitions that ensure the system's liveness and fault tolerance. 

The *propose* state is initiated by a designated proposer who broadcasts a proposal to all validators. This proposal can either be a new value or an old one from a previous round that did not reach consensus. Validators verify the proposal, and then move to the *prevote* state, where they broadcast their votes on the proposal, requiring a supermajority (`2F+1`) to proceed. If the necessary votes are not received within a certain timeframe, a timeout mechanism ensures the process continues to the next state or round. If a supermajority of valid `PREVOTE` messages is received, validators lock the proposal value and transition to the *precommit* state. This locking mechanism ensures that if consensus is not reached in the current round, the previously agreed-upon value can be proposed again, enhancing the algorithm’s robustness. During the *precommit* state, validators wait for another supermajority of `PRECOMMIT` messages. Upon receiving these, they commit the proposal locally, increment the height, and start a new consensus round.

The combination of these states and transitions, governed by timers and supermajority rules, creates a resilient and efficient consensus process. The Tendermint consensus engine's design allows it to handle both crash faults and Byzantine faults, making it suitable for various distributed systems, including blockchains.

## Taming asynchronous beasts with Golang

When we started working on [libtm](https://github.com/gnolang/libtm), we had a few goals in mind:
- create a consensus engine library that is easy to use, safe and performant, pluggable into systems we already have (the Gno blockchain, based on a fork of Tendermint Core)
- implement abstractions that allow the library to be also used outside the blockchain context, while also being modular enough to be used by existing blockchain projects
- maintain the principles of open-source software

We went about implementing #2 in a seemingly odd way:
- We did not want the library to handle networking, or enforce a networking protocol. We believe that the underlying transport layer that moves around messages is irrelevant, as long as it provides the properties of *eventual delivery*.
- We did not want the library to know how to sign, and verify signatures of arbitrary data. Signing protocols are something very use-case specific, and should not be enforced by the consensus engine. We do, however, provide the *data* that needs to be signed in the first place.
- We did not want the library to do validator set management, or to even have a notion of a “validator” and its properties. Validator sets have vastly different management techniques, a complication we wanted to avoid in the engine.

This might sound shocking, given how existing implementations of the Tendermint consensus engine operate today in modern blockchains, but it actually greatly reduces the complexity and allows for applications of the consensus library outside the blockchain context.

### Everything starts with the messaging system

Looking at how the Tendermint consensus state machine is structured, we can deduce that the main drivers of any kind of activity are — messages.

A message exchanged between validators, of any type, is the driving force behind the Tendermint consensus algorithm. It’s used to transfer information, but also to act as a notification system that preserves certain consensus properties like liveness.

Given the asynchronous nature of validator messages, we realized our messaging system would need to support the following properties:
- have buckets for each message type, to make messages easily fetch-able
- have a mechanism that avoids duplicates
- have a subscription mechanism we can tap into from any part of the codebase

We came to this conclusion after realizing our state implementations (discussed later) needed a more sophisticated way of fetching and filtering incoming messages, that didn’t involve constant read requests from the message log. These operations are very common, and they need to be performant in order to not delay the consensus process.

Luckily, we ended up utilizing the full power of Go generics for this use-case:
```go
// msgType is the combined message type interface,
// for easy reference and type safety
type msgType interface {
	types.ProposalMessage | types.PrevoteMessage | types.PrecommitMessage
}

type (
	// collection are the actual received messages.
	// Maps a unique identifier -> their message (of a specific type) to avoid duplicates.
	// Identifiers are derived from <sender ID, height, round>.
	// Each validator in the consensus needs to send at most 1 message of every type
	// (minus the PROPOSAL, which is only sent by the proposer),
	// so the message system needs to keep track of only 1 message per type, per validator, per view
	collection[T msgType] map[string]*T
)

// Collector is a single message type collector
type Collector[T msgType] struct {
	collection    collection[T]    // the message storage
	subscriptions subscriptions[T] // the active message subscriptions

	collectionMux    sync.RWMutex
	subscriptionsMux sync.RWMutex
}
```

The `libtm` library keeps a `Collector` for each message type, and that acts as a message log. Additionally, the message collector also keeps track of basic message subscriptions. These subscriptions (as we’ll discuss later), are for simple notifications that alert the subscriber a message appeared in the message log.

### Not all states are created equal

Once we’ve established the clear needs of our messaging system, implementing the actual state transitions around it was trivial.

Tendermint consensus states all share the following structure:
```go
// runX runs the X Tendermint consensus engine state
func (t *Tendermint) runX(ctx context.Context) {
	var (
		round              = t.state.getRound()

		expiredCh                   = make(chan struct{}, 1)
		timeoutCtx, cancelTimeoutFn = context.WithCancel(ctx)
		timeoutPrevote              = t.timeouts[prevote].CalculateTimeout(round)
	)

	// Defer the timeout timer cancellation (if running)
	defer cancelTimeoutFn()

	// Subscribe to all messages of a specific type
	// (=current height; unique; >= current round)
	ch, unsubscribeFn := t.store.subscribeToX()
	defer unsubscribeFn()

	for {
		select {
		case <-ctx.Done():
			// Outer consensus context cancelled
			return
		case <-expiredCh:
			// State timer triggered,
			// execute stateX teardown before returning

			// ...
			
			return
		case getMessagesFn := <-ch:
			// New valid message appeared in message log, parse it
			// ...

			// Check if conditions are satisfied for starting the state timer
			// ...

			// Check if conditions are satisfied for a state transition
			// ...
		}
	}
}
```

The pattern is clear:
- subscribe to messages of a specific type that are expected to appear in the current state
- verify those messages as they come in, and check if conditions are met for a state transition
- monitor the outer context, and the asynchronous timer context for stopping

Keeping this in mind, we can develop the main run loop:
```go
// runStates runs the consensus states, depending on the current step
func (t *Tendermint) runStates(ctx context.Context) *FinalizedProposal {
	for {
		currentStep := t.state.step.get()

		select {
		case <-ctx.Done():
			return nil
		default:
			switch currentStep {
			case propose:
				t.runPropose(ctx)
			case prevote:
				t.runPrevote(ctx)
			case precommit:
				return t.runPrecommit(ctx)
			}
		}
	}
}
```

The statement that *all* Tendermint consensus states share the outlined structure is not entirely true — remember the *propose* state?
In the *propose* state, there is a clear branching in logic, where non-proposers wait for and verify a proposal, while the proposer for the view needs to build it and broadcast it. The proposer, unlike other validators, *does not need to wait on their own proposal to appear*, but can instead move directly to the *prevote* state, having already accepted the proposal that was sent out.

```go
// ...
func (t *Tendermint) runPropose(ctx context.Context) {
	// ...
	
	// Check if the current process is the proposer for this view
	if t.verifier.IsProposer(t.node.ID(), height, round) {
		// Start the round by constructing and broadcasting a proposal
		t.startRound(height, round)

		return
	}

	// ...
}
// ...
```


### How does it all tie together?

Over the course of this article we’ve gone over the different Tendermint consensus engine states, and their caveats. We’ve also mentioned other asynchronous processes like round jumps, timers triggers — how is all of this managed in Go, in `libtm`?

The entry point into the Tendermint consensus process lies in `RunSequence`, which manages other Go routines that facilitate the previously-seen asynchronous `upon` blocks.
[![run-sequence](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/thumbs/run-sequence.png)](https://gnolang.github.io/blog/2024-06-13_reaching_consensus/src/run-sequence.png)

In the `libtm` library, this management structure looks like this, where `finalizeProposal` and `watchForRoundJumps` are Go routines:

```go
// RunSequence runs the Tendermint consensus sequence for a given height,
// returning only when a proposal has been finalized (consensus reached), or
// the context has been cancelled
func (t *Tendermint) RunSequence(ctx context.Context, h uint64) *FinalizedProposal {
	// Initialize the state before starting the sequence
	t.state.setHeight(h)

	// Grab the process view
	view := &types.View{
		Height: h,
		Round:  t.state.getRound(),
	}

	// Drop all old messages
	t.store.dropMessages(view)

	for {
		// set up the round context
		ctxRound, cancelRound := context.WithCancel(ctx)
		teardown := func() {
			cancelRound()
			t.wg.Wait()
		}

		select {
		case proposal := <-t.finalizeProposal(ctxRound):
			teardown()

			// Check if the proposal has been finalized
			if proposal != nil {
				return proposal
			}

			// 65: Function OnTimeoutPrecommit(height, round) :
			// 66: 	if height = hP ∧ round = roundP then
			// 67: 		StartRound(roundP + 1)
			t.state.increaseRound()
			t.state.step.set(propose)
		case recvRound := <-t.watchForRoundJumps(ctxRound):
			teardown()

			t.state.setRound(recvRound)
			t.state.step.set(propose)
		case <-ctx.Done():
			teardown()

			return nil
		}
	}
}
```

Go contexts, combined with minimal channel usage, allow us to construct a flow that is easily testable, readable and maintainable.

### Trust, but verify

One of the most important choices that we wanted to make while developing `libtm` concerned how we’d verify the functionality, and how frequently during our development process. This was especially amplified by the fact we were experiencing ongoing stability issues with the existing Tendermint BFT testing suite in Gno.

We wanted a testing suite that was minimal enough to verify individual state to state transitions, but also flexible enough to spin up full blown clusters and simulate a live consensus environment.

`libtm` has essentially 2 types of tests:
- regular unit tests that verify state -> state transitions. These tests have “pause” and “resume” functionality that allows us to exactly verify each step of the state machine
- cluster tests in which we simulate valid / byzantine environments. These tests spin up a multimode cluster, fully mocked, to simulate inner-node communication and consensus

We plan on expanding our testing suite even further, with more sophisticated cluster tests that cover an additional range of byzantine scenarios. Currently, our code coverage sits at ~96%.

## Conclusion

Our exploration of the Tendermint consensus engine has highlighted the meticulous design and robust functionality that make it a cornerstone of blockchain systems. The Tendermint consensus engine’s ability to maintain liveness and ensure fault tolerance through its well-defined states—*Propose*, *Prevote*, and *Precommit* — exemplifies the intricate balance needed in distributed systems, where validators can navigate the complexities of reaching consensus even in the face of network delays and potential Byzantine faults.

The implementation of Tendermint within the Gno blockchain, and the broader application potential of the [libtm library](https://github.com/gnolang/libtm), underscore the versatility and adaptability of this consensus mechanism. Our approach to developing `libtm` emphasized modularity and flexibility, ensuring it could be utilized both within and beyond blockchain contexts. By decoupling the library from specific networking protocols, signing mechanisms, and validator set management, we aimed to create a tool that is both easy to integrate and capable of supporting diverse use cases. It not only facilitates smooth blockchain operation but also opens up new possibilities for innovation in distributed systems. 

We look forward to the ongoing improvement of Tendermint, and its continued impact on the blockchain and open-source communities.