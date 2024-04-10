package tool

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/pkg/hash"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// HashIdsEncode HashIds 加密
	// @Tags API.tool
	// @Router /api/tool/hashids/encode/{id} [get]
	HashIdsEncode(*gin.Context)

	// HashIdsDecode HashIds 解密
	// @Tags API.tool
	// @Router /api/tool/hashids/decode/{id} [get]
	HashIdsDecode(*gin.Context)

	// SearchCache 查询缓存
	// @Tags API.tool
	// @Router /api/tool/cache/search [post]
	SearchCache(*gin.Context)

	// ClearCache 清空缓存
	// @Tags API.tool
	// @Router /api/tool/cache/clear [patch]
	ClearCache(*gin.Context)

	// Dbs 查询 DB
	// @Tags API.tool
	// @Router /api/tool/data/dbs [get]
	Dbs(*gin.Context)

	// Tables 查询 Table
	// @Tags API.tool
	// @Router /api/tool/data/tables [post]
	Tables(*gin.Context)

	// SearchMySQL 执行 SQL 语句
	// @Tags API.tool
	// @Router /api/tool/data/mysql [post]
	SearchMySQL(*gin.Context)

	// SendMessage 发送消息
	// @Tags API.tool
	// @Router /api/tool/send_message [post]
	SendMessage(*gin.Context)
}

type handler struct {
	hashids hash.Hash
}

func New() Handler {
	return &handler{
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
