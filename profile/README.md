# profile

Experimental profiling support package for Go.

* Based on the widely-used [`pkg/profile`](https://github.com/pkg/profile): mostly-compatible API
* Supports [multi-modal profiling](https://github.com/pkg/profile/issues/46): multiple profiles at once
* Configurable with idiomatic flags: `-cpuprofile`, `-memprofile`, ... just like `go test`
