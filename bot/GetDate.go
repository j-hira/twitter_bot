package bot

import (
	"time"
)

func GetDate() string {
	t := time.Now()
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	t = t.In(loc)
	date := t.Format("2006年01月02日 15時04分05秒")
	return date
}
