package sammi

import "fmt"

// SammiError

type SammiError struct {
	Err         string
	Description string `json:"description"`
}

var AuthError = &SammiError{
	"Authorization failed.",
	"No authorization header was found or wrong value was provided. Please verify your Authorization header matches your API Password in SAMMI Settings.",
}

var ButtonIdNotFoundError = &SammiError{
	"Button ID not found.",
	"You're trying to release a button that does not exist, or is not persistent.",
}

func (r *Response) Ok() bool {
	return r != nil && r.Error == ""
}

func (r *Response) Err() error {
	if r.Ok() {
		return nil
	}
	return &SammiError{
		r.Error,
		r.Description,
	}
}

func (s *SammiError) Error() string {
	return fmt.Sprintf("%s Details: %s", s.Err, s.Description)
}

func (s *SammiError) Is(target error) bool {
	sammiErrTarget, ok := target.(*SammiError)
	if !ok {
		return false
	}
	return s.Err == sammiErrTarget.Err
}

// serverError

type serverError struct {
	api string
}

var ServerError = serverError{""}

func (e serverError) Error() string {
	return fmt.Sprintf("Could not contact Sammi on %s", e.api)
}

func (e serverError) Is(target error) bool {
	_, ok := target.(serverError)
	return ok
}
