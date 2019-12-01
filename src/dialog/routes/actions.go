package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Arkiant/labxIII/src/dialog/dialogflow"

	"github.com/Arkiant/labxIII/src/dialog/internal"

	"github.com/Arkiant/labxIII/src/webhook/pkg"
)

type Server struct {
	df internal.Dialog
}

func NewServer(df internal.Dialog) *Server {
	return &Server{df: df}
}

func (s *Server) RequestHandler(w http.ResponseWriter, r *http.Request) {

	criteria, err := s.df.Convert(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	d, err := json.Marshal(criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := http.Post("http://labx.travelgatex.com:80/search", "application/json", bytes.NewBuffer(d))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: process response
	d, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(d))
	var rs pkg.SearchResponse
	err = json.Unmarshal(d, &rs)
	if err != nil {
		panic(err)
	}
	var response dialogflow.SearchResponse

	var outputContext dialogflow.OutputContext

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
