// Package wip provides utilities for interfacing with WIP's API
package wip

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Version is the version of this wrapper
const Version = "0.0.1"

const baseAddress = "https://wip.chat/graphql"

// Client is a client for working with the WIP API
type Client struct {
	http    *http.Client
	baseURL string
	apiKey  string
}

// NewClient creates a Client that will use the specified API key
func NewClient(apikey string) Client {
	return Client{
		baseURL: baseAddress,
		apiKey:  apikey,
	}

}

// doRequest send HTTP request.
func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprint("bearer", s.apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
