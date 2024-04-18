package sammi

type Response struct {
	Data        any    `json:"data"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

type ButtonTrigger struct {
	Request  string `json:"request"`
	ButtonId string `json:"buttonID"`
}

type SetVariable struct {
	Request  string                 `json:"request"`
	Name     string                 `json:"name"`
	Value    map[string]interface{} `json:"value"`
	ButtonId string                 `json:"buttonID"`
}

func NewSetVariable(name string, value map[string]interface{}, buttonId string) SetVariable {
	return SetVariable{
		"setVariable",
		name,
		value,
		buttonId,
	}
}
