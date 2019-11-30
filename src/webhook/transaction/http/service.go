package http

import (
	"context"
	"io"
)

type Service interface {
	Post(ctx context.Context, body io.Reader) (io.ReadCloser, error)
}
