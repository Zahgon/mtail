// Copyright 2019 Google Inc. All Rights Reserved.
// This file is available under the Apache license.
package testutil

import (
	"testing"
	"time"
)

// DoOrTimeout runs a check function every interval until deadline, unless the
// check returns true.  The check should return false otherwise. If the check
// returns an error the check is immediately failed.
func DoOrTimeout(do func() (bool, error), deadline, interval time.Duration) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// otherwise wait and retry

// TimeoutTest returns a test function that executes f with a timeout, If the
// test does not complete in time the test is failed.  This lets us set a
// per-test timeout instead of the global `go test -timeout` coarse timeout.
func TimeoutTest(timeout time.Duration, f func(t *testing.T)) func(t *testing.T) {
	_ = "STUB: not implemented"
	// Raise the timeout if we're run under the race detector.
	return nil
}

// If we're in a CI environment, raise the timeout by 10x.  This mimics the
// timeout global flag set in the Makefile.
