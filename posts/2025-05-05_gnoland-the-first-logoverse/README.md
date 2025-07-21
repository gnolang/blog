---
publication_date: 2025-05-05T00:00:00Z
slug: gnoland-the-first-logoverse
tags: [gnoland]
authors: [jaekwon]
---

# Gno.land - The First Logoverse, or Alternate Reality of Logical Objects

Tendermint changed the way blockchain developers think about blockchain
consensus algorithms. Gno.land will change the way developers think about
programming.

Gno.land represents a paradigm shift in multi-user programming that no other
solution offers. It is not just a smart contract platform and blockchain; it is
the world's first viable language-based multi-user operating system. Its
ultimate goal is to host the world's knowledge base for the new millennium.

## Why Gno.land?

Compare publishing a blog site in Gno.land to all prior smart contract systems.

```go
// Realm: gno.land/r/me/myblog

package gnoblog

import (
	"std"

	"gno.land/p/demo/blog"
)

var b = &blog.Blog{
	Title:  "gno.land's blog",
	Prefix: "/r/gnoland/blog:",
}

func AddComment(postSlug, comment string) {
    crossing()
	assertIsCommenter()
	assertNotInPause()

	caller := std.OriginCaller()
	err := b.GetPost(postSlug).AddComment(caller, comment)
	checkErr(err)
}

func Render(path string) string {
	return b.Render(path)
}
```

You can see the Gno source code that rendered this webpage by clicking
on "\<\/\>Source" on the top right of the webpage.

Q: Why is everything else so complicated?

A: Strangely difficult to answer, but ultimately because our languages,
compilers, interpreters, and programming paradigms are still evolving.

## Brief Evolution of Language

Written human language has only been around for a mere 6000 years, a blip in
our evolutionary history. Like living species, our language and writing have
evolved alongside us and within us. Adam was not the first Homo sapiens on
earth, but he may have been the first with written language, and thereby a new
kind of man. 

Programming languages likewise have been evolving rapidly, but only for a
handful of decades; it was in the 1970s when Alan Kay developed Smalltalk, the
first object-oriented programming language. In the 1990’s Brendan Eich of
Netscape invented JavaScript which forever transformed the World Wide Web; Sun
Microsystem made Java, and industries prospered greatly by these and similar
language technologies. 

## Gno vs Previous

Our languages, compilers & interpreters, and programs are today:
 - Nondeterministic - randomness is the norm in concurrent programming, but
   even Go randomizes map iteration. 
 - Disk Bound - programs need to be designed to save to disk —> SQL solutions;
   NOT native language
 - Dependent - running programs are owned by an owner; dependent on
   individuals, not self-sustaining
 - Ephemeral - running programs are expected to fail; no guarantee of
   presence.
 - Single User Realm - import of internal libraries are native, but
   interactions with external programs are NOT native; generally no `import
   “gno.land/r/external/realm”`, but leaky abstractions synthesized ie GRPC

Gno, GnoVM, and Gno.land is in contrast:
 - Deterministic - gno routines not yet supported, but even these will be
   deterministic.
 - Auto Persistent - all changes to instantiated Gno objects in the transaction
   are persisted transparently.
 - Self-Sustaining - every transaction locks $GNOT up for new storage allocated;
   CPU gas fees paid in any language.
 - Timeless - every Gno object that is referenced (not GC’d) remains forever.
 - Multi User Realm - all objects are stored in realm packages (namespaces). 

## Gno Language Innovation

All modern popular programming languages are designed for a single programmer
user. Programming languages support the importing of program libraries natively
for components of the single user's program, but this does not hold true for
interacting with components of another user's (other) program. Gno is an
extension of the Go language for multi-user programming. Gno allows a massive
number of programmers to iteratively and interactively develop a single shared
program such as Gno.land.

The added dimension of the program domain means the language should be extended
to best express the complexities of programming in the inter-realm (inter-user)
domain. In other words, Go is a restricted subset of the Gno language in the
single-user context. (In this analogy client requests for Go web servers don't
count as they run outside the server program).

Gno is Go plus:
 - [`cross(fn)(…)`](https://github.com/gnolang/gno/blob/master/docs/resources/gno-interrealm.md#crossfn-and-crossing-specification)
   calls `fn(…)` where fn is another realm.
 - `std.CurrentRealm()` and `std.PreviousRealm()` changes upon cross-calls.
 - `func fn() { crossing(); … }` signifies that fn is a crossing-function where
   std.CurrentRealm() returns the realm in which the function is declared.
   [Gno2 proposed syntax](https://github.com/gnolang/gno/issues/4223):
   `@fn(…)`, `@func @fn() { … }`. These are like verb (function) modifiers in
   [honorifics in Korean and Japanese](https://en.wikipedia.org/wiki/Honorifics_%28linguistics%29)
 - While all data is readable by other realms, dot.selector access
   across realms get [tainted with 'readonly' attribute](https://github.com/gnolang/gno/blob/master/docs/resources/gno-interrealm.md#readonly-taint-specification).
 - [`revive(fn)`](https://github.com/gnolang/gno/blob/master/docs/resources/gno-interrealm.md#panic-and-revivefn)
   for Software Transactional Memory (STM). 
 - Function/method return implies access without readonly taint.
 - Inter-realm type conversion limitations to prevent exploits.
 - More and refinements to come in Gno2.

These language innovations/extensions allow for safer multi-user application
development where many users are collaboratively programming a single timeless
(immortal) communal program.

## The Logoverse

Ἐν ἀρχῇ ἦν ὁ Λόγος καὶ ὁ Λόγος ἦν πρὸς τὸν Θεόν καὶ Θεὸς ἦν ὁ Λόγος.  In the
beginning was the Word (Logos), and the Word was with God, and the Word was
God. - John 1:1

Logos means “word, discourse; reason”, and shares its root with the word
“logic”. 

With these elements altogether you can derive a new property:
 - Gno expressions become "real" on Gno.land.
 - Ethereum comes close but isn't object-oriented and Solidity has no pronouns.
 - TBL's WWW, DOM model, HTTP verbs, Plan 9, Ethereum, and FB Meta are all
   attempts to arrive at the logoverse.
 - Gno.land is the first complete logoverse.

## Adoption Strategy

There are over a million Go developers and growing. Go as a language remains a
popular language for developers, an order of magnitude more than Rust
developers, on par with JavaScript developers but growing faster than
JavaScript.

[![TIOBE 2025](https://gnolang.github.io/blog/2025-05-05_gnoland-the-first-logoverse/src/thumbs/tiobe_2025.png)](https://gnolang.github.io/blog/2025-05-05_gnoland-the-first-logoverse/src/tiobe_2025.png)
[![GitHut2 2024](https://gnolang.github.io/blog/2025-05-05_gnoland-the-first-logoverse/src/thumbs/githut2_2024.png)](https://gnolang.github.io/blog/2025-05-05_gnoland-the-first-logoverse/src/githut2_2024.png)

Gno.land and its associated network of Gno VM chains, and AtomOne if it hosts
it, will become the nexus of human to human, human to machine, and machine to
machine coordination; but only after it finds a self-sustaining organic growth
cycle.

The best way to ensure success and to accelerate adoption is to seed the
initial community with the right community. There are many types of
communities, such as the crypto community, ethereum community, student community,
but since Bitcoin has gone mainstream, these communities aren't always in
agreement about the purpose of blockchain technology; because they aren't aware
of the history and fabric of the hidden power structures that run the
narrative--both mainstream AND controlled oppositions. They do not feel that
they need something, so their habits are not as obvious to change.

But the "free-thinking" and "conspiracy" and "anti-war" and "anti-Covid19-vax"
and even the "true Christian" communities feel an urgent need for
censorship-proof coordination and communication tools. These communities have
influencers who are kept hidden from the broader public; they have suffered
deplatforming, defamation, and even death.

Build tools, connections, and relations with these particular communities and
especially those influencers who are nuanced in their research and speech.
Even those that don't promote crypto will see the benefits uniquely offered by
Gno.land.

## Gno.land License

Anyone can make Gno VM powered chains derived from Gno.land according to the
viral copyleft license terms and strong attribution requirement. The Strong
Attribution clause of the Gno Network GPL license preserves the spirit of the
GNU AGPL license for the blockchain world.

## Tokenomics

$GNOT is the storage lock-up utility token, so Gno.land is to Gno England like
$GNOT is to presence in Gno England, where total storage is kept finite for
very-long-term existential purposes, and value is derived from the Gno
artifacts created by its users, and some new users competing for attention from
many existing users.

Gno.land may migrate to AtomOne ICS once it is support hard-fork upgrades.
There Gno.land would be one ICS shard, and many Gno VM shards may also exist,
each with their own namespace and probably each their own storage token unless
separate treaties are made between the main Gno.land chain (ICS shard) and
other Gno VM shards. Transaction fees for CPU usage may be paid in either $GNOT
or $PHOTON.

## Team

### New Tendermint, LLC

NewTendermint, LLC is the core maintainer of the GnoVM, Tendermint2, and 
at present Gno.land.

#### GnoVM & Gno.land Core Team

 * Jae Kwon before and after creating Tendermint and Cosmos always had a
   passion for programming languages and wrote multiple parsers and
   interpreters, and initially also wrote an EVM on top of the framework which
   became the Cosmos SDK. Gno.land is the result of two decades of search for
   the logoverse.

 * Manfred Touron, builder focused on open-source and resilient technologies;
   co-founded scaleway (cloud) and berty (p2p messaging), with contributions to
   900+ open-source projects.

 * Miloš Živković - Senior distributed systems engineer; passion for solving
   protocol-level problems in the blockchain space.

 * Morgan Bazalgette - Senior Go engineer; bringing the joy of developing Go to
   Gno.

 * Ray Qin - With over 15 years of experience in software development and
   building large-scale networks, I have a deep passion for Go programming
   language and blockchain technology.

 * Marc Vertes - Senior VM and hardware developer; more than 3 decades of
   experience, Co-founder of 3 companies (1 acquired by IBM), author of 34
   patents, author of the Yaegi Go interpreter.

 * Alexis Colin - Senior Frontend Engineer with 10+ years of experience
   building user-focused interfaces and exploring modern tech stacks. Driven to
   push forward efficient UX in the blockchain space through clean code and
   technical precision. Currently working on gnoweb.

#### Gno Studio Team

 * Ìlker Öztürk - Senior software architect; 17 years in building and designing
   products, distributed p2p systems, leadership and strategic vision.

 * Jerónimo Albi - Experienced full-stack systems engineer with attention on
   simplicity and minimalism, focused on Blockchain and Golang development

 * Salvatore Mazzarino - Site Reliability Engineer with over 10 years of
   experience in building and mantaining high and scalable distribute systems
   across different range of platforms. "If it is not monitored, it does not
   exist"

 * Danny Salman - Vast experience in blockchain developer relations, technical
   writing and education, and product, with a background in full-stack
   development, engineering, and policy.

 * Alan Soares - A Brazilian lost in middle earth.  Passionate coder with love
   for open-source and software craftsmanship. Taking the web forward for over
   16 years.

 * Lucio Caetano - Senior Frontend Engineer with 10+ years of experience
   building web applications, specializing in web3 and blockchain technologies.
   With a background in data analysis, providing data-driven decisions and
   reporting with large datasets.

### All in Bits, Inc 

Members of All in Bits, Inc also build applications and community on top of
Gno.land.

 * Kristov Atlas - Since 2012, I’ve been a crypto security engineer and
   researcher focusing on non-custodial wallets, CeFi exchanges, Bitcoin,
   Ethereum, and the Cosmos ecosystem.

 * Michelle Leech - Experienced marketing and ecosystem builder skilled at
   creating and driving strategic initiatives that foster relationship building
   and boost developer advocacy, engagement, education, and product
   utilization.

 * Lav Leon Hudak - DevRel Engineer with a strong background in blockchain
   development, documentation, and education.

 * Sean Casey - CFO (Chief Financial Gnome) - 10+ years’ experience in finance,
   aircraft leasing, and blockchain. Leads financial strategy, and treasury
   operations, ensuring capital discipline, regulatory compliance, and
   long-term value creation.

 * Jordan Magazine - Experienced General Counsel with over a decade of
   practice, specializing in designing frameworks that enable blockchain
   ecosystems to operate with clarity and compliance, empowering projects to
   scale confidently and with integrity.

 * Carolyn Pehrson - Paralegal keeping AIB,Inc and NT,LLC alive. Thinks you
   should get a pet mini pig.

### Onbloc Co., Ltd.

Onbloc Co., Ltd. is a core contributor to Gno, delivering essential tools to 
the Gno.land ecosystem—including the Adena Wallet, GnoScan Explorer, and GnoSwap 
AMM DEX—for both developers and everyday users.

 * Dongwon Shin - Blockchain evangelist and product‑driven builder, obsessed with
   user experience and ecosystem growth. Before founding Onbloc, he worked on several
   crypto initiatives, including a research firm, a crypto exchange, and a consulting firm.

 * Peter Yoon - Blockchain advocate and strategic specialist. A long-time supporter
   of Cosmos since its inception, he was an early member of Cosmos Korea and founded
   Lunamint, the first crypto wallet integrated with Telegram.

 * Jinwoo Choi - Full-stack engineer behind the Adena Wallet and GnoScan, with
   extensive experience building stable and scalable software systems. Passionate
   about designing robust architectures that endure scale and complexity.

 * Byeongjun Lee - Core engineer for Gno.land and the GnoSwap project, passionate about
   compilers and programming languages. He's dedicated to making developers’ lives easier.

 * Andrew Kang - Experienced product manager and researcher, skilled at identifying
   trends and driving strategic initiatives that advance the crypto ecosystem.
