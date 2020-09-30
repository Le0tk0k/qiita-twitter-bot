package qiita

import (
	"encoding/json"
	"errors"
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

// urlを作成する
func createUrl(u *url.URL, c *Client) string {
	q := u.Query()
	q.Set("page", "1")
	q.Set("per_page", "10")
	q.Set("query", "tag:"+c.Tag+" created:>"+c.CreatedAt)
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) GetQiitaArticles() (*[]Article, error) {
	e, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}

	u := createUrl(e, c)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var articles []Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	return &articles, nil
}
