package config

import (
	"go/token"
	"log"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

const minBusinessCode = 20000

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{logger: logger}
}

func (h *handler) Email() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("config_email", configs.Get())
	}
}

func (h *handler) Code() core.HandlerFunc {
	type codes struct {
		Code    int    `json:"code"`    // 错误码
		Message string `json:"message"` // 错误码信息
	}

	type codeViewResponse struct {
		SystemCodes   []codes
		BusinessCodes []codes
	}

	parsedFile, err := decorator.Parse(code.ByteCodeFile)
	if err != nil {
		log.Fatalf("parsing code.go: %s: %s\n", "ByteCodeFile", err)
	}

	var (
		systemCodes   []codes
		businessCodes []codes
	)

	dst.Inspect(parsedFile, func(n dst.Node) bool {
		// GenDecl 代表除函数以外的所有声明，即 import、const、var 和 type
		decl, ok := n.(*dst.GenDecl)
		if !ok || decl.Tok != token.CONST {
			return true
		}

		for _, spec := range decl.Specs {
			valueSpec, _ok := spec.(*dst.ValueSpec)
			if !_ok {
				continue
			}

			codeInt := cast.ToInt(valueSpec.Values[0].(*dst.BasicLit).Value)

			if codeInt < minBusinessCode {
				systemCodes = append(systemCodes, codes{
					Code:    codeInt,
					Message: code.Text(codeInt),
				})
			} else {
				businessCodes = append(businessCodes, codes{
					Code:    codeInt,
					Message: code.Text(codeInt),
				})
			}

		}

		return true
	})

	return func(ctx core.Context) {
		obj := new(codeViewResponse)
		obj.BusinessCodes = businessCodes
		obj.SystemCodes = systemCodes

		ctx.HTML("config_code", obj)
	}
}
