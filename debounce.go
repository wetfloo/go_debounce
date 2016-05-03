// Copyright © 2016 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package debounce provides a debouncer func. The most typical use case would be
// the user typing a text into a form; the UI needs an update, but let's wait for
// a break.
package debounce

import (
	"time"
)

// New returns a debounced function and a channel that can be closed to signal a stop
// of the goroutine.
// The function will, as long as it continues to be invoked, not be triggered.
// The function will be called after it stops being called for the given duration.
// Note that a stop signal means a full stop of the debouncer; there is no
// concept of flushing future invocations.
func New(d time.Duration) (func(f func()), chan struct{}) {
	in, out, quit := debounceChan(d)

	go func() {
		for {
			select {
			case f := <-out:
				f()
			case <-quit:
				close(out)
				close(in)
				return
			}
		}
	}()

	debounce := func(f func()) {
		in <- f
	}

	return debounce, quit
}

func debounceChan(interval time.Duration) (in, out chan func(), quit chan struct{}) {
	in = make(chan func(), 1)
	out = make(chan func())
	quit = make(chan struct{})

	go func() {
		var f func() = func() {}
		for {
			select {
			case f = <-in:
			case <-time.After(interval):
				out <- f
				<-in
				// new interval
			case <-quit:
				return
			}
		}
	}()

	return
}