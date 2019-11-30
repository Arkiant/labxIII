package transaction

type Service interface {
	Search(Criteria) SearchResponse
}

type Criteria struct{}
type SearchResponse struct{}

var _ Service = (*ServiceClient)(nil)

type ServiceClient struct {
}
