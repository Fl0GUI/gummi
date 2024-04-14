package gumroad

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) Ping() error {
	urlString := fmt.Sprintf("http://%s:%s", c.ip, c.port)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, urlString, nil)
	if err != nil {
		return fmt.Errorf("Incorrect ip and port: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SelfTestFailed{urlString}
	}

	if resp.StatusCode != 200 {
		return SelfTestFailed{urlString}
	}
	return nil
}
