package validator

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func TransInit(v *validator.Validate, local string) (err error) {
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(enT, zhT)
	tran, ok := uni.GetTranslator(local)
	trans = tran
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", local)
	}
	switch local {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		if err = zhTranslations.RegisterDefaultTranslations(v, trans); err != nil {
			return
		}
		err = extendZhTrans(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}

func extendZhTrans(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required_if",
			translation: "{0}为必填字段",
			override:    false,
			//customRegisFunc: func(ut ut.Translator) error {
			//	return ut.Add("required_if", "{0}为必填字段", false)
			//},
			//customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
			//	t, _ := ut.T("required_if", fe.Field())
			//	return t
			//},
		},
	}

	for _, t := range translations {
		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}
		if err != nil {
			return
		}
	}
	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("警告: 翻译字段错误: %#v", fe)
		return fe.(error).Error()
	}
	return t
}
