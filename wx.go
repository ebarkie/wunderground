// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import "github.com/ebarkie/wunderground/internal/request"

// WxObs is weather observation info for the associated timestamp.
type WxObs struct {
	Barometer            float64 `xml:"pressure_in"`
	DewPoint             float64 `xml:"dewpoint_f"`
	RainAccumulationDay  float64 `xml:"precip_today_in"`
	RainAccumulationHour float64 `xml:"precip_1hr_in"`
	OutdoorHumidity      int     `xml:"relative_humidity"`
	OutdoorTemperature   float64 `xml:"temp_f"`
	SolarRadiation       float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Time                 Time    `xml:"observation_time_rfc822"`
	UVIndex              float64 `xml:"UV"`
	WindDirection        int     `xml:"wind_degrees"`
	WindSpeed            float64 `xml:"wind_mph"`
	WindGust             float64 `xml:"wind_gust_mph"`
}

// Wx represents weather measurements.
type Wx struct {
	request.Obs
}

// Barometer is barometric pressure inches.
func (w *Wx) Barometer(f float64) {
	w.SetFloat("baromin", f)
}

// Clouds is SKC, FEW, SCT, BKN, OVC.
func (w *Wx) Clouds(s string) {
	w.SetString("clouds", s)
}

// DailyRain is rain inches so far today in local time.
func (w *Wx) DailyRain(f float64) {
	w.SetFloat("dailyrainin", f)
}

// DewPoint is F outdoor dewpoint.
func (w *Wx) DewPoint(f float64) {
	w.SetFloat("dewptf", f)
}

// IndoorHumidity is % indoor humidity 0-100%.
func (w *Wx) IndoorHumidity(i int) {
	w.SetInt("indoorhumudity", i)
}

// IndoorTemperature is indoor temperature in F.
func (w *Wx) IndoorTemperature(f float64) {
	w.SetFloat("indoortempf", f)
}

// LeafWetness is %.
func (w *Wx) LeafWetness(i int) {
	w.SetInt("leafwetnessX", i)
}

// OutdoorHumidity is % outdoor humidity 0-100%.
func (w *Wx) OutdoorHumidity(i int) {
	w.SetInt("humidity", i)
}

// OutdoorTemperature is outdoor temperature in F.
func (w *Wx) OutdoorTemperature(f float64) {
	w.SetFloat("tempXf", f)
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall in the past 60 min.
func (w *Wx) RainRate(f float64) {
	w.SetFloat("rainin", f)
}

// SolarRadiation is solar radiation in W/m^2.
func (w *Wx) SolarRadiation(i int) {
	w.SetInt("solarradiation", i)
}

// SoilMoisture is %.
func (w *Wx) SoilMoisture(i int) {
	w.SetInt("soilmoistureX", i)
}

// SoilTemperature is F soil temperature.
func (w *Wx) SoilTemperature(f float64) {
	w.SetFloat("soiltempXf", f)
}

// UVIndex is the UV index.
func (w *Wx) UVIndex(f float64) {
	w.SetFloat("UV", f)
}

// Visibility is nm visibility.
func (w *Wx) Visibility(i int) {
	w.SetInt("visibility", i)
}

// WindDirection is 0-360 instantaneous wind direction.
func (w *Wx) WindDirection(i int) {
	w.SetInt("winddir", i)
}

// WindGustDirection is 0-360 using software specific time period.
func (w *Wx) WindGustDirection(i int) {
	w.SetInt("windgustdir", i)
}

// WindGustDirection10m is 0-360 past 10 minutes wind gust direction.
func (w *Wx) WindGustDirection10m(i int) {
	w.SetInt("windgustdir_10m", i)
}

// WindGustSpeed is mph current wind gust, using software specific time period.
func (w *Wx) WindGustSpeed(f float64) {
	w.SetFloat("windgustmph", f)
}

// WindGustSpeed10m is mph past 10 minutes wind gust mph.
func (w *Wx) WindGustSpeed10m(f float64) {
	w.SetFloat("windgustmph_10m", f)
}

// WindSpeed is mph instantaneous wind speed.
func (w *Wx) WindSpeed(f float64) {
	w.SetFloat("windspeedmph", f)
}

// WindSpeedAverage2m is mph 2 minute average wind speed mph.
func (w *Wx) WindSpeedAverage2m(f float64) {
	w.SetFloat("windspdmph_avg2m", f)
}
