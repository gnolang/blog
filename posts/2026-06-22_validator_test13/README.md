---
publication_date: 2026-06-22T00:00:00Z
slug: validator-test13
tags: [gnoland, ecosystem, validators, blog]
authors: [michelleellen]
---

# **Invitation to Test13 Validators: How to Join**

Test13 is live, and we're introducing a single, standardized onboarding path for all external Testnet validators. Please note, with Test13 now running successfully, we'll be sunsetting Test11 on Wednesday, June 24th.

Validators looking toward joining the Gno.land Mainnet can begin with Test13. Registration through the Valoper Registry is required, and the accuracy of submitted information matters, since that data carries forward into the on-chain governance voting system.

There is no slashing or jailing at this stage. Performance is evaluated on uptime, correct configuration, and contribution to the network and its community. As the Test13 validator set matures, the plan is to coordinate a broader Mainnet phase with validators who have demonstrated high levels of system infrastructure and reliability, serving as a path towards potential Mainnet validation. How consistently you stay online now will matter for both future voting power increases and your standing for Mainnet validating.

Validators on Test13 will be voted on by a testnet GovDAO. For Test13, we'll execute batch GovDAO proposals to expedite onboarding, based on historical contributions, AtomOne alignment, and accurate secret key submissions. **We'll also be running a validator verification as part of the Test13 onboarding process on Discord**.

Here's exactly what you need to do to get registered, voted in, and active on the test network.

## **Step 1: Get Set Up on Test13**

Before registering as a validator, we've provided some information to help you get ready:

- General Test13 [documentation](https://github.com/gnolang/gno/blob/chain/test13/misc/deployments/test13.gno.land/README.md) and steps on how to [run a full node](https://github.com/gnolang/gno/blob/chain/test13/misc/deployments/test13.gno.land/VALIDATOR.md) on Test13. 
- A [guide](https://gnops.io/articles/guides/become-testnet-validator/) to becoming a testnet validator. 
- Test13 RPC connection points are [available](https://rpc.test13.testnets.gno.land/). 

## **Step 2: Register on the Valoper Registry**

Every validator needs a profile on the [Valoper Registry](https://test13.testnets.gno.land/r/gnops/valopers) for Test13. You can get testnet tokens on the [Faucet Hub](https://faucet.gno.land/) to make the registration transaction.

When filling out your profile:

- Double-check that you're submitting the correct keys.
- Include your properly configured gnoland secrets as required during registration.

Mistakes here are the most common reason approval gets declined or held up, and they're treated as a negative when evaluating validators.

## **Step 3: Take Action in Discord & Complete the Validator Verification**

Once your valoper registration is complete, head to the Gno.land Discord [#general-chat](https://discord.gg/WDsjtPHCG).

- Follow the pinned instructions; you should run /candidate-testnet.
- The bot will assign you ‘Testnet Validator Candidate’ role, and grant access to #testnet-onboarding.
- In the #testnet-onboarding channel, you’ll need to run /submit-request and provide your operator address, and two additional questions.
- Once you've completed the verification, the team will review it and you’ll move to the GovDAO voting phase.

## **Step 5: Get Voted In Through GovDAO**

Registration alone doesn't officially add you to the network, the Valoper Registry is a public directory. You also need to be voted in through GovDAO. After completing the verification step, we'll start adding validators to the voting system.

A few things to know about this process:

- We'll be running batch voting to move things along efficiently for Test13.
- Priority will go to validators who've completed the registry, verification, and are active contributors to the Gno.land and AtomOne ecosystems.
- Once the vote passes, your validator will be automatically added to the network's validator set. Please keep an eye out for this.
- On Discord, you'll be added to the #testnet-general (Test13) channel and given the associated ‘Testnet Validator’ role.

## **Step 6: Active Test13 Validator**

On Test13, all validators start out on equal footing:

- Every validator is assigned an initial voting power of 1.
- From there, voting power can increase via a GovDAO proposal, based on performance and uptime.
- This will help inform the selection process for the Mainnet validator set.

Thank you to all the validators who have contributed to the project and shown patience along the way, including block explorer work from [TechNodes](https://gno.satai.0xgen.online/validators) and [Oshvank](https://explorer-gnoland.oshvank.xyz/), [GnoDuty](https://github.com/AviaOne/GnoDuty) from AviaOne, and log support from Kalpa Tech and ProDelegators. If you're looking for ways to contribute to the Gno.land validator ecosystem, we'd love to hear from you.
