package transaction

import ()
import "github.com/travelgateX/go-io/log"
import "encoding/json"
import "compress/gzip"
import "io/ioutil"
import "context"
import "bytes"
import "fmt"

func (s *ServiceClient) Search(c SearchCriteria) (SearchResponse, error) {

	//ValidateRequest
	ret := SearchResponse{}
	paxes := `{"paxes":[`
	for index := 0; index < c.NumPaxes; index++ {
		paxes += `{"age":30}`
		if index != c.NumPaxes-1 {
			paxes += ","
		}
	}
	paxes += "]}"
	//buildRequest
	requestString := fmt.Sprintf(`{"query":"query ($criteriaSearch: HotelCriteriaSearchInput, $settings: HotelSettingsInput, $filter: FilterInput) {\n  hotelX {\n    search(criteria: $criteriaSearch, settings: $settings, filter: $filter) {\n      errors {\n        code\n        type\n        description\n      }\n      warnings {\n        code\n        type\n        description\n      }\n      options {\n        id\n        hotelName\n        price {\n          currency\n          gross\n        }\n      }\n    }\n  }\n}\n ","variables":{"criteriaSearch":{"checkIn":"%s","checkOut":"%s","occupancies":[%s],"language":"en","nationality":"GB","currency":"EUR","market":"ES","destinations":["%s"]},"settings":{"context":"HOTELTEST","client":"labx","auditTransactions":true,"testMode":true,"plugins":[{"step":"RESPONSE","pluginsType":[{"type":"AGGREGATION","name":"cheapest_price","parameters":[{"key":"primaryKey","value":"hotel,supplier,board"}]}]},{"step":"REQUEST","pluginsType":[{"type":"POST_STEP","name":"search_by_destination","parameters":[{"key":"accessID","value":"0"}]}]}]},"filter":null}}`, c.CheckIn.FormatToDashYYYYMMDD(), c.ChecOut.FormatToDashYYYYMMDD(), paxes, c.Destination)
	r := bytes.NewReader([]byte(requestString))

	//Do call
	resp, err := s.client.Post(context.TODO(), r)

	defer func() {
		if resp != nil {
			resp.Close()
		}
	}()
	if err != nil {
		return ret, err
	}

	rs := HotelxResponse{}

	//parseResponse

	re, err := gzip.NewReader(resp)

	if err != nil {
		return ret, err
	}

	bodyResponse, err := ioutil.ReadAll(re)

	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(bodyResponse, &rs)

	if err != nil {
		return ret, err
	}

	if len(rs.Data.HotelX.Search.Errors) > 0 {
		if len(rs.Data.HotelX.Search.Warnings) > 0 {
			log.Error(rs.Data.HotelX.Search.Warnings[0].Code + "   " + rs.Data.HotelX.Search.Warnings[0].Description)
		}
		log.Error(rs.Data.HotelX.Search.Errors[0].Code + "   " + rs.Data.HotelX.Search.Errors[0].Description)
		return ret, fmt.Errorf(rs.Data.HotelX.Search.Errors[0].Code + "   " + rs.Data.HotelX.Search.Errors[0].Description)
	}

	ret.Amount = rs.Data.HotelX.Search.Options[0].Price.Gross
	ret.Currency = rs.Data.HotelX.Search.Options[0].Price.Currency

	ret.HotelName = rs.Data.HotelX.Search.Options[0].HotelName
	ret.OptionID = rs.Data.HotelX.Search.Options[0].ID
	ret.Refundable = true // quitar pinchaco
	return ret, nil

	return SearchResponse{Amount: 100, Currency: "EUR", HotelName: "QueDiseResorts", OptionID: "uno", Refundable: true}, nil
}
