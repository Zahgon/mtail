// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package datum

import (
	"time"
)

// Int describes an integer value at a given timestamp.
type Int struct {
	BaseDatum
	Value int64
}

// Set sets the value of the Int to the value at timestamp.
func (d *Int) Set(value int64, timestamp time.Time) { _ = "STUB: not implemented"; return }

// IncBy increments the Int's value by the value provided, at timestamp.
func (d *Int) IncBy(delta int64, timestamp time.Time) { _ = "STUB: not implemented"; return }

// DecBy increments the Int's value by the value provided, at timestamp.
func (d *Int) DecBy(delta int64, timestamp time.Time) { _ = "STUB: not implemented"; return }

// Get returns the value of the Int.
func (d *Int) Get() int64 { _ = "STUB: not implemented"; return 0 }

// ValueString returns the value of the Int as a string.
func (d *Int) ValueString() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON returns a JSON encoding of the Int.
func (d *Int) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
