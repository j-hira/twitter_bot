package bot

import (
	"log"
	"time"

	"github.com/j-hira/twitter_bot/twitter"

	_ "github.com/joho/godotenv/autoload"
)

func Run() {
	api := twitter.GetApi()

	text := "現在時刻: "
	ticker := time.NewTicker(time.Second * 60 * 30) // 30分ごとに定期実行
	defer ticker.Stop()
	log.Println("ボットくん始動")
	for {
		select {
		case <-ticker.C:
			tweet, err := api.PostTweet(text+GetDate(), nil)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(tweet.Text)
		}
	}
}
