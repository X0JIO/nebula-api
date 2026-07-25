package xray

import (
	"context"
	"net/http"
)

type PingResponse struct {
	Success bool `json:"success"`
}

func (c *HTTPClient) Ping(
	ctx context.Context,
) error {

	var resp PingResponse

	err := c.do(
		ctx,
		http.MethodGet,
		"/health",
		nil,
		&resp,
	)

	if err != nil {
		return err
	}

	if !resp.Success {
		return ErrUnavailable
	}

	return nil
}
