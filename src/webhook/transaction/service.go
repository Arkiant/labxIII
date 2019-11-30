package transaction

import ()
import "github.com/Arkiant/labxIII/src/webhook/transaction/http"
import "github.com/Arkiant/labxIII/src/webhook/pkg"

type Criteria struct {
	Checkin     pkg.Date
	ChecOut     pkg.Date
	Destination string
	NumPaxes    int
}

type DestinationSearcherCriteria struct {
	Text    string
	Access  string
	MaxSize int
}

type SearchResponse struct {
}

type Service interface {
	Search(Criteria) (SearchResponse, error)
	DestinationSearcher(DestinationSearcherCriteria) (string, error)
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
	client http.Service
}
