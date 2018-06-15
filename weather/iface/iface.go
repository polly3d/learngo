package iface

import (
	"log"
	"time"
)

type WeatherCode int

const (
	CodeUnknown WeatherCode = iota
	CodeCloudy
	CodeFog
	CodeHeavyRain
	CodeHeavyShowers
	CodeHeavySnow
	CodeHeavySnowShowers
	CodeLightRain
	CodeLightShowers
	CodeLightSleet
	CodeLightSleetShowers
	CodeLightSnow
	CodeLightSnowShowers
	CodePartlyCloudy
	CodeSunny
	CodeThunderyHeavyRain
	CodeThunderyShowers
	CodeThunderySnowShowers
	CodeVeryCloudy
)

type Cond struct {
	Time                time.Time
	Code                WeatherCode
	Desc                string
	TempC               *float32
	FeelsLikeC          *float32
	ChanceOfRainPercent *int
	PrecipM             *float32
	VisibleDistM        *float32
	WindspeedKmph       *float32
	WindGustKmph        *float32
	WinddirDegree       *int
	Humidity            *int
}

type Astro struct {
	Moonrise time.Time
	Moonset  time.Time
	Sunrise  time.Time
	Sunset   time.Time
}

type Day struct {
	Date      time.Time
	Slots     []Cond
	Astronomy Astro
}

type LatLon struct {
	Latitude  float32
	Longitude float32
}

type Data struct {
	Current  Cond
	Forecast []Day
	Location string
	GeoLoc   *LatLon
}

type UnitSystem int

const (
	UnitsMetric UnitSystem = iota
	UnitsImperial
	UnitsSi
	UnitsMetricMs
)

func (u UnitSystem) Temp(tempC float32) (res float32, unit string) {
	if u == UnitsMetric || u == UnitsMetricMs {
		return tempC, "°C"
	} else if u == UnitsImperial {
		return tempC*1.8 + 32, "°F"
	} else if u == UnitsSi {
		return tempC + 273.16, "°K"
	}
	log.Fatalln("Unknown unit system:", u)
	return
}

func (u UnitSystem) Speed(spdKmph float32) (res float32, unit string) {
	if u == UnitsMetric {
		return spdKmph, "km/h"
	} else if u == UnitsImperial {
		return spdKmph / 1.609, "mph"
	} else if u == UnitsSi || u == UnitsMetricMs {
		return spdKmph / 3.6, "m/s"
	}
	log.Fatalln("Unknown unit system:", u)
	return
}

func (u UnitSystem) Distance(distM float32) (res float32, unit string) {
	if u == UnitsMetric || u == UnitsSi || u == UnitsMetricMs {
		if distM < 1 {
			return distM * 1000, "mm"
		} else if distM < 1000 {
			return distM, "m"
		} else {
			return distM / 1000, "km"
		}
	} else if u == UnitsImperial {
		res, unit = distM/0.0254, "in"
		if res < 3*12 { // 1yd = 3ft, 1ft = 12in
			return
		} else if res < 8*10*22*36 { //1mi = 8fur, 1fur = 10ch, 1ch = 22yd
			return res / 36, "yd"
		} else {
			return res / 8 / 10 / 22 / 36, "mi"
		}
	}
	log.Fatalln("Unknown unit system:", u)
	return
}

type Backend interface {
	Setup()
	Fetch(location string, numDays int) Data
}

type Frontend interface {
	Setup()
	Render(weather Data, unitSystem UnitSystem)
}

var (
	AllBackends  = make(map[string]Backend)
	AllFrontends = make(map[string]Frontend)
)
