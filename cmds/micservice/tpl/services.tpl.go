package tpl

const ServiceTmpl = `
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

//Handle .
func (u *{{.moduleName|lName|humpName}}Handler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
`

const ServiceFuncTmpl = `
func (u *{{.moduleName|lName|humpName}}Handler) {{.funcName}}Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
`
