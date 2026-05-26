// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package datum

import (
	"sync"
	"time"
)

type Range struct {
	Min float64
	Max float64
}

type BucketCount struct {
	Range Range
	Count uint64
}

func (r *Range) Contains(v float64) bool { _ = "STUB: not implemented"; return false }

// Buckets describes a floating point value at a given timestamp.
type Buckets struct {
	BaseDatum
	sync.RWMutex
	Buckets []BucketCount
	Count   uint64
	Sum     float64
}

func (d *Buckets) ValueString() string { _ = "STUB: not implemented"; return "" }

func (d *Buckets) Observe(v float64, ts time.Time) { _ = "STUB: not implemented"; return }

func (d *Buckets) GetCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *Buckets) GetSum() float64 { _ = "STUB: not implemented"; return 0 }

func (d *Buckets) AddBucket(r Range) { _ = "STUB: not implemented"; return }

func (d *Buckets) GetBuckets() map[Range]uint64 { _ = "STUB: not implemented"; return nil }

func (d *Buckets) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Range) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
