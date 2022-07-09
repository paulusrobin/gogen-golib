package v10

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
)

type (
	validationTranslation struct {
		translator locales.Translator
		register   func(translator ut.Translator) error
	}
)

// Translator getter function to implement ValidationTranslation.
func (v validationTranslation) Translator() locales.Translator {
	return v.translator
}

// Register getter function to implement ValidationTranslation.
func (v validationTranslation) Register(translator ut.Translator) error {
	return v.register(translator)
}

// NewValidationTranslation function to initialize ValidationTranslation interface.
func NewValidationTranslation(
	translator locales.Translator,
	register func(translator ut.Translator) error,
) validator.ValidationTranslation {
	return &validationTranslation{translator: translator, register: register}
}
