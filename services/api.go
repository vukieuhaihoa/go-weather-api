package services

import (
	"encoding/json"
	"net/http"
)

type OpenWeatherMapData struct {
	CityName string `json:"name"`
	Main     struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func GetWeatherByCityID(city string) (OpenWeatherMapData, error) {
	data := OpenWeatherMapData{}

	API := "9d8d1b67644586cc416058d7cbf323f1"

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + API)

	if err != nil || res.StatusCode != 200 {
		return data, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
