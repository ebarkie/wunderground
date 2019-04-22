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
	query.Data
}
```

Aq represents air quality observations.

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
func (p Pws) Encode(obs ...query.Values) string
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
func (p *Pws) Upload(obs ...query.Values) (err error)
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
	query.Data
}
```

Wx represents weather observations.

#### func (*Wx) Bar

```go
func (w *Wx) Bar(in float64)
```
Bar is barometric pressure in inches.

#### func (*Wx) Clouds

```go
func (w *Wx) Clouds(c string)
```
Clouds is the cloud cover specified as: SKC, FEW, SCT, BKN, OVC.

#### func (*Wx) DailyRain

```go
func (w *Wx) DailyRain(in float64)
```
DailyRain is rain so far today (local time) in inches.

#### func (*Wx) DewPoint

```go
func (w *Wx) DewPoint(f float64)
```
DewPoint is the outdoor dew point in degrees Fahrenheit.

#### func (*Wx) InHumidity

```go
func (w *Wx) InHumidity(p int)
```
InHumidity is the indoor humidity percentage (0-100).

#### func (*Wx) InTemp

```go
func (w *Wx) InTemp(f float64)
```
InTemp is the indoor temperature in degrees Fahrenheit.

#### func (*Wx) LeafWetness

```go
func (w *Wx) LeafWetness(p int)
```
LeafWetness is the leaf wetness percentage (0-100).

#### func (*Wx) OutHumidity

```go
func (w *Wx) OutHumidity(p int)
```
OutHumidity is the outdoor humidity percentage (0-100).

#### func (*Wx) OutTemp

```go
func (w *Wx) OutTemp(f float64)
```
OutTemp is outdoor temperature in degrees Fahrenheit.

#### func (*Wx) RainRate

```go
func (w *Wx) RainRate(in float64)
```
RainRate is rain inches over the past hour or the accumulated rainfall for the
past 60 minutes in inches.

#### func (*Wx) SoilMoist

```go
func (w *Wx) SoilMoist(p int)
```
SoilMoist is the soil moisture percentage (0-100)..

#### func (*Wx) SoilTemp

```go
func (w *Wx) SoilTemp(f float64)
```
SoilTemp is the soil temperature in degrees Fahrenheit.

#### func (*Wx) SolarRad

```go
func (w *Wx) SolarRad(wm2 int)
```
SolarRad is solar radiation in watts per square meter.

#### func (*Wx) UVIndex

```go
func (w *Wx) UVIndex(i float64)
```
UVIndex is the UltraViolet light index.

#### func (*Wx) Visibility

```go
func (w *Wx) Visibility(nm int)
```
Visibility is visibility in nautical miles.

#### func (*Wx) WindDir

```go
func (w *Wx) WindDir(deg int)
```
WindDir is instantaneous wind direction in degrees (0-359).

#### func (*Wx) WindGustDir

```go
func (w *Wx) WindGustDir(deg int)
```
WindGustDir is the software specific time period wind gust direction in degrees
(0-359).

#### func (*Wx) WindGustDir10m

```go
func (w *Wx) WindGustDir10m(deg int)
```
WindGustDir10m is the 10 minute wind gust direction in degrees (0-359).

#### func (*Wx) WindGustSpeed

```go
func (w *Wx) WindGustSpeed(mph float64)
```
WindGustSpeed is the software specific time period wind gust speed in miles per
hour.

#### func (*Wx) WindGustSpeed10m

```go
func (w *Wx) WindGustSpeed10m(mph float64)
```
WindGustSpeed10m is the 10 minute wind gust speed in miles per hour.

#### func (*Wx) WindSpeed

```go
func (w *Wx) WindSpeed(mph float64)
```
WindSpeed is the instantaneous wind speed in miles per hour.

#### func (*Wx) WindSpeedAvg2m

```go
func (w *Wx) WindSpeedAvg2m(mph float64)
```
WindSpeedAvg2m is the 2 minute average wind speed in miles per hour.

#### type WxObs

```go
type WxObs struct {
	Bar           float64 `xml:"pressure_in"`
	DewPoint      float64 `xml:"dewpoint_f"`
	RainAccumDay  float64 `xml:"precip_today_in"`
	RainAccumHour float64 `xml:"precip_1hr_in"`
	OutHumidity   int     `xml:"relative_humidity"`
	OutTemp       float64 `xml:"temp_f"`
	SolarRad      float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Time          Time    `xml:"observation_time_rfc822"`
	UVIndex       float64 `xml:"UV"`
	WindDir       int     `xml:"wind_degrees"`
	WindSpeed     float64 `xml:"wind_mph"`
	WindGust      float64 `xml:"wind_gust_mph"`
}
```

WxObs is weather observation info for the associated timestamp.
