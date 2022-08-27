package entities

import "errors"

type echoString struct {
	Message string `json:"message"`
}

func NewEchoString(str string) (*echoString, error) {
	if str == "" {
		return nil, errors.New("invalid parameters")
	}

	return &echoString{
		Message: str,
	}, nil

}

func (t *echoString) GetEchoString() (string, error) {
	return t.Message, nil
}
