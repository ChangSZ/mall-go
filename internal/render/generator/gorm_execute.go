package generator_handler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
)

type gormExecuteRequest struct {
	Tables string `form:"tables"`
}

func (h *handler) GormExecute(ctx *gin.Context) {
	dir, _ := os.Getwd()
	projectPath := strings.Replace(dir, "\\", "/", -1)
	gormgenSh := projectPath + "/scripts/gormgen.sh"
	gormgenBat := projectPath + "/scripts/gormgen.bat"

	req := new(gormExecuteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.ResponseOK(ctx, "参数传递有误")
		return
	}

	mysqlConf := configs.Get().MySQL.Read
	shellPath := fmt.Sprintf("%s %s %s %s %s %s", gormgenSh, mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name, req.Tables)
	batPath := fmt.Sprintf("%s %s %s %s %s %s", gormgenBat, mysqlConf.Addr, mysqlConf.User, mysqlConf.Pass, mysqlConf.Name, req.Tables)

	command := new(exec.Cmd)

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/C", batPath)
	} else {
		// runtime.GOOS = linux or darwin
		command = exec.Command("/bin/bash", "-c", shellPath)
	}

	var stderr bytes.Buffer
	command.Stderr = &stderr

	output, err := command.Output()
	if err != nil {
		api.ResponseOK(ctx, stderr.String())
		return
	}

	api.ResponseOK(ctx, string(output))
}
