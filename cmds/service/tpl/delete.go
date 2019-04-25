package tpl

//DeleteTpl
const DeleteTpl = `
//DeleteHandle 删除{{.desc}}数据
func (u *{{.name|cname}}Handler) DeleteHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------删除{{.desc}}数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check({{range $i,$c:=.pk -}}"{{$c.name}}"{{if $c.end}},{{end}}{{end -}}); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.{{.name|cname}}Lib.Delete({{range $i,$c:=.pk -}}ctx.Request.Get{{$c.type|ctype|cname}}("{{$c.name}}"){{if $c.end}},{{end}}{{end -}})
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
`
