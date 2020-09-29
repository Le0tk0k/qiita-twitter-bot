package qiita

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Endpoint  string
	CreatedAt string
	Tag       string
}

type Article struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
}

func (c *Client) GetQiitaArticles() {
	u, _ := url.Parse(c.Endpoint)
	q := u.Query()
	q.Set("page", "1")
	q.Set("per_page", "1")
	q.Set("query", "tag:"+c.Tag+" created:>"+c.CreatedAt)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

	req, _ := http.NewRequest("GET", u.String(), nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var articles []Article
	json.Unmarshal(body, &articles)
	fmt.Printf("%+v\n", articles)
}
