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
