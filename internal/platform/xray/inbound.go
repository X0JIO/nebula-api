package xray

import (
	"context"
	"fmt"
	"net/http"
)

type Inbound struct {
	ID       int    `json:"id"`
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type CreateInboundRequest struct {
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type UpdateInboundRequest struct {
	ID       int    `json:"id"`
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type CreateInboundResponse struct {
	Success bool `json:"success"`
}

type UpdateInboundResponse struct {
	Success bool `json:"success"`
}

type DeleteInboundResponse struct {
	Success bool `json:"success"`
}

func (c *HTTPClient) ListInbounds(
	ctx context.Context,
) ([]Inbound, error) {

	var resp []Inbound

	err := c.do(
		ctx,
		http.MethodGet,
		"/inbounds",
		nil,
		&resp,
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *HTTPClient) FindInboundByProtocol(
	ctx context.Context,
	protocol string,
) (*Inbound, error) {

	inbounds, err := c.ListInbounds(ctx)
	if err != nil {
		return nil, err
	}

	for _, inbound := range inbounds {

		if inbound.Protocol == protocol {
			return &inbound, nil
		}
	}

	return nil, fmt.Errorf("inbound for protocol %s not found", protocol)
}

func (c *HTTPClient) CreateInbound(
	ctx context.Context,
	req CreateInboundRequest,
) error {

	var resp CreateInboundResponse

	err := c.do(
		ctx,
		http.MethodPost,
		"/inbounds",
		req,
		&resp,
	)

	if err != nil {
		return err
	}

	if !resp.Success {
		return ErrRequestFailed
	}

	return nil
}

func (c *HTTPClient) UpdateInbound(
	ctx context.Context,
	req UpdateInboundRequest,
) error {

	var resp UpdateInboundResponse

	err := c.do(
		ctx,
		http.MethodPut,
		"/inbounds",
		req,
		&resp,
	)

	if err != nil {
		return err
	}

	if !resp.Success {
		return ErrRequestFailed
	}

	return nil
}

func (c *HTTPClient) DeleteInbound(
	ctx context.Context,
	id int,
) error {

	var resp DeleteInboundResponse

	err := c.do(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/inbounds/%d", id),
		nil,
		&resp,
	)

	if err != nil {
		return err
	}

	if !resp.Success {
		return ErrRequestFailed
	}

	return nil
}
