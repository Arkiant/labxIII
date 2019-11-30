package transaction

import (
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction/http"
)

type BookCriteria struct {
	CancelPolicy BookCancelPolicy `json:"cancelPolicy"`
	Price        BookPrice        `json:"price"`
	Status       string           `json:"status"`
	Reference    BookReference    `json:"reference"`
	Holder       BookHolder       `json:"holder"`
}

type BookPrice struct {
	Currency string `json:"currency"`
	Gross    int    `json:"gross"`
}

type BookReference struct {
	Client string `json:"client"`
}

type BookHolder struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type BookErrors struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type BookCancelPolicy struct {
	Refundable      bool                `json:"refundable"`
	CancelPenalties BookCancelPenalties `json:"cancelPenalties"`
}

type BookCancelPenalties struct {
	HoursBefore int     `json:"hoursBefore"`
	PenaltyType string  `json:"penaltyType"`
	Currency    string  `json:"currency"`
	Value       float32 `json:"value"`
}

type Service interface {
	Search(pkg.Criteria) pkg.SearchResponse
	DestinationSearcher(pkg.DestinationSearcherCriteria) pkg.DestinationSearcherResponse
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
	client http.Service
}
