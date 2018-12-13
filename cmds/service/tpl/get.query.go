package tpl

//GetAndQueryTpl GetHandle
const GetAndQueryTpl = `
//GetHandle 获取{{.desc}}单条数据
func (u *{{.name|cname}}Handler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取{{.desc}}单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check({{range $i,$c:=.pk -}}"{{$c.name}}"{{if $c.end}},{{end}}{{end -}}); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.{{.name|cname}}Lib.Get({{range $i,$c:=.pk -}}ctx.Request.Get{{$c.type|ctype|cname}}("{{$c.name}}"){{if $c.end}},{{end}}{{end -}})
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取{{.desc}}数据列表
func (u *{{.name|cname}}Handler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取{{.desc}}数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input {{.name|pname}}.Query{{.name|cname}}
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.{{.name|cname}}Lib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
`
