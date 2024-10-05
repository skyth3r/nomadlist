package nomadlist

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

var client *Client

func init() {
	if client != nil {
		return
	}

	client = &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewClient() *Client {
	return client
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	c.setHeaders(req)
	rsp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) setHeaders(r *http.Request) *http.Request {
	r.Header.Set("User-Agent", "Nomadlist Go Client (https://github.com/skyth3r/nomadlist)")
	return r
}
