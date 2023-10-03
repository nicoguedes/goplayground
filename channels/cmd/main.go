package main

import (
	"fmt"
	"github.com/nicoguedes/goplayground/channels/cmd/internal/weather"
)

func main() {
	weatherResult, err := weather.FetchWeather("Belo Horizonte")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(weatherResult)
}
