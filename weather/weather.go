package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	City    = "Malvern"
	State   = "PA"
	Country = "US"
)

type Weather struct {
	RWSync         *sync.RWMutex
	Temperature    Fahrenheit
	FeelsLike      Fahrenheit
	TemperatureMin Fahrenheit
	TemperatureMax Fahrenheit
	Pressure       int
	Humidity       int
	LastCall       time.Time
}

type Response struct {
	Weather []Description `json:"weather"`
	Main    Main          `json:"main"`
	Wind    Wind          `json:"wind"`
}

type Description struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temperature    Kelvin `json:"temp"`
	FeelsLike      Kelvin `json:"feels_like"`
	TemperatureMin Kelvin `json:"temp_min"`
	TemperatureMax Kelvin `json:"temp_max"`
	Pressure       int    `json:"pressure"`
	Humidity       int    `json:"humidity"`
}

type Wind struct {
	Speed  float64 `json:"speed"`
	Degree float64 `json:"deg"`
}

// NewWeather creates a new instance of weather. It requests the first fetch and also starts a Go Routine. The is
// concurrent safe.
func NewWeather() (*Weather, error) {
	weather := Weather{}

	if err := weather.FetchWeather(); err != nil {
		return &weather, err
	}

	go func(w *Weather) {
		for {
			log.Println("Fetching Weather...")
			time.Sleep(time.Minute * 1)

			if err := w.FetchWeather(); err != nil {
				return
			}
		}
	}(&weather)

	return &weather, nil
}

// FetchWeather is responsible for making the call to Open Weather and grabbing the most recent weather report.
// OS Env variable WEATHER_API_KEY containing the API Key secret is required for this to work.
func (w *Weather) FetchWeather() error {
	ApiKey := os.Getenv("WEATHER_API_KEY")

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + City + "," + State + "," + Country + "&appid=" + ApiKey) //nolint:lll
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var WeatherData Response

	if err := json.Unmarshal(body, &WeatherData); err != nil {
		return err
	}

	w.RWSync.Lock()
	w.Temperature = WeatherData.Main.Temperature.ConvtoF()
	w.FeelsLike = WeatherData.Main.FeelsLike.ConvtoF()
	w.TemperatureMax = WeatherData.Main.TemperatureMax.ConvtoF()
	w.TemperatureMin = WeatherData.Main.TemperatureMin.ConvtoF()
	w.Humidity = WeatherData.Main.Humidity
	w.Pressure = WeatherData.Main.Pressure

	w.LastCall = time.Now()
	w.RWSync.Unlock()

	return nil
}
