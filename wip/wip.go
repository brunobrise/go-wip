// Package wip provides utilities for interfacing with WIP's API
package wip

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// Version is the version of this wrapper
const Version = "0.1"

const baseAddress = "https://wip.chat/graphql"

// Client is a client for working with the WIP API
type Client struct {
	baseURL string
	apiKey  string

	graphql *graphql.Client
}

// NewClient creates a GraphQL client that uses the specified apiKey key
// to connect to server at baseAddress
func NewClient(apikey string) Client {
	return Client{
		apiKey:  apikey,
		graphql: graphql.NewClient(baseAddress, graphql.UseMultipartForm()),
	}

}

// do connect to GraphQL server, execute the specified request
// and mutate the provided variable
func (c *Client) do(req *graphql.Request, result interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	ctx := context.Background()

	err := c.graphql.Run(ctx, req, &result)
	if err != nil {
		return err
	}

	return nil
}
