package sammi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

type Client struct {
	api      string
	password string
	buttonId string
}

func NewClient(c *config.Config) *Client {
	conf := c.SammiConfig
	return &Client{
		fmt.Sprintf("http://%s:%s/api", conf.Host, conf.Port),
		conf.Password,
		conf.ButtonId,
	}
}

func (c *Client) Ping() error {
	err := c.doRequest(
		[]byte(fmt.Sprintf(`{"request": "releaseButton","buttonID": "%s"}`, c.buttonId)),
	)
	if err != nil {
		return fmt.Errorf("Could not release the button: %w", err)
	}
	return nil
}

func (c *Client) PushButton() error {
	req := ButtonTrigger{
		"triggerButton",
		c.buttonId,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("I could not create a sammi button trigger request: %w", err)
	}

	err = c.doRequest(data)

	if err != nil {
		return fmt.Errorf("I could not make a sammi button trigger request: %w", err)
	}
	return nil
}

func (c *Client) SetVariable(varName string, value any) error {
	req := NewSetVariable(varName, value, c.buttonId)
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("I could not set variable %s.%s to %s: %w", c.buttonId, varName, value, err)
	}
	return c.doRequest(data)
}

func (c *Client) doRequest(data []byte) error {
	req, err := http.NewRequest(http.MethodPost, c.api, bytes.NewReader(data))
	req.Header["Authorization"] = []string{c.password}
	if err != nil {
		return fmt.Errorf("Invalid request: %w", err)
	}
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
