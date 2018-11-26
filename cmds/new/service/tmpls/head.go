package tmpls

//HeadTpl .
const HeadTpl = `
package %s

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"%s"
)

//{{.name|cname}}Handler {{.desc}}接口
type {{.name|cname}}Handler struct {
	container component.IContainer
	{{.name|cname}}Lib   {{.name|pname}}.IDb{{.name|cname}}
}

//New{{.name|cname}}Handler 创建{{.desc}}对象
func New{{.name|cname}}Handler(container component.IContainer) (u *{{.name|cname}}Handler) {
	return &{{.name|cname}}Handler{
		container: container,
		{{.name|cname}}Lib:   {{.name|pname}}.NewDb{{.name|cname}}(container),
	}
}

`
