package pkg

type Criteria struct {
	Checkin     Date
	ChecOut     Date
	Destination string
	NumPaxes    int
}

type DestinationSearcherCriteria struct {
	Text    string
	Access  string
	MaxSize int
}

type Response struct {
	Errors []error
}

type DestinationSearcherResponse struct {
	Code      string
	Available bool
}

type SearchResponse struct {
	Response
	OptionID   string
	HotelName  string
	Amount     float32
	Currency   string
	Refundable bool
}

type BookResponse struct {
	Response
	BookingID string
}

type Service interface {
	Search(input Criteria) SearchResponse
	Book(optionID string) BookResponse
}
