package newserror

import "encoding/json"

const (
	InternalErr = iota
	ExternalErr = iota
	DbErr
	NofFoundErr
)

type ErrorNews struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewErrorNews(code int, message string, details string) error {
	return &ErrorNews{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (e ErrorNews) Error() string {
	errText, err := json.Marshal(&e)
	if err != nil {
		return ""
	}
	return string(errText)
}
