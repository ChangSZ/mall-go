package generator_handler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
)

type handlerExecuteRequest struct {
	Name string `form:"name"`
}

func (h *handler) HandlerExecute(ctx *gin.Context) {
	dir, _ := os.Getwd()
	projectPath := strings.Replace(dir, "\\", "/", -1)
	handlergenSh := projectPath + "/scripts/handlergen.sh"
	handlergenBat := projectPath + "/scripts/handlergen.bat"

	req := new(handlerExecuteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.ResponseOK(ctx, "参数传递有误")
		return
	}
	shellPath := fmt.Sprintf("%s %s", handlergenSh, req.Name)
	batPath := fmt.Sprintf("%s %s", handlergenBat, req.Name)

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
