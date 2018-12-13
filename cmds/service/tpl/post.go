package tpl

//PostTpl .
const PostTpl = `
//PostHandle 添加{{.desc}}数据
func (u *{{.name|cname}}Handler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加{{.desc}}数据--------")
	ctx.Log.Info("1.参数校验")

	var input {{.name|pname}}.Create{{.name|cname}}
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.{{.name|cname}}Lib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
`
