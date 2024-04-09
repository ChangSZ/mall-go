package authorized

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized_api"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int32, err error)
	List(ctx core.Context, searchData *SearchData) (listData []*authorized.Authorized, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*authorized.Authorized, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, id int32) (info *authorized.Authorized, err error)
	DetailByKey(ctx core.Context, key string) (data *CacheAuthorizedData, err error)

	CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error)
	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*authorized_api.AuthorizedApi, err error)
	DeleteAPI(ctx core.Context, id int32) (err error)
}

type service struct {
	db mysql.Repo
}

func New(db mysql.Repo) Service {
	return &service{
		db: db,
	}
}

func (s *service) i() {}
