package sammi

import "fmt"

type Response struct {
	Data        string `json:"data"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

type ButtonTrigger struct {
	Request  string `json:"request"`
	ButtonId string `json:"buttonID"`
}

type SetVariable struct {
	Request  string `json:"request"`
	Name     string `json:"name"`
	Value    any    `json:"value"`
	ButtonId string `json:"buttonID"`
}

func NewSetVariable(name string, value any, buttonId string) SetVariable {
	return SetVariable{
		"setVariable",
		name,
		value,
		buttonId,
	}
}

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
	return r != nil && r.Data == "Ok."
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
