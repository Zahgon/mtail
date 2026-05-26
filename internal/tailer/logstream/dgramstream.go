// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package logstream

import (
	"context"
	"net"
	"sync"

	"github.com/google/mtail/internal/waker"
)

type dgramStream struct {
	streamBase

	cancel context.CancelFunc

	scheme  string // Datagram scheme, either "unixgram" or "udp".
	address string // Given name for the underlying socket path on the filesystem or hostport.
}

func newDgramStream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, scheme, address string, oneShot OneShotMode) (LogStream, error) {
	_ = "STUB: not implemented"
	return *new(LogStream), nil
}

// The read buffer size for datagrams.
const datagramReadBufferSize = 131072

func (ds *dgramStream) stream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, oneShot OneShotMode) error {
	_ = "STUB: not implemented"
	return nil
}

// This is a test-only trick that says if we've already put this
// logstream in graceful shutdown, then a zero-byte read is
// equivalent to an "EOF" in connection and file oriented streams.

// No error implies more to read, so restart the loop.

// Yield and wait

// Exit after next read attempt.
// We may have started waiting here when the stop signal
// arrives, but since that wait the file may have been
// written to.  The file is not technically yet at EOF so
// we need to go back and try one more read.  We'll exit
// the stream in the zero byte handler above.

// sleep until next Wake()

// dgramConn wraps a PacketConn to add a Read method.
type dgramConn struct {
	net.PacketConn
}

// Read satisfies io.Reader
func (d *dgramConn) Read(p []byte) (count int, err error) { _ = "STUB: not implemented"; return 0, nil }
