package routes

import (
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

			var response SearchResponse
			var outputContext OutputContext

			response.FulfillmentText = "Para " + criteria.Destination + " tenemos un puti a 100 euros"
			outputContext.Name = "projects/hotelx-pjaswu/agent/sessions/55e8f133-bac9-542a-48e3-5574d9b30093/contexts/book"
			outputContext.LifespanCount = 5
			outputContext.Parameters.HotelName = "Hotel Prueba"
			outputContext.Parameters.Price = "100 euros"
			outputContext.Parameters.OptionID = 12345

			response.OutputContexts = append(response.OutputContexts, outputContext)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)

			fmt.Println(response)

		}
	}
}
