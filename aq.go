// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import "gitlab.com/ebarkie/http/query"

// Aq represents air quality observations.
type Aq struct {
	query.Data
}

// BC is BC (black carbon at 880 nm) UG/M3.
func (a *Aq) BC(ugm3 float64) {
	a.SetFloat("AqBC", ugm3)
}

// CO is CO (carbon monoxide), caentional ppm.
func (a *Aq) CO(ppm int) {
	a.SetInt("AqCO", ppm)
}

// COT is CO trace levels ppb.
func (a *Aq) COT(ppb int) {
	a.SetInt("AqCOT", ppb)
}

// EC is EC (elemental carbon) – PM2.5 UG/M3.
func (a *Aq) EC(ugm3 float64) {
	a.SetFloat("AqEC", ugm3)
}

// NO is NO (nitric oxide) ppb.
func (a *Aq) NO(ppb int) {
	a.SetInt("AqNO", ppb)
}

// NO2T is (nitrogen dioxide), true measure ppb.
func (a *Aq) NO2T(ppb int) {
	a.SetInt("AqNO2T", ppb)
}

// NO2 is NO2 computed, NOx-NO ppb.
func (a *Aq) NO2(ppb int) {
	a.SetInt("AqNO2", ppb)
}

// NO2Y is NO2 computed, NOy-NO ppb.
func (a *Aq) NO2Y(ppb int) {
	a.SetInt("AqNO2Y", ppb)
}

// NOX is NOx (nitrogen oxides) - ppb.
func (a *Aq) NOX(ppb int) {
	a.SetInt("AqNOX", ppb)
}

// NOY is NOy (total reactive nitrogen) - ppb.
func (a *Aq) NOY(ppb int) {
	a.SetInt("AqNOY", ppb)
}

// NO3 is NO3 ion (nitrate, not adjusted for ammonium ion) UG/M3.
func (a *Aq) NO3(ugm3 float64) {
	a.SetFloat("AqNO3", ugm3)
}

// OC is OC (organic carbon, not adjusted for oxygen and hydrogen) – PM2.5 UG/M3.
func (a *Aq) OC(ugm3 float64) {
	a.SetFloat("AqOC", ugm3)
}

// Ozone is Ozone - ppb.
func (a *Aq) Ozone(ppb int) {
	a.SetInt("AqOZONE", ppb)
}

// PM25 is PM2.5 mass - UG/M3.
func (a *Aq) PM25(ugm3 float64) {
	a.SetFloat("AqPM2.5", ugm3)
}

// PM10 s PM10 mass - PM10 mass.
func (a *Aq) PM10(ugm3 float64) {
	a.SetFloat("AqPM10", ugm3)
}

// SO4 is SO4 ion (sulfate, not adjusted for ammonium ion) UG/M3.
func (a *Aq) SO4(ugm3 float64) {
	a.SetFloat("AqSO4", ugm3)
}

// SO2 (sulfur dioxide), caentional ppb.
func (a *Aq) SO2(ppb int) {
	a.SetInt("AqSO2", ppb)
}

// SO2T is trace levels ppb.
func (a *Aq) SO2T(ppb int) {
	a.SetInt("AqSO2T", ppb)
}

// UVAETH is UV-AETH (second channel of Aethalometer at 370 nm) UG/M3.
func (a *Aq) UVAETH(ugm3 float64) {
	a.SetFloat("AqUV-AETH", ugm3)
}
