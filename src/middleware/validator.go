package middleware

import validator "github.com/go-playground/validator/v10"

// Validator 参数验证器
type Validator struct {
	Validator *validator.Validate
}

// Validate 验证参数
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// GetValidator 获得验证器
func GetValidator() *Validator {
	return &Validator{
		Validator: validator.New(),
	}
}
