package xray

import (
	"context"
	"net/http"
)

func (c *HTTPClient) AddUser(
	ctx context.Context,
	req AddUserRequest,
) error {

	var resp AddUserResponse

	err := c.do(
		ctx,
		http.MethodPost,
		"/users",
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

func (c *HTTPClient) UpdateUser(
	ctx context.Context,
	req UpdateUserRequest,
) error {

	var resp UpdateUserResponse

	err := c.do(
		ctx,
		http.MethodPut,
		"/users",
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

func (c *HTTPClient) DeleteUser(
	ctx context.Context,
	uuid string,
) error {

	var resp DeleteUserResponse

	err := c.do(
		ctx,
		http.MethodDelete,
		"/users/"+uuid,
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

func (c *HTTPClient) GetUserStats(
	ctx context.Context,
	uuid string,
) (*UserStats, error) {

	var stats UserStats

	err := c.do(
		ctx,
		http.MethodGet,
		"/users/"+uuid+"/stats",
		nil,
		&stats,
	)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}
