package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
)

func Validate(data interface{}) (string, error) {
	validate := validator.New()
	zhHansCN := zh_Hans_CN.New()
	uni := ut.New(zhHansCN, zhHansCN)
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return "", err
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		log.Println("aaa", err)
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Translate(trans))
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Println(v.Translate(trans))
			return v.Translate(trans), v
		}
	}
	return "", nil
}
