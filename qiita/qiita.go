package qiita

import (
	"encoding/json"
	"errors"
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

func (c *Client) GetQiitaArticles() error {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return err
	}

	q := u.Query()
	q.Set("page", "1")
	q.Set("per_page", "1")
	q.Set("query", "tag:"+c.Tag+" created:>"+c.CreatedAt)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var articles []Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return err
	}
	fmt.Printf("%+v\n", articles)

	return nil
}
