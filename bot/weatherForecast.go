package bot

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"

	"github.com/ChimeraCoder/anaconda"
)

// https://mholt.github.io/json-to-go/
type weatherJSON struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}

func weatherForecast(t time.Time, api *anaconda.TwitterApi, city string) {
	if t.Minute() != 0 {
		return
	}
	url := "http://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s"
	resp, err := http.Get(fmt.Sprintf(url, city, os.Getenv("WEATHER_API_KEY")))
	if err != nil {
		log.Println(err)
		return
	}
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	jsonBytes := ([]byte)(byteArray)
	data := new(weatherJSON)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Println(err)
		return
	}
	tweet, err := api.PostTweet(data.List[0].Weather[0].Main, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(tweet.Text)
}
