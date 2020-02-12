package bot

import (
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func timeSignal(t time.Time, api *anaconda.TwitterApi) {
	if min := t.Minute(); min == 0 || min == 30 || true {
		tweet, err := api.PostTweet("現在時刻: "+getDate(t), nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(tweet.Text)
	}
}

func getDate(t time.Time) string {
	date := t.Format("2006年01月02日 15時04分")
	return date
}
