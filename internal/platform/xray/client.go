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

	Timeout int
}

func NewHTTPClient(cfg *Config) *HTTPClient {
	return &HTTPClient{
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
		client: &http.Client{
			Timeout: 15 * time.Second,
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
			"%w: %s",
			ErrRequestFailed,
			string(raw),
		)
	}

	if response == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
