package httputils

import "fmt"

func Success() interface{} {
	return response{Code: OK.Code, Message: OK.Message, Data: emptyStruct}
}

func SuccessWithData(v interface{}) interface{} {
	return response{Code: OK.Code, Message: OK.Message, Data: v}
}

func Error(err error) interface{} {
	switch t := err.(type) {
	case response:
		return response{Code: t.Code, Message: t.Message, Data: emptyStruct}
	default:
		return response{Code: 500, Message: err.Error(), Data: emptyStruct}
	}
}

var emptyStruct struct{}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (e response) Error() string {
	return fmt.Sprintf("%d|%s", e.Code, e.Message)
}
