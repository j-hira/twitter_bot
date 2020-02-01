package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ChimeraCoder/anaconda"
)

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("API_SECRET_KEY"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func main() {
	api := getTwitterApi()
	text := "初めての投稿"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tweet.Text)
}
