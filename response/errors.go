package response

import (
	"encoding/json"
	"fmt"
)

type SillyHatError struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func (she SillyHatError) Error() string {
	sheJSON, err := json.Marshal(she)
	if err != nil {
		return fmt.Sprintf(`{ "code" : "%s", "data" : "%s" , "message" : %s }`, she.Code, she.Data, she.Msg)
	}
	return string(sheJSON)
}

func NewError(msg string) error {
	return SillyHatError{Code: ERROR, Data: nil, Msg: ""}
}

func NewCodeError(code int, msg string) error {
	return SillyHatError{Code: code, Data: nil, Msg: ""}
}
