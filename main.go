package main

import (
	"fmt"
	"github.com/Le0tk0k/qiita-twitter-bot/auth"
	"github.com/Le0tk0k/qiita-twitter-bot/qiita"
	"github.com/joho/godotenv"
	"os"
)

var tag = "go"

// .envの読み込み
func readEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := readEnv(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	c := qiita.Client{
		Endpoint:  "https://qiita.com/api/v2/items",
		CreatedAt: "2020-09-20",
		Tag:       tag,
	}

	api := auth.GetTwitterAPI()
	articles, err := c.GetQiitaArticles()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, i := range *articles {
		_, err = api.PostTweet(i.Title, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
