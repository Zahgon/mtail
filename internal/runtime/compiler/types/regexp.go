// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package types

import (
	"regexp/syntax"
)

// ParseRegexp ensures we use the same regexp syntax.Flags across all
// invocations of this method.
func ParseRegexp(pattern string) (re *syntax.Regexp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
