package authorized

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized"
)

type CreateAuthorizedData struct {
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

func (s *service) Create(ctx context.Context, authorizedData *CreateAuthorizedData) (id int64, err error) {
	buf := make([]byte, 10)
	io.ReadFull(rand.Reader, buf)
	secret := hex.EncodeToString(buf)

	model := authorized.NewModel()
	model.BusinessKey = authorizedData.BusinessKey
	model.BusinessSecret = secret
	model.BusinessDeveloper = authorizedData.BusinessDeveloper
	model.Remark = authorizedData.Remark
	model.CreatedUser = core.SessionUserInfo(ctx).UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	return
}
