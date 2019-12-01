package conversation

import (
	"github.com/Arkiant/labxIII/src/kit/date"
)

// Criteria for search request
type Criteria struct {
	Checkin     date.Date
	ChecOut     date.Date
	Destination string
	NumPaxes    int
}
