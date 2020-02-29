package bot

import (
	"time"

	"github.com/j-hira/twitter_bot/twitter"
)

func timeSignal(t time.Time) {
	if min := t.Minute(); min != 0 && min != 30 {
		return
	}
	twitter.TweetText("現在時刻: "+getDate(t))
}

func getDate(t time.Time) string {
	date := t.Format("2006年01月02日 15時04分")
	return date
}
