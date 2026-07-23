package dto

import (
	"net/http"
)

type Fail struct {
	Code    int
	Message string
}

func (f *Fail) Error() string {
	return f.Message
}

func FailClientNew(code int) error {
	return &Fail{
		Code:    code,
		Message: http.StatusText(code),
	}
}

func FailServerNew(message string, err error) error {
	if err != nil {
		message = message + ": " + err.Error()
	}
	return &Fail{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
