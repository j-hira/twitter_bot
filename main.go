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
	tweet, err := api.PostTweet(text+getDate(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tweet.Text)
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
