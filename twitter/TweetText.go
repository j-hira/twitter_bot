package twitter

import (
	"log"
)

func TweetText(text string) {
	tweet, err := GetApi().PostTweet(text, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(tweet.Text)
}
