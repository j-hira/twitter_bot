package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func GetApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("API_SECRET_KEY"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
