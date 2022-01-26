package validator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	gpValidator "github.com/go-playground/validator/v10"
	jatranslations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/hashicorp/go-multierror"
)

func Validate(target interface{}) error {
	translator := ja.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("ja")

	if !found {
		return nil
	}

	validate := gpValidator.New()

	// タグを日本語に変換
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("ja")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})

	err := jatranslations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		return err
	}

	err = validate.Struct(target)
	if err == nil {
		return nil
	}

	fmt.Println("sentinel5")
	// TODO(okubo): ここでエラー起きているので、詳細積める
	var result error
	fmt.Println("sentinel5.5")
	fmt.Println(err)
	for _, e := range err.(gpValidator.ValidationErrors) {
		fmt.Println("errors.New(e.Translate(trans))")
		fmt.Println(errors.New(e.Translate(trans)))
		result = multierror.Append(result, errors.New(e.Translate(trans)))
	}

	return result
}
