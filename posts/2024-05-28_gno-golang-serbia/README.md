---
publication_date: 2024-05-21T12:33:00Z
slug: gnomes-in-serbia
tags: [gnome, gnolocal, retreat, meetup, golang]
authors: [leohhhn, michelleellen]
---

# Gnomes Spotted in Belgrade, Serbia: Recap from the Engineering Retreat and Golang Serbia Meetup

Dobro jutro gnomes! Last week, the Gno core engineering team convened in Belgrade,
Serbia, to discuss various aspects of our technical roadmap, including Test4, 
blockchainless Gno, Gnoweb, governance and DAO structuring, GnoVM, and more. 
While the team was in town, we took the chance to host a Gno @ Golang Serbia 
meetup titled ‘Building Dynamic Applications with Interpreted Go’ - you can catch 
the recording of the full presentation [here](https://youtu.be/tNM1DHOxIQ8).

## Engineering Retreat Recap

Over five full days of coding and workshops, the Gno core engineering team delved 
into several critical topics and emerged with several key takeaways:
- Test4 Progress:
  - Test4 has been the primary effort for the team in the past couple of months, 
aiming to create a stable multi-node testnet which will serve as the last precursor
to Gno.land’s upcoming mainnet. Discussions revolved around chain initialization 
flows, node metrics & telemetry, versioning of binaries, adding a way to modify 
the validator set via a realm (r/sys/vals), and more. Currently, the 
[Test4 milestone](https://github.com/gnolang/gno/milestone/4) is 62% complete.
- Blockchainless Gno:
  - With the upcoming GopherCon EU & US conferences, the team discussed a 
different perspective on Gno - a fully functional Go interpreter with automatic 
state persistence. The team aims to modularize the GnoVM so that it can be
used in contexts outside the Gno.land blockchain, which will open up the
ecosystem to completely new use-cases, inviting even more tweb2 developers to 
join our mission.
- Gnoweb Enhancements:
  - [gnoweb](https://docs.gno.land/getting-started/local-setup/browsing-gnoland#2-browsing-gnoland) serves as one of the main tools to explore the Gno.land 
ecosystem. The team discussed changes to the UI as well as improving gnoweb’s 
functionalities such as the rendering of realm state & source code, and possible
approaches to add more ways to interact with on-chain apps.
- Governance and DAO Structuring:
  - Discussions relating to GovDAO and WorxDAO took place during the retreat, as 
they are the foundation for Gno.land’s consensus mechanism - Proof of Contribution.
Manfred Touron, Gno.land’s VP of Engineering, discussed the 
[initial implementation](https://github.com/gnolang/gno/pull/1945) of GovDAO and
its surrounding infrastructure.
- GnoVM & Gno.land Development:
  - The team took the opportunity to deal with priority PRs during the retreat, 
such as the [Gno Type Check PR](https://github.com/gnolang/gno/pull/1426) by
@itsmaxwell, which will add full type checking to the VM. The team has also
added GoReleaser to the monorepo in preparation for the upcoming Test4 milestone,
and the first 
[nightly Gno build](https://github.com/gnolang/gno/releases/tag/v0.1.0-nightly.20240523) was released.

## Gno @ Golang Serbia

While in Belgrade, we also hosted a new Gnolocal meetup with over 15 developers 
from the Golang Serbia community. During this event, we introduced Gno.land, with
a particular focus on the GnoVM (Virtual Machine) as a foundational layer for 
Gno - an interpreted version of Go. This was the first time we presented the
concept of Gno to Serbia’s Golang community, showcasing its potential to support
dynamic application development. It was an opportunity to identify areas of 
improvement and feedback from the user community.

[![banner](https://gnolang.github.io/blog/2024-05-28_gno-golang-serbia/src/thumbs/meetup.png)](https://gnolang.github.io/blog/2024-05-28_gno-golang-serbia/src/meetup.png)

[Dylan Boltz](https://github.com/deelawn), a Senior Golang Engineer from the core
team, led the presentation, demonstrating how the VM and Gno can be leveraged to 
design and build efficient, adaptable, and scalable applications. He also provided 
an overview of the stack-based VM architecture, illustrating how data storage 
mechanisms and execution traces operate within the VM.

By using the GnoVM outside of the blockchain context, relevant to the Go community,
developers can utilize the powerful features of Gno to build applications with 
automatic state persistence within a sandboxed environment. Dylan showcased how 
the GnoVM can be embedded in an HTTP server, allowing developers to write and 
persist Gno applications locally, and then share them with other users, all while
maintaining the security of a VM.

### The Meetup presentation highlights
- Virtual Machine Deep Dive: We provided a detailed understanding of the 
architectural setups of various VMs, and gave an overview of the current GnoVM.
- Hands-On Learning: We walked through how to embed a virtual machine into your
Go applications for dynamic code interpretation. The presentation covered 
practical techniques, including creating browser-based interfaces using interpreted Go.
- Interactive Demonstrations: We showed how to create browser-based interfaces
with interpreted Go as a foundation, demonstrating how this architecture enables
dynamic program execution while maintaining a structured, deterministic approach
to storage and state.
- This was the third meetup in our series of Gnolocal events. We’ll be popping
up in more cities around the globe to connect with gnomes, and spread the word 
about Gno.land. With a few gnomes based in Belgrade, it’s important to keep 
cultivating and building the Gno.land community locally. If you’re based in 
Serbia, you can find our regional based channel on Discord (TOBEADDED).

### The Feedback Loop
After the presentation, we gathered feedback from the attendees to assess the 
content presented, their interest in Gno.land, and in their interest in using 
the Gno interpreter tooling. Here are the key insights from form responses:
- 75% of participants said they were interested in learning more about Gno.
- 87% of participants found the presentation and concepts understandable.
- 38% expressed interest in using the embedded Go interpreter in their applications.
The feedback highlighted the need for us to clearly outline the problems Gno.land 
- is solving, how our technology addresses these issues, and how it can apply to
- real-world examples. This feedback is invaluable as it helps us refine our 
- approach and better engage with new contributors.

## Conclusion
The meetup was just one of several activities the team organized in Belgrade. 
In addition to the extensive technical sessions and workshops, we had the 
opportunity to experience and learn a bit about the local culture and Serbia,
visiting the Nikola Tesla Museum, sightseeing, and experiencing traditional 
Serbian music.
These in-person engineering retreats are some of the most important moments in
outlining priorities, troubleshooting the technology together, and brainstorming
ways to generally enhance and optimize the Gno.land blockchain and builder 
experience. 
