package search

import (
	"context"
	"encoding/json"
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction"
	"github.com/travelgateX/go-io/log"
	"io"
)

type SearchFactory struct {
	Transactioner transaction.Service
}

func (sf *SearchFactory) NewRunner() pkg.Runner {
	return &SearchService{
		transactioner: sf.Transactioner,
	}
}

type SearchService struct {
	transactioner transaction.Service
	rq            pkg.Criteria
}

var _ pkg.Runner = (*SearchService)(nil)

func (s *SearchService) Run(ctx context.Context, bodyRQ io.Reader) interface{} {
	var err error
	log.Debug("GetRequest")
	s.rq, err = s.getRequest(bodyRQ)
	if err != nil {
		log.Error(err.Error())
		return pkg.SearchResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}

	log.Debug("Send destination")
	code, err := s.transactioner.DestinationSearcher(
		transaction.DestinationSearcherCriteria{
			Text:    s.rq.Destination,
			MaxSize: 500,
			Access:  "0",
		},
	)
	if err != nil {
		log.Error(err.Error())
		return pkg.SearchResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}
	log.Debug("Send search with code: " + code)
	searchRS, err := s.transactioner.Search(
		transaction.SearchCriteria{
			ChecOut:     s.rq.ChecOut,
			CheckIn:     s.rq.Checkin,
			Destination: code,
			NumPaxes:    s.rq.NumPaxes,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return pkg.SearchResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}

	log.Debug("Send quote with id: " + searchRS.OptionID)
	quoteRS, err := s.transactioner.Quote(
		transaction.QuoteRequest{
			OptionRefId: searchRS.OptionID,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return pkg.SearchResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}

	ret := pkg.SearchResponse{
		Amount:     searchRS.Amount,
		Currency:   searchRS.Currency,
		HotelName:  searchRS.HotelName,
		OptionID:   quoteRS,
		Refundable: true,
	}

	return ret
}

func (s *SearchService) getRequest(bodyRQ io.Reader) (pkg.Criteria, error) {
	ret := pkg.Criteria{}
	err := json.NewDecoder(bodyRQ).Decode(&ret)
	return ret, err
}
