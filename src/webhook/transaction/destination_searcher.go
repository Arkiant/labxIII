package transaction

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (s *ServiceClient) DestinationSearcher(c DestinationSearcherCriteria) (string, error) {

	//ValidateRequest
	if c.Text == "" || c.Access == "" || c.MaxSize == 0 {
		return "", fmt.Errorf("Error params")
	}

	//buildRequest
	requestString := fmt.Sprintf(`{"query":"query ($criteria: HotelXDestinationSearcherInput!) {\n  hotelX {\n    destinationSearcher(criteria: $criteria) {\n      ... on DestinationData {\n        \n        code\n        available\n      }\n    }\n  }\n}\n\n","variables":
	{"criteria": {"text":"%s","access":"%s","maxSize":%d}}}`, c.Text, c.Access, c.MaxSize)

	r := bytes.NewReader([]byte(requestString))

	//Do call
	resp, err := s.client.Post(context.TODO(), r)

	defer func() {
		if resp != nil {
			resp.Close()
		}
	}()

	rs := HotelxResponse{}

	if err != nil {
		return "", fmt.Errorf("Error")
	}

	//parseResponse

	re, err := gzip.NewReader(resp)

	if err != nil {
		return "", fmt.Errorf("Error reader")
	}

	bodyResponse, err := ioutil.ReadAll(re)

	if err != nil {
		return "", fmt.Errorf("Error ReadAll")
	}

	println(string(bodyResponse))

	err = json.Unmarshal(bodyResponse, &rs)

	if err != nil {
		return "", fmt.Errorf("Error Unmarshal")
	}

	return rs.Data.HotelX.DestinationSearcher[0].Code, nil

	//errors control
}
