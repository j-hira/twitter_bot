package main

import (
	"log"
	"time"

	"github.com/j-hira/twitter_bot/twitter"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api := twitter.GetApi()

	text := "現在時刻: "
	ticker := time.NewTicker(time.Second * 60 * 30) // 30分ごとに定期実行
	defer ticker.Stop()
	tweet, err := api.PostTweet("ボットくん始動", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tweet.Text)
	for {
		select {
		case <-ticker.C:
			tweet, err := api.PostTweet(text+getDate(), nil)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(tweet.Text)
		}
	}
}

func getDate() string {
	t := time.Now()
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	t = t.In(loc)
	date := t.Format("2006年01月02日 15時04分05秒")
	return date
}
