package transaction

import ()
import "github.com/Arkiant/labxIII/src/webhook/transaction/http"
import "github.com/Arkiant/labxIII/src/webhook/pkg"

type SearchCriteria struct {
	CheckIn     pkg.Date `json:"checkIn"`
	ChecOut     pkg.Date `json:"checkOut"`
	Destination string   `json:"destination"`
	NumPaxes    int      `json:"numPaxes"`
}

type DestinationSearcherCriteria struct {
	Text    string `json:"text"`
	Access  string `json:"access"`
	MaxSize int    `json:"maxSize"`
}

type Price struct {
	Currency string  `json:"currency"`
	Gross    float32 `json:"gross"`
}
type Options struct {
	ID        string `json:"id"`
	HotelName string `json:"hotelName"`
	Price     Price  `json:"price"`
}
type Search struct {
	Errors   []Error   `json:"errors"`
	Warnings []Error   `json:"warnings"`
	Options  []Options `json:"options"`
}

type Error struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type HotelxResponse struct {
	Data Data `json:"data"`
}
type DestinationSearcher struct {
	Code      string `json:"code"`
	Available bool   `json:"available"`
}
type HotelX struct {
	Search              Search                `json:"search"`
	DestinationSearcher []DestinationSearcher `json:"destinationSearcher"`
}
type Data struct {
	HotelX HotelX `json:"hotelX"`
}

type SearchResponse struct {
	OptionID   string  `json:"optionID"`
	HotelName  string  `json:"hotelName"`
	Amount     float32 `json:"amount"`
	Currency   string  `json:"currency"`
	Refundable bool    `json:"refundable"`
}

type Service interface {
	Search(SearchCriteria) (SearchResponse, error)
	DestinationSearcher(DestinationSearcherCriteria) (string, error)
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
	client http.Service
}

func NewService(c http.Service) Service {
	return &ServiceClient{
		client: c,
	}
}
