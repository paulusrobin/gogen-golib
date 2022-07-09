package v10

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/paulusrobin/gogen-golib/validator"
	"github.com/rs/zerolog/log"
	"sync"
)

type (
	validation struct {
		validate            *validatorV10.Validate
		universalTranslator *ut.UniversalTranslator
		validators          map[string]Validator
		sync.Mutex
	}
)

// Validator function implement Validation interface to get locale validator.
func (v *validation) Validator(locale string) validator.Validator {
	v.Lock()
	defer v.Unlock()

	if localeValidator, found := v.validators[locale]; found {
		return localeValidator
	}

	translator, found := v.universalTranslator.GetTranslator(locale)
	if !found {
		log.Warn().Msgf("translator %s not found, use fallback translation")
	}

	response := &validator{
		validate:   v.validate,
		translator: translator,
	}
	v.validators[locale] = response
	return response
}

func (v *validation) registerTranslation(translation ValidationTranslation) {
	if err := v.universalTranslator.AddTranslator(translation.Translator(), true); err != nil {
		log.Warn().Msgf("failed to add %s translator", translation.Translator().Locale())
		return
	}

	translator, found := v.universalTranslator.GetTranslator(translation.Translator().Locale())
	if !found {
		log.Warn().Msgf("failed to find %s translator", translation.Translator().Locale())
		return
	}

	if err := translation.Register(v.validate, translator); err != nil {
		log.Warn().Msgf("failed to register %s translator", translation.Translator().Locale())
		return
	}
}

// NewValidation function to initialize Validation.
func NewValidation(translations ...ValidationTranslation) Validation {
	defaultTranslation := NewValidationTranslation(en.New(),
		func(validate *validatorV10.Validate, translator ut.Translator) error {
			return enTranslation.RegisterDefaultTranslations(validate, translator)
		})

	v := &validation{
		validate:            validatorV10.New(),
		universalTranslator: ut.New(defaultTranslation.Translator()),
	}
	v.registerTranslation(defaultTranslation)

	for _, translation := range translations {
		v.registerTranslation(translation)
	}
	return v
}
