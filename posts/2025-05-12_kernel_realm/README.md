---
publication_date: 2025-05-12T13:17:22Z
slug: kernel and realm
tags: [kernel, process, realm]
authors: [mvertes]
---

# Gno realms vs system kernels

Let's explore the analogy between Gno and existing operating systems.

In computers, the role of the operating system is to isolate each process from
the others, for memory integrity and security, and to control access to
system resources, including time sharing access to the CPU. It also lets processes
communicate with others, or externally through the network.

The operating system provides those capabilities through the kernel, which
executes system calls.  When invoked by a user process,  the kernel switches the
CPU into a privileged mode, to perform the services outside user reach. Each
resource located outside of the process means the private user memory must be manipulated
through system calls. This is done under the strict supervision of the kernel, which acts as
a neutral trusted facilitator between untrusted processes. On bare metal
systems, this is enforced by the hardware itself.

Gno is a distributed multi-user virtual computer implemented on top of the
blockchain, ensuring all operations execute deterministically. The distributed
storage layer of the blockchain provides Merkle trees and consensus-based verification
that makes program code and memory tamper-proof. 

In Gno, the equivalent of computer processes are *realms*. Each realm is like an
always-active process, running forever, like for example, a web server. In a
regular system, a live process is defined by its Process Identifier (PID, a number set by the system for the life of the process), its memory, and the set
of system resources it uses at a given time (files, connections, etc).
Similarly in Gno, a realm is defined by its identity (a unique crypto address),
its global memory state (the content of memory when the realm is at rest),
which may contain for example the amount of coins it retains.

A realm may provide services to other realms through the exported public
functions it declares, and it may use services provided by other realms by
importing the realm using an `import`  statement, and directly calling its functions
in code. The way of declaring exported functions and importing them is
exactly identical to how packages are defined in the Go language. A
realm is a Gno process, but is also a package (in the Go sense).

Functions in packages can be pure (in the functional programming sense: the
function has no side effects and operates only on temporary local variables
which are discarded at return), or not: some variable outside of the function
scope is modified. In that case, the modified variable must be defined in the
same package as the function. Functions can only write to global variables
declared within their own package space (assuming variables are unexported),
and static code analysis, as in Go, enforces this rule.

But because realms are also stateful processes with their own identity and
protected space, and because a process can write only in its own space, calling a
non-pure function imported from another realm means that the caller realm would
attempt to write in the callee realm space, which is forbidden!

To resolve this, the current realm identity is switched from the caller to the
callee until the function returns, like when crossing the user-kernel boundary
in a regular system call. From the system's point of view, the current realm is
set to the callee, the previous realm to the caller. At function return, the
current realm and previous realm are restored to their original value.

So non-pure exported functions act exactly like a system call: they provide the
ability to write outside the calling realm space, by crossing the realm
boundary. A realm doesn't need to trust the calling realm: only itself can write in
its own space. A realm doesn't need to trust an external kernel: it is its own
kernel, it decides exactly how its data can be accessed and/or modified by
caller realms.

Gno unifies the concepts of Go packages and Unix processes in a single simple
concept: the realm. The kernel is decentralized and put back in control of
package developers. The operating system itself becomes transparent: realm
processes are both resource users and providers, with full control and
accountability.

In Gno, programs are processes are packages are realms. Gno rethinks the traditional boundaries between processes, packages, and system calls. By making each realm self-contained and in control of its own state, it replaces the need for a central kernel with a decentralized model where programs define their own access rules. It’s a practical and minimal approach to building secure, composable software on-chain.
