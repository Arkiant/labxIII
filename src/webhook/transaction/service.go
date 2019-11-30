package transaction

import (
	"github.com/Arkiant/labxIII/src/webhook/pkg"
)

type Service interface {
	Search(pkg.Criteria) pkg.SearchResponse
}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
}
