package dialogflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Arkiant/labxIII/src/conversation"
	"github.com/Arkiant/labxIII/src/kit/date"

	"github.com/Jeffail/gabs"
)

const URL = "http://labx.travelgatex.com:80/search"

type DialogFlow struct{}

var _ conversation.Dialog = DialogFlow{}

func NewDialog() *DialogFlow {
	return &DialogFlow{}
}

// Convert a read closer into a pkg.Criteria
func (d DialogFlow) Convert(body io.ReadCloser) (*conversation.Criteria, error) {
	container, err := parse(body)
	if err != nil {
		return nil, err
	}
	intent := getIntent(container)
	switch intent {
	case conversation.SEARCH:
		params, err := getSearchParameters(container)
		if err != nil {
			return nil, err
		}
		return params, nil
	case conversation.BOOK:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("not found intent")
	}
}

func (d DialogFlow) Send(criteria *conversation.Criteria) (io.ReadCloser, error) {
	c, err := json.Marshal(criteria)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(URL, "application/json", bytes.NewBuffer(c))
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (d DialogFlow) Speak(destination string, hotelName string, amount string, optionID string) ([]byte, error) {

	outputContext := OutputContext{
		Name:          "projects/hotelx-pjaswu/agent/sessions/55e8f133-bac9-542a-48e3-5574d9b30093/contexts/book",
		LifespanCount: 5,
		Parameters: Parameters{
			HotelName: hotelName,
			Price:     amount,
			OptionID:  optionID,
		},
	}

	response := SearchResponse{
		FulfillmentText: "Para " + destination + " tenemos un " + hotelName + " a " + amount + " euros",
	}

	response.OutputContexts = append(response.OutputContexts, outputContext)

	res, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func parse(body io.ReadCloser) (*gabs.Container, error) {

	v, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	res, err := gabs.ParseJSON(v)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getIntent(container *gabs.Container) conversation.Intent {

	dName, ok := container.Path("queryResult.intent.displayName").Data().(string)

	if ok {
		return conversation.Intent(dName)
	}

	return ""
}

func getSearchParameters(jsonParsed *gabs.Container) (*conversation.Criteria, error) {
	p, err := jsonParsed.S("queryResult", "parameters").ChildrenMap()

	if err != nil {
		return nil, err
	}

	checkIn, err := dateFormat(p["checkIn"].Data().(string))
	if err != nil {
		return nil, err
	}

	checkOut, err := dateFormat(p["checkOut"].Data().(string))
	if err != nil {
		return nil, err
	}

	criteria := conversation.Criteria{
		Checkin:     checkIn,
		CheckOut:    checkOut,
		Destination: p["destination"].Data().(string),
		NumPaxes:    p["pax"].Data().(int),
	}

	return &criteria, nil
}

func dateFormat(dateToFormat string) (date.Date, error) {
	idx := strings.Index(dateToFormat, "T")
	v, err := date.DateFromDashYYYYMMDDFormat(dateToFormat[0:idx])
	if err != nil {
		return date.DateNow(), err
	}
	return v, nil
}
