package bot

import (
	"time"

	"github.com/j-hira/twitter_bot/twitter"
)

func timeSignal() {
	twitter.TweetText("現在時刻: "+getDate(currentTime))
}

func getDate(t time.Time) string {
	date := t.Format("2006年01月02日 15時04分")
	return date
}
