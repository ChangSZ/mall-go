package role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 批量删除角色
// @Summary 批量删除角色
// @Description 批量删除角色
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /role/delete [post]
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
