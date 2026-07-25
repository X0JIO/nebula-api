package xray

import (
	"context"
	"net/http"
)

func (c *HTTPClient) Health(
	ctx context.Context,
) error {

	return c.do(
		ctx,
		http.MethodGet,
		"/health",
		nil,
		nil,
	)
}
