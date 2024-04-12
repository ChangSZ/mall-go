package install

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/proposal/tablesqls"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type initExecuteRequest struct {
	Language  string `form:"language" `  // 语言包
	RedisAddr string `form:"redis_addr"` // 连接地址，例如：127.0.0.1:6379
	RedisPass string `form:"redis_pass"` // 连接密码
	RedisDb   string `form:"redis_db"`   // 连接 db

	MySQLAddr string `form:"mysql_addr"`
	MySQLUser string `form:"mysql_user"`
	MySQLPass string `form:"mysql_pass"`
	MySQLName string `form:"mysql_name"`
}

func (h *handler) Execute(ctx *gin.Context) {

	installTableList := map[string]map[string]string{
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

	req := new(initExecuteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError)
		return
	}

	// region 验证 version
	versionStr := runtime.Version()
	version := cast.ToFloat32(versionStr[2:6])
	if version < configs.MinGoVersion {
		api.Response(ctx, http.StatusBadRequest, code.GoVersionError)
		return
	}
	// endregion

	// region 验证 Redis 配置
	cfg := configs.Get()
	redisClient := redis.NewClient(&redis.Options{
		Addr:         req.RedisAddr,
		Password:     req.RedisPass,
		DB:           cast.ToInt(req.RedisDb),
		MaxRetries:   cfg.Redis.MaxRetries,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	})

	if err := redisClient.Ping().Err(); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.RedisConnectError)
		return
	}

	defer redisClient.Close()

	outPutString := "已检测 Redis 配置可用。\n"
	// endregion

	// region 验证 MySQL 配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		req.MySQLUser,
		req.MySQLPass,
		req.MySQLAddr,
		req.MySQLName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MySQLConnectError)
		return
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	dbClient, _ := db.DB()
	defer dbClient.Close()

	outPutString += "已检测 MySQL 配置可用。\n"
	// endregion

	// region 写入配置文件
	viper.Set("language.local", req.Language)

	viper.Set("redis.addr", req.RedisAddr)
	viper.Set("redis.pass", req.RedisPass)
	viper.Set("redis.db", req.RedisDb)

	viper.Set("mysql.read.addr", req.MySQLAddr)
	viper.Set("mysql.read.user", req.MySQLUser)
	viper.Set("mysql.read.pass", req.MySQLPass)
	viper.Set("mysql.read.name", req.MySQLName)

	viper.Set("mysql.write.addr", req.MySQLAddr)
	viper.Set("mysql.write.user", req.MySQLUser)
	viper.Set("mysql.write.pass", req.MySQLPass)
	viper.Set("mysql.write.name", req.MySQLName)

	if viper.WriteConfig() != nil {
		api.Response(ctx, http.StatusBadRequest, code.WriteConfigError)
		return
	}

	outPutString += "语言包 " + req.Language + " 配置成功。\n"
	outPutString += "配置项 Redis、MySQL 配置成功。\n"
	// endregion

	// region 初始化表结构 + 默认数据
	for k, v := range installTableList {
		if v["table_sql"] != "" {
			// region 初始化表结构
			if err = db.Exec(v["table_sql"]).Error; err != nil {
				log.WithTrace(ctx).Error(err)
				api.Response(ctx, http.StatusBadRequest, code.MySQLConnectError, err)
				return
			}

			outPutString += "初始化 MySQL 数据表：" + k + " 成功。\n"
			// endregion

			// region 初始化默认数据
			if v["table_data_sql"] != "" {
				if err = db.Exec(v["table_data_sql"]).Error; err != nil {
					log.WithTrace(ctx).Error(err)
					api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, err)
					return
				}

				outPutString += "初始化 MySQL 数据表：" + k + " 默认数据成功。\n"
			}
			// endregion
		}
	}
	// endregion

	// region 生成 install 完成标识
	f, err := os.Create(configs.ProjectInstallMark)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MySQLExecError, err)
		return
	}
	defer f.Close()
	// endregion

	api.ResponseOK(ctx, outPutString)
}
