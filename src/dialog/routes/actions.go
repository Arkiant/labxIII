package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Jeffail/gabs"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error leyendo la solicitud",
			http.StatusInternalServerError)
	}

	jsonParsed, err := gabs.ParseJSON(body)

	if err != nil {
		http.Error(w, "Estás enviando un JSON?",
			http.StatusInternalServerError)
	}

	var value string
	var ok bool

	value, ok = jsonParsed.Path("queryResult.intent.displayName").Data().(string)

	if err != nil {
		panic(err)
	}

	if ok == true {
		if value == "search" {

			values, err := jsonParsed.S("queryResult", "parameters").ChildrenMap()

			if err != nil {
				panic(err)
			}

			var criteria pkg.Criteria

			checkIn := values["checkIn"].Data().(string)
			idx := strings.Index(checkIn, "T")

			criteria.Checkin, err = pkg.DateFromDashYYYYMMDDFormat(checkIn[0:idx])

			if err != nil {
				panic(err)
			}

			checkOut := values["checkOut"].Data().(string)
			idx = strings.Index(checkOut, "T")

			criteria.ChecOut, err = pkg.DateFromDashYYYYMMDDFormat(checkOut[0:idx])

			if err != nil {
				panic(err)
			}

			criteria.Destination = values["destination"].Data().(string)
			criteria.NumPaxes = int(values["pax"].Data().(float64))

			d, err := json.Marshal(criteria)
			println(string(d))
			if err != nil {
				panic(err)
			}

			r, err := http.Post("http://labx.travelgatex.com:80/search", "application/json", bytes.NewBuffer(d))

			if err != nil {
				panic(err)
			}
			d, err = ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			println(string(d))
			var rs pkg.SearchResponse
			err = json.Unmarshal(d, &rs)
			if err != nil {
				panic(err)
			}
			var response SearchResponse

			var outputContext OutputContext

			amountEuros := fmt.Sprintf("%v", rs.Amount)
			response.FulfillmentText = "Para " + criteria.Destination + " tenemos un " + rs.HotelName + " a " + amountEuros + " euros"
			outputContext.Name = "projects/hotelx-pjaswu/agent/sessions/55e8f133-bac9-542a-48e3-5574d9b30093/contexts/book"
			outputContext.LifespanCount = 5
			outputContext.Parameters.HotelName = rs.HotelName
			outputContext.Parameters.Price = amountEuros
			outputContext.Parameters.OptionID = rs.OptionID

			response.OutputContexts = append(response.OutputContexts, outputContext)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)

			fmt.Println(response)

		}
	}
}
