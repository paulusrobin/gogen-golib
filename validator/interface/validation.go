package validator

type Validation interface {
	Validator(locale string) Validator
}
