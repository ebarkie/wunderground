// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package wunderground implements the Weather Underground Personal Weather
// Station (PWS) protocol.
package wunderground

// Refer to:
// http://wiki.wunderground.com/index.php/PWS_-_Upload_Protocol

import (
	"net/http"
	"time"
)

// Base download and upload URL's.
var (
	DownloadURL = "https://www.wunderground.com/weatherstation"
	UploadURL   = "https://rtupdate.wunderground.com/weatherstation"
)

var httpClient = &http.Client{
	Timeout: time.Duration(8 * time.Second),
}

// Pws represents a Personal Weather Station.
type Pws struct {
	ID           string        // Station name
	Password     string        // Station key
	SoftwareType string        // Software type
	Interval     time.Duration // Upload interval (default is 2s)
	Time         time.Time     // Upload time (default is now)
	last         lastUpload
}
