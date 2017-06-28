// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadBadURL(t *testing.T) {
	a := assert.New(t)

	w := New("Kssssssnn", "deadbeef")
	w.URL = "http://localhost/updateweatherstation.php"

	w.Wx.Barometer(29.86)
	_, err := w.Upload()
	a.NotNil(err)
}

func TestUploadRequest(t *testing.T) {
	a := assert.New(t)

	w := New("KTEST0", "deadbeef")

	w.Aq.NO2(10)
	w.Wx.Barometer(29.86)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=KTEST0&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337", w.String())

	w.Wx.OutdoorTemperature(88.44)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=KTEST0&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337&tempf=88.44", w.String())
}

func TestUploadSkip(t *testing.T) {
	a := assert.New(t)

	w := New("KTEST0", "deadbeef")

	w.Wx.Barometer(29.86)
	skipped, err := w.Upload()
	a.Equal(false, skipped)
	a.Nil(err)

	w.Wx.Barometer(29.86)
	skipped, err = w.Upload()
	a.Equal(true, skipped)
	a.Nil(err)

	w.Wx.Barometer(29.87)
	skipped, err = w.Upload()
	a.Equal(false, skipped)
	a.Nil(err)
}
