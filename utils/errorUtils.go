package utils

import "encoding/json"

type ErrorStruct struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type GenericError struct {
	Errors []ErrorStruct `json:"errors"`
	Data   *string       `json:"data"`
}

func GetError(message string, code int) []byte {

	myErr := ErrorStruct{
		Message: message,
		Code:    code,
	}

	genericErrorObj := GenericError{
		Errors: []ErrorStruct{myErr},
		Data:   nil,
	}

	errJson, _ := json.Marshal(genericErrorObj)

	return errJson

}
