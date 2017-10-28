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

import "github.com/ebarkie/wunderground/internal/request"

// DownloadDaily downloads all PWS observations for a given day.
func (p Pws) DownloadDaily(t time.Time) (obs []WxObs, err error) {
	// Build HTTP request.
	req, _ := http.NewRequest("GET", DownloadURL+"/WXDailyHistory.asp", nil)
	q := req.URL.Query()
	q.Add("ID", p.ID)
	q.Add("year", request.Int(t.Year()))
	q.Add("month", request.Int(int(t.Month())))
	q.Add("day", request.Int(t.Day()))
	q.Add("format", "XML")
	req.URL.RawQuery = q.Encode()

	// Initiate HTTP request.
	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
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
	body, _ := ioutil.ReadAll(resp.Body)
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
