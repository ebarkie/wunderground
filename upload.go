// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package wunderground implements the Weather Underground Personal Weather
// Station (PWS) protocol.
package wunderground

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const rapidFireThreshold = 5 * time.Minute

// Wunderground sets up a PWS for uploading.
type Wunderground struct {
	ID           string        // Station name
	Password     string        // Station key
	SoftwareType string        // Software type
	Interval     time.Duration // Upload interval
	URL          string        // Upload URL
	Timestamp    time.Time     // Upload time
	Aq           AqObs         // Air quality observations
	Wx           WxObs         // Weather observations
	p            params        // Map of HTTP request parameters
	last         lastUpload    // Last upload -- so uploads can skipped if data is unchanged
}

// Map of HTTP request parameters.
type params map[string]string

// obs is an artificially introduced nesting to make things more
// organized.  It allows us to expose air quality observations as
// Wunderground.Aq and weather as Wundeground.Wx but under the covers
// we point the parameter maps back to the root Wunderground struct.
type obs struct {
	p params
}

// AqObs represents air quality observations.
type AqObs obs

// WxObs represents weather observations.
type WxObs obs

// lastUpload tracks the last successful upload time and payload that
// was sent.  This is used to determine if anything has changed so
// unnecessary RapidFire uploads can be avoided.
type lastUpload struct {
	time time.Time
	p    params
}

// New sets up a new Wunderground connection.
func New(id string, pass string) (w Wunderground) {
	const (
		rapidFireURL = "https://rtupdate.wunderground.com/weatherstation/updateweatherstation.php"
		uploadURL    = "https://weatherstation.wunderground.com/weatherstation/updateweatherstation.php"
	)

	// Initialize parameters
	w.Clear()

	// User settings
	w.ID = id
	w.Password = pass

	// Defaults
	w.Interval = 2 * time.Second
	if w.isRapidFire() {
		w.URL = rapidFireURL
	} else {
		w.URL = uploadURL
	}
	w.SoftwareType = "GoWunder 1337"

	return
}

// isRapidFire returns if we are running in RapidFire mode or not.
func (w Wunderground) isRapidFire() bool {
	return w.Interval < rapidFireThreshold
}

// request creates the HTTP request.
func (w Wunderground) createRequest() (req *http.Request) {
	req, _ = http.NewRequest("GET", w.URL, nil)

	// Create GET query parameters
	q := req.URL.Query()
	q.Add("action", "updateraw")
	q.Add("ID", w.ID)
	q.Add("PASSWORD", w.Password)
	q.Add("softwaretype", w.SoftwareType)
	if w.Timestamp.IsZero() {
		q.Add("dateutc", "now")
	} else {
		q.Add("dateutc", w.Timestamp.In(time.UTC).Format("2006-01-02 15:04:05"))
	}
	if w.isRapidFire() {
		q.Add("realtime", "1")
		q.Add("rtfreq", fmt.Sprintf("%d", w.Interval/time.Second))
	}
	for k, v := range w.p {
		q.Add(k, v)
	}

	// Do a '+' -> %20 replacement here since WU doesn't decode the '+'
	// properly and we use it in softwaretype, maybe elsewhere too.
	req.URL.RawQuery = strings.Replace(q.Encode(), "+", "%20", -1)

	return
}

// Clear clears all fields in the payload.
func (w *Wunderground) Clear() {
	w.Timestamp = time.Time{}

	w.p = make(params)
	w.Aq.p = w.p
	w.Wx.p = w.p
}

// UploadRequest returns the request URL.  This is for debugging.
func (w Wunderground) String() string {
	return w.createRequest().URL.String()
}

// Upload performs the actual upload to Wunderground.
func (w *Wunderground) Upload() (skipped bool, err error) {
	defer w.Clear()

	// Skip RapidFire upload if nothing has changed
	if w.isRapidFire() &&
		reflect.DeepEqual(w.last.p, w.p) &&
		((time.Since(w.last.time) + w.Interval) < rapidFireThreshold) {
		skipped = true

		return
	}

	// Initiate HTTP request
	client := &http.Client{}
	resp, err := client.Do(w.createRequest())
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
	if strings.TrimRight(string(body), "\n\r") != "success" {
		err = fmt.Errorf("HTTP request returned bad body: %s", body)
		return
	}

	// In RapidFire mode track last successful payload to avoid
	// unnecessary uploads.
	if w.isRapidFire() {
		w.last.p = w.p
		w.last.time = time.Now()
	}

	return
}
