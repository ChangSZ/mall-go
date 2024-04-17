package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// RegisterTagName 调整报错信息中的字段提示
func RegisterTagName(v *validator.Validate, name string) {
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(name)
	})
}
