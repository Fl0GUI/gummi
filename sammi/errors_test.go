package sammi

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	if !errors.Is(ServerError, ServerError) {
		t.Errorf("Expected ServerError to be equal to a ServerError.")
	}
	if !errors.Is(AuthError, AuthError) {
		t.Errorf("Expected AuthError to be equal to a AuthError.")
	}
	if !errors.Is(ButtonIdNotFoundError, ButtonIdNotFoundError) {
		t.Errorf("Expected ButtonIdNotFoundError to be equal to a ButtonIdNotFoundError.")
	}
	if errors.Is(AuthError, ButtonIdNotFoundError) {
		t.Errorf("Expected AuthError to NOT be equal to a ButtonIdNotFoundError.")
	}

	r := Response{
		"",
		"Authorization failed.",
		"No authorization header was found or wrong value was provided. Please verify your Authorization header matches your API Password in SAMMI Settings.",
	}

	if r.Ok() {
		t.Errorf("Expected test response to be not ok.")
	}

	err := r.Err()
	if !errors.Is(err, AuthError) {
		t.Errorf("Expected test response to be a AuthError.")
	}
	if errors.Is(err, ButtonIdNotFoundError) {
		t.Errorf("Expected test response to NOT be a ButtonIdNotFoundError.")
	}
}
