package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiKey = "d1e6e5e749e9453cafe231524230310"

func FetchWeather(city string) (string, error) {
	escapedCity := url.QueryEscape(city)
	reqUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", escapedCity, apiKey)
	fmt.Println(reqUrl)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	// TODO: decode the JSON response instead of just returning the raw body string

	weather := string(body)
	return fmt.Sprintf("Weather in %s is %s", city, weather), nil
}
