# Wunderground

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](http://choosealicense.com/licenses/mit/)
[![Build Status](https://travis-ci.org/ebarkie/wunderground.svg?branch=master)](https://travis-ci.org/ebarkie/wunderground)

Go package for uploading and downloading [Weather Underground](http://www.wunderground.com) Personal Weather Station observations.

## Installation

```
$ go get github.com/ebarkie/wunderground
```

## Usage

See [USAGE](USAGE.md).

## Example

```go
package main

import (
	"log"
	"time"

	"github.com/ebarkie/wunderground"
)

func main() {
	w := wunderground.New("Kssssssnn", "deadbeef")

	w.Wx.Barometer(30.0)
	w.Wx.OutdoorHumidity(78)
	w.Wx.OutdoorTemperature(92.3)

	_, err := w.Upload()
	if err != nil {
		log.Printf("Upload failed: %s", err)
	}
}
```

## License

Copyright (c) 2016-2017 Eric Barkie. All rights reserved.  
Use of this source code is governed by the MIT license
that can be found in the [LICENSE](LICENSE) file.
