package logentriesapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client Type and constructor

type Client struct {
	accountKey string
	baseURL    *url.URL
}

func NewClient(accountKey string) (le *Client) {
	le = new(Client)
	le.accountKey = accountKey
	le.baseURL, _ = url.Parse("http://api.logentries.com/")
	return
}

// Base Types

type APIResponse struct {
	Response string `json:"response"`
	Reason   string `json:"reason,omitempty"`
}

func (r *APIResponse) AssertOk() error {
	if r.Response == "ok" {
		return nil
	}
	return fmt.Errorf("API returned non-okay status: %s with reason: %s", r.Response, r.Reason)
}

type APIObject struct {
	Type string `json:"object"`
}

func (o *APIObject) AssertType(t string) error {
	if o.Type == t {
		return nil
	}
	return fmt.Errorf("Logentries API: Expected type %s but recieved type %s", t, o.Type)
}

// low level helpers

func (c *Client) getBytes(path string) ([]byte, error) {
	reqURL := *c.baseURL
	reqURL.Path = c.accountKey + path

	res, err := http.Get(reqURL.String())
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}
