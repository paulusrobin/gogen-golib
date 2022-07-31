package goPlaygroundV10

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
	"github.com/rs/zerolog/log"
	"sync"
)

type (
	validation struct {
		validate            *validatorV10.Validate
		universalTranslator *ut.UniversalTranslator
		validators          map[string]validator.Validator
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

	response := &validatorInstance{
		validate:   v.validate,
		translator: translator,
	}
	v.validators[locale] = response
	return response
}

func (v *validation) registerTranslation(translation validator.ValidationTranslation) {
	if err := v.universalTranslator.AddTranslator(translation.Translator(), true); err != nil {
		log.Warn().Msgf("failed to add %s translator", translation.Translator().Locale())
		return
	}

	translator, found := v.universalTranslator.GetTranslator(translation.Translator().Locale())
	if !found {
		log.Warn().Msgf("failed to find %s translator", translation.Translator().Locale())
		return
	}

	if err := translation.Register(translator); err != nil {
		log.Warn().Msgf("failed to register %s translator", translation.Translator().Locale())
		return
	}
}

// NewValidation function to initialize Validation.
func NewValidation(translations ...validator.ValidationTranslation) validator.Validation {
	validate := validatorV10.New()
	defaultLocaleTranslator := en.New()
	universalTranslator := ut.New(defaultLocaleTranslator)
	defaultTranslation := NewValidationTranslation(en.New(),
		func(translator ut.Translator) error {
			return enTranslation.RegisterDefaultTranslations(validate, translator)
		})

	v := &validation{
		validate:            validate,
		universalTranslator: universalTranslator,
		validators:          make(map[string]validator.Validator),
	}
	v.registerTranslation(defaultTranslation)

	for _, translation := range translations {
		v.registerTranslation(translation)
	}
	return v
}
