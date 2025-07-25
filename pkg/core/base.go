package core

import (
	"fmt"
	"net/url"

	"github.com/Sunhill666/goalex/internal/model"
)

// ListEntities retrieves a paginated list of entities from the specified endpoint.
func ListEntities[T any](
	c *Client,
	endpoint string,
	params *QueryParams,
) (*model.PaginatedResponse[T], error) {
	q := url.Values{}
	if params != nil {
		q = params.ToQuery()
	}
	urlWithParams := endpoint
	if encoded := q.Encode(); encoded != "" {
		urlWithParams += "?" + encoded
	}

	var resp model.PaginatedResponse[T]
	err := c.Get(urlWithParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEntity retrieves a single entity by ID from the specified endpoint.
func GetEntity[T any](c *Client, endpoint, id string) (*T, error) {
	var entity T
	err := c.Get(fmt.Sprintf("%s/%s", endpoint, id), &entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
