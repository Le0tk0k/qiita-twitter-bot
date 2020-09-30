package main

import (
	"fmt"
	"github.com/Le0tk0k/qiita-twitter-bot/auth"
	"github.com/Le0tk0k/qiita-twitter-bot/qiita"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	creds := auth.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	api := auth.GetTwitterAPI(&creds)

	_, err = api.PostTweet("testtt", nil)
	if err != nil {
		log.Println(err)
	}

	c := qiita.Client{
		Endpoint:  "https://qiita.com/api/v2/items",
		CreatedAt: "2020-09-27",
		Tag:       "go",
	}

	err = c.GetQiitaArticles()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
