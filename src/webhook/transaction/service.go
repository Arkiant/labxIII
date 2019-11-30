package transaction

import ()
import "github.com/Arkiant/labxIII/src/webhook/pkg"

type Criteria struct {
	Checkin     pkg.Date
	ChecOut     pkg.Date
	Destination string
	NumPaxes    int
}
type SearchResponse struct {
}

type Service interface {
	Search(Criteria) (SearchResponse, error)
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
}
