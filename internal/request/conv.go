// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package request

import "strconv"

// Float converts a float to a string.
func Float(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Int converts a integer to a string.
func Int(i int) string {
	return strconv.Itoa(i)
}
