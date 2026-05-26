// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package metrics

import (
	"math/rand"
	"reflect"
)

// Type describes the type of value stored in a Datum.
type Type int

const (
	// Int indicates this metric is an integer metric type.
	Int Type = iota
	// Float indicates this metric is a floating-point metric type.
	Float
	// String indicates this metric contains printable string values.
	String
	// Buckets indicates this metric is a histogram metric type.
	Buckets

	endType // end of enumeration for testing
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// Generate implements the quick.Generator interface for Type.
func (Type) Generate(rand *rand.Rand, _ int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
