---
publication_date: 2024-12-05T00:00:00Z
slug: audit-proposal-request
tags: [gnoland, gnovm, security, audit]
authors: [kristovatlas]
---

# Call for Security Auditors - gno.land Audit

In the gno.land ecosystem, security is one of our top priorities as we continue to develop our blockchain focused on smart contracts and decentralized applications. As we move closer to the beta launch, we're calling on experienced security auditors to help us ensure the robustness and security of our codebase.
We are now accepting proposals for the auditing of the gno repository on GitHub. This is a significant and complex task, so we are looking for qualified firms or individuals with experience auditing blockchain systems, virtual machines, and distributed systems.
You can find the full gno.land repository here:

https://github.com/gnolang/gno

## What's in Scope?

We've already done some work to narrow the scope of the code that requires auditing. For a more detailed list of which files are in scope and which are not, please refer to the [Google Spreadsheet](https://docs.google.com/spreadsheets/d/1rAvzvCH1TBZAykWCzKpefJnbx0SaCEgnP6IypzX8Xo4/edit?usp=sharing) that outlines the specifics.

### Scope Breakdown by Directory

Here's the estimated breakdown of the lines of code that are in scope for auditing as of November 11th. This code is largely written in Go:

* **gnovm**: 43,452 lines (25% of total)
* **gno.land**: 18,677 lines (10% of total)
* **tm2**: 108,020 lines (65% of total)

Total lines in scope: 170,149 lines

This won't change drastically before the audit, so these line counts may help auditors estimate time costs. Many proposals we've received so far are simply time-boxing for a 4 engineer-week audit.

## Auditing Approach

Due to the size and complexity of the audit, we are considering breaking the task into sub-components. We may hire multiple firms to audit different parts of the codebase simultaneously. The proposed divisions are:

1. Gno Virtual Machine (GnoVM) - contained in the "gnovm" subdirectory
2. The Remaining Components - including the "gno.land" and "tm2" subdirectories

### Key Focus Areas for Auditors

While the full codebase needs a thorough review, we've identified several areas of potential security concern that require special attention:
1. **Determinism in GnoVM**  
    As a blockchain designed for deterministic execution, ensuring that the GnoVM executes contracts consistently across all nodes is crucial. Our goal is to eliminate non-deterministic components from Go, such as using AVL trees instead of Go maps. However, we may still have lingering issues that could lead to non-deterministic behavior. A prime example is the module within `gnovm/pkg/gnolang/values_string.go`, which should be carefully reviewed for any such issues.  
    **Why this matters**: Non-determinism can lead to chain halts or splits, which could be exploited by attackers.
2. **Other GnoVM Challenges**  
    Gno.land contributor Morgan has detailed some additional areas of concern of the Virtual Machine here: https://github.com/gnolang/gno/issues/2886#issuecomment-2400274812 
3. **Security in Realms (Smart Contracts)**  
    Developers deploy smart contracts, called "Realms," to the chain. Malicious Realms could attempt to inject harmful content that could affect other users of the chain, particularly in the `Render` function or supporting tools like **Gnoweb**, which displays Realms to end users.  
    **Potential risk**: Cross-site scripting (XSS) and other injection attacks.
4. **Security Risks Found in Other Blockchain VMs**  
    Many blockchain VMs, such as Ethereum's EVM, have faced high-profile security issues. We expect that similar vulnerabilities could be targeted in the GnoVM, so it's crucial to audit and mitigate these risks in advance.
5. **Key Management (gnokey)**  
    Although this is a lower priority than some previously mentioned, auditors will need to review the gnokey package, which handles key generation and signing, to ensure that security best practices are being followed; for example, ensure our Ledger hardware wallet integration with gnokey uses the correct build flags.

## How to Submit a Proposal

We're looking for proposals that include the following line items:

- **Cost of auditing the entire "gno" repository** -- covering all directories and files in scope.
- **Cost of auditing the Gno Virtual Machine (gnovm)** -- focused on the **gnovm** subdirectory.
- **Cost of auditing the other components** -- covering the **gno.land** and **tm2** subdirectories.
- **Fuzzing efforts** -- auditors are encouraged to include fuzzing as part of their code review process, though it should be listed as a separate line item.

Please include the timelines for auditing and your current availability beginning December 23rd. 

## Timeline & Contact Information
We expect the audit to consume at least 4 engineer weeks and conclude by the end of the day January 20th, giving us a week to remediate any high- or critical-severity issues by January 28th. Audit teams may decide to dedicate multiple auditors in parallel to meet our desired timeline. We are open to discussions with auditing teams to answer any questions about the scope of the project.

Please send your proposals via email to **security [at] tendermint [dot] com**. We are happy to meet with potential auditing teams to further discuss the details and answer any questions.


## Conclusion

The gno.land community is committed to building a secure and resilient blockchain platform. We believe that involving experienced auditors in our development process is essential to ensure that we can deliver a secure environment for dApp developers. If you or your team have the expertise and are interested in contributing to gno.land's security, we encourage you to submit a proposal.

Let's work together to make gno.land the most secure blockchain for smart contracts and decentralized applications!
