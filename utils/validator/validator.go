package validator

import (
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptTans "github.com/go-playground/validator/v10/translations/pt"
	"log"
	"reflect"
)

func Validate(data interface{}) (string, error) {
	validate := validator.New()
	uni := ut.New(pt.New())
	trans, _ := uni.GetTranslator("pt")
	err := ptTans.RegisterDefaultTranslations(validate, trans)
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
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), v
		}
	}
	return "", nil
}
