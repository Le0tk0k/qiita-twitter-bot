package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Le0tk0k/qiita-twitter-bot/oauth"
)

type Article struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

func getQiitaArticles() {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=tag:go", nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var article []Article
	json.Unmarshal(body, &article)
	fmt.Printf("%+v\n", article)
}

func main() {
	creds := oauth.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	client := oauth.GetClient(&creds)

	_, resp, err := client.Statuses.Update("test test test", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)

	getQiitaArticles()
}
