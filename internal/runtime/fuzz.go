// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

//go:build gofuzz

package runtime

import (
	"flag"
)

// U+2424 SYMBOL FOR NEWLINE
const SEP = "␤"

// Enable this when debugging with a fuzz crash artifact; it slows the fuzzer down when enabled.
const dumpDebug = false

func Fuzz(data []byte) int {
	_ = "STUB: not implemented"
	// Data contains the program and sample input, separated by SEP.
	return 0
}

// If no SEP, then append one and an empty line of input.

// false

func init() {
	// We need to successfully parse flags to initialize the glog logger used
	// by the compiler, but the fuzzer gets called with flags captured by the
	// libfuzzer main, which we don't want to intercept here.
	flag.CommandLine.Parse([]string{})
}
