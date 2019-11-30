package transaction

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (s *ServiceClient) Quote(c QuoteRequest) (string, error) {

	//buildRequest
	requestString := fmt.Sprintf(`{
		"query": "query ($criteriaQuote: HotelCriteriaQuoteInput!, $settings: HotelSettingsInput) {\n  hotelX {\n    quote(criteria: $criteriaQuote, settings: $settings) {\n      optionQuote {\n        optionRefId\n        status\n      }\n      errors {\n        code\n        type\n        description\n      }\n      warnings {\n        code\n        type\n        description\n      }\n    }\n  }\n}\n",
		"variables": {
			"criteriaQuote": {
				"optionRefId": "%s"
			},
			"settings": {
				"context": "HOTELTEST",
				"client": "labx",
				"auditTransactions": true,
				"testMode": true,
				"plugins": []
			}
		}
	}`, c.OptionRefId)

	r := bytes.NewReader([]byte(requestString))

	println(requestString)
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

	return rs.Data.HotelX.Quote.OptionQuote.OptionRefID, nil

	//errors control
}
