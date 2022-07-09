package v10

import (
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	validator2 "github.com/paulusrobin/gogen-golib/validator/interface"
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

	var err = validator2.ValidationError{
		Message: defaultMessage,
		Details: make([]string, 0),
		detail:  make(map[string]string),
	}

	errs := validationErr.(v10.ValidationErrors)
	for _, e := range errs {
		translated := e.Translate(v.translator)
		err.detail[e.StructNamespace()] = translated
		err.Details = append(err.Details, translated)
	}
	return err
}
