package tool

import (
	"fmt"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type tablesRequest struct {
	DbName string `form:"db_name"` // 数据库名称
}

type tablesResponse struct {
	List []tableData `json:"list"` // 数据表列表
}

type tableData struct {
	Name    string `json:"table_name"`    // 数据表名称
	Comment string `json:"table_comment"` // 数据表备注
}

// Tables 查询 Table
// @Summary 查询 Table
// @Description 查询 Table
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param db_name formData string true "数据库名称"
// @Success 200 {object} tablesResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/tables [post]
// @Security LoginToken
func (h *handler) Tables(ctx *gin.Context) {
	req := new(tablesRequest)
	res := new(tablesResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", req.DbName)

	// TODO 后期支持查询多个数据库
	rows, err := mysql.DB().GetDbR().Raw(sqlTables).Rows()
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var info tableData
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
			continue
		}

		res.List = append(res.List, info)
	}

	api.ResponseOK(ctx, res)
}
