package interceptor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/services/ums_user"
	"github.com/ChangSZ/mall-go/pkg/errors"
	"github.com/ChangSZ/mall-go/pkg/jwt"
)

var (
	jwtConfig    = configs.Get().Jwt
	jwtTokenUtil = jwt.NewJwtTokenUtil(jwtConfig.Secret, jwtConfig.Expiration, jwtConfig.TokenHead)
)

func (i *interceptor) CheckToken(ctx core.Context) (userInfo proposal.UmsUserInfo, err core.BusinessError) {
	token := ctx.GetHeader(jwtConfig.TokenHeader)
	if token == "" {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("header中缺少token参数"))
		ctx.AbortWithError(err)
		return
	}
	token = strings.TrimPrefix(token, jwtConfig.TokenHead)
	username := jwtTokenUtil.GetUserNameFromToken(token)
	userDetails, loadErr := ums_user.DefalutService.LoadUserByUsername(ctx, username)
	if loadErr != nil {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(fmt.Errorf("未找到用户: %v, %w", username, loadErr))
		ctx.AbortWithError(err)
		return
	}
	if validateErr := jwtTokenUtil.ValidateToken(token, userDetails.GetUsername()); validateErr != nil {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(validateErr)
		ctx.AbortWithError(err)
		return
	}
	userInfo.Token = token
	userInfo.UserName = userDetails.GetUsername()
	return
}
