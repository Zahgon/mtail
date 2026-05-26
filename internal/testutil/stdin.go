package testutil

import (
	"os"
	"testing"
)

// OverrideStdin replaces standard input with another *os.File and restores
// it at the end of the test.
func OverrideStdin(tb testing.TB, fakeStdin *os.File) { _ = "STUB: not implemented"; return }
