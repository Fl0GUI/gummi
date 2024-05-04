package gumroad

import (
	"testing"

	"j322.ica/gumroad-sammi/config"
)

func TestUrlCreationAndCheck(t *testing.T) {
	config := config.GumroadConfig{true, "access_token"}
	pathSecret := pathSecret(&config)
	c := Client{"N6PXhi0Fy9qs-jQxX2bs5AjJZEssvqEn6ukFQCZNNUU", "localhost", "1234", "", pathSecret}
	url := c.subscriptionURL()
	t.Logf("Path secret is %#v\n", pathSecret)
	t.Logf("Url is %#v\n", url)
	if err := c.isMyUrl(url); err != nil {
		t.Fatalf("Expected %s to be my url, since I just created it. Reason: %s", url, err)
	}
}
