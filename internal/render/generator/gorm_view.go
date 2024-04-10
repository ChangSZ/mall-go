package generator_handler

import (
	"fmt"
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func (h *handler) GormView(ctx *gin.Context) {

	type tableInfo struct {
		Name    string `db:"table_name"`    // name
		Comment string `db:"table_comment"` // comment
	}

	var tableCollect []tableInfo
	mysqlConf := configs.Get().MySQL.Read
	sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", mysqlConf.Name)
	rows, err := mysql.DB().GetDbR().Raw(sqlTables).Rows()
	if err != nil {
		log.WithTrace(ctx).Error("rows err: ", err)
		ctx.HTML(http.StatusOK, "generator_gorm.html", tableCollect)
		return
	}

	err = rows.Err()
	if err != nil {
		log.WithTrace(ctx).Error("rows err", zap.Error(err))
		ctx.HTML(http.StatusOK, "generator_gorm.html", tableCollect)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var info tableInfo
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
			continue
		}
		tableCollect = append(tableCollect, info)
	}

	ctx.HTML(http.StatusOK, "generator_gorm.html", tableCollect)
}
