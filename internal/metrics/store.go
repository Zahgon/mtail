// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package metrics

import (
	"context"
	"io"
	"sync"
	"time"
)

// Store contains Metrics.
type Store struct {
	searchMu sync.RWMutex // read for iterate and insert, write for delete
	insertMu sync.Mutex   // locked for insert and delete, unlocked for iterate
	Metrics  map[string][]*Metric
}

// NewStore returns a new metric Store.
func NewStore() (s *Store) { _ = "STUB: not implemented"; return nil }

// Add is used to add one metric to the Store.
func (s *Store) Add(m *Metric) error { _ = "STUB: not implemented"; return nil }

// To avoid duplicate metrics:
// - copy old LabelValues into new metric;
// - discard old metric.

// If a set of label keys has changed, discard
// old metric completely, w/o even copying old
// data, as they are now incompatible.

// Otherwise, copy everything into the new metric

// We're in modify mode now so lock out search

// FindMetricOrNil returns a metric in a store, or returns nil if not found.
func (s *Store) FindMetricOrNil(name, prog string) *Metric { _ = "STUB: not implemented"; return nil }

// ClearMetrics empties the store of all metrics.
func (s *Store) ClearMetrics() { _ = "STUB: not implemented"; return }

// MarshalJSON returns a JSON byte string representing the Store.
func (s *Store) MarshalJSON() (b []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Range calls f sequentially for each Metric present in the store.
// The Metric is not locked when f is called.
// If f returns non nil error, Range stops the iteration.
// This looks a lot like sync.Map, ay.
func (s *Store) Range(f func(*Metric) error) error { _ = "STUB: not implemented"; return nil }

// Gc iterates through the Store looking for metrics that can be tidied up,
// if they are passed their expiry or sized greater than their limit.
func (s *Store) Gc() error { _ = "STUB: not implemented"; return nil }

// StartGcLoop runs a permanent goroutine to expire metrics every duration.
func (s *Store) StartGcLoop(ctx context.Context, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// WriteMetrics dumps the current state of the metrics store in JSON format to
// the io.Writer.
func (s *Store) WriteMetrics(w io.Writer) error { _ = "STUB: not implemented"; return nil }
