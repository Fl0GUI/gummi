package gumroad

import (
	"path"
	"testing"

	"j322.ica/gumroad-sammi/config"
)

func TestUrlCreationAndCheck(t *testing.T) {
	config := config.GumroadConfig{true, "access_token"}
	pathSecret := pathSecret(&config)
	c := Client{"access_token", "localhost", "1234", "subscription", pathSecret}
	url := c.subscriptionURL()
	t.Logf("Path secret is %#v\n", pathSecret)
	t.Logf("Url is %#v\n", url)
	if err := c.isMyUrl(url); err != nil {
		t.Fatalf("Expected %#v to be my url, since I just created it. Reason: %s", url, err)
	}
	_, secret := path.Split(url)
	t.Logf("secret is %#v\n", secret)
	if err := c.isMyUrlSecret(secret); err != nil {
		t.Fatalf("Expected %#v to be my url secret, but it is not. Reason: %s", secret, err)
	}
}
