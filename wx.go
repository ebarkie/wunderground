// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import "github.com/ebarkie/http/query"

// WxObs is weather observation info for the associated timestamp.
type WxObs struct {
	Bar           float64 `xml:"pressure_in"`
	DewPoint      float64 `xml:"dewpoint_f"`
	RainAccumDay  float64 `xml:"precip_today_in"`
	RainAccumHour float64 `xml:"precip_1hr_in"`
	OutHumidity   int     `xml:"relative_humidity"`
	OutTemp       float64 `xml:"temp_f"`
	SolarRad      float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Time          Time    `xml:"observation_time_rfc822"`
	UVIndex       float64 `xml:"UV"`
	WindDir       int     `xml:"wind_degrees"`
	WindSpeed     float64 `xml:"wind_mph"`
	WindGust      float64 `xml:"wind_gust_mph"`
}

// Wx represents weather observations.
type Wx struct {
	query.Data
}

// Bar is barometric pressure inches.
func (w *Wx) Bar(in float64) {
	w.SetFloat("baromin", in)
}

// Clouds is the cloud cover specified as: SKC, FEW, SCT, BKN, OVC.
func (w *Wx) Clouds(c string) {
	w.SetString("clouds", c)
}

// DailyRain is rain so far today (local time) in inches.
func (w *Wx) DailyRain(in float64) {
	w.SetFloat("dailyrainin", in)
}

// DewPoint is the outdoor dew point in degrees Fahrenheit.
func (w *Wx) DewPoint(f float64) {
	w.SetFloat("dewptf", f)
}

// InHumidity is the indoor humidity percentage (0-100).
func (w *Wx) InHumidity(p int) {
	w.SetInt("indoorhumudity", p)
}

// InTemp is the indoor temperature in degrees Fahrenheit.
func (w *Wx) InTemp(f float64) {
	w.SetFloat("indoortempf", f)
}

// LeafWetness is the leaf wetness percentage (0-100).
func (w *Wx) LeafWetness(p int) {
	w.SetIntf(query.Indexed{Format: "leafwetness#", Begin: 1, Zero: 1}, p)
}

// OutHumidity is the outdoor humidity percentage (0-100).
func (w *Wx) OutHumidity(p int) {
	w.SetInt("humidity", p)
}

// OutTemp is outdoor temperature in degrees Fahrenheit.
func (w *Wx) OutTemp(f float64) {
	w.SetFloatf(query.Indexed{Format: "temp#f", Begin: 1, Zero: 1}, f)
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall for the past 60 minutes in inches.
func (w *Wx) RainRate(in float64) {
	w.SetFloat("rainin", in)
}

// SolarRad is solar radiation in watts per square meter.
func (w *Wx) SolarRad(wm2 int) {
	w.SetInt("solarradiation", wm2)
}

// SoilMoist is the soil moisture percentage (0-100)..
func (w *Wx) SoilMoist(p int) {
	w.SetIntf(query.Indexed{Format: "soilmoisture#", Begin: 1, Zero: 1}, p)
}

// SoilTemp is the soil temperature in degrees Fahrenheit.
func (w *Wx) SoilTemp(f float64) {
	w.SetFloatf(query.Indexed{Format: "soiltemp#f", Begin: 1, Zero: 1}, f)
}

// UVIndex is the UltraViolet light index.
func (w *Wx) UVIndex(i float64) {
	w.SetFloat("UV", i)
}

// Visibility is visibility in nautical miles.
func (w *Wx) Visibility(nm int) {
	w.SetInt("visibility", nm)
}

// WindDir is instantaneous wind direction in degrees (0-359).
func (w *Wx) WindDir(deg int) {
	w.SetInt("winddir", deg)
}

// WindGustDir is the software specific time period wind gust
// direction in degrees (0-359).
func (w *Wx) WindGustDir(deg int) {
	w.SetInt("windgustdir", deg)
}

// WindGustDir10m is the 10 minute wind gust direction in
// degrees (0-359).
func (w *Wx) WindGustDir10m(deg int) {
	w.SetInt("windgustdir_10m", deg)
}

// WindGustSpeed is the software specific time period wind gust
// speed in miles per hour.
func (w *Wx) WindGustSpeed(mph float64) {
	w.SetFloat("windgustmph", mph)
}

// WindGustSpeed10m is the 10 minute wind gust speed in miles per
// hour.
func (w *Wx) WindGustSpeed10m(mph float64) {
	w.SetFloat("windgustmph_10m", mph)
}

// WindSpeed is the instantaneous wind speed in miles per hour.
func (w *Wx) WindSpeed(mph float64) {
	w.SetFloat("windspeedmph", mph)
}

// WindSpeedAvg2m is the 2 minute average wind speed in
// miles per hour.
func (w *Wx) WindSpeedAvg2m(mph float64) {
	w.SetFloat("windspdmph_avg2m", mph)
}
