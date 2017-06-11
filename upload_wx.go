// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

// Refer to:
// http://wiki.wunderground.com/index.php/PWS_-_Upload_Protocol

// Barometer is barometric pressure inches.
func (w WxObs) Barometer(v float64) {
	w.p["baromin"] = f2s(v)
}

// Clouds is SKC, FEW, SCT, BKN, OVC.
func (w *WxObs) Clouds(v string) {
	w.p["clouds"] = v
}

// DailyRain is rain inches so far today in local time.
func (w *WxObs) DailyRain(v float64) {
	w.p["dailyrainin"] = f2s(v)
}

// DewPoint is F outdoor dewpoint.
func (w *WxObs) DewPoint(v float64) {
	w.p["dewptf"] = f2s(v)
}

// IndoorHumidity is % indoor humidity 0-100%.
func (w *WxObs) IndoorHumidity(v int) {
	w.p["indoorhumudity"] = i2s(v)
}

// IndoorTemperature is indoor temperature in F.
func (w *WxObs) IndoorTemperature(v float64) {
	w.p["indoortempf"] = f2s(v)
}

// LeafWetness is %.
func (w *WxObs) LeafWetness(v int) {
	w.p[incrKey(w.p, "leafwetnessX")] = i2s(v)
}

// OutdoorHumidity is % outdoor humidity 0-100%.
func (w *WxObs) OutdoorHumidity(v int) {
	w.p["humidity"] = i2s(v)
}

// OutdoorTemperature is outdoor temperature in F.
func (w *WxObs) OutdoorTemperature(v float64) {
	w.p[incrKey(w.p, "tempXf")] = f2s(v)
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall in the past 60 min.
func (w *WxObs) RainRate(v float64) {
	w.p["rainin"] = f2s(v)
}

// SolarRadiation is solar radiation in W/m^2.
func (w *WxObs) SolarRadiation(v int) {
	w.p["solarradiation"] = i2s(v)
}

// SoilMoisture is %.
func (w *WxObs) SoilMoisture(v int) {
	w.p[incrKey(w.p, "soilmoistureX")] = i2s(v)
}

// SoilTemperature is F soil temperature.
func (w *WxObs) SoilTemperature(v float64) {
	w.p[incrKey(w.p, "soiltempXf")] = f2s(v)
}

// UVIndex is the UV index.
func (w *WxObs) UVIndex(v float64) {
	w.p["UV"] = f2s(v)
}

// Visibility is nm visibility.
func (w *WxObs) Visibility(v int) {
	w.p["visibility"] = i2s(v)
}

// WindDirection is 0-360 instantaneous wind direction.
func (w *WxObs) WindDirection(v int) {
	w.p["winddir"] = i2s(v)
}

// WindGustDirection is 0-360 using software specific time period.
func (w *WxObs) WindGustDirection(v int) {
	w.p["windgustdir"] = i2s(v)
}

// WindGustDirection10m is 0-360 past 10 minutes wind gust direction.
func (w *WxObs) WindGustDirection10m(v int) {
	w.p["windgustdir_10m"] = i2s(v)
}

// WindGustSpeed is mph current wind gust, using software specific time period.
func (w *WxObs) WindGustSpeed(v float64) {
	w.p["windgustmph"] = f2s(v)
}

// WindGustSpeed10m is mph past 10 minutes wind gust mph.
func (w *WxObs) WindGustSpeed10m(v float64) {
	w.p["windgustmph_10m"] = f2s(v)
}

// WindSpeed is mph instantaneous wind speed.
func (w *WxObs) WindSpeed(v float64) {
	w.p["windspeedmph"] = f2s(v)
}

// WindSpeedAverage2m is mph 2 minute average wind speed mph.
func (w *WxObs) WindSpeedAverage2m(v float64) {
	w.p["windspdmph_avg2m"] = f2s(v)
}
