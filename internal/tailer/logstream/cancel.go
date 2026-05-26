package logstream

import (
	"context"
	"time"
)

// ReadDeadliner has a SetReadDeadline function to be used for interrupting reads.
type ReadDeadliner interface {
	SetReadDeadline(t time.Time) error
}

// SetReadDeadlineOnDone waits for the context to be done, and then sets an
// immediate read deadline on the flie descriptor `d`.  This causes any blocked
// reads on that descriptor to return with an i/o timeout error.
func SetReadDeadlineOnDone(ctx context.Context, d ReadDeadliner) { _ = "STUB: not implemented"; return }

// IsExitableError returns true if a stream should exit because of this error.
func IsExitableError(err error) bool { _ = "STUB: not implemented"; return false }

// https://github.com/golang/go/issues/4373
