package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	City    = "Malvern"
	State   = "PA"
	Country = "US"
)

type Weather struct {
	Temperature    Fahrenheit
	FeelsLike      Fahrenheit
	TemperatureMin Fahrenheit
	TemperatureMax Fahrenheit
	Pressure       int
	Humidity       int
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

	w.Temperature = WeatherData.Main.Temperature.ConvtoF()
	w.FeelsLike = WeatherData.Main.FeelsLike.ConvtoF()
	w.TemperatureMax = WeatherData.Main.TemperatureMax.ConvtoF()
	w.TemperatureMin = WeatherData.Main.TemperatureMin.ConvtoF()
	w.Humidity = WeatherData.Main.Humidity
	w.Pressure = WeatherData.Main.Pressure

	return nil
}
