// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package mtail

// BuildInfo records the compile-time information for use when reporting the mtail version.
type BuildInfo struct {
	Branch   string
	Version  string
	Revision string
}

func (b BuildInfo) String() string { _ = "STUB: not implemented"; return "" }
