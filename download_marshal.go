// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"encoding/xml"
	"time"
)

// Time is an anonymous struct representing a standard time.Time
// so we can create a custom unmarshaler for WU's timestamp format.
type Time struct {
	time.Time
}

// UnmarshalXML method for converting a WU timestamp string to a time.Time
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	ts, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", v)
	if err != nil {
		return err
	}
	*t = Time{ts.Local()}
	return nil
}
