package install

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/pkg/file"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) View(ctx *gin.Context) {
	type viewResponse struct {
		Config       configs.Config
		MinGoVersion float64
		GoVersion    string
	}

	if _, ok := file.IsExists(configs.ProjectInstallMark); ok {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}

	obj := new(viewResponse)
	obj.Config = configs.Get()
	obj.MinGoVersion = configs.MinGoVersion
	obj.GoVersion = runtime.Version()
	ctx.HTML(200, "install_view.html", obj)
}
