// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package waker

import (
	"context"
	"sync"
)

// A testWaker is used to manually signal to idle routines it's time to look
// for new work.  It works by synchronising several client goroutines
// ("wakees"), waiting for a certain number to have called `Wake()`, and then
// sending them a wakeup signal together.  It then waits in a loop for more
// goroutines to call `Wake` again before returning to the test.
type testWaker struct {
	Waker

	ctx context.Context

	name string

	wakeeReady chan struct{}
	wakeeDone  chan struct{}
	waiting    chan struct{}

	mu   sync.Mutex // protects following fields
	wake chan struct{}
}

// WakeFunc describes a function used by tests to trigger a wakeup of blocked
// idle goroutines under test.  It takes as first parameter the number of
// goroutines to wake up, and the second parameter is the number of goroutines
// to wait to call Wake before returning.
type WakeFunc func(int, int)

// NewTest creates a new Waker to be used in tests, returning it and a function to trigger a wakeup.
// `wait` says how many wakees are expected to be waiting before the first `wakeFunc` call.
// `name` gives it a name for debug log messages
func NewTest(ctx context.Context, wait int, name string) (Waker, WakeFunc) {
	_ = "STUB: not implemented"
	return *new(Waker), *new(WakeFunc)
}

// awaken issues a wakeup signal to the "wakees", those clients who've used
// the `Wake` call.  wake is the number of wakees we expect to wake up,
// wait is the number of wakees to wait for before returning.

// First wait for `t.n` wakees to have called `Wake`, synchronising them.

// Now `awaken` blocks here, as we wait for them in turn to return to another call to Wake, in their polling loops.  We wait for only a count of `after` routines this time, as some may exit.

// Wake satisfies the Waker interface.
func (t *testWaker) Wake() (w <-chan struct{}) { _ = "STUB: not implemented"; return nil }

// Background this so we can return the wake channel.
// The wakeFunc won't close the channel until this completes.

// Signal we've reentered Wake.  wakeFunc can't return until we do this.

// Block wakees here until a subsequent wakeFunc is called.

// Signal we've got the wake chan, telling wakeFunc it can now issue a broadcast.

func (t *testWaker) broadcastWakeAndReset() { _ = "STUB: not implemented"; return }

// alwaysWaker never blocks the wakee.
type alwaysWaker struct {
	wake chan struct{}
}

func NewTestAlways() Waker { _ = "STUB: not implemented"; return *new(Waker) }

func (w *alwaysWaker) Wake() <-chan struct{} { _ = "STUB: not implemented"; return nil }
