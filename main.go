package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Le0tk0k/qiita-twitter-bot/oauth"
	"github.com/Le0tk0k/qiita-twitter-bot/qiita"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	creds := oauth.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	client := oauth.GetClient(&creds)

	_, resp, err := client.Statuses.Update("teeeessssstttttt", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)

	qiita.GetQiitaArticles()
}
