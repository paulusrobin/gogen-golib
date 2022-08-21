package goPlaygroundV10

import (
	v10 "github.com/go-playground/validator/v10"
	validator "github.com/paulusrobin/gogen-golib/validator/interface"
)

type (
	validationOptions struct {
		translations          map[string]validator.ValidationTranslation
		customFieldValidator  map[string]func(fl v10.FieldLevel) bool
		customStructValidator map[string]func(sl v10.StructLevel)
	}
	ValidationOption interface {
		Apply(options *validationOptions)
	}
)

var defaultOption = validationOptions{
	translations:          make(map[string]validator.ValidationTranslation),
	customFieldValidator:  make(map[string]func(fl v10.FieldLevel) bool),
	customStructValidator: make(map[string]func(sl v10.StructLevel)),
}

func (option *validationOptions) addTranslations(translations ...validator.ValidationTranslation) {
	for _, translation := range translations {
		option.translations[translation.Translator().Locale()] = translation
	}
}

type withTranslation struct {
	validationTranslation validator.ValidationTranslation
}

// Apply implement ValidationOption interface function.
func (w withTranslation) Apply(options *validationOptions) {
	options.addTranslations(w.validationTranslation)
}

// WithTranslation function for add new translation ValidationOption.
func WithTranslation(translation validator.ValidationTranslation) ValidationOption {
	return withTranslation{translation}
}

type withTranslations struct {
	validationTranslations []validator.ValidationTranslation
}

// Apply implement ValidationOption interface function.
func (w withTranslations) Apply(options *validationOptions) {
	options.addTranslations(w.validationTranslations...)
}

// WithTranslations function for add multiple new translation ValidationOption.
func WithTranslations(translations ...validator.ValidationTranslation) ValidationOption {
	return withTranslations{validationTranslations: translations}
}

type withCustomFieldValidator struct {
	name string
	fn   func(fl v10.FieldLevel) bool
}

// Apply implement ValidationOption interface function.
func (w withCustomFieldValidator) Apply(options *validationOptions) {
	options.customFieldValidator[w.name] = w.fn
}

// WithCustomFieldValidator function for add custom field validator ValidationOption.
func WithCustomFieldValidator(tagName string, fn func(fl v10.FieldLevel) bool) ValidationOption {
	return withCustomFieldValidator{
		name: tagName,
		fn:   fn,
	}
}

type withCustomStructValidator struct {
	name string
	fn   func(fl v10.StructLevel)
}

// Apply implement ValidationOption interface function.
func (w withCustomStructValidator) Apply(options *validationOptions) {
	options.customStructValidator[w.name] = w.fn
}

// WithCustomStructValidator function for add custom struct validator ValidationOption.
func WithCustomStructValidator(validationName string, fn func(fl v10.StructLevel)) ValidationOption {
	return withCustomStructValidator{
		name: validationName,
		fn:   fn,
	}
}
