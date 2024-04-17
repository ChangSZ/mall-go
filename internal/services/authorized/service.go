package authorized

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized_api"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx context.Context, authorizedData *CreateAuthorizedData) (id int64, err error)
	List(ctx context.Context, searchData *SearchData) (listData []*authorized.Authorized, err error)
	PageList(ctx context.Context, searchData *SearchData) (listData []*authorized.Authorized, err error)
	PageListCount(ctx context.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx context.Context, id int64, used int32) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Detail(ctx context.Context, id int64) (info *authorized.Authorized, err error)
	DetailByKey(ctx context.Context, key string) (data *CacheAuthorizedData, err error)

	CreateAPI(ctx context.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int64, err error)
	ListAPI(ctx context.Context, searchAPIData *SearchAPIData) (listData []*authorized_api.AuthorizedApi, err error)
	DeleteAPI(ctx context.Context, id int64) (err error)
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
