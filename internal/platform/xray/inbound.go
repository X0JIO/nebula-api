package xray

import (
	"context"
	"fmt"
	"net/http"
)

type Inbound struct {
	ID       int    `json:"id"`
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Listen   string `json:"listen"`
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

func (c *HTTPClient) ListInbounds(
	ctx context.Context,
) ([]Inbound, error) {

	var resp []Inbound

	if err := c.do(
		ctx,
		http.MethodGet,
		"/inbounds",
		nil,
		&resp,
	); err != nil {
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

	return nil, fmt.Errorf("inbound for protocol %q not found", protocol)
}

func (c *HTTPClient) CreateInbound(
	ctx context.Context,
	req CreateInboundRequest,
) error {

	return c.do(
		ctx,
		http.MethodPost,
		"/inbounds",
		req,
		nil,
	)
}

func (c *HTTPClient) UpdateInbound(
	ctx context.Context,
	req UpdateInboundRequest,
) error {

	return c.do(
		ctx,
		http.MethodPut,
		fmt.Sprintf("/inbounds/%d", req.ID),
		req,
		nil,
	)
}

func (c *HTTPClient) DeleteInbound(
	ctx context.Context,
	id int,
) error {

	return c.do(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/inbounds/%d", id),
		nil,
		nil,
	)
}
