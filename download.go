// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// dayObs is the parent struct for all observations.  It's
// only here for XML parsing.
type dayObs struct {
	Obs []Ob `xml:"current_observation"`
}

// Ob is the observation info for the associated timestamp.
type Ob struct {
	Barometer            float64 `xml:"pressure_in"`
	Dewpoint             float64 `xml:"dewpoint_f"`
	RainAccumulationDay  float64 `xml:"precip_today_in"`
	RainAccumulationHour float64 `xml:"precip_1hr_in"`
	OutdoorHumidity      int     `xml:"relative_humidity"`
	OutdoorTemperature   float64 `xml:"temp_f"`
	SolarRadiation       float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Timestamp            Time    `xml:"observation_time_rfc822"`
	UVIndex              float64 `xml:"UV"`
	WindDirection        int     `xml:"wind_degrees"`
	WindSpeed            float64 `xml:"wind_mph"`
	WindGust             float64 `xml:"wind_gust_mph"`
}

// DownloadDailyHistory downloads all PWS observations for a given day.
func (w Wunderground) DownloadDailyHistory(t time.Time) (obs []Ob, err error) {
	const downloadURL = "https://www.wunderground.com/weatherstation"

	// Build HTTP request
	req, _ := http.NewRequest("GET", downloadURL+"/WXDailyHistory.asp", nil)
	q := req.URL.Query()
	q.Add("ID", w.ID)
	q.Add("year", i2s(t.Year()))
	q.Add("month", i2s(int(t.Month())))
	q.Add("day", i2s(t.Day()))
	q.Add("format", "XML")
	req.URL.RawQuery = q.Encode()

	// Initiate HTTP request
	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check HTTP return status code
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP request returned non-OK status code %d", resp.StatusCode)
		return
	}

	// Parse response
	body, _ := ioutil.ReadAll(resp.Body)
	var day dayObs
	err = xml.Unmarshal(body, &day)
	if err != nil {
		return
	}
	obs = day.Obs

	return
}
