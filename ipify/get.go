package ipify

import (
	"fmt"
	"io"
	"net/http"
)

func Get() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", fmt.Errorf("I could not get your public ip: %w", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("I could not get your public ip: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("I could not get your public ip: %w", err)
	}

	return string(data), nil
}
