// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package testutil

import (
	"expvar"
	"testing"
	"time"
)

// TestGetExpvar fetches the expvar metric `name`, and returns the expvar.
// Callers are responsible for type assertions on the returned value.
func TestGetExpvar(tb testing.TB, name string) expvar.Var {
	_ = "STUB: not implemented"
	return *new(expvar.Var)
}

const defaultDoOrTimeoutDeadline = 10 * time.Second

// ExpectExpvarDeltaWithDeadline returns a deferrable function which tests if the expvar metric with name has changed by delta within the given deadline, once the function begins.  Before returning, it fetches the original value for comparison.
func ExpectExpvarDeltaWithDeadline(tb testing.TB, name string, want int64) func() {
	_ = "STUB: not implemented"
	return nil
}

// ExpectMapExpvarMetricDeltaWithDeadline returns a deferrable function which tests if the expvar map metric with name and key has changed by delta within the given deadline, once the function begins.  Before returning, it fetches the original value for comparison.
func ExpectMapExpvarDeltaWithDeadline(tb testing.TB, name, key string, want int64) func() {
	_ = "STUB: not implemented"
	return nil
}
