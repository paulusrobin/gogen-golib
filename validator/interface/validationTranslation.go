package validator

import (
	"github.com/go-playground/locales"
)

type ValidationTranslation interface {
	Translator() locales.Translator
	Register() func() error
}
