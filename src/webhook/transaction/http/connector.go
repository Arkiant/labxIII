package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type connector struct {
	httpClient http.Client
	url        string
	auth       string
}

var _ Service = (*connector)(nil)

func NewService(client http.Client, url, auth string) Service {
	return &connector{httpClient: client, url: url, auth: auth}
}

func (c *connector) Post(ctx context.Context, body io.Reader) (io.ReadCloser, error) {

	req, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Authorization", c.auth)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error servidor %s", res.Status)
	}
	return res.Body, nil
}
