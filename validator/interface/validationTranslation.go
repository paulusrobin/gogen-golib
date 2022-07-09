package validator

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
)

type ValidationTranslation interface {
	Translator() locales.Translator
	Register(translator ut.Translator) error
}
