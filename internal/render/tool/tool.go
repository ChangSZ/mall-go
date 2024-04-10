package tool

import (
	"encoding/json"
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/pkg/file"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/gin-gonic/gin"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Cache(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool_cache.html", nil)
}

func (h *handler) Data(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool_data.html", nil)
}

func (h *handler) HashIds(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool_hashids.html", configs.Get())
}

func (h *handler) Websocket(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool_websocket.html", nil)
}

func (h *handler) Log(ctx *gin.Context) {
	type logData struct {
		Level       string  `json:"level"`
		Time        string  `json:"time"`
		Path        string  `json:"path"`
		HTTPCode    int     `json:"http_code"`
		Method      string  `json:"method"`
		Msg         string  `json:"msg"`
		TraceID     string  `json:"trace_id"`
		Content     string  `json:"content"`
		CostSeconds float64 `json:"cost_seconds"`
	}

	type logsViewResponse struct {
		Logs []logData `json:"logs"`
	}

	type logParseData struct {
		Level        string  `json:"level"`
		Time         string  `json:"time"`
		Caller       string  `json:"caller"`
		Msg          string  `json:"msg"`
		Domain       string  `json:"domain"`
		Method       string  `json:"method"`
		Path         string  `json:"path"`
		HTTPCode     int     `json:"http_code"`
		BusinessCode int     `json:"business_code"`
		Success      bool    `json:"success"`
		CostSeconds  float64 `json:"cost_seconds"`
		TraceID      string  `json:"trace_id"`
	}

	readLineFromEnd, err := file.NewReadLineFromEnd(configs.ProjectLogFile)
	if err != nil {
		log.WithTrace(ctx).Error("NewReadLineFromEnd err: ", err)
	}

	logSize := 100

	obj := new(logsViewResponse)
	obj.Logs = make([]logData, logSize)

	for i := 0; i < logSize; i++ {
		content, _ := readLineFromEnd.ReadLine()
		if string(content) != "" {
			var logParse logParseData
			err = json.Unmarshal(content, &logParse)
			if err != nil {
				log.WithTrace(ctx).Error("NewReadLineFromEnd json Unmarshal err: ", err)
			}

			data := logData{
				Content:     string(content),
				Level:       logParse.Level,
				Time:        logParse.Time,
				Path:        logParse.Path,
				Method:      logParse.Method,
				Msg:         logParse.Msg,
				HTTPCode:    logParse.HTTPCode,
				TraceID:     logParse.TraceID,
				CostSeconds: logParse.CostSeconds,
			}

			obj.Logs[i] = data
		}
	}
	ctx.HTML(http.StatusOK, "tool_logs.html", obj)
}
