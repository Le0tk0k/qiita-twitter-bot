package auth

import (
	"github.com/ChimeraCoder/anaconda"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetTwitterAPI(creds *Credentials) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(creds.ConsumerKey)
	anaconda.SetConsumerSecret(creds.ConsumerSecret)
	api := anaconda.NewTwitterApi(creds.AccessToken, creds.AccessTokenSecret)

	return api
}
