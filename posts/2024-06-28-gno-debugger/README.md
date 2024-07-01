---
title: "Debugging Gno Programs"
publication_date: 2024-06-28T13:37:00Z
slug: gno-debugger-blog-post
tags: [blog, post, tutorial, gno, debugger]
authors: [mvertes]
---
# Debugging Gno Programs
 
In this article, we introduce the new Gno debugger feature and show how it can be used to better understand gno programs and help fixing bugs.

## Motivation for a Gno debugger

> Debugging is twice as hard as writing code.
> Brian Kerninghan, "The Elements of Programming Style"

> On average, you spend about eight to ten times debugging as you do writing code.
> Anonymous

Having a good debugger is important. But the Gno language is almost Go, and gno.land itself is entirely written in Go. Could I just use the existing Go tools, i.e.  the [delve] debugger, to take control and debug my Gno programs?

You cannot debug your *Gno* program this way because doing so would entail debugging the Gno virtual machine rather than your own program. The relevant state information would be opaque and would need to be reversed and reconstructed from internal Gno virtual machine data structures.

The Gno debugger addresses this issue by displaying the state of the Gno program memory symbolically. It allows for control of program execution at the source code level, regardless of the virtual machine implementation.

## Setting up

The Gno debugger is fully integrated in the [gno] single binary, which is the tool required to build, test and run gno programs locally.

There is no need to install a specific tool. You just have to install gno itself by:

```shell=
git clone https://github.com/gnolang/gno
cd gno
go install ./gnovm/cmd/gno
```

We are now ready to play with gno programs. Let's consider a simple classical example as a target program, the computation of [Fibonacci numbers]:

```go=
// fib.gno
package main

// fib returns the nth number in the Fibonacci sequence.
func fib(n int) int {
        if n < 2 {
                return n
        }
        return fib(n-2) + fib(n-1)
}

func main() {
        println(fib(4))
}

```
To execute this program, we run the command `gno run ./fib.gno`. To activate the debugger, we just pass the `-debug` flag: `gno run -debug ./fib.gno`. Use `gno run -help` to get more options if needed.

## Quick tour of the debugger

When you start a program under debug, you are greeted by a prompt allowing you to interact through the terminal:
```shell=
$ gno run -debug ./fib.gno
Welcome to the GnoVM debugger. type 'help' for list of commands.
dbg>
```

Entering `help` gives you the list of available commands and their short usage:

```shell
dbg> help
The following commands are available:

break|b [locspec]         Set a breakpoint.
breakpoints|bp            Print out info for active breakpoints.
clear [id]                Delete breakpoint (all if no id).
continue|c                Run until breakpoint or program termination.
detach                    Close debugger and resume program.
down [n]                  Move the current frame down by n (default 1).
exit|quit|q               Exit the debugger and program.
help|h [command]          Print the help message.
list|l [locspec]          Show source code.
print|p <expression>      Print a variable or expression.
stack|bt                  Print stack trace.
step|s                    Single step through program.
stepi|si                  Single step a single VM instruction.
up [n]                    Move the current frame up by n (default 1).

Type help followed by a command for full documentation.
dbg>
```

If you have already used a debugger before, like [gdb] of [lldb] for C/C++ programs, or [delve] for Go programs, the Gno debugger should look familiar: the commands are similar in their syntax and their use.

Those commands can be classified in the following categories:
- managing breakpoints: `break`, `breakpoints`, `clear`,
- controlling execution: `step`, `stepi`, `continue`,
- browse code, data and stack: `list`, `print`, `stack`,
- navigate the stack: `up`, `down`,
- quit the debugger: `detach`, `exit`.

## Controlling and exploring the program state

Let's go back to our Fibonacci program, still paused. We `step` a first time, which instructs the GnoVM to execute a single statement and give back control to the user:

```shell
dbg> s
> main.main() main/./fib.gno:11:1
      7: 		return n
      8: 	}
      9: 	return fib(n-2) + fib(n-1)
     10: }
     11: 
=>   12: func main() {
     13: 	println(fib(4))
     14: }
```

The first output line `> main.main() main/./fib.gno:11:1` indicates the precise current location in source code, followed by a short source listing around this location. The current line is indicated by the cursor `=>`.

From there, we could repeat `step` commands to progress, but that would be too tedious. Instead, we set a breakpoint at an interesting line in the `fib` function, and `continue` to it directly:

```shell
dbg> b 7
Breakpoint 0 at main main/./fib.gno:7:1
dbg> c
> main.fib() main/./fib.gno:7:10
      2: package main
      3: 
      4: // fib returns the nth number in the Fibonacci sequence.
      5: func fib(n int) int {
      6: 	if n < 2 {
=>    7: 		return n
      8: 	}
      9: 	return fib(n-2) + fib(n-1)
     10: }
     11: 
dbg>
```

Note that we have used the short alias of commands: `b` for `break` and `c` for `continue`. We could also just provide the target line number instead of the full location including the file name, because the stay in the same source file.

We can now examine the caller stack which indicates the successive nested function calls up to the current location:

```shell
dbg> stack
0	in main.fib
	at main/./fib.gno:7:10
1	in main.fib
	at main/./fib.gno:9:20
2	in main.fib
	at main/./fib.gno:9:20
3	in main.main
	at main/./fib.gno:13:2
dbg>
```
We see a call stack of depth 4, with call frames (local function contexts) numbered from 0 to 3, 0 being the current call level (the deepest). This information is crucial especially when debugging recursive functions like `fib`. We know that the caller and its caller were both `fib`.

We want now to examine the value of the local parameter `n`, for each call level:
```shell!
dbg> print n
(0 int)
dbg> up
> main.fib() main/./fib.gno:7:10
Frame 1: main/./fib.gno:9:20
      4: // fib returns the nth number in the Fibonacci sequence.
      5: func fib(n int) int {
      6: 	if n < 2 {
      7: 		return n
      8: 	}
=>    9: 	return fib(n-2) + fib(n-1)
     10: }
     11: 
     12: func main() {
     13: 	println(fib(4))
dbg> print n
(2 int)
dbg> up
> main.fib() main/./fib.gno:7:10
Frame 2: main/./fib.gno:9:20
      4: // fib returns the nth number in the Fibonacci sequence.
      5: func fib(n int) int {
      6: 	if n < 2 {
      7: 		return n
      8: 	}
=>    9: 	return fib(n-2) + fib(n-1)
     10: }
     11: 
     12: func main() {
     13: 	println(fib(4))
dbg> print n
(4 int)
dbg>
```
We see that the local value `n` is 0 at current frame 0, 2 at frame 1 and 4 at frame 2, which corresponds to the nested calls of `fib` expressed at line 9.

With the stack navigation commands `up` and `down` the debugger is able to display the value of local function variables and parameters for the whole call chain.

In this example, the `n` variable is simply an integer, but the `print` command is also able to handle more complex expressions to uncover the content of arbitrary maps, struct, arrays, etc using the same syntax as Go, for example as in `print a.b[n]` with a being a value of type:
```go=
var a struct {
    b []string
}
```
For security reasons, `print` command will only evaluate expressions with no side effects on the virtual machine state. For example it is not possible to perform an arithmetic operation like `print a + 2`, or to call a function like `printf f(6)`.

## Conclusion

We have introduced the new Gno debugger and presented its main capabilities. 

This is just the start of a new project, with a lot of room for improvement. The whole Gno project being open source, you are welcome not only to provide feedbacks and suggestions, but also to contribute at https://github.com/gnolang/gno.


[delve]: https://github.com/go-delve/delve
[gno]: https://github.com/gnolang/gno/tree/master/gnovm/cmd/gno
[Fibonacci numbers]: https://simple.wikipedia.org/wiki/Fibonacci_number
[gdb]: https://sourceware.org/gdb/
[lldb]: https://lldb.llvm.org
