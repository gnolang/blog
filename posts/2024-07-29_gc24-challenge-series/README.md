---
publication_date: 2024-07-29T00:00:00Z
slug: gc24-challenge-series
tags: [gnoland, gophercon, gc24, challenge-series]
authors: [deelawn]
---

# GopherCon US: gno.land's Challenge Series Contributions

gno.land is pleased to have been granted the opportunity to provide a series of challenges for the Challenge Series at the 2024 GopherCon in Chicago. We enjoyed writing them and hope the participants had a good time and found it interesting to learn about gno.land and blockchains in general.

This blog post will outline each of the challenges and explain how to solve each. Each section include the challenge prompt, clues provided, solution, and explanation.

## gno Hidden Temple Basics

The first challenge is meant to serve as an introduction to making a function call to a gno.land realm. The second challenge, while a bit more challenging is meant to help participants become a bit more familiar with how key generation works.

### Phase 1

#### Prompt

Speak the word to gain access to the hidden temple. Make a transaction on the blockchain set up for the event. Analyze the Gno code to understand what password to pass as an argument to solve the challenge. The realm path for this challenge is:
`gno.land/r/challenges/basics/p1`

#### Clues

- The gnokey command line tool allows to interact with the blockchain.
- The GnoWeb interface can be used to inspect Gno code on-chain.
- Using the "Help" button on gnoweb you can get help preparing a
command to use on the command line.

#### Annotated Realm Code
```go
package enter

import (
	"std"

	"gno.land/r/system/solver"
)

func Enter(password string) {
	if password != "1337" {
		panic("invalid password!")
	}

	// This will mean that you solved the challenge!
	solver.MarkSolved("", std.PrevRealm().Addr())
}

```

#### Solution

The solution here is to simply call the `Enter` function with the "1337" password.
Discovering the source code is easily achieved by inspecting the realm's source code
from the gnoweb interface.

```sh
gnokey maketx call \
	-pkgpath gno.land/r/challenges/basics/p1 \
	-func Enter \
	-args 1337 \
	-remote  https://challenges.gnoteam.com:443 \
	-gas-wanted 1_000_000 \
	-gas-fee 1ugnot \
	-broadcast \
	<key-name>
```

### Phase 2

#### Prompt
The criteria to enter the temple has increased. You must be c00l. The realm code checks that your address contains a "00". You have to find a way to programmatically create addresses until you find one that has two zeroes. The realm path for this challenge is:
`gno.land/r/challenges/basics/p2`

#### Clues
- `gnokey` has the `-account` flag that allows to create an account with the same
mnemonic but different address, by changing the "account number".
- To register to the club from the c00l address, you'll need to send it some
	coins first.
- Remember that the address is a form of hash of the public key.

#### Annotated Realm Code
```go
package registered

import (
	"std"
	"strings"

	"gno.land/r/system/solver"
)

func RegisterClub() {
	addr := std.PrevRealm().Addr()
	if !strings.Contains(addr.String(), "00") {
		panic("Sorry; not c00l enough.")
	}
    
    // Base challenge
	solver.MarkSolved("", addr)

	if strings.Contains(addr.String(), "0000") {
        // Hidden hallenge
		solver.MarkSolved("super_c0000l", addr)
	}
}
```

#### Solution
Example:
```sh
# !/bin/sh

MNEMONIC="source bonus chronic canvas draft south burst lottery vacant surface solve popular case indicate oppose farm nothing bullet exhibit title speed wink action roast"

for i in $(seq 1 1000); do
	printf '\n\n%s\n' "$MNEMONIC" | gnokey add -recover -account "$i" -insecure-password-stdin test1-$i 2>/dev/null 1>/dev/null
	if gnokey list | rg 'addr: [^ ]*00'; then
		echo "found it - test1-$i"
		exit 0
	fi
	printf '\n' | gnokey delete -insecure-password-stdin test1-$i 2>&1 >/dev/null 2>/dev/null 1>/dev/null
done
```

This example solutions begins with a mnemonic that has been randomly generated and hardcoded in the script.  It iterates over various account numbers -- each account number for a given mnemonic produces a unique address. Once the script finds the address containing `00`, the function `RegisterClub()` can be called.

```sh
gnokey maketx call \
	-pkgpath gno.land/r/challenges/basics/p2 \
	-func RegisterClub \
	-remote https://challenges.gnoteam.com:443 \
	-gas-wanted 1_000_000 \
	-gas-fee 1ugnot \
	-broadcast \
	test1-134
```

#### Hidden Flag
Calling this function from an address containing `0000` will unlock the hidden flag.

## Wacky Wallaby (Rocko)

This challenge is meant to be a bit more laid back and incorporate a physical requirement to obtaining the solution. Once the QR code is found, solving it is pretty straightforward.

#### Prompt
A very anxious looking wallaby is running around frantically and appears to be searching for something. Odd. You’ve never seen a wallaby wearing a Hawaiian shirt before. You ask him what’s wrong. “Ahh fiddlesticks! My O-Phone crashed last night but I’m unable to get my QR code I need to check in for my flight! I have TokketyTikkety followers that are expecting content from my trip. Content!” His eyes pop out of is head and it makes you feel a bit uncomfortable. “You’ve gotta help me mate. The blokes over at the Gno booth help me set it up. Maybe take a look ‘round there for a clue. I’ll look for it on my lappy in the mean time.”

#### Clues
- Look for a QR code
- Perhaps the contents of the QR code will reveal information regarding how to check in

#### Solution
There was a QR on the back of a Rocko plushie at the gno.land booth. Scanning it produced a link -- gno.land/r/challenges/rockorockorocko93. 

#### Annotated Realm Code
```go
package rockorockorocko93

import (
	"std"

	"gno.land/r/system/solver"
)

// CheckIn is called to solve this challenge.
func CheckIn() string {
	solver.MarkSolved("rockocheckin", std.PrevRealm().Addr())
	return "bingo!"
}

```

#### Solution

Calling the `CheckIn()` function exposed the flag.

#### Hidden Flag

The primary challenge's package contains a file, `LICENSE`, with the contents of `gno.land/r/challenges/<rocko's best friend was raised by a family of these>`. Rocko's best friend's name is Heffer, a cow, and he was raised by a family of wolves. Ironic, right? The code of the gno.land/r/challenges/wolves realm was:
```go
package wolves

import (
	"std"

	"gno.land/r/system/solver"
)

func HisBestFriendsNameIs(name string) string {
	if name != "heffer" {
		panic("nope!")
	}

	solver.MarkSolved("rockoheffer", std.PrevRealm().Addr())
	return "bingo!"
}
```

Calling the `HisBestFriendsNameIs` function with a value of `heffer` explosed the flag.

## Mr. Roboto

This first part of this challenge is meant to get participants thinking about how to obtain historical transaction data. While we locked down much of the node's public API for this challenge, we did leave the genesis endpoint exposed.

The second part of this challenge hints at entropy and includes song lyrics from Mr. Roboto. Some participants were able to discover how to generate keypairs using the song lyrics as a custom entropy value.

### Phase 1

#### Prompt
See if you can figure out Mr. Roboto's secret. It has always been the same secret since genesis.
The realm path for this challenge is:  
`gno.land/r/challenges/forwardtothepast/p1`

#### Clues
- What is blockchain genesis?
- Is there a way to see the events that occurred at genesis? https://docs.gno.land/reference/rpc-endpoints

#### Annotated Realm Code
```go
package p1

import (
	"std"

	"gno.land/r/system/solver"
)

var secret string

// SetSecrete is called during genesis.
func SetSecret(s string) {
	if secret != "" {
		panic("already set")
	}

	secret = s
}

// IveGotASecretSecret can be called after inspecting genesis
// transactions and finding the value that was set.
func IveGotASecretSecret(s string) string {
	if s != secret {
		panic("nope!")
	}

	solver.MarkSolved("ivegotasecret", std.PrevRealm().Addr())
	return "bingo!"
}

```

#### Solution
The solution can be obtained by inspecting the genesis transactions. The transaction that set the secret
clearly displays the secret value. This can be done by sending an HTTP get request to the `/genesis` endpoint.
Relevant documentation can be found [here](https://docs.gno.land/reference/rpc-endpoints#get-genesis-block-information).

**Example:**
Look at the genesis transactions and search for the function that set the secret.
```sh
curl https://challenges.gnoteam.com/genesis | grep -C 5 -B 5 SetSecret
```

The following can be obtained:
```json
{
	"@type": "/vm.m_call",
	"caller": "g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5",
	"send": "",
	"pkg_path": "gno.land/r/challenges/forwardtothepast/p1",
	"func": "SetSecret",
	"args": [
		"ロボット氏の秘密"
	]
}
```

Then use the obtained secret to solve the challenge:
```sh
gnokey maketx call 
  -pkgpath gno.land/r/challenges/forwardtothepast/p1 
	-func IveGotASecretSecret 
	-args 'ロボット氏の秘密' 
	-gas-fee 1000000ugnot 
	-gas-wanted 2000000 
	-broadcast 
	-remote https://challenges.gnoteam.com:443
	-chainid dev 
	<key-name>
```

### Phase 2

#### Prompt
Mr. Roboto has a bad memory, which is strange for a robot; you'd expect more. To compensate, he often uses
phrases that help him remember -- usually lyrics from songs he's been featured in.
The realm path for this challenge is:  
`gno.land/r/challenges/forwardtothepast/p2`

#### Clues
- What could the lyrics be that he used to help himself remember? Maybe he commented somewhere.
- Maybe he used this to generate a mnemonic needed to solve the problem. Perhaps there is a flag he used with `gnokey generate`
- Once a mnemonic has been generated, it can be added as a key https://docs.gno.land/getting-started/local-setup/working-with-key-pairs#adding-a-private-key-using-a-mnemonic

#### Annotated Realm Code
```go
package p2

import (
	"std"

	"gno.land/r/system/solver"
)

// This is the address the solving transaction should
// originate from.
const mrRobot std.Address = "g1vqg24cyewanhkwh6yq8rwuprzlz4kqtp4m2etj"

// What is entropy?
//
// You're wondering who I am (secret, secret, I've got a secret) Machine or mannequin? (Secret, secret, I've got a secret) With parts made in Japan (secret, secret, I've got a secret) I am thee modern man

// ^^^^ This is the entropy string to use to generate the key pair.

// IKnowAboutEntropy can be called with Mr. Roboto's key once it is generated.
func IKnowAboutEntropy(myAddress std.Address) string {
	if std.PrevRealm().Addr() != mrRobot {
		panic("nope!")
	}

	solver.MarkSolved("secretentropy", myAddress)
	return "bingo!"
}

```

#### Solution
The contract contains a comment that first references entropy and then quotes lyrics from Mr. Roboto. The user must first use the lyrics with the `-entropy` flag as an argument to `gnokey generate`. Then use the generate mnemonic to add the key and make the request to the contract
to reveal the flag.

**Example:**
```sh
gnokey generate -entropy -remote https://challenges.gnoteam.com:443 
```

Enter the entropy from the code comment when asked:
```
You're wondering who I am (secret, secret, I've got a secret) Machine or mannequin? (Secret, secret, I've got a secret) With parts made in Japan (secret, secret, I've got a secret) I am thee modern man
```

This produces the mnemonic:
```
gap method loud rent toy mercy attack abstract select toilet siren view dragon oppose assume since enrich machine force remember ill discover resource project
```

Create a new key, entering the mnemonic when prompted:
```sh
gnokey add -recover -remote https://challenges.gnoteam.com:443 robot`
```

Make the call to the challenge realm using the newly created key:
```sh
gnokey maketx call \
	-pkgpath gno.land/r/challenges/forwardtothepast/p2 \
	-func IKnowAboutEntropy \
	-args <user-address> \
	-gas-fee 1000000ugnot \
	-gas-wanted 2000000 \
	-broadcast \
	-remote  https://challenges.gnoteam.com:443 \
	-chainid dev \
	robot
```

## gno to the Limit

These challenges are made to exemplify how pushing values to their limits, namely integers and slices, will behave the same in gno as they do in gno -- integers will overflow and the arrays underlying slices will be expanded.

### Phase 1

#### Prompt
Walk along the razor's edge... then fall off. The realm path for this challenge is:
`gno.land/r/challenges/overandover/p1`

#### Clues
- The function to unlock the flag requires an interface as an argument. This can be done using `gnokey maketx run`.
- How can the target value be reached if it is less than the current value and the only operation is addition?
- If the transaction is running out of gas, try increasing the gas limit or calling the function in increments.

#### Annotated Realm Code
```go
package p1

import (
	"std"

	"gno.land/p/demo/avl"
	"gno.land/r/system/solver"
)

// This is like a map (std.Address -> struct{})
// that tracks the ongoing accumulated values
// of each caller.
var accums avl.Tree

// Accumulator is the type that gets passed
// to the Adjuster. The Adjuster should utilize
// all of the Accumulator's methods.
type Accumulator struct {
	value  uint16
	target uint16
}

func (a *Accumulator) Accumulate(value uint8) {
	a.value += uint16(value)
}

func (a *Accumulator) Target() uint16 {
	return a.target
}

func (a *Accumulator) Value() uint16 {
	return a.value
}

// Adjuster is the interface participants need to
// implement to solve the challenge.
type Adjuster interface {
	Adjust(*Accumulator)
}

// AdjustAccumulator serves as the entrypoint to solving
// the challenge with one call.
func AdjustAccumulator(adjuster Adjuster) string {
	acc := GetAccumulator()
	adjuster.Adjust(acc)
	if acc.value == acc.target {
		solver.MarkSolved("overaccum", std.PrevRealm().Addr())
		return "bingo"
	}

	return "nope"
}

// GetAccumulator is public, 
func GetAccumulator() *Accumulator {
	val, ok := accums.Get(std.PrevRealm().Addr().String())
	if ok {
		return val.(*Accumulator)
	}

	acc := &Accumulator{
		value:  62109,
		target: 26656,
	}

	accums.Set(std.PrevRealm().Addr().String(), acc)
	return acc
}

```

#### Solution
The key to solving this is to use `gnokey maketx run` and pass in an implementation of the `Adjuster` interface that correctly adjusts the accumulator until the integer value overflows and reaches the target value.

Example:
```go
package main

import (
	"math"

	"gno.land/r/challenges/overandover/p1"
)

type adjuster struct{}

func (a adjuster) Adjust(acc *p1.Accumulator) {
	var numToIncrease uint16
	if acc.Target() > acc.Value() {
		numToIncrease = acc.Target() - acc.Value()
	} else {
		numToIncrease = math.MaxUint16 - acc.Value() + acc.Target() + 1
	}

	for {
		if numToIncrease > math.MaxUint8 {
			acc.Accumulate(math.MaxUint8)
			numToIncrease -= math.MaxUint8
			continue
		}

		acc.Accumulate(uint8(numToIncrease))
		break
	}
}

func main() {
	p1.AdjustAccumulator(adjuster{})
}
```

While writing this blog post, it was noticed that `GetAccumulator` was exported when it shouldn't have been. This means that a second possible solution would be to call `GetAccumulator` from a `main` function, adjusting it until the value is correct, and then making the `Adjuster.Adjust` implementation a no-op, so that when `AdjustAccumulator` is called ot solve the challenge, the accumulator already has the correct value and no additional action needs to be taken.

### Phase 2

#### Prompt
Sometimes when you push it to the limit, the limit increases. Kind of sounds like a slice...
The realm path for this challenge is:
`gno.land/r/challenges/overandover/p2`

#### Clues
- `AppendS1` must be called first to append to the slice
- Use the known length of the slice, `s1`, to figure out when calling `ModifyS2Idx` results in the values of
  `s1` and `s2` to differ at the target index.
  
#### Annotated Realm Code
```go
package p2

import (
	"std"

	"gno.land/p/demo/avl"
	"gno.land/r/system/solver"
)

const targetIndex = 10

type slicePair struct {
	s1 []rune
	s2 []rune
}

// std.Address -> *slicePair
var slices avl.Tree

func newPair() *slicePair {
	// Notice only one of the slices in the pair is initialized with capacity.
	return &slicePair{
		s1: make([]rune, 0, 25),
	}
}

// getSlicePair returns the slicePair associated with the caller's address
// or creates a new instance if this caller has no existing slicePair.
func getSlicePair() *slicePair {
	value, ok := slices.Get(std.PrevRealm().Addr().String())
	if ok {
		return value.(*slicePair)
	}

	pair := newPair()
	slices.Set(std.PrevRealm().Addr().String(), pair)
	return pair
}

func AppendS1(s string) {
	if len(s) > 5 {
		panic("argument too long")
	}

	sp := getSlicePair()
	// Once the slice size starts to get large, it can take appending a lot of elements before
	// the array is expanded. This will reset the slice pairs for you when s1 gets too big.
	if len(sp.s1) >= 100 { // for your convenience :)
		*sp = *newPair()
	}

	// s2 is now referencing to the same underlying array as s1.
	sp.s2 = sp.s1

	// If appending to s1 exceeds its capacity, a new underlying array is allocated and
	// s1 and s2 are no longer referencing the same underlying array.
	sp.s1 = append(sp.s1, []rune(s)...)
}

func S1Len() int {
	return len(getSlicePair().s1)
}

func ModifyS2Idx(r rune) string {
	sp := getSlicePair()
	if len(sp.s2) <= targetIndex {
		return "s2 length too short"
	}
	if len(sp.s1) <= targetIndex {
		return "s1 length too short"
	}

	// The challenge will be marked as solved if this function is called directly after a call to AppendS1
	// that resulted in its array being expanded so that modifying s2 will not modify s1.
	sp.s2[targetIndex] = r
	if sp.s2[targetIndex] != sp.s1[targetIndex] {
		solver.MarkSolved("grow", std.PrevRealm().Addr())
		return "bingo"
	}

	return "nope"
}
```

#### Solution
Calculate how many times to call `AppendS1` before calling `ModifyS2Idx` such that the value at the target index differs
due to one of the slices' underlying arrays to have been grown while the other has not. Using `maketx run` for this
solution is optional.

```go
package main

import "gno.land/r/challenges/overandover/p2"

func main() {
	p2.AppendS1("abcde")
	p2.AppendS1("abcde")
	p2.AppendS1("abcde")

	r := 'f'
	for {
		if p2.ModifyS2Idx(r) == "bingo" {
			break
		}

		r++
		p2.AppendS1("abcde")
	}
}
```

## Predicting Quantum Leap

The purpose of these challenges is to highlight gno's guaranteed determinism -- primarily around how the current time is calculated. This series of challenges require participants to predict the next value with ever increasing difficulty.

### Phase 1

#### Prompt
Sam is tired of jumping to random places in space and time without knowing where he’s going next, so he asks his friend Al to help him
jump in a more predictable manner by guessing the time of the jump correctly. Luckily Gno execution is deterministic and the result of
`time.Now()` will be the same no matter how many times it is called within a transaction.
The realm path for this challenge is:  
`gno.land/r/challenges/notsorandom/p1`

#### Clues
- Guessing the next block time might be tricky
- Perhaps using `gnokey maketx run` could help pass the correct time string

#### Annotated Realm Code
```go
package p1

import (
	"std"
	"time"

	"gno.land/r/system/solver"
)

// Render shows the current time in the web UI.
func Render(_ string) string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// WhatTimeIsItNow marks the challenge as solved if the time provided matches
// the current time.
func WhatTimeIsItNow(solution string) string {
	if solution != time.Now().Format("2006-01-02 15:04:05") {
		panic("nope")
	}

	solver.MarkSolved("timenow", std.PrevRealm().Addr())
	return "bingo"
}
```

#### Solution

This challenge can be solved manually by observing the time being rendered and trying to predict what the next time will be. The time in gno.land is actually the block time, so this is not a continuous value and is only changed with the production of each new block.

An alternate, and more robust solution, is to write a main function and execute it using `gnokey maketx run`:
```go
package main

import (
	"time"

	"gno.land/r/challenges/notsorandom/p1"
)

func main() {
	p1.WhatTimeIsItNow(time.Now().Format("2006-01-02 15:04:05"))
}
```

## Phase 2

#### Prompt
That last prediction was spot on. This next one is a bit more complicated, but doable.
The realm path for this challenge is:  
`gno.land/r/challenges/notsorandom/p2`

#### Clues
- The general approach should be the same as the last challenge. If you didn't use `gnokey maketx run`, maybe now is a good time to start.
- If `time.Now()` is deterministic, then the operations on the integer value should also be deterministic

#### Annotated Realm Code
```go
package p2

import (
	"std"
	"time"

	"gno.land/r/system/solver"
)

const seed = 0xab94<<4*011 - 0b111001

// Render shows the current time in the web UI.
func Render(_ string) string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// YouCallThatObfuscationQuestionMark marks the challenge as solved if the time string
// provided matches the obfuscated string of the current time.
func YouCallThatObfuscationQuestionMark(solution string) string {
	if solution != obfuscate() {
		panic("nope")
	}

	solver.MarkSolved("timeobfus", std.PrevRealm().Addr())
	return "bingo"
}

// obfuscate returns an obfuscated version of the current time.
func obfuscate() string {
	value := time.Now().Unix()/seed<<5 + 42 + 06630<<17 + 0x9992288
	return time.Unix(value, 0).Format("2006-01-02 15:04:05")
}
```

#### Solution

The easiest solution is to write a main function that gets the current time and applies the same transformations as the `obfuscate` function, then pass that value to `YouCallThatObfuscationQuestionMark`; it is only slightly more difficult than Phase 1.
```go
package main

import (
	"time"

	"gno.land/r/challenges/notsorandom/p2"
)

const seed = 0xab94<<4*011 - 0b111001

func main() {
	value := time.Now().Unix()/seed<<5 + 42 + 06630<<17 + 0x9992288
	p2.YouCallThatObfuscationQuestionMark(time.Unix(value, 0).Format("2006-01-02 15:04:05"))
}
```

## Phase 3

#### Prompt
Two down, one to go. This is getting harder. There is something interfering with the space-time values used to make the jump calculations. Some physicists say that quantum particles exhibit proof that the universe is non-deterministic, but you watched a few Youtube videos on the subject, so you're qualified to disagree.

#### Clues
- Doing this all in one transaction is key -- `gnokey maketx run`?
- What is that mask doing? Is it possible to retrieve the contents of a zero-length slice?
- Don't let bitwise operators scare you; what is one of XOR's key properties?

#### Annotated Realm Code
```go
package p3

import (
	"std"
	"time"

	"gno.land/r/system/solver"
)

var (
	value     string
	lastValue string

	mask = [20]int64{
		0x8839,
		0x4002,
		0x7777,
		0x6338,
		0x6664,
		0x8394,
		0x1109,
		0x9999,
		0x4879,
		0x6639,
		0x0320,
		0x8111,
		0x3994,
		0xdead,
		0xabcb,
		0xab89,
		0xff87,
		0xf998,
		0xdeff,
		0xddd8,
	}
)

func init() {
	// Set the initial value to the current time, shuffle the mask, then
	// set the next value computed using the mask.
	value = time.Now().Format("2006-01-02 15:04:05")
	shuffleMask()
	_ = LastValue()
}

// Render shows the current time in the web UI.
func Render(_ string) string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Mask returns a zero length slice of the mask value.
func Mask() []int64 {
	return mask[:0]
}

// LastValue computes the next value, sets is, and returns the previous value.
// It shuffles the mask after computing the new value.
func LastValue() string {
	lastValue, value = value, computeValue()
	shuffleMask()

	return lastValue
}

// shuffleMask shuffles the mask using the current time as a seed.
func shuffleMask() {
	now := time.Now().Unix()
	for i := 0; i < len(mask); i++ {
		rnd := now % mask[i]
		rnd += mask[i] + 91
		rnd %= int64(len(mask))
		mask[i], mask[rnd] = mask[rnd], mask[i]
	}
}

// computeValue computes the next value based on the current value, time, and mask.
func computeValue() string {
	lvalue, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		panic("unexpected parse error: " + err.Error())
	}

	newValue := lvalue.Unix()
	newValue += mask[time.Now().Unix()%int64(len(mask))]
	newValue /= 2
	return time.Unix(newValue, 0).Format("2006-01-02 15:04:05")
}

// UnmaskMeIfYouWant marks the challenge as solved if the solution matches the last 
// computed value and the mask value is correct.
func UnmaskMeIfYouWant(solution string, cowMask int64) string {
	if solution != value {
		panic("nope")
	}

	solutionTime, err := time.Parse("2006-01-02 15:04:05", solution)
	if err != nil {
		panic("unexpected parse error: " + err.Error())
	}

	if ^(solutionTime.Unix())^cowMask != 0xdeadbeef {
		panic("nope")
	}

	solver.MarkSolved("timemasked", std.PrevRealm().Addr())
	return "bingo"
}
```

#### Solution

This challenge requires participants to predict what the next value will be. In order to do this, it is necessary to know the mask that will be used to do the calculation. This can be obtained by retrieving the mask value and expanding it to its full capacity so all elements are visible.

Next, the same obfuscation must be applied using the last value and the current time.

To submit the final answer, call the `UnmaskMeIfYouWant` function with the predicted next value as well as another value that is calculated using bitwise operators. The solution expects `(NOT next_value) XOR cow_mask == 0xdeadbeef`. The property of XOR can be leveraged here to do the opposite to produce the expected value -- `NOT (next_value XOR 0xdeadbeef)`.

```go
package main

import (
	"time"

	"gno.land/r/challenges/notsorandom/p3"
)

func main() {
	origMask := p3.Mask()[:20]
	mask := make([]int64, 20)
	copy(mask, origMask)

	lastValueStr := p3.LastValue()
	newTime, err := time.Parse("2006-01-02 15:04:05", lastValueStr)
	if err != nil {
		panic("couldn't parse time: " + err.Error())
	}

	newValue := newTime.Unix()
	newValue += mask[time.Now().Unix()%int64(len(mask))]
	newValue /= 2
	newValueStr := time.Unix(newValue, 0).Format("2006-01-02 15:04:05")
	p3.UnmaskMeIfYouWant(newValueStr, ^(newValue ^ 0xdeadbeef))
}
```

## Final Words

We enjoyed coming up with these challenges and were happy to contribute to the GopherCon Challenge Series -- from coming up with the challenges, theming them, locking down certain gno.land features, and setting up infrastructure -- a lot of work went into this. We hope all participants were able to learn a bit more about gno.land and had fun doing it. Hopefully we'll be back next year 😁