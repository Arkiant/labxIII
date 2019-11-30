package routes

import (
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

			//fmt.Println(criteria)

		}
	}
}
