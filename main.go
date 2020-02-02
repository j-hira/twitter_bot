package main

import (
	"log"

	"github.com/j-hira/twitter_bot/twitter"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api := twitter.GetApi()
	text := "初めての投稿"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tweet.Text)
}
