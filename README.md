TermLog v0.1.0
==============

Simple level based terminal logger


Features
--------

* Conditional Level Logging: uses a different level depending if the condition is `true` or `false`
* Custom datetime formatting
* Custom Fatal level exit code
* Custom Panic level exit code
* Forcing ANSI color on or off for level labels
* Logger namespacing
* Outputs to `stderr` by default but allows assignment of any `io.Writter`


Usage Examples
--------------

```go
package main

import "github.com/runeimp/termlog"

func main() {
	log := termlog.New()

	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
}
```


### Conditional Level Logging

```go
package main

import "github.com/runeimp/termlog"

func main() {
	value := 0

	log := termlog.New()
	conditionTrueLevel := termlog.WarnLevel
	conditionFalseLevel := termlog.DebugLevel
	log.ConditionalLevel(value == 0, conditionTrueLevel, conditionFalseLevel, "If the value equals zero pay attention!")
	// Output: 2023-02-06 19:20:14.192617 WARN  If the value equals zero pay attention!
}
```

### If You Wish to Control All Things and Use a Namespace

```go
package main

import "github.com/runeimp/termlog"

func main() {
	log := termlog.New("my-namespace")
	log.ForceColor = termlog.ForceColorOff // Force off ANSI color rendering on supported systems
	log.Level = termlog.WarnLevel          // TermLog defaults to DebugLevel
	log.TimeFormat = "2006-01-02 15:04:05" // Default format but with no fractional seconds
	log.FatalExitCode = 13
	log.PanicExitCode = 42

	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
	log.Fatal("Fatal message")
	log.Panic("Panic message")
}
```


### You Can Force Color if You Desire

If You Run Windows and don't use MS Terminal you can force TermLog to output ANSI color codes. If your system already has color support turned on it will simply work. If not you'll see the ANSI codes as text. ANSI color support should be active by default for later versions of Windows 10, all versions of Windows 11 and on. I just haven't built in testing for that support yet. So you have to force it to see if it's supported or not and to use it outside of MS Terminal currently.

```go
package main

import "github.com/runeimp/termlog"

func main() {
	log := termlog.New()
	log.ForceColor = termlog.ForceColorOn

	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
	log.Fatal("Fatal message")
	log.Panic("Panic message")
}
```


Example Output
--------------

```sh
$ tester; echo "Exit Code: $?"
==> Testing Logger Up to Error Level With Info Level Set and ForceColor is ForceColorOff

2023-02-06 19:20:14.192486 INFO  Informational message
2023-02-06 19:20:14.192594 WARN  Warning message
2023-02-06 19:20:14.192596 ERROR Error message

==> Testing Namespaced Logger with Custom TimeFormat Up to Error Level With Default Settings

2023-02-06 19:20:14 my-namespace DEBUG Debugging message
2023-02-06 19:20:14 my-namespace INFO  Informational message
2023-02-06 19:20:14 my-namespace WARN  Warning message
2023-02-06 19:20:14 my-namespace ERROR Error message

==> Testing Logger with the ConditionalLevel method

2023-02-06 19:20:14.192609 INFO  This is condition: true  # Should be Info Log Level in this test
2023-02-06 19:20:14.192617 WARN  This is condition: false # Should be Warn Log Level in this test

==> Testing Logger Up to Panic Level With Custom Exit Codes and ForceColor is ForceColorOn

2023-02-06 19:20:14.192621 DEBUG Debugging message
2023-02-06 19:20:14.192623 INFO  Informational message
2023-02-06 19:20:14.192624 WARN  Warning message
2023-02-06 19:20:14.192626 ERROR Error message
2023-02-06 19:20:14.192627 FATAL Fatal message
exit status 13
```

Above represents the output when the library is used on an unsupported system i.e.: Windows without using MS Terminal or MSYS (Git Bash, MinGW, MinGW-64, etc.). On supported systems (everything else I've tested thus far) the level colors <span style="background: black; color: white; padding: 0.2em">(<span style="color:silver; font-weight:bold;">DEBUG</span>, <span style="color:cyan; font-weight:bold;">INFO</span>, <span style="color:yellow; font-weight:bold;">WARN</span>, <span style="color:red; font-weight:bold;">ERROR</span>, <span style="color:red; font-weight:bold;">FATAL</span>, and <span style="color:red; font-weight:bold;">PANIC</span>)</span> are displayed on non-Windows systems and on Windows when using MS Terminal or ANSI colors are forced on. I may eventually add testing for color support for PowerShell, etc. on Windows when used outside of MS Terminal as well. For now you can test this with `ForceColor` as noted above in the Usage Examples.




[MinGW 32/64 enable ANSI color sequences on Windows 10]: https://gist.github.com/fleroviux/8343879d95a72140274535dc207f467d

