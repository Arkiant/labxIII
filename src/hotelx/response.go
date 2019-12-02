package hotelx

type Search struct {
	OptionID   string
	HotelName  string
	Amount     float32
	Currency   string
	Refundable bool
}

type Book struct {
	BookingID string
	Status    string
}

type Destination struct {
	Code      string
	Available bool
}
