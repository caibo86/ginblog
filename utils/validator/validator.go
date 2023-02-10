package validator

import (
	"errors"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// Validate 根据tag验证结构体
func Validate(data interface{}) error {
	validate := validator.New()
	zhHansCN := zh_Hans_CN.New()
	uni := ut.New(zhHansCN, zhHansCN)
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		ts := errs.Translate(trans)
		var ss []string
		for _, v := range ts {
			ss = append(ss, v)
		}
		msg := strings.Join(ss, ",")
		return errors.New(msg)
	}
	return nil

}
