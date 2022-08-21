package goPlaygroundV10

import validator "github.com/paulusrobin/gogen-golib/validator/interface"

type (
	validationOptions struct {
		translations map[string]validator.ValidationTranslation
	}
	ValidationOption interface {
		Apply(options *validationOptions)
	}
)

var defaultOption = validationOptions{
	translations: make(map[string]validator.ValidationTranslation),
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
