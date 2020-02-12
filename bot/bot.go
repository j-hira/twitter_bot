package bot

import (
	"log"
	"time"

	"github.com/j-hira/twitter_bot/twitter"

	_ "github.com/joho/godotenv/autoload"
)

func Run() {
	api := twitter.GetApi()

	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	log.Println("ボットくん始動")
	for {
		select {
		case <-ticker.C:
			t := initTime()
			timeSignal(t, api)
		}
	}
}
