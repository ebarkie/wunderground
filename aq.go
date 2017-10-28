// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import "github.com/ebarkie/wunderground/internal/request"

// Aq represents air quality measurements.
type Aq struct {
	request.Obs
}

// BC is BC (black carbon at 880 nm) UG/M3.
func (a *Aq) BC(f float64) {
	a.SetFloat("AqBC", f)
}

// CO is CO (carbon monoxide), caentional ppm.
func (a *Aq) CO(i int) {
	a.SetInt("AqCO", i)
}

// COT is CO trace levels ppb.
func (a *Aq) COT(i int) {
	a.SetInt("AqCOT", i)
}

// EC is EC (elemental carbon) – PM2.5 UG/M3.
func (a *Aq) EC(f float64) {
	a.SetFloat("AqEC", f)
}

// NO is NO (nitric oxide) ppb.
func (a *Aq) NO(i int) {
	a.SetInt("AqNO", i)
}

// NO2T is (nitrogen dioxide), true measure ppb.
func (a *Aq) NO2T(i int) {
	a.SetInt("AqNO2T", i)
}

// NO2 is NO2 computed, NOx-NO ppb.
func (a *Aq) NO2(i int) {
	a.SetInt("AqNO2", i)
}

// NO2Y is NO2 computed, NOy-NO ppb.
func (a *Aq) NO2Y(i int) {
	a.SetInt("AqNO2Y", i)
}

// NOX is NOx (nitrogen oxides) - ppb.
func (a *Aq) NOX(i int) {
	a.SetInt("AqNOX", i)
}

// NOY is NOy (total reactive nitrogen) - ppb.
func (a *Aq) NOY(i int) {
	a.SetInt("AqNOY", i)
}

// NO3 is NO3 ion (nitrate, not adjusted for ammonium ion) UG/M3.
func (a *Aq) NO3(f float64) {
	a.SetFloat("AqNO3", f)
}

// OC is OC (organic carbon, not adjusted for oxygen and hydrogen) – PM2.5 UG/M3.
func (a *Aq) OC(f float64) {
	a.SetFloat("AqOC", f)
}

// Ozone is Ozone - ppb.
func (a *Aq) Ozone(i int) {
	a.SetInt("AqOZONE", i)
}

// PM25 is PM2.5 mass - UG/M3.
func (a *Aq) PM25(f float64) {
	a.SetFloat("AqPM2.5", f)
}

// PM10 s PM10 mass - PM10 mass.
func (a *Aq) PM10(f float64) {
	a.SetFloat("AqPM10", f)
}

// SO4 is SO4 ion (sulfate, not adjusted for ammonium ion) UG/M3.
func (a *Aq) SO4(f float64) {
	a.SetFloat("AqSO4", f)
}

// SO2 (sulfur dioxide), caentional ppb.
func (a *Aq) SO2(i int) {
	a.SetInt("AqSO2", i)
}

// SO2T is trace levels ppb.
func (a *Aq) SO2T(i int) {
	a.SetInt("AqSO2T", i)
}

// UVAETH is UV-AETH (second channel of Aethalometer at 370 nm) UG/M3.
func (a *Aq) UVAETH(f float64) {
	a.SetFloat("AqUV-AETH", f)
}
