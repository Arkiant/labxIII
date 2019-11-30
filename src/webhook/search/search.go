package search

import (
	"context"
	"encoding/json"
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction"
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
	s.rq, err = s.getRequest(bodyRQ)
	if err != nil {
		return pkg.SearchResponse{Response: pkg.Response{Errors: []error{err}}}
	}
	code, err := s.transactioner.DestinationSearcher(s.rq.Destination)

	searchRS, err := s.transactioner.Search(
		transaction.Criteria{
			ChecOut:     s.rq.ChecOut,
			Checkin:     s.rq.Checkin,
			Destination: code,
			NumPaxes:    s.rq.NumPaxes,
		},
	)

	return searchRS
}

func (s *SearchService) getRequest(bodyRQ io.Reader) (pkg.Criteria, error) {
	ret := pkg.Criteria{}
	err := json.NewDecoder(bodyRQ).Decode(ret)
	return ret, err
}
