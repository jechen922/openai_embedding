package validator

import (
	"strconv"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CustomCase struct {
	Tag string
	validator.Func
	validator.RegisterTranslationsFunc
}

var customValidateCases = []CustomCase{
	{"id", ValidateID, IDTranslationsFunc},
}

// ValidateID 驗證 ID
func ValidateID(fl validator.FieldLevel) bool {
	idUint64, err := strconv.ParseUint(fl.Field().Interface().(string), 10, 64)
	if err != nil {
		return false
	}
	if idUint64 == 0 {
		return false
	}
	return true
}

// IDTranslationsFunc 轉換為編號錯誤
func IDTranslationsFunc(ut ut.Translator) error {
	return ut.Add("id", "{0} is invalid", false)
}
