// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package metrics

type MetricSlice []*Metric

func (s MetricSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s MetricSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s MetricSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func Less(m1, m2 *Metric) bool { _ = "STUB: not implemented"; return false }

// if lv.Value < m2.LabelValues[x].Value {
// 	return true
// }
