// Copyright 2024 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package logstream

import (
	"context"
	"expvar"
	"io"
	"time"

	"github.com/google/mtail/internal/logline"
)

// logLines counts the number of lines read per log file.
var logLines = expvar.NewMap("log_lines_total")

// LineReader reads lines from input and sends lines through the channel
type LineReader struct {
	sourcename string                  // name of owner, for sending loglines
	lines      chan<- *logline.LogLine // not owned
	f          io.Reader               // not owned
	cancel     context.CancelFunc
	staleTimer *time.Timer // call CancelFunc if no read in 24h

	size int
	buf  []byte
	off  int // tracks the start of the next line in buf
}

// NewLineReader creates a new LineReader
func NewLineReader(sourcename string, lines chan<- *logline.LogLine, f io.Reader, size int, cancel context.CancelFunc) *LineReader {
	_ = "STUB: not implemented"
	return nil
}

// ReadAndSend reads bytes from f, attempts to find line endings in the bytes read, and sends them to the lines channel.  It manages the read buffer size to make sure we can always read size bytes.
func (lr *LineReader) ReadAndSend(ctx context.Context) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// reslice to set len

// reslice to drop earlier bytes

// send sends the line and resets the buffer offset
func (lr *LineReader) send(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// No newlines in the latest bytes, wait for next read

// excluding delim
// len of delim char

// Most file-based log sources will end with \n on Unixlike systems.  On
// Windows they appear to be both \r\n.  syslog disallows \r (and \t and
// others) and writes them escaped, per syslog(7).  [RFC
// 3164](https://www.ietf.org/rfc/rfc3164.txt) disallows newlines in the
// message: "The MSG part of the syslog packet MUST contain visible
// (printing) characters."  Thus if the previous char was a \r then ignore
// it as well.

// move past delim

// Finish sends the current accumulated line to the end of the buffer, despite
// there being no closing newline.
func (lr *LineReader) Finish(ctx context.Context) { _ = "STUB: not implemented"; return }
