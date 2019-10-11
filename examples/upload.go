package main

import (
	"log"

	"gitlab.com/ebarkie/wunderground"
)

func main() {
	w := wunderground.Pws{ID: "Kssssssnn", Password: "deadbeef"}

	wx := &wunderground.Wx{}
	wx.Bar(30.0)
	wx.OutHumidity(78)
	wx.OutTemp(92.3)

	err := w.Upload(wx)
	if err != nil {
		log.Printf("Upload failed: %s", err)
	}
}
