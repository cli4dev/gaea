package tmpls

const handlingTmpl = `package main
{{$empty := "" -}}
import (
	'github.com/micro-plat/hydra/context'
	{{if ne .login $empty -}}
	'{{.projectName}}/modules/member'
	{{- end}}
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *{{.projectName|lName}}) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {		
	{{if .jwt -}}
	//handling.jwt#//
	jwt, err := ctx.Request.GetJWTConfig() 
	if err != nil {
		return err
	}
	for _, u := range jwt.Exclude { 
		if u == ctx.Service {
			return nil
		}
	}
	//#handling.jwt//
	{{- else -}}
	//handling.jwt#//
	//#handling.jwt//
	{{- end}}

	{{if eq .login $empty -}}
	{{- else -}}
	//缓存用户信息
	var m member.LoginState
	if err := ctx.Request.GetJWT(&m); err != nil {
		return context.NewError(context.ERR_FORBIDDEN, err)
	}
	member.Save(ctx, &m)
	{{- end}}
		return nil
	})
}
`
