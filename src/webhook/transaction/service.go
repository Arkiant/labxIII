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

type AutoGenerated struct {
	Data Data `json:"data"`
}
type DestinationSearcher struct {
	Code      string `json:"code"`
	Available bool   `json:"available"`
}
type HotelX struct {
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

type BookCriteria struct {
	Input    Input    `json:"input"`
	Settings Settings `json:"settings"`
}
type DeltaPrice struct {
	Amount    int  `json:"amount"`
	Percent   int  `json:"percent"`
	ApplyBoth bool `json:"applyBoth"`
}
type Holder struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type Paxes struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}
type Rooms struct {
	OccupancyRefID int     `json:"occupancyRefId"`
	Paxes          []Paxes `json:"paxes"`
}
type Input struct {
	OptionRefID     string     `json:"optionRefId"`
	ClientReference string     `json:"clientReference"`
	DeltaPrice      DeltaPrice `json:"deltaPrice"`
	Holder          Holder     `json:"holder"`
	Rooms           []Rooms    `json:"rooms"`
}
type Parameters struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type PluginsType struct {
	Type       string       `json:"type"`
	Name       string       `json:"name"`
	Parameters []Parameters `json:"parameters"`
}
type Plugins struct {
	Step        string        `json:"step"`
	PluginsType []PluginsType `json:"pluginsType"`
}
type Settings struct {
	Context           string  `json:"context"`
	Client            string  `json:"client"`
	AuditTransactions bool    `json:"auditTransactions"`
	TestMode          bool    `json:"testMode"`
	Plugins           Plugins `json:"plugins"`
}

type BookResponse struct {
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
