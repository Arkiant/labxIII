package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ClientService interface {
	Do(request io.Reader) (io.ReadCloser, error)
}

type connector struct {
	httpClient http.Client
	url        string
}

type DestinationSearcherRequest struct {
	Text    string
	Access  string
	MaxSize int
}

type DestinationSearcherResponse struct {
	Code      string
	Available bool
}

func main() {
	client := http.Client{}
	c := &connector{httpClient: client, url: "https://api.travelgatex.com"}

	request := DestinationSearcherRequest{Text: "Barcelona", Access: "0", MaxSize: 5000}

	destinationSearcherResponse, err := c.DestinationSearcher(request)

	if err != nil {
		panic(err)
	}

	fmt.Printf(destinationSearcherResponse.Code)
}

func (c *connector) Do(ctx context.Context, body io.Reader) (io.ReadCloser, error) {

	req, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Authorization", "Apikey 4a0bed9a-0ad0-4319-4c1f-0d681b829c14")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error servidor %s", res.Status)
	}
	return res.Body, nil
}

func (c *connector) DestinationSearcher(request DestinationSearcherRequest) (DestinationSearcherResponse, error) {

	requestString := fmt.Sprintf(`{"query":"query ($criteria: HotelXDestinationSearcherInput!) {\n  hotelX {\n    destinationSearcher(criteria: $criteria) {\n      ... on DestinationData {\n        \n        code\n        available\n      }\n    }\n  }\n}\n\n","variables":
					{"criteria": {"text":"%s","access":"%s","maxSize":%d}}}`, request.Text, request.Access, request.MaxSize)

	r := bytes.NewReader([]byte(requestString))

	rs := DestinationSearcherResponse{}

	resp, err := c.Do(context.TODO(), r)

	defer func() {
		if resp != nil {
			resp.Close()
		}
	}()

	if err != nil {
		return rs, err
	}

	re, err := gzip.NewReader(resp)
	bodyResponse, err := ioutil.ReadAll(re)

	if err != nil {
		return rs, fmt.Errorf("failed to open gzip reader: %v", err)
	}

	if err != nil {
		return rs, err
	}

	println(string(bodyResponse))

	err = json.Unmarshal(bodyResponse, &rs)

	if err != nil {
		return rs, err
	}
	return rs, nil
}
