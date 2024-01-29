---
title: "5 Things I Learned While Porting Flippando From Solidity to Gno"
publication_date: 2024-01-24T00:00:00Z
slug: porting-flippando-gno
tags: [gnoland, ecosystem, updates, flippando]
authors: [dragos]
---

Last year, while visiting Seoul, South Korea, I decided, on a whim, to sign up for a hackathon called Glitch. The project I was going to present was a tiny little game, written in Solidity, called Flippando. It started as a weekend project to help me learn Solidity (I had no prior experience with this language). To my surprise, my tiny little game won the first prize on the Polygon track of the Glitch hackathon.

Encouraged and even more curious now, I started attending side events during Buidl.Asia. One was about Gno, a smart contract platform written in Go. After the presentation, which was really great, I started a light conversation with the team. One thing led to another, and I ended up showing them Flippando. 

Just for context, Flippando is a non-degen, deceptively simple memory game. You start with an empty matrix and flip tiles to see what’s “underneath.” If the tiles match, they remain uncovered; if not, they are briefly shown, and you have to memorize their color until the entire matrix is uncovered. The end result can be minted as an NFT, and you can later assemble all the boards into bigger, more complex NFTs, basically “painting” with the uncovered tiles.

The Gno team seemed to like it, and they suggested I should apply for a grant to port it to Go/Gno. I had no prior experience in Go either, so I thought this would be a good opportunity to learn more. To my surprise, again, my grant submission was accepted.

Fast forward a few months until now: the Gno version of Flippando is live (in testnet beta) at [https://gno.flippando.xyz](https://gno.flippando.xyz). What follows sums up my experience porting the game from Solidity to Gno. This blog post is a mix of technical and not-so-technical takeaways.

## 1. Being Early Pays Off

Solidity has been around for some time now, and there is already a solid tooling ecosystem for it. I used Hardhat for my development, and I got really comfortable with it. When I started to port Flippando, though, I was quite surprised to see there was almost no tooling in Gno. Developing was mostly TDD (test-driven development) against a local VM, and deploying realms on the actual chain was more complicated than I expected. 

My first feedback rounds to the team revolved almost exclusively around this topic. Very soon, I started to receive signals that my feedback was not only heard but taken into account and processed, and there were actual projects built aiming to improve the developer experience. In just two or three months, two full projects were finished: gnodev, and Gno Playground. 

Gnodev makes development very similar to Metro in React Native: there is a watchdog on the file system, and your changes to the realm code are reloaded every time you save. It’s almost like deploying in real time; no need to stop the chain, wipe the state, restart the chain, and redeploy your modifications. Gno Playground is a sandbox-like environment, which helps tremendously with quick testing and even deploying packages on-chain. Both projects were finished, as I said, in just two to three months.

Being early pays off because you get to shape your development environment much faster than in a solidified (pun intended!) environment. You may have to deal with a little chaos in the beginning, but the benefits are well worth it.

## 2. TDD All Day Long

As I said above, developing realms in Gno consists mainly of writing and testing your code with another code. It’s called TDD and it’s a very useful developing strategy, in general. I used it, at my day job, in all my projects consistently, but only in the initial stages. Once the codebase was more stable, I was relying more on regression tests from the Q&A team.

Mind you, there was no Q&A team this time; I was just coding alone, and I was forced to comply more and more with this TDD approach. In the end, I have to admit that, while slower and a bit boring, this approach is more effective, especially in a volatile environment, where patches are added literally every day, and the environment changes continuously.

## 3. Marshal and Unmarshal

The current GnoVM doesn’t yet have an API standard for formatting. You can’t put a setting somewhere that will make the response be automatically translated into JSON. You have to write these JSON objects yourself for every payload you return from your realm. 

In Solidity, all this is hidden under the event mechanism and handled by existing libraries, like ether.js, which take care of all this nitpicking. It soon became obvious that development time would be significantly longer in Gno because, on top of the logic, I also had to write the formatted response “by hand.”

But as with every other thing that seemed weird in the beginning, eventually, I came to appreciate it. It forced me to prototype more carefully not only the actual response but all the objects needed in my game. Eventually, it resulted in simpler and more flexible code.

## 4. Eating Your Own Dog Food

When developing in Solidity, most of the time, you just import OpenZeppelin contracts for ERC20 and ERC721 tokens (which are battle-tested, bug-free, and relatively easy to understand) and focus on your own contract logic. No mingling with low-level token implementation details; these are already packaged and ready to use.

While porting Flippando to Gno, I realized I had to deal with these low-level details upfront simply because there was no equivalent of the OpenZeppeling contracts. Moreover, some current GRCs (the Gno equivalent of ERC) were incomplete. 

So, I had to make a PR for a GRC721 implementation that was missing the SetTokenURI functionality, and this PR ended up being merged into the main Gno codebase (that felt really good, to be honest). 

## 5. Being Early Pays Off. Did I Say That Already?

Yes, but this time it’s about something else. It’s not about the satisfaction of shaping the development environment in the early days. It’s about the privilege of witnessing something coming to life from literally nothing. Gno has been in development for almost two years now, and it is several months before its mainnet. It’s literally on the verge of coming “alive.”

Every day new commits are added, and new decisions are made. There are new contributors constantly joining, and new projects prototyped and launched faster and faster. Every day the ecosystem is coagulating itself into something more and more visible, more and more alive.

Being able to witness this from the inside is a rare privilege and something I’m very grateful for.

## Final Thoughts 

So, these are, in a nutshell, my five top takeaways from porting Flippando from Solidity to Gno. There are many others, of course, and Gno is (did I already say this?) still very early. If you’re interested in learning more, please visit the official repo, look at the docs, and try interacting with the devs. You’ll never gno what can grow out of it! And be sure to play [Flippando](https://gno.flippando.xyz) today live in testnet beta and share your flips.

## Here’s How to Play Flippando

The game presents a 16 tiles (4x4) or 64 tiles (8x8) matrix. These tiles are “covering” a board of various colors and gradients or shapes, like dice or hexagrams. Clicking two tiles consecutively “flips” them, showing what’s underneath. If they match, they remain uncovered; if not, they are briefly shown, and the player needs to remember their position. Once an entire board is flipped, revealing its random combination of colors, the player can choose to mint it as an NFT.

When minting a solved board as an NFT, the game also mints a fungible token, FLIP, which is “locked” inside the NFT. This is the player's “reward.” But the token can only be unlocked if someone else uses that NFT in a larger project.

These larger projects, or “artworks,” can be assembled in the Flippando Playground. All minted basic NFTs are displayed here in an area from where the player can drag and drop them onto a canvas, creating a much bigger and more complex NFT. Once the canvas is fully filled and the player is satisfied with what’s in there, these new “artwork” NFTs can also be minted. This unlocks all the FLIP tokens for the NFTs used inside the artwork and sends them to their initial players. Furthermore, these complex artworks can be listed and traded in a marketplace, closing the circle of a virtual economy of goods.

Start playing Flippando and share your Flips with Gno.land on [Twitter/X](https://x.com/_gnoland?lang=en) by tagging #gnoflip. 


