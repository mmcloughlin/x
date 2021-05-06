# profile

Experimental profiling support package for Go.

* Based on the widely-used [`pkg/profile`](https://github.com/pkg/profile):
  mostly-compatible API
* Supports [multi-modal profiling](https://github.com/pkg/profile/issues/46):
  multiple profiles at once
* Configurable with idiomatic flags: `-cpuprofile`, `-memprofile`, ... just
  like `go test`

# Usage

`profile` should be mostly compatible with
[`pkg/profile`](https://github.com/pkg/profile), so examples for that package
should work here as well.

The following example shows the additional features of `profile`, namely
multi-modal profiles and flag configuration.

[embedmd]:# (example/main.go)
```go
package main

import (
	"flag"
	"log"

	"github.com/mmcloughlin/x/profile"
)

func main() {
	log.SetPrefix("example: ")
	log.SetFlags(0)

	// Setup profiler.
	p := profile.New(
		profile.CPUProfile,
		profile.MemProfile,
	)

	// Configure flags.
	n := flag.Int("n", 1000000, "sum the integers 1 to `n`")
	p.SetFlags(flag.CommandLine)
	flag.Parse()

	// Start profiler.
	defer p.Start().Stop()

	// Sum 1 to n.
	sum := 0
	for i := 1; i <= *n; i++ {
		sum += i
	}
	log.Printf("sum: %d", sum)
}
```

See the registered flags:

```
$ go run ./example/ -h
...
  -cpuprofile file
    	write a cpu profile to file
  -memprofile file
    	write an allocation profile to file
  -memprofilerate rate
    	set memory allocation profiling rate (see runtime.MemProfileRate)
  -n n
    	sum the integers 1 to n (default 1000000)
```

Profile the application:

```
$ go run ./example/ -n 1000000000 -cpuprofile cpu.out -memprofile mem.out
example: cpu profile: started
example: mem profile: started
example: sum: 500000000500000000
example: cpu profile: stopped
example: mem profile: stopped

$ ls *.out
cpu.out	mem.out
```
