package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwtTokenUtil(t *testing.T) {
	tests := []struct {
		name     string
		username string
	}{
		{
			name:     "test1",
			username: "jack",
		},
	}
	var jwtUtil = NewJwtTokenUtil("123456", 2, "Bearer ")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := jwtUtil.GenerateToken(tt.username)
			assert.Equal(t, nil, err)

			username := jwtUtil.GetUserNameFromToken(token)
			assert.Equal(t, tt.username, username, "应该解析出用户名")

			expiration := jwtUtil.getExpiredDateFromToken(token)
			assert.NotEqual(t, time.Time{}, expiration, "应该解析出过期时间")

			err = jwtUtil.ValidateToken(token, "nick")
			assert.NotNil(t, err)

			err = jwtUtil.ValidateToken(token, "jack")
			assert.Nil(t, err)

			newToken, err := jwtUtil.RefreshHeadToken(token, 1)
			assert.Equal(t, nil, err)
			assert.Equal(t, token, newToken, "无需刷新, 返回的应该是原始token")

			time.Sleep(time.Second * 1)
			newToken, err = jwtUtil.RefreshHeadToken(token, 1)
			assert.Equal(t, nil, err)
			assert.NotEqual(t, token, newToken, "token应该刷新")

			// token过期失效
			time.Sleep(time.Second * 2)
			err = jwtUtil.ValidateToken(token, "jack")
			assert.NotNil(t, err, "token应该过期")
		})
	}
}
