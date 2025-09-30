---
publication_date: 2025-05-XXXT00:00:00Z
slug: XXX
tags: [gnovm]
authors: [manfred]
---

# Interrealm Calls vs. Traditional Microservices: A Paradigm Shift

Understanding Gno's interrealm call mechanism can be simplified by comparing it
to the familiar concept of microservices in traditional distributed systems.
While both aim to modularize functionality, Gno offers a fundamentally simpler,
more robust, and integrated approach.

## The Traditional Microservice Experience

In typical microservice architectures (e.g., using gRPC, REST APIs), developers
face several inherent complexities:

1. **Connection Management:** Explicit steps are needed to establish connections
   (e.g., `grpc.Dial`, HTTP client setup). These connections can fail due to
   network issues, configuration errors, or service unavailability, requiring 
   dedicated error handling.
2. **Network Failures:** A call to a remote service can fail *not only* because
   the service's logic encountered an error, but also due to transient network
   problems, timeouts, or infrastructure issues. Distinguishing between these
   failure modes adds complexity.
3. **Data Serialization:** Data exchanged between services must often be
   explicitly serialized and deserialized into primitive types or specific
   formats (e.g., JSON, Protobuf). Passing complex objects or pointers directly
   is generally not possible.
4. **Manual Transaction Management:** Ensuring atomicity across multiple service
   calls is notoriously difficult. If a sequence of calls needs to succeed or
   fail together, developers must implement complex rollback or compensation
   logic. For example, if Service A succeeds but a subsequent call to Service B
   fails, the developer is responsible for:
    * Rolling back local state changes.
    * Potentially calling Service A again with a compensating action to undo its
      previous operation.
    This manual coordination is error-prone and can lead to inconsistent states,
    often requiring complex distributed locking mechanisms.
5. **Language/Protocol Proliferation:** Developers typically need to master the
   primary programming language *plus* various data formats (JSON, Protobuf,
   XML) and potentially interface definition languages (IDLs).

```ascii
 Traditional Microservices:

 [Your App] ------> Connect() -----> [Network] ------> [Service A]
     |                 |                  |                 |
     |---< Fail? <-----|                  |                 |
     |                 |                  |                 |
     +-----> Call(primitive) -> Marshal -> + -> Unmarshal -> +---> Logic A ---+
     |                 |                  |                 |                 |
     |<---- Fail? <----|<---- Fail? <-----|                 |<--- Logic OK ---+
     |                 |                  |                 |
     +-----> Call(primitive) -> Marshal -> + -> Unmarshal -> [Service B] ---> Logic B ---+
     |                 |                  |                 |                            |
     |<---- Fail? <----|<---- Fail? <-----|                 |<--------- Logic FAIL! -----+
     |                                                      |
     +---> // Now what? Manually rollback state?            |
     |     // Call Service A again to compensate?           |
     +------------------------------------------------------+
```

## The Gno Interrealm Experience

Gno's interrealm calls, facilitated by the GnoVM, provide a seamless and
powerful alternative, akin to an idealized microservice architecture:

1. **Transparent Connection:** The "connection" between realms is managed
   implicitly by the GnoVM. There's no separate connection step to manage or
   fail.
2. **Reliable Calls (Logic-Centric Failures):** Interrealm calls abstract away
   network failures. A call fails *only* if the logic within the called realm's
   function panics or returns an error. The underlying transport is guaranteed
   by the VM and the consensus layer.
3. **Rich Data Types:** Gno allows passing complex data types, including slices,
   maps, structs, and even pointers (with specific rules), directly between
   realms without manual serialization/deserialization.
4. **Automatic Atomic Transactions:** This is a major advantage. All operations
   within a Gno transaction, including multiple interrealm calls across
   different realms, are atomic. If *any* part of the transaction fails (e.g.,
   a panic in one of the called realms), the *entire* transaction is
   automatically rolled back by the GnoVM. There is no need for manual
   compensation logic; consistency is guaranteed at the VM level.
5. **Unified Language:** Developers only need to know Gno. The language used for
   writing realm logic is the same language used for making interrealm calls.
   No separate IDLs or data formats are required.

```ascii
 Gno Interrealm Call:

 [Realm A] ----> call_other_realm(complex_type) ----> [GnoVM Manages Interaction] ----> [Realm B]
     |                     ^                                                            |
     |                     |                                                            |
     |                     +-------------------- Logic OK ------------------------------+
     |                                                                                  |
     +----------------------------------- Logic OK -------------------------------------+

 // OR, if Realm B fails:

 [Realm A] ----> call_other_realm(complex_type) ----> [GnoVM Manages Interaction] ----> [Realm B]
     |                     ^                                                            |
     |                     |                                                            X<--- Logic Panic!
     |                     +--------------------- Panic Bubbles ------------------------+
     |
     X<----------- GnoVM Automatically Rolls Back ENTIRE Transaction -------------------+

```

In essence, Gno provides the "best microservice experience" imaginable: simple,
highly composable, robust error handling, and automatic transactional
consistency, all within a single language and runtime environment. This opens
the door to powerful new patterns of code reuse and composability directly on
the blockchain.

## Understanding `cross` and `crossing`

While basic interrealm calls provide access to another realm's functions and
data (read-only), the optional `cross` keyword and `crossing()` function
modifier introduce control over the execution *context*.

* **Calling without `cross` (Non-crossing call):** When you call a function in
  another realm *without* using `cross`, you are essentially executing that
  remote code within your *current* realm's context and memory space. Direct
  modification of the remote realm's state is restricted. Think of it like
  borrowing a library from another service but running it locally. Trying to
  modify memory of another realm without `cross` is forbidden and will result in
  a panic. For example, if Realm A calls a function in Realm B that attempts
  `realmB.someState = "new value"`, this operation will fail if not done within
  a `crossing()` call.

    ```ascii
     Non-crossing Call (Realm A calls Realm B's function):

     [Realm A Context]------------+
     |                            |
     | Executes Realm B's code    |
     | (within Realm A's context) |
     |                            |
     +----------------------------+
                       |
                       | Reads data from
                       V
                    [Realm B State] (Read-Only Access)
    ```

* **Calling with `cross` (Crossing call):** When you call a `crossing()`
  function using the `cross(realmB.Function)(...)` syntax, you explicitly shift
  the execution context *into* the target realm (Realm B). The function now
  executes with Realm B's permissions, environment (`std.CurrentRealm()`
  changes), and can directly modify Realm B's state. This is like truly entering
  the other microservice's environment to perform an operation; Realm A is
  allowed to modify Realm B's memory, but this modification happens within Realm
  B's context.

    ```ascii
     Crossing Call (Realm A calls Realm B's crossing function):

     [Realm A Context] --cross()--> [Realm B Context]-----------+
                                    |                           |
                                    | Executes Realm B's code   |
                                    | (within Realm B's context)|
                                    | Can modify Realm B State  |
                                    |                           |
     [Realm A State] <----return----+- std.CurrentRealm() is B -+
                                    |                           |
                                    +---------------------------+
    ```

A common question is: What if Realm A crosses into Realm B, and then Realm B
attempts to cross back into Realm A within the same transaction? While a direct
circular import between realm packages is disallowed at compile time, a circular
*call* sequence (A -> B -> A) during a crossing call is more nuanced. When B
calls back into A, the context shifts *again*. It's not as if A never crossed;
rather, it's a nested context change. The execution is now in Realm A's context,
but this is a *new* context initiated by Realm B. `std.CurrentRealm()` would
reflect Realm A, and modifications would apply to Realm A's state. However, the
overall transaction's atomicity is still managed by the GnoVM. If any part of
this A -> B -> A sequence fails, the entire set of operations across all
involved realms is rolled back.