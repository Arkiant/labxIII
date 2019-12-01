package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Arkiant/labxIII/src/hotelx"
	"io/ioutil"
	"net/http"

	"github.com/Arkiant/labxIII/src/conversation"
)

type Server struct {
	df conversation.Dialog
}

func NewServer(df conversation.Dialog) *Server {
	return &Server{df: df}
}

func (s *Server) RequestHandler(w http.ResponseWriter, r *http.Request) {

	criteria, err := s.df.Convert(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := s.df.Send(criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	d, err := ioutil.ReadAll(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var rs hotelx.Search
	err = json.Unmarshal(d, &rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	amountEuros := fmt.Sprintf("%v", rs.Amount)
	response, err := s.df.Speak(criteria.Destination, rs.HotelName, amountEuros, rs.OptionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return

	fmt.Println(response)

}
