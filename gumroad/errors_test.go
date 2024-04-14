package gumroad

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	if !errors.Is(SelfTestFailed{"url"}, SelfTestFailed{"lru"}) {
		t.Errorf("Expected SelfTestError to be equal to SelfTestError regardless of url")
	}
}
