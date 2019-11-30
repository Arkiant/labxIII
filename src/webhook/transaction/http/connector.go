package http

import (
	"context"
	"io"
)

type connector struct{}

var _ Service = (*connector)(nil)

func (c *connector) Post(ctx context.Context, body io.Reader) (io.ReadCloser, error) {
	return nil, nil
}
