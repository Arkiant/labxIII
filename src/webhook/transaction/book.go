package transaction

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (s *ServiceClient) Book(c BookCriteria) (BookResponse, error) {

	rs := BookResponse{}

	//ValidateRequest
	if c.OptionRefID == "" {
		return rs, fmt.Errorf("Error OptionRefID")
	}

	//buildRequest
	requestString := fmt.Sprintf(`{{"query":"mutation ($input: HotelBookInput!, $settings: HotelSettingsInput) {\n  hotelX {\n    book(input: $input, settings: $settings) {\n      booking {\n        cancelPolicy {\n          refundable\n        }\n        price {\n          currency\n          gross\n        }\n        status\n        reference {\n          client\n        }\n        holder {\n          name\n          surname\n        }\n      }\n      errors {\n        code\n        type\n        description\n      }\n      warnings {\n        code\n        type\n        description\n      }\n    }\n  }\n}\n","variables":
	{"input":{"optionRefId":"%s","clientReference":"%s,"deltaPrice":{"amount":10,"percent":10,"applyBoth":true},"holder":{"name":"Name1","surname":"Surname1"},"rooms":[{"occupancyRefId":1,"paxes":[{"name":"Name1","surname":"Surname1","age":18},{"name":"Name2","surname":"Surname2","age":30}]}]},"settings":{"context":"HOTELTEST","client":"labx","auditTransactions":true,"testMode":true,"plugins":{"step":"RESPONSE_OPTION","pluginsType":[{"type":"COMMISSION","name":"commissionX","parameters":[{"key":"default","value":"50"}]}]}}}}`, c.OptionRefID, c.ClientReference)

	r := bytes.NewReader([]byte(requestString))

	//Do call
	resp, err := s.client.Post(context.TODO(), r)

	defer func() {
		if resp != nil {
			resp.Close()
		}
	}()

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
