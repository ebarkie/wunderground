// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"fmt"
	"strings"
)

// f2s converts parameter float values to a string.
func f2s(f float64) string {
	return fmt.Sprintf("%g", f)
}

// i2s converts parameter integer values to a string.
func i2s(i int) string {
	return fmt.Sprintf("%d", i)
}

// incrKey is used to fill in the template replacement character X with
// either an empty string or the appropriate incrementor.  This follows
// the Weather Underground pattern for parameters that can be specified
// several times.
func incrKey(p params, k string) (key string) {
	key = strings.Replace(k, "X", "", 1)

	i := 2
	for {
		if _, present := p[key]; !present {
			return
		}
		key = strings.Replace(k, "X", fmt.Sprintf("%d", i), 1)
		i++
	}
}
