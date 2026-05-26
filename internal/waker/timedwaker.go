// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package waker

import (
	"context"
	"sync"
	"time"
)

// A timedWaker wakes callers on a regular interval.
type timedWaker struct {
	Waker

	t    *time.Ticker
	mu   sync.Mutex // protects following fields
	wake chan struct{}
}

// NewTimed returns a new timedWaker that is shut down when the context is cancelled.
func NewTimed(ctx context.Context, interval time.Duration) Waker {
	_ = "STUB: not implemented"
	return *new(Waker)
}

// Wake implements the Waker interface.
func (t *timedWaker) Wake() (w <-chan struct{}) { _ = "STUB: not implemented"; return nil }
