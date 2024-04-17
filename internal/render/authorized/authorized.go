package authorized

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorized_add.html", nil)
}

func (h *handler) Demo(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorized_demo.html", nil)
}

func (h *handler) List(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorized_list.html", nil)
}

func (h *handler) Api(ctx *gin.Context) {
	type apiRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type apiResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	req := new(apiRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	obj := new(apiResponse)
	obj.HashID = req.Id

	ctx.HTML(http.StatusOK, "authorized_api.html", obj)
}
