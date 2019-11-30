package transaction

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (s *ServiceClient) Book(c BookCriteria) (Booking, error) {
	rs := Booking{}

	//ValidateRequest
	if c.OptionRefID == "" {
		return rs, fmt.Errorf("Error OptionRefID")
	}

	//buildRequest
	requestString := fmt.Sprintf(`{"query":"mutation ($input: HotelBookInput!, $settings: HotelSettingsInput) {\n  hotelX {\n    book(input: $input, settings: $settings) {\n      booking {\n        status\n        reference {\n          client\n        }\n      }\n      errors {\n        code\n        type\n        description\n      }\n      warnings {\n        code\n        type\n        description\n      }\n    }\n  }\n}\n","variables":{"input":{"optionRefId":"%s","clientReference":"%s","deltaPrice":{"amount":10,"percent":10,"applyBoth":true},"holder":{"name":"Name1","surname":"Surname1"},"rooms":[{"occupancyRefId":1,"paxes":[{"name":"Name1","surname":"Surname1","age":30},{"name":"Name2","surname":"Surname2","age":30}]}]},"settings":{"context":"HOTELTEST","client":"labx","auditTransactions":true,"testMode":true,"plugins":{"step":"RESPONSE_OPTION","pluginsType":[{"type":"COMMISSION","name":"commissionX","parameters":[{"key":"default","value":"50"}]}]}}}}`, c.OptionRefID, c.ClientReference)
	println(requestString)
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

	hxrs := HotelxResponse{}
	err = json.Unmarshal(bodyResponse, &hxrs)

	if err != nil {
		return rs, fmt.Errorf("Error Unmarshal")
	}

	rs.Status = hxrs.Data.HotelX.Book.Booking.Status
	rs.Reference = hxrs.Data.HotelX.Book.Booking.Reference

	return rs, nil

	//errors control
}
