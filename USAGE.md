# wunderground

```go
import "github.com/ebarkie/wunderground"
```

Package wunderground implements the Weather Underground Personal Weather Station
(PWS) protocol.

## Usage

#### type AqObs

```go
type AqObs obs
```

AqObs represents air quality observations.

#### func (*AqObs) BC

```go
func (a *AqObs) BC(v float64)
```
BC is BC (black carbon at 880 nm) UG/M3.

#### func (*AqObs) CO

```go
func (a *AqObs) CO(v int)
```
CO is CO (carbon monoxide), conventional ppm.

#### func (*AqObs) COT

```go
func (a *AqObs) COT(v int)
```
COT is CO trace levels ppb.

#### func (*AqObs) EC

```go
func (a *AqObs) EC(v float64)
```
EC is EC (elemental carbon) – PM2.5 UG/M3.

#### func (*AqObs) NO

```go
func (a *AqObs) NO(v int)
```
NO is NO (nitric oxide) ppb.

#### func (*AqObs) NO2

```go
func (a *AqObs) NO2(v int)
```
NO2 is NO2 computed, NOx-NO ppb.

#### func (*AqObs) NO2T

```go
func (a *AqObs) NO2T(v int)
```
NO2T is (nitrogen dioxide), true measure ppb.

#### func (*AqObs) NO2Y

```go
func (a *AqObs) NO2Y(v int)
```
NO2Y is NO2 computed, NOy-NO ppb.

#### func (*AqObs) NO3

```go
func (a *AqObs) NO3(v float64)
```
NO3 is NO3 ion (nitrate, not adjusted for ammonium ion) UG/M3.

#### func (*AqObs) NOX

```go
func (a *AqObs) NOX(v int)
```
NOX is NOx (nitrogen oxides) - ppb.

#### func (*AqObs) NOY

```go
func (a *AqObs) NOY(v int)
```
NOY is NOy (total reactive nitrogen) - ppb.

#### func (*AqObs) OC

```go
func (a *AqObs) OC(v float64)
```
OC is OC (organic carbon, not adjusted for oxygen and hydrogen) – PM2.5 UG/M3.

#### func (*AqObs) Ozone

```go
func (a *AqObs) Ozone(v int)
```
Ozone is Ozone - ppb.

#### func (*AqObs) PM10

```go
func (a *AqObs) PM10(v float64)
```
PM10 s PM10 mass - PM10 mass.

#### func (*AqObs) PM25

```go
func (a *AqObs) PM25(v float64)
```
PM25 is PM2.5 mass - UG/M3.

#### func (*AqObs) SO2

```go
func (a *AqObs) SO2(v int)
```
SO2 (sulfur dioxide), conventional ppb.

#### func (*AqObs) SO2T

```go
func (a *AqObs) SO2T(v int)
```
SO2T is trace levels ppb.

#### func (*AqObs) SO4

```go
func (a *AqObs) SO4(v float64)
```
SO4 is SO4 ion (sulfate, not adjusted for ammonium ion) UG/M3.

#### func (*AqObs) UVAETH

```go
func (a *AqObs) UVAETH(v float64)
```
UVAETH is UV-AETH (second channel of Aethalometer at 370 nm) UG/M3.

#### type Ob

```go
type Ob struct {
	Barometer            float64 `xml:"pressure_in"`
	Dewpoint             float64 `xml:"dewpoint_f"`
	RainAccumulationDay  float64 `xml:"precip_today_in"`
	RainAccumulationHour float64 `xml:"precip_1hr_in"`
	OutdoorHumidity      int     `xml:"relative_humidity"`
	OutdoorTemperature   float64 `xml:"temp_f"`
	SolarRadiation       float64 `xml:"solar_radiation"` // Should be an int but WU returns floats sometimes
	Timestamp            Time    `xml:"observation_time_rfc822"`
	UVIndex              float64 `xml:"UV"`
	WindDirection        int     `xml:"wind_degrees"`
	WindSpeed            float64 `xml:"wind_mph"`
	WindGust             float64 `xml:"wind_gust_mph"`
}
```

Ob is the observation info for the associated timestamp.

#### type Time

```go
type Time struct {
	time.Time
}
```

Time is an anonymous struct representing a standard time.Time so we can create a
custom unmarshaler for WU's timestamp format.

#### func (*Time) UnmarshalXML

```go
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```
UnmarshalXML method for converting a WU timestamp string to a time.Time

#### type Wunderground

```go
type Wunderground struct {
	ID           string        // Station name
	Password     string        // Station key
	SoftwareType string        // Software type
	Interval     time.Duration // Upload interval
	URL          string        // Upload URL
	Timestamp    time.Time     // Upload time
	Aq           AqObs         // Air quality observations
	Wx           WxObs         // Weather observations
}
```

Wunderground sets up a PWS for uploading.

#### func  New

```go
func New(id string, pass string) (w Wunderground)
```
New sets up a new Wunderground connection.

#### func (*Wunderground) Clear

```go
func (w *Wunderground) Clear()
```
Clear clears all fields in the payload.

#### func (Wunderground) DownloadDailyHistory

```go
func (w Wunderground) DownloadDailyHistory(t time.Time) (obs []Ob, err error)
```
DownloadDailyHistory downloads all PWS observations for a given day.

#### func (Wunderground) String

```go
func (w Wunderground) String() string
```
UploadRequest returns the request URL. This is for debugging.

#### func (*Wunderground) Upload

```go
func (w *Wunderground) Upload() (skipped bool, err error)
```
Upload performs the actual upload to Wunderground.

#### type WxObs

```go
type WxObs obs
```

WxObs represents weather observations.

#### func (WxObs) Barometer

```go
func (w WxObs) Barometer(v float64)
```
Barometer is barometric pressure inches.

#### func (*WxObs) Clouds

```go
func (w *WxObs) Clouds(v string)
```
Clouds is SKC, FEW, SCT, BKN, OVC.

#### func (*WxObs) DailyRain

```go
func (w *WxObs) DailyRain(v float64)
```
DailyRain is rain inches so far today in local time.

#### func (*WxObs) DewPoint

```go
func (w *WxObs) DewPoint(v float64)
```
DewPoint is F outdoor dewpoint.

#### func (*WxObs) IndoorHumidity

```go
func (w *WxObs) IndoorHumidity(v int)
```
IndoorHumidity is % indoor humidity 0-100%.

#### func (*WxObs) IndoorTemperature

```go
func (w *WxObs) IndoorTemperature(v float64)
```
IndoorTemperature is indoor temperature in F.

#### func (*WxObs) LeafWetness

```go
func (w *WxObs) LeafWetness(v int)
```
LeafWetness is %.

#### func (*WxObs) OutdoorHumidity

```go
func (w *WxObs) OutdoorHumidity(v int)
```
OutdoorHumidity is % outdoor humidity 0-100%.

#### func (*WxObs) OutdoorTemperature

```go
func (w *WxObs) OutdoorTemperature(v float64)
```
OutdoorTemperature is outdoor temperature in F.

#### func (*WxObs) RainRate

```go
func (w *WxObs) RainRate(v float64)
```
RainRate is rain inches over the past hour or the accumulated rainfall in the
past 60 min.

#### func (*WxObs) SoilMoisture

```go
func (w *WxObs) SoilMoisture(v int)
```
SoilMoisture is %.

#### func (*WxObs) SoilTemperature

```go
func (w *WxObs) SoilTemperature(v float64)
```
SoilTemperature is F soil temperature.

#### func (*WxObs) SolarRadiation

```go
func (w *WxObs) SolarRadiation(v int)
```
SolarRadiation is solar radiation in W/m^2.

#### func (*WxObs) UVIndex

```go
func (w *WxObs) UVIndex(v float64)
```
UVIndex is the UV index.

#### func (*WxObs) Visibility

```go
func (w *WxObs) Visibility(v int)
```
Visibility is nm visibility.

#### func (*WxObs) WindDirection

```go
func (w *WxObs) WindDirection(v int)
```
WindDirection is 0-360 instantaneous wind direction.

#### func (*WxObs) WindGustDirection

```go
func (w *WxObs) WindGustDirection(v int)
```
WindGustDirection is 0-360 using software specific time period.

#### func (*WxObs) WindGustDirection10m

```go
func (w *WxObs) WindGustDirection10m(v int)
```
WindGustDirection10m is 0-360 past 10 minutes wind gust direction.

#### func (*WxObs) WindGustSpeed

```go
func (w *WxObs) WindGustSpeed(v float64)
```
WindGustSpeed is mph current wind gust, using software specific time period.

#### func (*WxObs) WindGustSpeed10m

```go
func (w *WxObs) WindGustSpeed10m(v float64)
```
WindGustSpeed10m is mph past 10 minutes wind gust mph.

#### func (*WxObs) WindSpeed

```go
func (w *WxObs) WindSpeed(v float64)
```
WindSpeed is mph instantaneous wind speed.

#### func (*WxObs) WindSpeedAverage2m

```go
func (w *WxObs) WindSpeedAverage2m(v float64)
```
WindSpeedAverage2m is mph 2 minute average wind speed mph.
