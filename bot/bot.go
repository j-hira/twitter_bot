package bot

import (
	"log"
	"time"
)

var currentTime time.Time

// Run : botを起動する関数
func Run() {
	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	log.Println("ボットくん始動")
	for {
		select {
		case <-ticker.C:
			currentTime = initTime()
			exeFuncAtSpecifiedTime(timeSignal, []int{}, []int{0, 30})
			exeFuncAtSpecifiedTime(weatherForecastNagoya, []int{6}, []int{0})
		}
	}
}

func weatherForecastNagoya() {
	weatherForecast("Nagoya")
}
