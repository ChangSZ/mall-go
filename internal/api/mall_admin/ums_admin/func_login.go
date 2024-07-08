package ums_admin

import (
	"net"
	"strings"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}

// Login 登录以后返回token
// @Summary 登录以后返回token
// @Description 登录以后返回token
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} code.Success{data=loginResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/login [post]
func (h *handler) Login(ctx *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	token, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Token = token
	res.TokenHead = configs.Get().Jwt.TokenHead
	h.service.InsertLoginLog(ctx, req.Username, getRequestIp(ctx))
	api.Success(ctx, res)
}

// getRequestIp 获取请求的真实IP地址
func getRequestIp(c *gin.Context) string {
	// 尝试从X-Forwarded-For头部获取IP
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" || strings.ToLower(ip) == "unknown" {
		// 尝试从Proxy-Client-IP头部获取IP
		ip = c.GetHeader("Proxy-Client-IP")
	}
	if ip == "" || strings.ToLower(ip) == "unknown" {
		// 尝试从WL-Proxy-Client-IP头部获取IP
		ip = c.GetHeader("WL-Proxy-Client-IP")
	}
	if ip == "" || strings.ToLower(ip) == "unknown" {
		// 从远程地址获取IP
		ip = c.ClientIP()
		// 如果是本地地址，则获取本地IP
		if ip == "127.0.0.1" || ip == "::1" {
			if localIp, err := getLocalIp(); err == nil {
				ip = localIp
			}
		}
	}

	// 如果通过多个代理转发，获取第一个IP
	if ip != "" && len(ip) > 15 {
		if index := strings.Index(ip, ","); index > 0 {
			ip = ip[:index]
		}
	}
	return ip
}

// getLocalIp 获取本地IP地址
func getLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", nil
}
