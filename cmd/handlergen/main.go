package main

import (
	"flag"
	"fmt"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

var handlerName string

func init() {
	handler := flag.String("handler", "", "请输入需要生成的 handler 名称\n")
	flag.Parse()

	handlerName = strings.ToLower(*handler)
}

func getLastSegment(s string) string {
	index := strings.LastIndex(s, "/")
	if index == -1 {
		return s
	}
	return s[index+1:]
}

func main() {
	fs := token.NewFileSet()
	filePath := fmt.Sprintf("./internal/api/%s", handlerName)
	parsedFile, err := decorator.ParseFile(fs, filePath+"/handler.go", nil, 0)
	if err != nil {
		log.Fatalf("parsing package: %s: %s\n", filePath, err)
	}

	files, _ := os.ReadDir(filePath)
	if len(files) == 0 {
		log.Fatalf("请先确保 %s 目录中，至少含有 handler.go 文件。", filePath)
	}
	existFilesMap := make(map[string]bool) // 已存在的文件将会跳过
	for _, v := range files {
		existFilesMap[v.Name()] = true
	}

	dst.Inspect(parsedFile, func(n dst.Node) bool {
		decl, ok := n.(*dst.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}

		for _, spec := range decl.Specs {
			typeSpec, _ok := spec.(*dst.TypeSpec)
			if !_ok {
				continue
			}

			var interfaceType *dst.InterfaceType
			if interfaceType, ok = typeSpec.Type.(*dst.InterfaceType); !ok {
				continue
			}

			for _, v := range interfaceType.Methods.List {
				if len(v.Names) > 0 {
					if v.Names[0].String() == "i" {
						continue
					}

					filepath := "./internal/api/" + handlerName
					name := fmt.Sprintf("func_%s.go", strings.ToLower(v.Names[0].String()))
					filename := fmt.Sprintf("%s/%s", filepath, name)
					if _, ok := existFilesMap[name]; ok { // func文件已存在，直接跳过即可
						continue
					}
					funcFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
					if err != nil {
						fmt.Printf("create and open func file error %v\n", err.Error())
						continue
					}

					if funcFile == nil {
						fmt.Printf("func file is nil \n")
						continue
					}

					fmt.Println("  └── file : ", filename)

					funcContent := fmt.Sprintf("package %s\n\n", getLastSegment(handlerName))
					funcContent += "import (\n\"github.com/ChangSZ/mall-go/internal/api\"\n\n"
					funcContent += `"github.com/gin-gonic/gin"`
					funcContent += "\n)\n"
					funcContent += fmt.Sprintf("\n\ntype %sRequest struct {}\n\n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("type %sResponse struct {}\n\n", Lcfirst(v.Names[0].String()))

					// 首行注释
					funcContent += fmt.Sprintf("%s\n", v.Decorations().Start.All()[0])

					nameArr := strings.Split(v.Decorations().Start.All()[0], v.Names[0].String())
					funcContent += fmt.Sprintf("// @Summary%s \n", nameArr[1])
					funcContent += fmt.Sprintf("// @Description%s \n", nameArr[1])
					// Tags
					funcContent += fmt.Sprintf("%s \n", v.Decorations().Start.All()[1])
					funcContent += "// @Accept application/x-www-form-urlencoded \n"
					funcContent += "// @Produce json \n"
					funcContent += fmt.Sprintf("// @Param Request body %sRequest true \"请求信息\" \n", Lcfirst(v.Names[0].String()))
					funcContent += fmt.Sprintf("// @Success 200 {object} code.Success{data=%sResponse} \n", Lcfirst(v.Names[0].String()))
					funcContent += "// @Failure 400 {object} code.Failure \n"
					// Router
					funcContent += fmt.Sprintf("%s \n", v.Decorations().Start.All()[2])
					funcContent += fmt.Sprintf("func (h *handler) %s(ctx *gin.Context) { \n api.Success(ctx, nil) \n }", v.Names[0].String())

					funcFile.WriteString(funcContent)
					funcFile.Close()
				}
			}
		}
		return true
	})
	genServices(handlerName)
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// genServices 生成services目录及文件
func genServices(svcName string) {
	svcDirPath := "./internal/services/" + svcName
	pkgName := getLastSegment(svcName)

	// 检查文件夹是否存在
	if _, err := os.Stat(svcDirPath); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(svcDirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory: %v\n", err)
			return
		}
	}

	// 检查service.go文件是否存在
	serviceFilePath := filepath.Join(svcDirPath, "service.go")
	if _, err := os.Stat(serviceFilePath); os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err := os.Create(serviceFilePath)
		if err != nil {
			fmt.Printf("Failed to create service.go: %v\n", err)
			return
		}
		content := fmt.Sprintf("package %s\n\n", pkgName)
		content += "type service struct{}\n\n"
		content += "func New() Service {\nreturn &service{}\n}\n\n"
		content += "func (s *service) i() {}"

		file.WriteString(content)
		file.Close()
	}

	// 检查interface.go文件是否存在
	interfaceFilePath := filepath.Join(svcDirPath, "interface.go")
	if _, err := os.Stat(interfaceFilePath); os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err := os.Create(interfaceFilePath)
		if err != nil {
			fmt.Printf("Failed to create interface.go: %v\n", err)
			return
		}
		content := fmt.Sprintf("package %s\n\n", pkgName)
		content += "var _ Service = (*service)(nil)\n\n"
		content += "type Service interface {\ni()\n}"

		file.WriteString(content)
		file.Close()
	}
}
