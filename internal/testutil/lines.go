// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package testutil

import (
	"testing"

	"github.com/google/mtail/internal/logline"
)

func LinesReceived(lines <-chan *logline.LogLine) (r []*logline.LogLine) {
	_ = "STUB: not implemented"
	return nil
}

func ExpectLinesReceivedNoDiff(tb testing.TB, wantLines []*logline.LogLine, gotLines <-chan *logline.LogLine) func() {
	_ = "STUB: not implemented"
	return nil
}
