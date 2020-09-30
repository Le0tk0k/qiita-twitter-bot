package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Le0tk0k/qiita-twitter-bot/auth"
	"github.com/Le0tk0k/qiita-twitter-bot/qiita"
	"github.com/aws/aws-lambda-go/lambda"
)

var tag = "go"

func post() {
	c := qiita.Client{
		Endpoint:  "https://qiita.com/api/v2/items",
		CreatedAt: time.Now().Format("2006-01-02"),
		Tag:       tag,
	}

	api := auth.GetTwitterAPI()
	articles, err := c.GetQiitaArticles()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	t := time.Now().Add(time.Duration((-1) * time.Hour))

	for _, i := range *articles {
		if i.CreatedAt.After(t) {
			fmt.Println(i.CreatedAt)
			_, err = api.PostTweet(i.Title+"\n#golang\n"+i.URL, nil)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func main() {
	lambda.Start(post)
}
