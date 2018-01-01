// Copyright (c) 2016-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import "github.com/ebarkie/request"

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

// Wx represents weather observations.
type Wx struct {
	request.Data
}

// Barometer is barometric pressure inches.
func (w *Wx) Barometer(in float64) {
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

// IndoorHumidity is the indoor humidity percentage (0-100).
func (w *Wx) IndoorHumidity(p int) {
	w.SetInt("indoorhumudity", p)
}

// IndoorTemperature is the indoor temperature in degrees Fahrenheit.
func (w *Wx) IndoorTemperature(f float64) {
	w.SetFloat("indoortempf", f)
}

// LeafWetness is the leaf wetness percentage (0-100).
func (w *Wx) LeafWetness(p int) {
	w.SetIntf(request.Indexed{Format: "leafwetness#", Begin: 1, Zero: 1}, p)
}

// OutdoorHumidity is the outdoor humidity percentage (0-100).
func (w *Wx) OutdoorHumidity(p int) {
	w.SetInt("humidity", p)
}

// OutdoorTemperature is outdoor temperature in degrees Fahrenheit.
func (w *Wx) OutdoorTemperature(f float64) {
	w.SetFloatf(request.Indexed{Format: "temp#f", Begin: 1, Zero: 1}, f)
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall for the past 60 minutes in inches.
func (w *Wx) RainRate(in float64) {
	w.SetFloat("rainin", in)
}

// SolarRadiation is solar radiation in watts per square meter.
func (w *Wx) SolarRadiation(wm2 int) {
	w.SetInt("solarradiation", wm2)
}

// SoilMoisture is the soil moisture percentage (0-100)..
func (w *Wx) SoilMoisture(p int) {
	w.SetIntf(request.Indexed{Format: "soilmoisture#", Begin: 1, Zero: 1}, p)
}

// SoilTemperature is the soil temperature in degrees Fahrenheit.
func (w *Wx) SoilTemperature(f float64) {
	w.SetFloatf(request.Indexed{Format: "soiltemp#f", Begin: 1, Zero: 1}, f)
}

// UVIndex is the UltraViolet light index.
func (w *Wx) UVIndex(i float64) {
	w.SetFloat("UV", i)
}

// Visibility is visibility in nautical miles.
func (w *Wx) Visibility(nm int) {
	w.SetInt("visibility", nm)
}

// WindDirection is instantaneous wind direction in degrees (0-359).
func (w *Wx) WindDirection(deg int) {
	w.SetInt("winddir", deg)
}

// WindGustDirection is the software specific time period wind gust
// direction in degrees (0-359).
func (w *Wx) WindGustDirection(deg int) {
	w.SetInt("windgustdir", deg)
}

// WindGustDirection10m is the 10 minute wind gust direction in
// degrees (0-359).
func (w *Wx) WindGustDirection10m(deg int) {
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

// WindSpeedAverage2m is the 2 minute average wind speed in
// miles per hour.
func (w *Wx) WindSpeedAverage2m(mph float64) {
	w.SetFloat("windspdmph_avg2m", mph)
}
