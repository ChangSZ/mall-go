package core

import (
	"bytes"
	stdctx "context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/pkg/trace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerFunc func(c Context)

type Trace = trace.T

const (
	_Alias           = "_alias_"
	_TraceName       = "_trace_"
	_LoggerName      = "_logger_"
	_BodyName        = "_body_"
	_PayloadName     = "_payload_"
	_SessionUserInfo = "_session_user_info"
	_AbortErrorName  = "_abort_error_"
	_IsRecordMetrics = "_is_record_metrics_"
	_UmsUserInfo     = "_ums_user_info_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(context)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*context)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*context)
	c.ctx = nil
	contextPool.Put(c)
}

var _ Context = (*context)(nil)

type Context interface {
	init()

	// ShouldBindQuery 反序列化 querystring
	// tag: `form:"xxx"` (注：不要写成 query)
	ShouldBindQuery(obj interface{}) error

	// ShouldBindPostForm 反序列化 postform (querystring会被忽略)
	// tag: `form:"xxx"`
	ShouldBindPostForm(obj interface{}) error

	// ShouldBindForm 同时反序列化 querystring 和 postform;
	// 当 querystring 和 postform 存在相同字段时，postform 优先使用。
	// tag: `form:"xxx"`
	ShouldBindForm(obj interface{}) error

	ShouldBind(obj interface{}) error

	// ShouldBindJSON 反序列化 postjson
	// tag: `json:"xxx"`
	ShouldBindJSON(obj interface{}) error

	// ShouldBindURI 反序列化 path 参数(如路由路径为 /user/:name)
	// tag: `uri:"xxx"`
	ShouldBindURI(obj interface{}) error

	// Redirect 重定向
	Redirect(code int, location string)

	// Trace 获取 Trace 对象
	Trace() Trace
	setTrace(trace Trace)
	disableTrace()

	// Logger 获取 Logger 对象
	Logger() *zap.Logger
	setLogger(logger *zap.Logger)

	// Payload 正确返回
	Payload(payload interface{})
	getPayload() interface{}

	// HTML 返回界面
	HTML(name string, obj interface{})

	// AbortWithError 错误返回
	AbortWithError(err BusinessError)
	abortError() BusinessError

	// Header 获取 Header 对象
	Header() http.Header
	// GetHeader 获取 Header
	GetHeader(key string) string
	// SetHeader 设置 Header
	SetHeader(key, value string)

	// SessionUserInfo 当前用户信息
	SessionUserInfo() proposal.SessionUserInfo
	setSessionUserInfo(info proposal.SessionUserInfo)

	// GetUmsUserInfo ums用户信息
	GetUmsUserInfo() proposal.UmsUserInfo
	setUmsUserInfo(proposal.UmsUserInfo)

	// Alias 设置路由别名 for metrics path
	Alias() string
	setAlias(path string)

	// disableRecordMetrics 设置禁止记录指标
	disableRecordMetrics()
	ableRecordMetrics()
	isRecordMetrics() bool

	// RequestInputParams 获取所有参数
	RequestInputParams() url.Values
	// RequestPostFormParams  获取 PostForm 参数
	RequestPostFormParams() url.Values
	// Request 获取 Request 对象
	Request() *http.Request
	// RawData 获取 Request.Body
	RawData() []byte
	// Method 获取 Request.Method
	Method() string
	// Host 获取 Request.Host
	Host() string
	// Path 获取 请求的路径 Request.URL.Path (不附带 querystring)
	Path() string
	// URI 获取 unescape 后的 Request.URL.RequestURI()
	URI() string
	// RequestContext 获取请求的 context (当 client 关闭后，会自动 canceled)
	RequestContext() StdContext

	// ResponseWriter 获取 ResponseWriter 对象
	ResponseWriter() gin.ResponseWriter
}

type context struct {
	gin.Context
}

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
}

func (c *context) init() {
	body, err := c.GetRawData()
	if err != nil {
		panic(err)
	}

	c.Set(_BodyName, body)                               // cache body是为了trace使用
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // re-construct req body
}

// Redirect 重定向
func (c *context) redirect(code int, location string) {
	c.Redirect(code, location)
}

func (c *context) Trace() Trace {
	t, ok := c.Get(_TraceName)
	if !ok || t == nil {
		return nil
	}

	return t.(Trace)
}

func (c *context) setTrace(trace Trace) {
	c.Set(_TraceName, trace)
}

func (c *context) disableTrace() {
	c.setTrace(nil)
}

func (c *context) Logger() *zap.Logger {
	logger, ok := c.Get(_LoggerName)
	if !ok {
		return nil
	}

	return logger.(*zap.Logger)
}

func (c *context) setLogger(logger *zap.Logger) {
	c.Set(_LoggerName, logger)
}

func (c *context) getPayload() interface{} {
	if payload, ok := c.Get(_PayloadName); ok {
		return payload
	}
	return nil
}

func (c *context) Payload(payload interface{}) {
	c.Set(_PayloadName, payload)
}

func (c *context) HTML(name string, obj interface{}) {
	c.HTML(200, name+".html", obj)
}

func (c *context) Header() http.Header {
	header := c.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
}

func (c *context) GetHeader(key string) string {
	return c.GetHeader(key)
}

func (c *context) SetHeader(key, value string) {
	c.Header(key, value)
}

func (c *context) SessionUserInfo() proposal.SessionUserInfo {
	val, ok := c.Get(_SessionUserInfo)
	if !ok {
		return proposal.SessionUserInfo{}
	}

	return val.(proposal.SessionUserInfo)
}

func (c *context) setSessionUserInfo(info proposal.SessionUserInfo) {
	c.Set(_SessionUserInfo, info)
}

func (c *context) GetUmsUserInfo() proposal.UmsUserInfo {
	val, ok := c.Get(_UmsUserInfo)
	if !ok {
		return proposal.UmsUserInfo{}
	}
	return val.(proposal.UmsUserInfo)
}

func (c *context) setUmsUserInfo(info proposal.UmsUserInfo) {
	c.Set(_UmsUserInfo, info)
}

func (c *context) AbortWithError(err BusinessError) {
	if err != nil {
		httpCode := err.HTTPCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		c.AbortWithStatus(httpCode)
		c.Set(_AbortErrorName, err)
	}
}

func (c *context) abortError() BusinessError {
	err, _ := c.Get(_AbortErrorName)
	return err.(BusinessError)
}

func (c *context) Alias() string {
	path, ok := c.Get(_Alias)
	if !ok {
		return ""
	}

	return path.(string)
}

func (c *context) setAlias(path string) {
	if path = strings.TrimSpace(path); path != "" {
		c.Set(_Alias, path)
	}
}

func (c *context) isRecordMetrics() bool {
	isRecordMetrics, ok := c.Get(_IsRecordMetrics)
	if !ok {
		return false
	}

	return isRecordMetrics.(bool)
}

func (c *context) ableRecordMetrics() {
	c.Set(_IsRecordMetrics, true)
}

func (c *context) disableRecordMetrics() {
	c.Set(_IsRecordMetrics, false)
}

// RequestInputParams 获取所有参数
func (c *context) RequestInputParams() url.Values {
	_ = c.Request.ParseForm()
	return c.Request.Form
}

// RequestPostFormParams 获取 PostForm 参数
func (c *context) RequestPostFormParams() url.Values {
	_ = c.Request.ParseForm()
	return c.Request.PostForm
}

func (c *context) RawData() []byte {
	body, ok := c.Get(_BodyName)
	if !ok {
		return nil
	}

	return body.([]byte)
}

// Method 请求的method
func (c *context) Method() string {
	return c.Request.Method
}

// Host 请求的host
func (c *context) Host() string {
	return c.Request.Host
}

// Path 请求的路径(不附带querystring)
func (c *context) Path() string {
	return c.Request.URL.Path
}

// URI unescape后的uri
func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.Request.URL.RequestURI())
	return uri
}

// RequestContext (包装 Trace + Logger) 获取请求的 context (当client关闭后，会自动canceled)
func (c *context) RequestContext() StdContext {
	return StdContext{
		//c.Request.Context(),
		stdctx.Background(),
		c.Trace(),
		c.Logger(),
	}
}

// ResponseWriter 获取 ResponseWriter
func (c *context) ResponseWriter() gin.ResponseWriter {
	return c.Writer
}
