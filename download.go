// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

import "strconv"

// DownloadDaily downloads all PWS observations for a given day.
func (p Pws) DownloadDaily(t time.Time) (obs []WxObs, err error) {
	// Build HTTP request.
	req, _ := http.NewRequest("GET", DownloadURL+"/WXDailyHistory.asp", nil)
	q := req.URL.Query()
	q.Add("ID", p.ID)
	q.Add("year", strconv.Itoa(t.Year()))
	q.Add("month", strconv.Itoa(int(t.Month())))
	q.Add("day", strconv.Itoa(t.Day()))
	q.Add("format", "XML")
	req.URL.RawQuery = q.Encode()

	// Initiate HTTP request.
	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check HTTP return status code.
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP request returned non-OK status code %d", resp.StatusCode)
		return
	}

	// Parse response.
	body, _ := io.ReadAll(resp.Body)
	day := struct {
		Current []WxObs `xml:"current_observation"`
	}{}
	err = xml.Unmarshal(body, &day)
	if err != nil {
		return
	}
	obs = day.Current

	return
}
