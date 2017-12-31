# wunderground

```go
import "github.com/ebarkie/wunderground"
```

Package wunderground implements the Weather Underground Personal Weather Station
(PWS) protocol.

## Usage

```go
var (
	DownloadURL = "https://www.wunderground.com/weatherstation"
	UploadURL   = "https://rtupdate.wunderground.com/weatherstation"
)
```
Base download and upload URL's.

#### type Aq

```go
type Aq struct {
	request.Data
}
```

Aq represents air quality measurements.

#### func (*Aq) BC

```go
func (a *Aq) BC(ugm3 float64)
```
BC is BC (black carbon at 880 nm) UG/M3.

#### func (*Aq) CO

```go
func (a *Aq) CO(ppm int)
```
CO is CO (carbon monoxide), caentional ppm.

#### func (*Aq) COT

```go
func (a *Aq) COT(ppb int)
```
COT is CO trace levels ppb.

#### func (*Aq) EC

```go
func (a *Aq) EC(ugm3 float64)
```
EC is EC (elemental carbon) – PM2.5 UG/M3.

#### func (*Aq) NO

```go
func (a *Aq) NO(ppb int)
```
NO is NO (nitric oxide) ppb.

#### func (*Aq) NO2

```go
func (a *Aq) NO2(ppb int)
```
NO2 is NO2 computed, NOx-NO ppb.

#### func (*Aq) NO2T

```go
func (a *Aq) NO2T(ppb int)
```
NO2T is (nitrogen dioxide), true measure ppb.

#### func (*Aq) NO2Y

```go
func (a *Aq) NO2Y(ppb int)
```
NO2Y is NO2 computed, NOy-NO ppb.

#### func (*Aq) NO3

```go
func (a *Aq) NO3(ugm3 float64)
```
NO3 is NO3 ion (nitrate, not adjusted for ammonium ion) UG/M3.

#### func (*Aq) NOX

```go
func (a *Aq) NOX(ppb int)
```
NOX is NOx (nitrogen oxides) - ppb.

#### func (*Aq) NOY

```go
func (a *Aq) NOY(ppb int)
```
NOY is NOy (total reactive nitrogen) - ppb.

#### func (*Aq) OC

```go
func (a *Aq) OC(ugm3 float64)
```
OC is OC (organic carbon, not adjusted for oxygen and hydrogen) – PM2.5 UG/M3.

#### func (*Aq) Ozone

```go
func (a *Aq) Ozone(ppb int)
```
Ozone is Ozone - ppb.

#### func (*Aq) PM10

```go
func (a *Aq) PM10(ugm3 float64)
```
PM10 s PM10 mass - PM10 mass.

#### func (*Aq) PM25

```go
func (a *Aq) PM25(ugm3 float64)
```
PM25 is PM2.5 mass - UG/M3.

#### func (*Aq) SO2

```go
func (a *Aq) SO2(ppb int)
```
SO2 (sulfur dioxide), caentional ppb.

#### func (*Aq) SO2T

```go
func (a *Aq) SO2T(ppb int)
```
SO2T is trace levels ppb.

#### func (*Aq) SO4

```go
func (a *Aq) SO4(ugm3 float64)
```
SO4 is SO4 ion (sulfate, not adjusted for ammonium ion) UG/M3.

#### func (*Aq) UVAETH

```go
func (a *Aq) UVAETH(ugm3 float64)
```
UVAETH is UV-AETH (second channel of Aethalometer at 370 nm) UG/M3.

#### type Pws

```go
type Pws struct {
	ID           string        // Station name
	Password     string        // Station key
	SoftwareType string        // Software type
	Interval     time.Duration // Upload interval (default is 2s)
	Time         time.Time     // Upload time (default is now)
}
```

Pws represents a Personal Weather Station.

#### func (Pws) DownloadDaily

```go
func (p Pws) DownloadDaily(t time.Time) (obs []WxObs, err error)
```
DownloadDaily downloads all PWS observations for a given day.

#### func (Pws) Encode

```go
func (p Pws) Encode(ob ...obs) string
```
Encode returns the request URL for the specified observations. This is generally
used for testing and debugging.

#### func (Pws) Skipped

```go
func (p Pws) Skipped() bool
```
Skipped indicates if the last upload was skipped or not.

#### func (*Pws) Upload

```go
func (p *Pws) Upload(ob ...obs) (err error)
```
Upload uploads the specified observations.

#### type Time

```go
type Time struct {
	time.Time
}
```

Time represents a time that can be unmarshaled from a Weather Underground XML
response.

#### func (*Time) UnmarshalXML

```go
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```
UnmarshalXML method for converting a Weather Underground time string to a
time.Time.

#### type Wx

```go
type Wx struct {
	request.Data
}
```

Wx represents weather measurements.

#### func (*Wx) Barometer

```go
func (w *Wx) Barometer(in float64)
```
Barometer is barometric pressure inches.

#### func (*Wx) Clouds

```go
func (w *Wx) Clouds(c string)
```
Clouds is SKC, FEW, SCT, BKN, OVC.

#### func (*Wx) DailyRain

```go
func (w *Wx) DailyRain(in float64)
```
DailyRain is rain inches so far today in local time.

#### func (*Wx) DewPoint

```go
func (w *Wx) DewPoint(f float64)
```
DewPoint is F outdoor dewpoint.

#### func (*Wx) IndoorHumidity

```go
func (w *Wx) IndoorHumidity(p int)
```
IndoorHumidity is % indoor humidity 0-100%.

#### func (*Wx) IndoorTemperature

```go
func (w *Wx) IndoorTemperature(f float64)
```
IndoorTemperature is indoor temperature in F.

#### func (*Wx) LeafWetness

```go
func (w *Wx) LeafWetness(p int)
```
LeafWetness is %.

#### func (*Wx) OutdoorHumidity

```go
func (w *Wx) OutdoorHumidity(p int)
```
OutdoorHumidity is % outdoor humidity 0-100%.

#### func (*Wx) OutdoorTemperature

```go
func (w *Wx) OutdoorTemperature(f float64)
```
OutdoorTemperature is outdoor temperature in F.

#### func (*Wx) RainRate

```go
func (w *Wx) RainRate(in float64)
```
RainRate is rain inches over the past hour or the accumulated rainfall in the
past 60 min.

#### func (*Wx) SoilMoisture

```go
func (w *Wx) SoilMoisture(p int)
```
SoilMoisture is %.

#### func (*Wx) SoilTemperature

```go
func (w *Wx) SoilTemperature(f float64)
```
SoilTemperature is F soil temperature.

#### func (*Wx) SolarRadiation

```go
func (w *Wx) SolarRadiation(wm2 int)
```
SolarRadiation is solar radiation in W/m^2.

#### func (*Wx) UVIndex

```go
func (w *Wx) UVIndex(i float64)
```
UVIndex is the UV index.

#### func (*Wx) Visibility

```go
func (w *Wx) Visibility(nm int)
```
Visibility is nm visibility.

#### func (*Wx) WindDirection

```go
func (w *Wx) WindDirection(deg int)
```
WindDirection is 0-360 instantaneous wind direction.

#### func (*Wx) WindGustDirection

```go
func (w *Wx) WindGustDirection(deg int)
```
WindGustDirection is 0-360 using software specific time period.

#### func (*Wx) WindGustDirection10m

```go
func (w *Wx) WindGustDirection10m(deg int)
```
WindGustDirection10m is 0-360 past 10 minutes wind gust direction.

#### func (*Wx) WindGustSpeed

```go
func (w *Wx) WindGustSpeed(mph float64)
```
WindGustSpeed is mph current wind gust, using software specific time period.

#### func (*Wx) WindGustSpeed10m

```go
func (w *Wx) WindGustSpeed10m(mph float64)
```
WindGustSpeed10m is mph past 10 minutes wind gust mph.

#### func (*Wx) WindSpeed

```go
func (w *Wx) WindSpeed(mph float64)
```
WindSpeed is mph instantaneous wind speed.

#### func (*Wx) WindSpeedAverage2m

```go
func (w *Wx) WindSpeedAverage2m(mph float64)
```
WindSpeedAverage2m is mph 2 minute average wind speed mph.

#### type WxObs

```go
type WxObs struct {
	Barometer            float64 `xml:"pressure_in"`
	DewPoint             float64 `xml:"dewpoint_f"`
	RainAccumulationDay  float64 `xml:"precip_today_in"`
	RainAccumulationHour float64 `xml:"precip_1hr_in"`
	OutdoorHumidity      int     `xml:"relative_humidity"`
	OutdoorTemperature   float64 `xml:"temp_f"`
	SolarRadiation       float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Time                 Time    `xml:"observation_time_rfc822"`
	UVIndex              float64 `xml:"UV"`
	WindDirection        int     `xml:"wind_degrees"`
	WindSpeed            float64 `xml:"wind_mph"`
	WindGust             float64 `xml:"wind_gust_mph"`
}
```

WxObs is weather observation info for the associated timestamp.
