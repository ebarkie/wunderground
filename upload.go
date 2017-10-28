// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package wunderground

// Refer to:
// http://wiki.wunderground.com/index.php/PWS_-_Upload_Protocol

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const rapidFireDuration = 5 * time.Minute

// lastUpload tracks the last successful upload time and request.  It
// is used to determine if anything has changed so unnecessary RapidFire
// uploads can be avoided.
type lastUpload struct {
	time    time.Time
	query   string
	skipped bool
}

type obs interface {
	Clear()
	Values() map[string]string
}

// rapidFire indicates if the interval is RapidFire.
func (p Pws) rapidFire() bool {
	return p.Interval < rapidFireDuration
}

// createRequest builds the HTTP request.
func (p Pws) createRequest(ob ...obs) *http.Request {
	req, _ := http.NewRequest("GET", UploadURL+"/updateweatherstation.php", nil)

	// Create mandatory GET query parameters.
	q := req.URL.Query()
	q.Add("action", "updateraw")
	q.Add("ID", strings.ToUpper(p.ID))
	q.Add("PASSWORD", p.Password)
	if p.SoftwareType == "" {
		q.Add("softwaretype", "GoWunder 1337")
	} else {
		q.Add("softwaretype", p.SoftwareType)
	}
	if p.Time.IsZero() {
		q.Add("dateutc", "now")
	} else {
		q.Add("dateutc", p.Time.In(time.UTC).Format("2006-01-02 15:04:05"))
	}
	if p.rapidFire() {
		q.Add("realtime", "1")
		if p.Interval == 0 {
			q.Add("rtfreq", "2")
		} else {
			q.Add("rtfreq", fmt.Sprintf("%d", p.Interval/time.Second))
		}
	}

	// Add observations to  GET query parameters.
	for _, o := range ob {
		for k, v := range o.Values() {
			q.Add(k, v)
		}
	}

	// Do a '+' -> %20 replacement here since WU doesn't decode the '+'
	// properly and we use it in softwaretype, maybe elsewhere too.
	req.URL.RawQuery = strings.Replace(q.Encode(), "+", "%20", -1)

	return req
}

// Encode returns the request URL for the specified observations.  This
// is generally used for testing and debugging.
func (p Pws) Encode(ob ...obs) string {
	return p.createRequest(ob...).URL.String()
}

// Skipped indicates if the last upload was skipped or not.
func (p Pws) Skipped() bool {
	return p.last.skipped
}

// Upload uploads the specified observations.
func (p *Pws) Upload(ob ...obs) (err error) {
	// Clear payload(s) after upload attempt.
	defer func() {
		for _, o := range ob {
			o.Clear()
		}
	}()

	// Skip RapidFire upload if nothing has changed.
	if p.rapidFire() &&
		p.last.query == p.Encode(ob...) &&
		(time.Since(p.last.time)+p.Interval < rapidFireDuration) {
		p.last.skipped = true
		return
	}
	p.last.skipped = false

	// Build fake HTTP response if test station was specified, otherwise
	// initiate the real HTTP request.
	client := &http.Client{}
	var resp *http.Response
	if p.ID == "KTEST0" {
		resp = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("success\r\n")),
		}
	} else {
		resp, err = client.Do(p.createRequest(ob...))
	}
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
	if strings.TrimRight(string(body), "\n\r") != "success" {
		err = fmt.Errorf("HTTP request returned bad body: %s", body)
		return
	}

	// Record last successful upload time and query.
	p.last.time = time.Now()
	p.last.query = p.Encode(ob...)

	return
}
