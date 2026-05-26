// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package datum

import (
	"time"
)

// Datum is an interface for metric datums, with a type, value and timestamp to be exported.
type Datum interface {
	// // Type returns the Datum type.
	// Type() metrics.Type

	// ValueString returns the value of a Datum as a string.
	ValueString() string

	// TimeString returns the timestamp of a Datum as a string.
	TimeString() string

	// Time returns the timestamp of the Datum as time.Time in UTC
	TimeUTC() time.Time
}

// BaseDatum is a struct used to record timestamps across all Datum implementations.
type BaseDatum struct {
	Time int64 // nanoseconds since unix epoch
}

var zeroTime time.Time

func (d *BaseDatum) stamp(timestamp time.Time) { _ = "STUB: not implemented"; return }

// TimeString returns the timestamp of this Datum as a string.
func (d *BaseDatum) TimeString() string { _ = "STUB: not implemented"; return "" }

func (d *BaseDatum) TimeUTC() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewInt creates a new zero integer datum.
func NewInt() Datum { _ = "STUB: not implemented"; return *new(Datum) }

// NewFloat creates a new zero floating-point datum.
func NewFloat() Datum { _ = "STUB: not implemented"; return *new(Datum) }

// NewString creates a new zero string datum.
func NewString() Datum { _ = "STUB: not implemented"; return *new(Datum) }

// NewBuckets creates a new zero buckets datum.
func NewBuckets(buckets []Range) Datum { _ = "STUB: not implemented"; return *new(Datum) }

// MakeInt creates a new integer datum with the provided value and timestamp.
func MakeInt(v int64, ts time.Time) Datum { _ = "STUB: not implemented"; return *new(Datum) }

// MakeFloat creates a new floating-point datum with the provided value and timestamp.
func MakeFloat(v float64, ts time.Time) Datum { _ = "STUB: not implemented"; return *new(Datum) }

// MakeString creates a new string datum with the provided value and timestamp.
func MakeString(v string, ts time.Time) Datum { _ = "STUB: not implemented"; return *new(Datum) }

// MakeBuckets creates a new bucket datum with the provided list of ranges and
// timestamp.  If no +inf bucket is provided, one is created.
func MakeBuckets(buckets []Range, _ time.Time) Datum { _ = "STUB: not implemented"; return *new(Datum) }

// GetInt returns the integer value of a datum, or error.
func GetInt(d Datum) int64 { _ = "STUB: not implemented"; return 0 }

// GetFloat returns the floating-point value of a datum, or error.
func GetFloat(d Datum) float64 { _ = "STUB: not implemented"; return 0 }

// GetString returns the string of a datum, or error.
func GetString(d Datum) string { _ = "STUB: not implemented"; return "" }

// SetInt sets an integer datum to the provided value and timestamp, or panics if the Datum is not an IntDatum.
func SetInt(d Datum, v int64, ts time.Time) { _ = "STUB: not implemented"; return }

// SetFloat sets a floating-point Datum to the provided value and timestamp, or panics if the Datum is not a FloatDatum.
func SetFloat(d Datum, v float64, ts time.Time) { _ = "STUB: not implemented"; return }

// SetString sets a string Datum to the provided value and timestamp, or panics if the Datym is not a String Datum.
func SetString(d Datum, v string, ts time.Time) { _ = "STUB: not implemented"; return }

// IncIntBy increments an integer Datum by the provided value, at time ts, or panics if the Datum is not an IntDatum.
func IncIntBy(d Datum, v int64, ts time.Time) { _ = "STUB: not implemented"; return }

// DecIntBy increments an integer Datum by the provided value, at time ts, or panics if the Datum is not an IntDatum.
func DecIntBy(d Datum, v int64, ts time.Time) { _ = "STUB: not implemented"; return }

func GetBuckets(d Datum) *Buckets { _ = "STUB: not implemented"; return nil }

// Observe records an observation v at time ts in d, or panics if d is not a BucketsDatum.
func Observe(d Datum, v float64, ts time.Time) { _ = "STUB: not implemented"; return }

// GetBucketCount returns the total count of observations in d, or panics if d is not a BucketsDatum.
func GetBucketsCount(d Datum) uint64 { _ = "STUB: not implemented"; return 0 }

// GetBucketsSum returns the sum of observations in d, or panics if d is not a BucketsDatum.
func GetBucketsSum(d Datum) float64 { _ = "STUB: not implemented"; return 0 }

// GetBucketsCumByMax returns a map of cumulative bucket observations by their
// upper bonds, or panics if d is not a BucketsDatum.
func GetBucketsCumByMax(d Datum) map[float64]uint64 { _ = "STUB: not implemented"; return nil }
