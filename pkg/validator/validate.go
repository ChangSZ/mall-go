package validator

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := TransInit(v, "zh"); err != nil {
			return
		}
		RegisterTagName(v, "alias")
	}
}

func GetError(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		TransErrs := errs.Translate(trans)
		firstErr := ""
		for _, v := range TransErrs {
			firstErr = v
			break
		}
		return errors.New(firstErr)
	} else {
		return err
	}
}

func GetErrors(err error) []error {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		TransErrs := errs.Translate(trans)
		resErrs := make([]error, 0)
		for _, v := range TransErrs {
			resErrs = append(resErrs, errors.New(v))
		}
		return resErrs
	} else {
		return []error{err}
	}
}
