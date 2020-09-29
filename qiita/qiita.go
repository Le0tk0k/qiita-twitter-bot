package qiita

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Article struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

func GetQiitaArticles() {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=tag:go", nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var article []Article
	json.Unmarshal(body, &article)
	fmt.Printf("%+v\n", article)
}
