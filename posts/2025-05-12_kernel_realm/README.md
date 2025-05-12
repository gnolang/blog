---
publication_date: 2025-05-12T13:17:22Z
slug: kernel and realm
tags: [kernel, process, realm]
authors: [mvertes]
---

# Gno realms vs system kernels

Let's explore the analogy between Gno and operating systems.

In computers, the role of the operating system is to isolate each process from
each other, for memory integrity and security, and to control access to
system resources, including time sharing access to CPU. It also let processes
to communicate with others, or externally through the network.

The operating system provides those capabilities through the kernel, which
execute system calls.  When invoked by a user process,  the kernel switches the
CPU into a priviledged mode, to perform the services outside user reach. Each
resource located outside of the process private user memory must be manipulated
through system calls, under the strict supervision of the kernel, which acts as
a neutral trusted facilitator between untrusted processes. On bare metal
systems, this is enforced by the hardware itself.

Gno is a distributed multi-user virtual computer implemented on top of the
blockchain, so that every operation occurs in a deterministic way. User program
code and memory states are made temper proof by the use of merkle trees and
consensus based verification, all provided by the underlying blockchain storage
layer, itself distributed.

In Gno, the equivalent of computer processes are *realms*. Each realm is like an
always active process, running for ever, like for example a web server. In a
regular system, a live process is defined by its PID, its memory, and the set
of system resources it uses at a given time (files, connections, etc).
Similarly in Gno, a realm is defined by its identity (a unique crypto address),
its global memory state (the content of memory when the realm is at rest),
which may contain for example the amount of coins it detains.

A realm may provide services to other realms through the exported public
functions it declares, and it may use services provided by other realms by
importing another realm functions and directly calling the other realm function
in its code. The way of declaring exported function, and importing them is
exactly identical to the way that packages are defined in the Go language. A
realm is a Gno process, but is a also a package (in the Go sense).

Functions in packages can be pure (in the functional programming sense: the
function has no side effect and operates only on temporary local variables
which are disgarded at return), or not: some variable outside of the function
scope is modified. In that case, the modified variable must be defined in the
same package as the function. Functions can only write global variables
declared inside their own package space (assuming variables are unexported),
and it is always inforced, as in Go, by static code analysis.

But because realms are also stateful processes with their own identity and
protected space, and that a process can only write in its own space, calling a
non-pure function imported from another realm means that the caller realm would
attempt to write in the callee realm space, which is forbidden!

To resolve this, the current realm identity is switched from the caller to the
callee until the function returns, like when crossing the user-kernel boundary
in a regular system call. From the system point of view, the current realm is
set to the callee, the previous realm to the caller. At function return, the
current realm and previous realm are restored to their original value.

So non-pure exported functions act exactly like a system call: they provide the
ability to write outside the calling realm space, by crossing the realm
boundary. A realm doesn't need to trust a caller realm: only itself can write in
its own space. A realm doesn`t need to trust an external kernel: it is its own
kernel, it decides exactly how its data can be accessed and or modified by
caller realms.

Gno unifies the concepts of Go package and Unix process in a single simple
concept: the realm. The kernel is decentralized and put back in control of
package developers. The operating system itself becomes transparent: realm
processes are both resource users and providers, with full control and
accountability.

In Gno, programs are processes are packages are realms.
