package book

import (
	"context"
	"encoding/json"
	"io"

	"github.com/Arkiant/labxIII/src/hotelx"
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction"
	"github.com/marstr/guid"
	"github.com/travelgateX/go-io/log"
)

type BookFactory struct {
	Transactioner transaction.Service
}

func (sf *BookFactory) NewRunner() pkg.Runner {
	return &BookService{
		transactioner: sf.Transactioner,
	}
}

type BookService struct {
	transactioner transaction.Service
	rq            pkg.BookRQ
}

var _ pkg.Runner = (*BookService)(nil)

func (s *BookService) Run(ctx context.Context, bodyRQ io.Reader) interface{} {
	var err error
	log.Debug("GetRequest")
	s.rq, err = s.getRequest(bodyRQ)
	if err != nil {
		log.Error(err.Error())
		return pkg.BookResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}
	bookrefID := guid.NewGUID().String()
	BookRS, err := s.transactioner.Book(
		transaction.BookCriteria{
			OptionRefID:     s.rq.OptionID,
			ClientReference: bookrefID,
		},
	)
	if err != nil {
		log.Error(err.Error())
		return pkg.BookResponse{Response: pkg.Response{Errors: []pkg.Error{pkg.Error{Code: "101", Description: err.Error()}}}}
	}

	ret := pkg.BookResponse{
		Book: hotelx.Book{
			BookingID: bookrefID,
			Status:    BookRS.Status,
		},
	}

	return ret
}

func (s *BookService) getRequest(bodyRQ io.Reader) (pkg.BookRQ, error) {
	ret := pkg.BookRQ{}
	err := json.NewDecoder(bodyRQ).Decode(&ret)
	return ret, err
}
