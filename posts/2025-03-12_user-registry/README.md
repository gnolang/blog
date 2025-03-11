---
publication_date: 2025-03-12T00:00:00Z
slug: user-registry
tags: [gnoland, usernames, dapps]
authors: [leohhhn]
---

# The Revamped gno.land User Registry

The user registration system has been around since the beginnings of gno.land, serving
as one of the core components of the chain. It allowed users to register custom
names for their addresses and claim matching namespaces to which they can deploy 
Gno packages.

While the username registration system existed as far back as April 2022 ([r/demo/users](https://github.com/gnolang/gno/commit/914f267dd31c0382a472b5fcf98fcfc53129a32d))
the namespace permission deployment went into effect with Test4 for the first time
in July 2024. 

With all the changes to the language, to the chain, and many more functionalities,
the user registry deserved a full refactor. This refactor introduces a set of realms,
which integrate directly with the core of the gno.land blockchain providing the
aforementioned functionality, with full upgradeability via GovDAO proposals.

## System Architecture

The user registration system, paired with the namespace system, comprise the 
logic that allows custom usernames and 

`r/sys/users` - core logic of the system
`r/gnolang/users/vX` - user-facing application
`r/sys/names` - deployment permission system

Let's dive into what each part of the system does.

### `r/sys/users`

This is a system realm that defines the core of the user registry. This realm is
not upgradeable and acts as the core layer including the needed types, storage 
variables, and functions to modify the aforementioned storage. This includes 
creating new usernames, updating old usernames, and deleting usernames.

This realm is callable only by realms which are whitelisted via GovDAO proposals.
This makes sure that the only entity that has access to the username store is
strictly reviewed and voted in via governance.






## `r/gnoland/users/v1`

The first iteration of the gno.land user registry allows users to register a
name following a specific pattern - starts with three letters, ends with three
numbers, and is not longer than 20 characters. The registration costs 


