package weather

import "net/http"

const (
	ApiKey  = "3df348241c25ac51065232df76efe127"
	City    = "Malvern"
	State   = "PA"
	Country = "US"
)

type Weather struct{}

func FetchWeather() error {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + City + "," + State + "," + Country + "&appid=" + ApiKey)
	if err != nil {
		return err
	}

	return nil
}
