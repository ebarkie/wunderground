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

	"github.com/ebarkie/wunderground"
)

func main() {
	w := wunderground.Pws{ID: "Kssssssnn", Password: "deadbeef"}

	wx := &wunderground.Wx{}
	wx.Barometer(30.0)
	wx.OutdoorHumidity(78)
	wx.OutdoorTemperature(92.3)

	err := w.Upload(wx)
	if err != nil {
		log.Printf("Upload failed: %s", err)
	}
}
```

## License

Copyright (c) 2016-2017 Eric Barkie. All rights reserved.  
Use of this source code is governed by the MIT license
that can be found in the [LICENSE](LICENSE) file.
