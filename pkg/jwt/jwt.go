package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
 * JwtToken生成的工具类
 * JWT token的格式：header.payload.signature
 * header的格式（算法、token的类型）：
 * {"alg": "HS512","typ": "JWT"}
 * payload的格式（用户名、创建时间、生成时间）：
 * {"sub":"wang","created":1489079981393,"exp":1489684781}
 * signature的生成算法：
 * HMACSHA512(base64UrlEncode(header) + "." +base64UrlEncode(payload),secret)
 */

const (
	CLAIM_KEY_USERNAME = "sub"
	CLAIM_KEY_CREATED  = "created"
	CLAIM_KEY_EXP      = "exp"
)

type JwtTokenUtil struct {
	Secret     string
	Expiration int64
	TokenHead  string
}

func NewJwtTokenUtil(secret string, expiration int64, tokenHead string) *JwtTokenUtil {
	return &JwtTokenUtil{
		Secret:     secret,
		Expiration: expiration,
		TokenHead:  tokenHead,
	}
}

// GenerateToken 根据用户信息生成JWT的token
func (j *JwtTokenUtil) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		CLAIM_KEY_USERNAME: username,
		CLAIM_KEY_CREATED:  time.Now().Unix(),
		CLAIM_KEY_EXP:      time.Now().Add(time.Duration(j.Expiration) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(j.Secret))
}

// GetClaimsFromToken 从token中获取JWT中的负载
func (j *JwtTokenUtil) GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("Token is invalid", jwt.ValidationErrorSignatureInvalid)
	}
}

// ValidateToken 验证token是否还有效
func (j *JwtTokenUtil) ValidateToken(tokenString string, username string) (bool, error) {
	claims, err := j.GetClaimsFromToken(tokenString)
	if err != nil {
		return false, err
	}

	user := claims[CLAIM_KEY_USERNAME].(string)
	expiration := time.Unix(int64(claims[CLAIM_KEY_EXP].(float64)), 0)

	return user == username && time.Now().Before(expiration), nil
}
