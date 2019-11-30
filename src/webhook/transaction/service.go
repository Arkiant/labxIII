package transaction

import (
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction/http"
)

type Service interface {
	Search(pkg.Criteria) pkg.SearchResponse
	DestinationSearcher(pkg.DestinationSearcherCriteria) pkg.DestinationSearcherResponse
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
	client http.Service
}
