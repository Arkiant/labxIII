package pkg

type Criteria struct {
	Checkin     Date
	ChecOut     Date
	Destination string
	NumPaxes    int
}

type Response struct {
	Errors []error
}

type SearchRespone struct {
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
	Search(input Criteria) SearchRespone
	Book(optionID string) BookResponse
}
