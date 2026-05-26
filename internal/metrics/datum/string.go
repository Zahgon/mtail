// Copyright 2018 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package datum

import (
	"sync"
	"time"
)

// String describes a string value at a given timestamp.
type String struct {
	BaseDatum
	mu    sync.RWMutex
	Value string
}

// Set sets the value of the String to the value at timestamp.
func (d *String) Set(value string, timestamp time.Time) { _ = "STUB: not implemented"; return }

// Get returns the value of the String.
func (d *String) Get() string { _ = "STUB: not implemented"; return "" }

// ValueString returns the value of the String as a string.
func (d *String) ValueString() string {
	_ = "STUB: not implemented"

	// MarshalJSON returns a JSON encoding of the String.
	return ""
}

func (d *String) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
