# Go Debounce

[![Tests on Linux, MacOS and Windows](https://github.com/wetfloo/go_debounce/workflows/Test/badge.svg)](https://github.com/wetfloo/go_debounce/actions?query=workflow:Test)
[![GoDoc](https://godoc.org/github.com/wetfloo/go_debounce?status.svg)](https://godoc.org/github.com/wetfloo/go_debounce)
[![Go Report Card](https://goreportcard.com/badge/github.com/wetfloo/go_debounce)](https://goreportcard.com/report/github.com/wetfloo/go_debounce)
[![codecov](https://codecov.io/gh/wetfloo/go_debounce/branch/master/graph/badge.svg)](https://codecov.io/gh/wetfloo/go_debounce)
[![Release](https://img.shields.io/github/release/wetfloo/go_debounce.svg?style=flat-square)](https://github.com/wetfloo/go_debounce/releases/latest)

## Example

```go
func ExampleNew() {
	var counter uint64

	f := func() {
		atomic.AddUint64(&counter, 1)
	}

	debounced := debounce.New(100 * time.Millisecond)

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			debounced(f)
		}

		time.Sleep(200 * time.Millisecond)
	}

	c := int(atomic.LoadUint64(&counter))

	fmt.Println("Counter is", c)
	// Output: Counter is 3
}
```

