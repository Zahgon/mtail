// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package golden

import (
	"io"
	"regexp"

	"github.com/google/mtail/internal/metrics"
)

var varRe = regexp.MustCompile(`^(counter|gauge|timer|text|histogram) ([^ ]+)(?: {([^}]+)})?(?: (\S+))?(?: (.+))?`)

// ReadTestData loads a "golden" test data file from a programfile and returns as a slice of Metrics.
func ReadTestData(file io.Reader, programfile string) metrics.MetricSlice {
	_ = "STUB: not implemented"
	return *new(metrics.MetricSlice)
}

// Now we have enough information to get or create a metric.

// Initialize to zero at the zero time.

/* #nosec G104 -- Always returns nil. nolint:errcheck */
