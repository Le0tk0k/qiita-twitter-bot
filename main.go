package main

import (
	"fmt"
	"github.com/Le0tk0k/qiita-twitter-bot/auth"
	"github.com/Le0tk0k/qiita-twitter-bot/qiita"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// .envの読み込み
func readEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

var tag = "go"

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

	_, err := api.PostTweet("tesaattt", nil)
	if err != nil {
		log.Println(err)
	}

	articles, err := c.GetQiitaArticles()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, i := range *articles {
		fmt.Printf("%+v\n", i)
	}
}
