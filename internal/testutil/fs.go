// Copyright 2019 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package testutil

import (
	"os"
	"testing"
)

// TestTempDir creates a temporary directory for use during tests, returning the pathname.
func TestTempDir(tb testing.TB) string { _ = "STUB: not implemented"; return "" }

// TestOpenFile creates a new file called name and returns the opened file.
func TestOpenFile(tb testing.TB, name string) *os.File { _ = "STUB: not implemented"; return nil }

// OpenLogFile creates a new file that emulates being a log.
func OpenLogFile(tb testing.TB, name string) *os.File { _ = "STUB: not implemented"; return nil }

// Chdir changes current working directory, and registers a cleanup function
// to return to the previous directory.
func Chdir(tb testing.TB, dir string) { _ = "STUB: not implemented"; return }
