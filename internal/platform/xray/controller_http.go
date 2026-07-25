package xray

import (
	"context"
	"net/http"
)

func (c *HTTPClient) Restart(ctx context.Context) error {
	var resp struct {
		Success bool `json:"success"`
	}

	if err := c.do(ctx, http.MethodPost, "/restart", nil, &resp); err != nil {
		return err
	}

	if !resp.Success {
		return ErrRequestFailed
	}

	return nil
}

func (c *HTTPClient) Reload(ctx context.Context) error {
	var resp struct {
		Success bool `json:"success"`
	}

	if err := c.do(ctx, http.MethodPost, "/reload", nil, &resp); err != nil {
		return err
	}

	if !resp.Success {
		return ErrRequestFailed
	}

	return nil
}

func (c *HTTPClient) Version(ctx context.Context) (string, error) {
	var resp struct {
		Version string `json:"version"`
	}

	if err := c.do(ctx, http.MethodGet, "/version", nil, &resp); err != nil {
		return "", err
	}

	return resp.Version, nil
}

func (c *HTTPClient) Start(ctx context.Context) error {
	// Пока Xray запускается systemd/docker.
	// Реализация появится позже.
	return nil
}

func (c *HTTPClient) Stop(ctx context.Context) error {
	// Пока Xray запускается systemd/docker.
	return nil
}

func (c *HTTPClient) Validate(ctx context.Context) error {
	var resp struct {
		Valid bool `json:"valid"`
	}

	if err := c.do(ctx, http.MethodGet, "/validate", nil, &resp); err != nil {
		return err
	}

	if !resp.Valid {
		return ErrRequestFailed
	}

	return nil
}
