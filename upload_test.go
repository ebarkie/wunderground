// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadRequest(t *testing.T) {
	a := assert.New(t)

	w := New("Kssssssnn", "deadbeef")

	w.Aq.NO2(10)
	w.Wx.Barometer(29.86)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=Kssssssnn&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337", w.String())
	w.Wx.OutdoorTemperature(88.44)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=Kssssssnn&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337&tempf=88.44", w.String())
}
