package sammi

import (
	"encoding/json"
	"fmt"

	"j322.ica/gumroad-sammi/config"
)

type ButtonClient struct {
	Client
	buttonId string
}

func NewButtonClient(c *config.SammiConfig, buttonId string) *ButtonClient {
	client := NewClient(c)
	return &ButtonClient{*client, buttonId}
}

func (c *ButtonClient) PushButton() error {
	req := ButtonTrigger{
		"triggerButton",
		c.buttonId,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("I could not create a sammi button trigger request: %w", err)
	}

	err = c.doPostRequest(data)

	if err != nil {
		return fmt.Errorf("I could not make a sammi button trigger request: %w", err)
	}
	return nil
}

func (c *ButtonClient) Ping() error {
	err := c.doPostRequest(
		[]byte(fmt.Sprintf(`{"request": "releaseButton","buttonID": "%s"}`, c.buttonId)),
	)
	if err != nil {
		return fmt.Errorf("Could not release button '%s': %w", c.buttonId, err)
	}
	return nil
}

func (c *ButtonClient) SetVariable(varName string, value any) error {
	req := NewSetVariable(varName, value, c.buttonId)
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("I could not set variable %s.%s to %s: %w", c.buttonId, varName, value, err)
	}
	return c.doPostRequest(data)
}
