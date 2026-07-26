package xray

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client interface {
	Ping(ctx context.Context) error

	AddUser(ctx context.Context, req AddUserRequest) error
	UpdateUser(ctx context.Context, req UpdateUserRequest) error
	DeleteUser(ctx context.Context, uuid string) error
	GetUserStats(ctx context.Context, uuid string) (*UserStats, error)

	ListInbounds(ctx context.Context) ([]Inbound, error)
	FindInboundByProtocol(ctx context.Context, protocol string) (*Inbound, error)

	CreateInbound(ctx context.Context, req CreateInboundRequest) error
	UpdateInbound(ctx context.Context, req UpdateInboundRequest) error
	DeleteInbound(ctx context.Context, id int) error

	Reload(ctx context.Context) error
	Restart(ctx context.Context) error
	Version(ctx context.Context) (string, error)

	Validate(ctx context.Context) error
}

type HTTPClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

type Config struct {
	Enabled bool
	BaseURL string
	APIKey  string

	Timeout time.Duration
}

func NewHTTPClient(cfg *Config) *HTTPClient {

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 15 * time.Second
	}

	return &HTTPClient{
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPClient) do(
	ctx context.Context,
	method string,
	path string,
	request any,
	response any,
) error {

	var body io.Reader

	if request != nil {

		data, err := json.Marshal(request)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		body,
	)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	if request != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {

		raw, _ := io.ReadAll(resp.Body)

		return fmt.Errorf(
			"%w (%d): %s",
			ErrRequestFailed,
			resp.StatusCode,
			string(raw),
		)
	}

	if response == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *HTTPClient) Reload(
	ctx context.Context,
) error {

	var resp struct {
		Success bool `json:"success"`
	}

	return c.do(
		ctx,
		http.MethodPost,
		"/reload",
		nil,
		&resp,
	)
}

func (c *HTTPClient) Restart(
	ctx context.Context,
) error {

	var resp struct {
		Success bool `json:"success"`
	}

	return c.do(
		ctx,
		http.MethodPost,
		"/restart",
		nil,
		&resp,
	)
}

func (c *HTTPClient) Version(
	ctx context.Context,
) (string, error) {

	var resp struct {
		Version string `json:"version"`
	}

	err := c.do(
		ctx,
		http.MethodGet,
		"/version",
		nil,
		&resp,
	)

	if err != nil {
		return "", err
	}

	return resp.Version, nil
}

func (c *HTTPClient) Validate(
	ctx context.Context,
) error {

	var resp struct {
		Valid bool `json:"valid"`
	}

	if err := c.do(
		ctx,
		http.MethodGet,
		"/validate",
		nil,
		&resp,
	); err != nil {
		return err
	}

	if !resp.Valid {
		return ErrRequestFailed
	}

	return nil
}
