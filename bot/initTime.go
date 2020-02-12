package bot

import (
	"time"
)

func initTime() time.Time {
	t := time.Now()
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	t = t.In(loc)
	return t
}
