// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

// Refer to:
// http://wiki.wunderground.com/index.php/PWS_-_Upload_Protocol

// BC is BC (black carbon at 880 nm) UG/M3.
func (a *AqObs) BC(v float64) {
	a.p["AqBC"] = f2s(v)
}

// CO is CO (carbon monoxide), conventional ppm.
func (a *AqObs) CO(v int) {
	a.p["AqCO"] = i2s(v)
}

// COT is CO trace levels ppb.
func (a *AqObs) COT(v int) {
	a.p["AqCOT"] = i2s(v)
}

// EC is EC (elemental carbon) – PM2.5 UG/M3.
func (a *AqObs) EC(v float64) {
	a.p["AqEC"] = f2s(v)
}

// NO is NO (nitric oxide) ppb.
func (a *AqObs) NO(v int) {
	a.p["AqNO"] = i2s(v)
}

// NO2T is (nitrogen dioxide), true measure ppb.
func (a *AqObs) NO2T(v int) {
	a.p["AqNO2T"] = i2s(v)
}

// NO2 is NO2 computed, NOx-NO ppb.
func (a *AqObs) NO2(v int) {
	a.p["AqNO2"] = i2s(v)
}

// NO2Y is NO2 computed, NOy-NO ppb.
func (a *AqObs) NO2Y(v int) {
	a.p["AqNO2Y"] = i2s(v)
}

// NOX is NOx (nitrogen oxides) - ppb.
func (a *AqObs) NOX(v int) {
	a.p["AqNOX"] = i2s(v)
}

// NOY is NOy (total reactive nitrogen) - ppb.
func (a *AqObs) NOY(v int) {
	a.p["AqNOY"] = i2s(v)
}

// NO3 is NO3 ion (nitrate, not adjusted for ammonium ion) UG/M3.
func (a *AqObs) NO3(v float64) {
	a.p["AqNO3"] = f2s(v)
}

// OC is OC (organic carbon, not adjusted for oxygen and hydrogen) – PM2.5 UG/M3.
func (a *AqObs) OC(v float64) {
	a.p["AqOC"] = f2s(v)
}

// Ozone is Ozone - ppb.
func (a *AqObs) Ozone(v int) {
	a.p["AqOZONE"] = i2s(v)
}

// PM25 is PM2.5 mass - UG/M3.
func (a *AqObs) PM25(v float64) {
	a.p["AqPM2.5"] = f2s(v)
}

// PM10 s PM10 mass - PM10 mass.
func (a *AqObs) PM10(v float64) {
	a.p["AqPM10"] = f2s(v)
}

// SO4 is SO4 ion (sulfate, not adjusted for ammonium ion) UG/M3.
func (a *AqObs) SO4(v float64) {
	a.p["AqSO4"] = f2s(v)
}

// SO2 (sulfur dioxide), conventional ppb.
func (a *AqObs) SO2(v int) {
	a.p["AqSO2"] = i2s(v)
}

// SO2T is trace levels ppb.
func (a *AqObs) SO2T(v int) {
	a.p["AqSO2T"] = i2s(v)
}

// UVAETH is UV-AETH (second channel of Aethalometer at 370 nm) UG/M3.
func (a *AqObs) UVAETH(v float64) {
	a.p["AqUV-AETH"] = f2s(v)
}
