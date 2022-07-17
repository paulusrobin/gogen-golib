package goPlaygroundV10

import (
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
)

const (
	defaultMessage = "validation error"
)

type (
	validatorInstance struct {
		validate   *v10.Validate
		translator ut.Translator
	}
)

// Struct validates a structs exposed fields.
func (v validatorInstance) Struct(s interface{}) error {
	validationErr := v.validate.Struct(s)
	if validationErr == nil {
		return nil
	}

	var err = validator.ValidationError{
		Message: defaultMessage,
		Details: make([]string, 0),
	}

	errs := validationErr.(v10.ValidationErrors)
	for _, e := range errs {
		translated := e.Translate(v.translator)
		err.SetDetail(e.StructNamespace(), translated)
		err.Details = append(err.Details, translated)
	}
	return err
}
