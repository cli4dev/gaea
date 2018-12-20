package tpl

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
{{$empty := "" -}}

{{if ne .di $empty -}}
//GetDictionary 获取数据字典 
func (u *{{.name|cname}}Handler) GetDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取{{.desc}}数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.{{.name|cname}}Lib.Get{{.name|cname}}Dictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}
{{- end}}
`
