package sammi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"j322.ica/gumroad-sammi/config"
)

type Client struct {
	api      string
	webhook  string
	password string
}

func NewClient(c *config.SammiConfig) *Client {
	return &Client{
		fmt.Sprintf("http://%s:%s/api", c.Host, c.Port),
		fmt.Sprintf("http://%s:%s/webhook", c.Host, c.Port),
		c.Password,
	}
}

func (c *Client) Ping() error {
	query := url.Values{}
	query.Add("request", "getVariable")
	query.Add("name", "api_server_opened")
	data, err := c.doGetRequest(query.Encode())
	if err != nil {
		return err
	}
	res := Response{}
	err = json.NewDecoder(data).Decode(&res)
	if err != nil {
		return err
	}
	return res.Err()
}

func (c *Client) Trigger(trigger string, data map[string]interface{}) error {
	data["trigger"] = trigger
	sData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.doPostRequestTo(sData, c.webhook)
}

func (c *Client) doPostRequest(data []byte) error {
	return c.doPostRequestTo(data, c.api)
}

func (c *Client) doPostRequestTo(data []byte, api string) error {
	req, err := http.NewRequest(http.MethodPost, api, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("Invalid request: %w", err)
	}
	req.Header["Authorization"] = []string{c.password}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res := &Response{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return fmt.Errorf("Could not decode response: %w", err)
	}
	return res.Err()
}

func (c *Client) doGetRequest(query string) (io.ReadCloser, error) {
	url, err := url.Parse(c.api)
	if err != nil {
		panic(err)
	}
	url.RawQuery = query

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Invalid request: %w", err)
	}
	req.Header["Authorization"] = []string{c.password}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, serverError{c.api}
	}
	return resp.Body, nil
}
