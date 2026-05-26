package testutil

import (
	"io"
	"testing"
)

func WriteString(tb testing.TB, f io.StringWriter, str string) int {
	_ = "STUB: not implemented"
	return 0
}

// If this is a regular file (not a pipe or other StringWriter) then ensure
// it's flushed to disk, to guarantee the write happens-before this
// returns.
