package api

const serviceTmpl = `
package {{.moduleName|spkgName|lName}}

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type {{.moduleName|lName|humpName}}Handler struct {
	container component.IContainer
}

func New{{.moduleName|lName|humpName}}Handler(container component.IContainer) (u *{{.moduleName|lName|humpName}}Handler) {
	return &{{.moduleName|lName|humpName}}Handler{container: container}
}

{{if .restful}}
func (u *{{.moduleName|lName|humpName}}Handler) GetHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *{{.moduleName|lName|humpName}}Handler) PostHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *{{.moduleName|lName|humpName}}Handler) PutHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
func (u *{{.moduleName|lName|humpName}}Handler) DeleteHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
{{else}}
func (u *{{.moduleName|lName|humpName}}Handler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
{{end}}




`
