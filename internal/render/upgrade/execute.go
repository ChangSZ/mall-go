package upgrade

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/proposal/tablesqls"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type upgradeExecuteRequest struct {
	TableName string `form:"table_name"` // 数据表
	Op        string `form:"op"`         // 操作类型
}

func (h *handler) UpgradeExecute(ctx *gin.Context) {

	upgradeTableList := map[string]map[string]string{
		"authorized": {
			"table_sql":      tablesqls.CreateAuthorizedTableSql(),
			"table_data_sql": tablesqls.CreateAuthorizedTableDataSql(),
		},
		"authorized_api": {
			"table_sql":      tablesqls.CreateAuthorizedAPITableSql(),
			"table_data_sql": tablesqls.CreateAuthorizedAPITableDataSql(),
		},
		"admin": {
			"table_sql":      tablesqls.CreateAdminTableSql(),
			"table_data_sql": tablesqls.CreateAdminTableDataSql(),
		},
		"admin_menu": {
			"table_sql":      tablesqls.CreateAdminMenuTableSql(),
			"table_data_sql": tablesqls.CreateAdminMenuTableDataSql(),
		},
		"menu": {
			"table_sql":      tablesqls.CreateMenuTableSql(),
			"table_data_sql": tablesqls.CreateMenuTableDataSql(),
		},
		"menu_action": {
			"table_sql":      tablesqls.CreateMenuActionTableSql(),
			"table_data_sql": tablesqls.CreateMenuActionTableDataSql(),
		},
		"cron_task": {
			"table_sql":      tablesqls.CreateCronTaskTableSql(),
			"table_data_sql": "",
		},
	}

	upgradeTableOp := map[string]bool{
		"table":      true,
		"table_data": true,
	}

	req := new(upgradeExecuteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	outPutString := ""
	db := mysql.DB().GetDbW()

	if upgradeTableList[req.TableName] == nil {
		log.WithTrace(ctx).Error("数据表不存在")
		api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, "数据表不存在")
		return
	}

	if !upgradeTableOp[req.Op] {
		log.WithTrace(ctx).Error("非法操作")
		api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, "非法操作")
		return
	}

	if req.Op == "table" {
		if err := db.Exec(upgradeTableList[req.TableName]["table_sql"]).Error; err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, err)
			return
		}

		outPutString = "初始化 MySQL 数据表：" + req.TableName + " 成功。"
	} else if req.Op == "table_data" {
		if err := db.Exec(upgradeTableList[req.TableName]["table_data_sql"]).Error; err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, err)
			return
		}

		outPutString = "初始化 MySQL 数据表：" + req.TableName + " 默认数据成功。"
	}

	api.ResponseOK(ctx, outPutString)
}
