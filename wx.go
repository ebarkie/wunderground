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

// Wx represents weather measurements.
type Wx struct {
	request.Data
}

// Barometer is barometric pressure inches.
func (w *Wx) Barometer(in float64) {
	w.SetFloat("baromin", in)
}

// Clouds is SKC, FEW, SCT, BKN, OVC.
func (w *Wx) Clouds(c string) {
	w.SetString("clouds", c)
}

// DailyRain is rain inches so far today in local time.
func (w *Wx) DailyRain(in float64) {
	w.SetFloat("dailyrainin", in)
}

// DewPoint is F outdoor dewpoint.
func (w *Wx) DewPoint(f float64) {
	w.SetFloat("dewptf", f)
}

// IndoorHumidity is % indoor humidity 0-100%.
func (w *Wx) IndoorHumidity(p int) {
	w.SetInt("indoorhumudity", p)
}

// IndoorTemperature is indoor temperature in F.
func (w *Wx) IndoorTemperature(f float64) {
	w.SetFloat("indoortempf", f)
}

// LeafWetness is %.
func (w *Wx) LeafWetness(p int) {
	w.SetIntf(request.Indexed{Format: "leafwetness#", Begin: 1, Zero: 1}, p)
}

// OutdoorHumidity is % outdoor humidity 0-100%.
func (w *Wx) OutdoorHumidity(p int) {
	w.SetInt("humidity", p)
}

// OutdoorTemperature is outdoor temperature in F.
func (w *Wx) OutdoorTemperature(f float64) {
	w.SetFloatf(request.Indexed{Format: "temp#f", Begin: 1, Zero: 1}, f)
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall in the past 60 min.
func (w *Wx) RainRate(in float64) {
	w.SetFloat("rainin", in)
}

// SolarRadiation is solar radiation in W/m^2.
func (w *Wx) SolarRadiation(wm2 int) {
	w.SetInt("solarradiation", wm2)
}

// SoilMoisture is %.
func (w *Wx) SoilMoisture(p int) {
	w.SetIntf(request.Indexed{Format: "soilmoisture#", Begin: 1, Zero: 1}, p)
}

// SoilTemperature is F soil temperature.
func (w *Wx) SoilTemperature(f float64) {
	w.SetFloatf(request.Indexed{Format: "soiltemp#f", Begin: 1, Zero: 1}, f)
}

// UVIndex is the UV index.
func (w *Wx) UVIndex(i float64) {
	w.SetFloat("UV", i)
}

// Visibility is nm visibility.
func (w *Wx) Visibility(nm int) {
	w.SetInt("visibility", nm)
}

// WindDirection is 0-360 instantaneous wind direction.
func (w *Wx) WindDirection(deg int) {
	w.SetInt("winddir", deg)
}

// WindGustDirection is 0-360 using software specific time period.
func (w *Wx) WindGustDirection(deg int) {
	w.SetInt("windgustdir", deg)
}

// WindGustDirection10m is 0-360 past 10 minutes wind gust direction.
func (w *Wx) WindGustDirection10m(deg int) {
	w.SetInt("windgustdir_10m", deg)
}

// WindGustSpeed is mph current wind gust, using software specific time period.
func (w *Wx) WindGustSpeed(mph float64) {
	w.SetFloat("windgustmph", mph)
}

// WindGustSpeed10m is mph past 10 minutes wind gust mph.
func (w *Wx) WindGustSpeed10m(mph float64) {
	w.SetFloat("windgustmph_10m", mph)
}

// WindSpeed is mph instantaneous wind speed.
func (w *Wx) WindSpeed(mph float64) {
	w.SetFloat("windspeedmph", mph)
}

// WindSpeedAverage2m is mph 2 minute average wind speed mph.
func (w *Wx) WindSpeedAverage2m(mph float64) {
	w.SetFloat("windspdmph_avg2m", mph)
}
