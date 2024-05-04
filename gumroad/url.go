package gumroad

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"j322.ica/gumroad-sammi/config"
)

var baseUrl = "gumroad"

func (c *Client) isMyUrl(urlString string) error {
	u, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("not an url: %v", err)
	}
	path := u.EscapedPath()
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return fmt.Errorf("url path has %#v parts instead of 3", len(parts))
	}
	if parts[1] != baseUrl {
		return fmt.Errorf("url does not begin with %#v, but with %#v", baseUrl, parts[0])
	}
	return nil
}

func (c *Client) isMyUrlSecret(urlSecret string) error {
	if urlSecret != c.pathSecret {
		return fmt.Errorf("url does not match the path secret")
	}
	return nil
}

func (c *Client) subscriptionURL() string {
	return fmt.Sprintf("http://%s:%s/%s/%s", c.ip, c.port, baseUrl, c.pathSecret)
}

func pathSecret(c *config.GumroadConfig) string {
	// salted hash
	token := c.AccessToken
	secret, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	// encode to base64
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(secret)
}
