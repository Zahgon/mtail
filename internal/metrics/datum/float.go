// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package datum

import (
	"time"
)

// Float describes a floating point value at a given timestamp.
type Float struct {
	BaseDatum
	Valuebits uint64
}

// ValueString returns the value of the Float as a string.
func (d *Float) ValueString() string { _ = "STUB: not implemented"; return "" }

// Set sets value of the Float at the timestamp ts.
func (d *Float) Set(v float64, ts time.Time) { _ = "STUB: not implemented"; return }

// Get returns the floating-point value.
func (d *Float) Get() float64 { _ = "STUB: not implemented"; return 0 }

// MarshalJSON returns a JSON encoding of the Float.
func (d *Float) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
