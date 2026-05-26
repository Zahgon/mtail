// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// Package position implements a data structure for storing source code positions.
package position

// A Position is the location in the source program that a token appears.  It
// can specify a single character in the pinput, in which case the start and
// end columns are the same, or a span of sequential characters on one line.
type Position struct {
	Filename string // Source filename in which this token appears.
	Line     int    // Line in the source for this token.
	Startcol int    // Starting and ending columns in the source for this token.
	Endcol   int
}

// String formats a position to be useful for printing messages associated with
// this position, e.g. compiler errors.
func (p Position) String() string { _ = "STUB: not implemented"; return "" }

// MergePosition returns the union of two positions such that the result contains both inputs.
func Merge(a, b *Position) *Position { _ = "STUB: not implemented"; return nil }

// TODO(jaq): handle multi-line positions
