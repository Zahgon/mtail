// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package logstream

import (
	"context"
	"net"
	"sync"

	"github.com/google/mtail/internal/waker"
)

type socketStream struct {
	streamBase

	cancel context.CancelFunc

	oneShot OneShotMode
	scheme  string // URL Scheme to listen with, either tcp or unix
	address string // Given name for the underlying socket path on the filesystem or host/port.
}

func newSocketStream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, scheme, address string, oneShot OneShotMode) (LogStream, error) {
	_ = "STUB: not implemented"
	return *new(LogStream), nil
}

// stream starts goroutines to read data from the stream socket, until Stop is called or the context is cancelled.
func (ss *socketStream) stream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker) error {
	_ = "STUB: not implemented"
	return nil
}

// signals when a connection has been opened

// tracks connection handling routines

// Set up for shutdown

// If oneshot, wait only for the one conn handler to start, otherwise
// wait for context Done or stopChan.

func (ss *socketStream) handleConn(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, c net.Conn) {
	_ = "STUB: not implemented"
	return
}

// No error implies more to read, so restart the loop.

// Yield and wait

// Exit after next read attempt.

// sleep until next Wake()
