---
publication_date: 2025-03-12T00:00:00Z
slug: user-registry
tags: [gnoland, usernames, namespaces, dapps]
authors: [leohhhn]
---

# Introducing the New gno.land User Registry

The user registration system has been around since the early days of gno.land, 
serving as one of the foundational components of the chain. It allows users to 
register custom names linked to their Gno addresses and claim matching namespaces 
for their Gno package deployments.

While the username registration system existed as far back as April 2022 ([`r/demo/users`](https://github.com/gnolang/gno/commit/914f267dd31c0382a472b5fcf98fcfc53129a32d))
the namespace permission deployment went into effect with Test4 for the first time
in July 2024.

With changes to the Gno language, chain infrastructure, as well as the addition
of governance and other new functionality, the user registry was due for a major
refactor. This new version introduces a set of realms that integrate directly 
with the core of the gno.land blockchain, providing similar but more flexible
functionality and upgradeability through GovDAO proposals.

## Register your gno.land username now!

We invite you to register your gno.land username with the new User Registry on
the Portal Loop network.

Here's how to register in a few steps:
- Create a Gno address with `gnokey` or a third-party wallet such as Adena
- Fetch Portal Loop GNOT from the [Faucet Hub](https://faucet.gno.land)
- Visit the [User Registry realm](/r/gnoland/users/v1) 
- If you're using Adena, you can use [GnoStudio Connect](https://gno.studio/connect/view/gno.land/r/gnoland/users/v1?network=portal-loop#Register)
to send a transaction via the web. 
  - Under advanced options, make sure to set the "send" amount to `1000000ugnot` 
  to pay for the registration fee.
- If you're using `gnokey`, click on "Register" to get a pre-built `gnokey` command
  - Make sure to also set `-send "1000000ugnot"` to pay for the registration fee
- Enter your desired username which follows the naming rules outlined in the 
User Registry realm.
- Call the function and see your name show up in the "Latest registrations" section!

> Make sure to specify the exact amount of GNOT for the registration fee (1 GNOT).

## System Architecture

For developers that are looking to use the new User Registry system in their 
realms, below is an outline of the main components of the system including a 
practical example found under `r/docs`. 

The new user registration system, paired with the namespace system, enables
custom usernames and namespace permission checking upon package deployment. It
consists of three key realms:

- [`r/sys/users`](/r/sys/users) - Core logic of the system, handling storage and resolution.
- [`r/gnolang/users/v1`](/r/gnoland/users/v1) - The user-facing application where users register names.
- [`r/sys/names`](/r/sys/names) - The deployment permission system for namespaces.

Let's dive into what each part of the system does.

### r/sys/users

This is a system realm that defines the core logic of the user registry. It is
not upgradeable and serves as the foundational layer that includes key data 
structures, storage variables, and functions for managing usernames.

Only whitelisted realms (approved via GovDAO proposals) can write to this storage,
ensuring governance oversight over any entity modifying the registry.

The `UserData` structure defines how user information is stored:

```go
type UserData struct {
    addr     std.Address  // User's blockchain address
    username string       // Latest registered username
    deleted  bool         // Flag for soft deletion
}
```

This realm also provides user-facing functions for resolving names and addresses.

#### Resolving Usernames & Addresses

To enable applications and smart contracts to look up user information, 
`r/sys/users` provides functions to resolve both usernames and addresses:

```go
// ResolveName returns the latest UserData of a specific user by name or alias
func ResolveName(name string) (data *UserData, isCurrent bool) {...}

// ResolveAddress returns the latest UserData of a specific user by address
func ResolveAddress(addr std.Address) *UserData {...}
```

See how to use these functions in a practical example [here](/r/docs/users), and 
check out the source code [here](/r/sys/users$source&file=users.gno).

### r/gnoland/users/v1

The first iteration of the gno.land user registry allows users to register a
name following a specific pattern - starts with three letters, ends with three
numbers, and is not longer than 20 characters. The registration costs 1 GNOT
and the fee is modifiable via a GovDAO proposal.

This is the user-facing realm where users interact with the registry. It 
provides an interface for registering usernames while enforcing certain constraints:

- Usernames must start with three letters.
- They must end with three numbers.
- The total length must not exceed 20 characters.
- Only `_` is allowed as a special character.
- The registration fee is 1 GNOT (modifiable via GovDAO proposals).

This realm communicates with `r/sys/users` to ensure new registrations follow the 
system-wide rules and that only valid usernames are stored. Future versions 
(v2, v3, etc.) will follow different patterns, such as vanity usernames, username
auctions, and more can be introduced without modifying the core storage logic.

Find all releases for `r/gnoland/users` [here](/r/gnoland/users).

### r/sys/names

The deployment permission system ensures that only the rightful owner of a 
namespace can deploy Gno packages under it.

This check is performed at the GnoVM Keeper level, which communicates with 
`r/sys/names` realm before allowing a deployment to a certain package patch.
This functionality is enabled in the genesis block, with an ownership-dropping 
mechanism ensuring no single entity retains control over the system.

Check out the source code [here](/r/sys/names$source&file=verifier.gno).