package dialogflow

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/Arkiant/labxIII/src/conversation"
	"github.com/Arkiant/labxIII/src/kit/date"

	"github.com/Arkiant/labxIII/src/dialog/internal"

	"github.com/Jeffail/gabs"
)

type DialogFlow struct{}

var _ internal.Dialog = DialogFlow{}

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
	case internal.SEARCH:
		params, err := getSearchParameters(container)
		if err != nil {
			return nil, err
		}
		return params, nil
	case internal.BOOK:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("not found intent")
	}
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

func getIntent(container *gabs.Container) internal.Intent {

	dName, ok := container.Path("queryResult.intent.displayName").Data().(string)

	if ok {
		return internal.Intent(dName)
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
		ChecOut:     checkOut,
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
