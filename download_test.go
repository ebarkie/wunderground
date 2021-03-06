// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDownloadDaily(t *testing.T) {
	a := assert.New(t)

	p := Pws{ID: "KNCCARY89"}

	day := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
	o, err := p.DownloadDaily(day)

	if err != nil {
		a.FailNow(err.Error())
	}

	a.Equal(284, len(o), "Observation count")
	a.Equal(30.15, o[0].Bar, "Pressure")
}
