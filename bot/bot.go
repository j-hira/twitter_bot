package bot

import (
	"log"
	"time"

	"github.com/j-hira/twitter_bot/twitter"
)

// Run : botを起動する関数
func Run() {
	api := twitter.GetApi()

	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	log.Println("ボットくん始動")
	for {
		select {
		case <-ticker.C:
			t := initTime()
			// 時間による処理分けを別に切り出す
			timeSignal(t, api)
			weatherForecast(t, api, "Nagoya")
		}
	}
}
