package libs

import (
	"fmt"
	"strings"
)

type ValidationBuilder struct {
	Errors []ValidationError
}


func NewValidator() ValidationBuilder {
	return ValidationBuilder{}
}

func (v ValidationBuilder) NotEmpty(fieldName string, value string) ValidationBuilder {
	if value != "" {
		return ValidationBuilder{
			Errors: v.Errors,
		}
	}

	err := ValidationError{
		FieldName:    fieldName,
		ErrorMessage: fmt.Sprintf("%s can't be empty", fieldName),
	}

	return ValidationBuilder{
		Errors: append(v.Errors, err),
	}
}

func (v ValidationBuilder) Must(fieldName string, errorMsg string, condition func() bool) ValidationBuilder {
	if condition() {
		return ValidationBuilder{
			Errors: v.Errors,
		}
	}

	err := ValidationError{
		FieldName:    fieldName,
		ErrorMessage: errorMsg,
	}

	return ValidationBuilder{
		Errors: append(v.Errors, err),
	}
}

func (v ValidationBuilder) Ok() bool {
	return len(v.Errors) == 0
}

type ValidationError struct {
	FieldName    string
	ErrorMessage string
}

func Valid() []ValidationError  {
	return []ValidationError{}
}

func SingleValidationError(fieldName string, msg string) []ValidationError {
	err := ValidationError{
		FieldName:    fieldName,
		ErrorMessage: msg,
	}
	return []ValidationError{err}
}

func ToResponse(errs []ValidationError) string {
	builder := strings.Builder{}
	builder.WriteString("validation errors: ")

	for _, err := range errs {
		_, e := builder.WriteString(fmt.Sprintf("(%s: %s)", err.FieldName, err.ErrorMessage))
		if e != nil {
			panic(fmt.Sprintf("can't write validation erros to response: %s", err))
		}
	}

	return builder.String()
}