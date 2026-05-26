// Copyright 2018 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// Reimport the go-cmp package as the name 'cmp' conflicts with the cmp
// instruction in the vm.
package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Diff(a, b interface{}, opts ...cmp.Option) string { _ = "STUB: not implemented"; return "" }

func IgnoreUnexported(types ...interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

func AllowUnexported(types ...interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

func IgnoreFields(typ interface{}, names ...string) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

func SortSlices(lessFunc interface{}) cmp.Option {
	_ = "STUB: not implemented"
	return *new(cmp.Option)
}

// ExpectNoDiff tests to see if the two interfaces have no diff.
// If there is no diff, the retrun value is true.
// If there is a diff, it is logged to tb and an error is flagged, and the return value is false.
func ExpectNoDiff(tb testing.TB, a, b interface{}, opts ...cmp.Option) bool {
	_ = "STUB: not implemented"
	return false
}
