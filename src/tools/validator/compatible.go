package validator

import (
	"errors"
	"fmt"
	"openaigo/src/lib/logger"
	"reflect"
	"strings"

	localeEn "github.com/go-playground/locales/en"
	translationsEn "github.com/go-playground/validator/v10/translations/en"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func newCustomValidate(log logger.ILogger) *CustomValidator {
	v := new(CustomValidator)
	v.validate = validator.New()
	v.validate.SetTagName("validate")
	v.trans = ut.New(localeEn.New()).GetFallback()
	v.log = log
	if err := translationsEn.RegisterDefaultTranslations(v.validate, v.trans); err != nil {
		log.ApplicationErrorSimple("register en translation err", err)
	}

	// 自定義驗證器 https://pkg.go.dev/github.com/go-playground/validator/v10
	for _, item := range customValidateCases {
		if err := v.RegisterValidate(item.Tag, item.Func, v.trans,
			item.RegisterTranslationsFunc, v.translateFunc); err != nil {
			log.ApplicationErrorSimple("register custom validate case err", err)
			return nil
		}
	}
	return v
}

type CustomValidator struct {
	validate *validator.Validate
	trans    ut.Translator
	log      logger.ILogger
}

// Parse : 給資料格式驗證， do fiber BodyParser & validate struct
func (v *CustomValidator) Parse(ctx *fiber.Ctx, vs interface{}) error {
	if err := ctx.BodyParser(vs); err != nil {
		return err
	}
	return v.validateStruct(vs)
}

func (v *CustomValidator) validateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			errs := err.(validator.ValidationErrors)
			messages := make([]string, len(errs))
			for i, e := range errs {
				messages[i] = e.Translate(v.trans)
			}
			return errors.New(strings.Join(messages, ", "))
		}
	}
	return nil
}

// RegisterValidate 註冊自定義驗證器
func (v *CustomValidator) RegisterValidate(tag string,
	fn validator.Func,
	trans ut.Translator,
	registerFn validator.RegisterTranslationsFunc,
	translationFn validator.TranslationFunc) (err error) {

	// 註冊參數驗證
	if err = v.validate.RegisterValidation(tag, fn); err != nil {
		return err
	}
	// 註冊翻譯
	if err = v.validate.RegisterTranslation(tag, trans, registerFn, translationFn); err != nil {
		return err
	}
	return err
}

// 翻譯字串
func (v *CustomValidator) translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		v.log.ApplicationError(
			fmt.Sprintf("validator translateFunc err, fe: %v, err: %v",
				fe,
				err),
		)
		return fe.(error).Error()
	}
	return t
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
