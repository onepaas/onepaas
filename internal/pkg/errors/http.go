package errors

import "fmt"

func NewErrorsPayload(errors ...*ErrorObject) *ErrorsPayload {
	errorsPayload := new(ErrorsPayload)
	errorsPayload.Errors = errors

	return errorsPayload
}

// ErrorsPayload is a serializer struct for representing a errors payload.
type ErrorsPayload struct {
	Errors []*ErrorObject `json:"errors"`
}

func NewErrorObject(title string, detail string, code int) *ErrorObject {
	errorObject := ErrorObject{
		Title: title,
		Detail: detail,
		Code: code,
	}

	return &errorObject
}

// ErrorObject is an `Error` implementation as well as an implementation of the error object.
type ErrorObject struct {
	// Title is a short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `json:"title,omitempty"`

	// Detail is a human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized.
	Detail string `json:"detail,omitempty"`

	// Code is an application-specific error code, expressed as a int value.
	Code int `json:"code,omitempty"`
}

// Error implements the `Error` interface.
func (e *ErrorObject) Error() string {
	return fmt.Sprintf("Error: %s %s\n", e.Title, e.Detail)
}
