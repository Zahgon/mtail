// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package logstream

import (
	"context"
	"os"
	"sync"

	"github.com/google/mtail/internal/waker"
)

type fifoStream struct {
	streamBase

	cancel context.CancelFunc

	pathname string // Given name for the underlying named pipe on the filesystem
}

// newFifoStream creates a new stream reader for Unix Fifos.
// `pathname` must already be verified as clean.
func newFifoStream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, pathname string, fi os.FileInfo) (LogStream, error) {
	_ = "STUB: not implemented"
	return *new(LogStream), nil
}

func fifoOpen(pathname string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// Open in nonblocking mode because the write end of the fifo may not have started yet; this also gives us the ability to set a read deadline when the context is cancelled. https://github.com/golang/go/issues/24842
// #nosec G304 -- path already validated by caller

// The read buffer size for fifos.
//
// Before Linux 2.6.11, the capacity of a fifo was the same as the
// system page size (e.g., 4096 bytes on i386).  Since Linux 2.6.11,
// the fifo capacity is 16 pages (i.e., 65,536 bytes in a system
// with a page size of 4096 bytes).  Since Linux 2.6.35, the default
// fifo capacity is 16 pages, but the capacity can be queried and
// set using the fcntl(2) F_GETPIPE_SZ and F_SETPIPE_SZ operations.
// See fcntl(2) for more information.
//
// https://man7.org/linux/man-pages/man7/pipe.7.html
const defaultFifoReadBufferSize = 131072

func (ps *fifoStream) stream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, _ os.FileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// No error implies there is more to read so restart the loop.

// `pipe(7)` tells us "If all file descriptors referring to the
// write end of a fifo have been closed, then an attempt to
// read(2) from the fifo will see end-of-file (read(2) will
// return 0)."  To avoid shutting down the stream at startup
// before any writer has connected to the fifo, condition on
// having read any bytes previously.

// Test to see if we should exit.

// Because we've opened in nonblocking mode, this Read can return
// straight away.  If there are no writers, it'll return EOF (per
// `pipe(7)` and `read(2)`.)  This is expected when `mtail` is
// starting at system init as the writer may not be ready yet.

// Wait for wakeup or termination.

// Exit after next read attempt.

// sleep until next Wake()
