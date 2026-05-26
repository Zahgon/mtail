// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package logstream

import (
	"context"
	"expvar"
	"os"
	"sync"

	"github.com/google/mtail/internal/waker"
)

// fileTruncates counts the truncations of a file stream.
var fileTruncates = expvar.NewMap("file_truncates_total")

// fileStream streams log lines from a regular file on the file system.  These
// log files are appended to by another process, and are either rotated or
// truncated by that (or yet another) process.  Rotation implies that a new
// inode with the same name has been created, the old file descriptor will be
// valid until EOF at which point it's considered completed.  A truncation means
// the same file descriptor is used but the file offset will be reset to 0.
// The latter is potentially lossy as far as mtail is concerned, if the last
// logs are not read before truncation occurs.  When an EOF is read, the
// goroutine tests for both truncation and inode change and resets or spins off
// a new goroutine and closes itself down.  The shared context is used for
// cancellation.
type fileStream struct {
	streamBase

	cancel context.CancelFunc

	pathname string // Given name for the underlying file on the filesystem
}

// newFileStream creates a new log stream from a regular file.
func newFileStream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, pathname string, fi os.FileInfo, oneShot OneShotMode) (LogStream, error) {
	_ = "STUB: not implemented"
	return *new(LogStream), nil
}

// Stream from the start of the file when in one shot mode.

func (fs *fileStream) stream(ctx context.Context, wg *sync.WaitGroup, waker waker.Waker, fi os.FileInfo, oneShot OneShotMode, streamFromStart bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Normal operation for first stream is to ignore the past, and seek to
// EOF immediately to start tailing.

// Blocking read but regular files will return EOF straight away.

// No error implies there is more to read so restart the loop.

// TODO: This could be generalised to check for any retryable
// errors, and end on unretriables; e.g. ESTALE looks
// retryable.

// streamFromStart always true on a stream reopen

// Close this stream.

// If we have read no bytes and are at EOF, check for truncation and rotation.

// Both rotation and truncation need to stat, so check for
// rotation first.  It is assumed that rotation is the more
// common change pattern anyway.

// If this is a NotExist error, then we should wrap up this
// goroutine. The Tailer will create a new logstream if the
// file is in the middle of a rotation and gets recreated
// in the next moment.  We can't rely on the Tailer to tell
// us we're deleted because the tailer can only tell us to
// cancel.

// Stream from start always true on a stream reopen

// We're at EOF so there's nothing left to read here.

// We know that newfi is from the current file.  Truncation can
// only be detected if the new file is currently shorter than
// the current seek offset.  In test this can be a race, but in
// production it's unlikely that a new file writes more bytes
// than the previous after rotation in the time it takes for
// mtail to notice.

// About to lose all remaining data because of the truncate so flush the accumulator.

// If we get here it's because we've stalled.  First test to see if it's
// time to exit.

// Exit now, because oneShot means read only to EOF.

// keep going

// Don't exit, instead yield and wait for a termination signal or
// wakeup.

// Exit after next read attempt.
// We may have started waiting here when the cancellation
// arrives, but since that wait the file may have been
// written to.  The file is not technically yet at EOF so
// we need to go back and try one more read.  We'll exit
// the stream in the select stanza above. This makes tests stable, but
// could argue exiting immediately is less surprising.
// Assumption is that this doesn't make a difference in
// production.

// sleep until next Wake()
