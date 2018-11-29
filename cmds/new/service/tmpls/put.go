package tmpls

//PutTpl .
const PutTpl = `
//PutHandle 更新{{.desc}}数据
func (u *{{.name|cname}}Handler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新{{.desc}}数据--------")
	ctx.Log.Info("1.参数校验")

	var input {{.name|pname}}.Update{{.name|cname}}
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.{{.name|cname}}Lib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
`
