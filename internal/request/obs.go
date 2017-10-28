// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package request

import (
	"fmt"
	"strings"
)

// Obs represents a set of observations (e.g. air quality or weather).
type Obs struct {
	vals map[string]string
}

// Clear clears all measurements that were previously set.
func (o *Obs) Clear() {
	o.vals = make(map[string]string)
}

// SetFloat adds a float measurement.
func (o *Obs) SetFloat(k string, f float64) {
	o.SetString(k, Float(f))
}

// SetInt adds an integer measurement.
func (o *Obs) SetInt(k string, i int) {
	o.SetString(k, Int(i))
}

// SetString adds a string measurement
func (o *Obs) SetString(k string, s string) {
	if o.vals == nil {
		o.Clear()
	}
	o.vals[o.genKey(k)] = s
}

// Values returns all measurements that were previously set.
func (o Obs) Values() map[string]string {
	return o.vals
}

// genKey returns a key string by replacing the X character in the template
// with either an empty string or an appropriate incremented value.
func (o Obs) genKey(t string) string {
	if !strings.Contains(t, "X") {
		return t
	}

	k := strings.Replace(t, "X", "", 1)
	i := 2
	for {
		if _, exists := o.vals[k]; !exists {
			return k
		}
		k = strings.Replace(t, "X", fmt.Sprintf("%d", i), 1)
		i++
	}
}
