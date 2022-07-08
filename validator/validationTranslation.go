package validator

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
)

type (
	ValidationTranslation interface {
		Translator() locales.Translator
		Register(validate *v10.Validate, translator ut.Translator) error
	}

	validationTranslation struct {
		translator locales.Translator
		register   func(validate *v10.Validate, translator ut.Translator) error
	}
)

// Translator getter function to implement ValidationTranslation.
func (v validationTranslation) Translator() locales.Translator {
	return v.translator
}

// Register getter function to implement ValidationTranslation.
func (v validationTranslation) Register(validate *v10.Validate, translator ut.Translator) error {
	return v.register(validate, translator)
}

// NewValidationTranslation function to initialize ValidationTranslation interface.
func NewValidationTranslation(
	translator locales.Translator,
	register func(validate *v10.Validate, translator ut.Translator) error,
) ValidationTranslation {
	return &validationTranslation{translator: translator, register: register}
}
