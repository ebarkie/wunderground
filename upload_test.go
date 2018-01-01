// Copyright (c) 2016-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadEncode(t *testing.T) {
	a := assert.New(t)

	p := Pws{ID: "KTEST0", Password: "deadbeef"}

	aq := &Aq{}
	aq.NO2(10)
	wx := &Wx{}
	wx.Bar(29.86)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=KTEST0&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337", p.Encode(aq, wx))

	wx.OutTemp(88.44)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?AqNO2=10&ID=KTEST0&PASSWORD=deadbeef&action=updateraw&baromin=29.86&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337&tempf=88.44", p.Encode(aq, wx))
}

func TestUploadSkip(t *testing.T) {
	a := assert.New(t)

	p := Pws{ID: "KTEST0", Password: "deadbeef"}

	wx := &Wx{}
	wx.Bar(29.86)
	err := p.Upload(wx)
	a.Equal(false, p.Skipped())
	a.Nil(err)
	a.Equal("https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?ID=KTEST0&PASSWORD=deadbeef&action=updateraw&dateutc=now&realtime=1&rtfreq=2&softwaretype=GoWunder%201337", p.Encode(wx))

	wx.Bar(29.86)
	err = p.Upload(wx)
	a.Equal(true, p.Skipped())
	a.Nil(err)

	wx.Bar(29.87)
	err = p.Upload(wx)
	a.Equal(false, p.Skipped())
	a.Nil(err)
}
