// Package victorops provides a Go client for the VictorOps API
package victorops

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.victorops.com/api-public/v1"

// Client for the VictorOps API
type Client struct {
	*http.Client
	User, ID, Key string
}

// New VictorOps client for the given API id and key
func New(user, apiID, apiKey string) *Client {
	return &Client{
		Client: http.DefaultClient,
		User:   user,
		ID:     apiID,
		Key:    apiKey,
	}
}

// request adds headers and does some basic error handling
func (c *Client) request(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", apiURL, path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-VO-Api-Id", c.ID)
	req.Header.Add("X-VO-Api-Key", c.Key)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Do(req)
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		bts, _ := ioutil.ReadAll(resp.Body)
		return resp, fmt.Errorf("Got a %v error: %v", resp.StatusCode, string(bts))
	}
	return resp, err
}
