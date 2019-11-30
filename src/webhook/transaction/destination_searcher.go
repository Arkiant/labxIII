package transaction

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Arkiant/labxIII/src/webhook/pkg"
)

func (s *ServiceClient) DestinationSearcher(c pkg.DestinationSearcherCriteria) (pkg.DestinationSearcherResponse, error) {

	//ValidateRequest
	if c.Text == "" || c.Access == "" || c.MaxSize == 0 {
		return pkg.DestinationSearcherResponse{}, fmt.Errorf("Error params")
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

	rs := pkg.DestinationSearcherResponse{}

	if err != nil {
		return rs, fmt.Errorf("Error")
	}

	//parseResponse

	re, err := gzip.NewReader(resp)

	if err != nil {
		return rs, fmt.Errorf("Error reader")
	}

	bodyResponse, err := ioutil.ReadAll(re)

	if err != nil {
		return rs, fmt.Errorf("Error ReadAll")
	}

	println(string(bodyResponse))

	err = json.Unmarshal(bodyResponse, &rs)

	if err != nil {
		return rs, fmt.Errorf("Error Unmarshal")
	}
	return rs, nil

	//errors control
}
