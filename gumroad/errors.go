package gumroad

import (
	"errors"
	"fmt"
)

type SelfTestFailed struct {
	url string
}

func (e SelfTestFailed) Error() string {
	return fmt.Sprintf("I could not ping my own server over on %s", e.url)
}

func (e SelfTestFailed) Is(target error) bool {
	_, ok := target.(SelfTestFailed)
	return ok
}

var Unauthorized = errors.New("Gumroad access is unauthorized because of missing or incorrect Access Token.")
