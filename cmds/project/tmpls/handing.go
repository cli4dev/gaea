package tmpls

const handlingTmpl = `package main

import (

	'github.com/micro-plat/hydra/context'
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *{{.projectName|lName}}) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {		
		{{if .jwt}}
		//handling.jwt#//
		jwt, err := ctx.Request.GetJWTConfig() //获取jwt配置
		if err != nil {
			return err
		}
		for _, u := range jwt.Exclude { //排除指定请求
			if u == ctx.Service {
				return nil
			}
		}
		//#handling.jwt//
		{{else}}
		//handling.jwt#//
		//#handling.jwt//
		{{end}}
		return nil
	})
}
`
