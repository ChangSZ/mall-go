package jwt

import (
	"errors"
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

// generateToken 根据负责生成JWT的token
func (j *JwtTokenUtil) generateToken(claims jwt.MapClaims) (string, error) {
	claims[CLAIM_KEY_EXP] = time.Now().Add(time.Duration(j.Expiration) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(j.Secret))
}

// GenerateToken 根据用户信息生成JWT的token
func (j *JwtTokenUtil) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		CLAIM_KEY_USERNAME: username,
		CLAIM_KEY_CREATED:  time.Now().Unix(),
	}
	return j.generateToken(claims)
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
		return nil, jwt.ErrTokenSignatureInvalid
	}
}

// ValidateToken 验证token是否还有效
func (j *JwtTokenUtil) ValidateToken(tokenString string, username string) bool {
	user := j.GetUserNameFromToken(tokenString)
	return user == username && !j.isTokenExpired(tokenString)
}

// isTokenExpired 判断token是否已经失效
func (j *JwtTokenUtil) isTokenExpired(tokenString string) bool {
	expiration := j.getExpiredDateFromToken(tokenString)
	return time.Now().After(expiration)
}

// getExpiredDateFromToken 从token中获取过期时间
func (j *JwtTokenUtil) getExpiredDateFromToken(tokenString string) time.Time {
	claims, err := j.GetClaimsFromToken(tokenString)
	if err != nil {
		return time.Time{}
	}
	expUnix, ok := claims[CLAIM_KEY_EXP].(float64)
	if !ok {
		return time.Time{}
	}
	expiration := time.Unix(int64(expUnix), 0)
	return expiration
}

// GetUserNameFromToken 从token中获取登录用户名
func (j *JwtTokenUtil) GetUserNameFromToken(tokenString string) string {
	var username string
	claims, err := j.GetClaimsFromToken(tokenString)
	if err != nil {
		return ""
	}

	username, ok := claims[CLAIM_KEY_USERNAME].(string)
	if !ok {
		return ""
	}
	return username
}

// RefreshHeadToken 根据旧令牌刷新新令牌
func (j *JwtTokenUtil) RefreshHeadToken(oldToken string, interval int) (string, error) {
	if oldToken == "" {
		return oldToken, errors.New("token is empty")
	}

	// 解析旧令牌
	claims, err := j.GetClaimsFromToken(oldToken)
	if err != nil {
		return oldToken, err
	}

	// 检查令牌是否过期
	if j.isTokenExpired(oldToken) {
		return oldToken, jwt.ErrTokenExpired
	}

	// 检查令牌是否在30分钟内刚刷新过
	if j.tokenRefreshJustBefore(oldToken, interval) {
		return oldToken, nil // 返回原令牌
	}

	claims[CLAIM_KEY_CREATED] = time.Now().Unix()
	// 生成新令牌
	newTokenString, err := j.generateToken(claims)
	if err != nil {
		return "", err
	}
	return newTokenString, nil
}

/**
 * 判断token在指定时间内是否刚刚刷新过
 * @param token 原token
 * @param offsetSecond 指定时间（秒）
 */
func (j *JwtTokenUtil) tokenRefreshJustBefore(token string, offsetSecond int) bool {
	claims, err := j.GetClaimsFromToken(token)
	if err != nil {
		return false
	}

	createdUnix, ok := claims[CLAIM_KEY_CREATED].(float64)
	if !ok {
		return false
	}
	created := time.Unix(int64(createdUnix), 0)
	refreshDate := time.Now()
	// 刷新时间在创建时间的指定时间内
	if refreshDate.After(created) && refreshDate.Before(created.Add(time.Duration(offsetSecond)*time.Second)) {
		return true
	}
	return false
}
