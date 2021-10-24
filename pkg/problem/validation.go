package problem

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewValidationProblem(errors validator.ValidationErrors, opts ...Option) *Problem {
	validationErrors := make([]map[string]string, 0)

	for _, err := range errors {
		validationError := map[string]string{
			"pointer": err.Field(),
			"detail":  err.(error).Error(),
		}

		validationErrors = append(validationErrors, validationError)
	}

	opts = append(
		[]Option{WithExtension("validationErrors", validationErrors)},
		opts...,
	)

	return NewStatusProblem(http.StatusUnprocessableEntity, opts...)
}
