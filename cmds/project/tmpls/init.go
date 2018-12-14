package tmpls

const initTmpl = `
{{$empty := "" -}}
package main
import (
	{{if .appconf -}}
	"github.com/asaskevich/govalidator"
	{{- else -}}
	{{- end}}
	"github.com/micro-plat/hydra/component"
	{{range $i,$m:=.pkgs}}"{{$.projectName}}/{{$m}}"
	{{end}}

	{{if ne .login $empty -}}
	"fmt"
	"github.com/asaskevich/govalidator"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/services/member"
	{{- end -}}
)

{{if .appconf -}}
//appconf.struct#//
//AppConf 应用程序配置
type AppConf struct {
}
//#appconf.struct//
{{- else -}}
	//appconf.struct#//
//#appconf.struct//
{{- end}}


//init 检查应用程序配置文件，并根据配置初始化服务
func (r *{{.projectName|lName}}) init() {
	r.Initializing(func(c component.IContainer) error {
	{{if .appconf -}}
	//appconf.func#//
	var conf app.Conf
	if err := c.GetAppConf(&conf); err != nil {
		return err
	}
	if b, err := govalidator.ValidateStruct(&conf); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	app.SaveConf(c, &conf)
	//appconf.func#//
	{{- else -}}
	{{if eq .login $empty -}}
	//appconf.func#//
	//#appconf.func//
	{{- else -}}
	//appconf.func#//
	var conf app.Conf
	if err := c.GetAppConf(&conf); err != nil {
		return err
	}
	if b, err := govalidator.ValidateStruct(&conf); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	app.SaveConf(c, &conf)
	//#appconf.func//
	{{- end -}}
	{{- end}}

	{{if eq .db $empty -}}
	//db.init#//
	//#db.init//
	{{- else -}}
	//db.init#//
	//检查db配置是否正确
	if _, err := c.GetDB(); err != nil {
		return err
	}
	//#db.init//
	{{- end}}

	{{if .cahce -}}
	//cache.init#//
	//检查cache配置是否正确
	if _, err := c.GetCache(); err != nil {
		return err
	}
	//#cache.init//
	{{- else -}}
	//cache.init#//
	//#cache.init//
	{{- end}}

	{{if .queue -}}
	//queue.init#//
	//检查queue配置是否正确
	if _, err := c.GetQueue(); err != nil {
		return err
	}
	//#queue.init//
	{{- else -}}
	//queue.init#//
	//#queue.init//
	{{- end}}

	
	{{if eq .login $empty -}}
	//login.router#//
	//#login.router//
	{{- else -}}
	//login.router#//
	r.Micro("/member/login", member.NewLoginHandler, "*")     //登录系统
	r.Micro("/member/update", member.NewUpdateHandler, "*")   //修改密码
	r.Micro("/member/menu/get", member.NewMenuHandler, "*")   //获取菜单
	r.Micro("/member/user", member.NewUserHandler, "*")  //获取用户
	r.Micro("/member/getsysinfo", member.NewInfoHandler, "*") //获取系统信息
	//#login.router//
	{{- end}}

	//service.router#//
	//#service.router//

	return nil
	})
}`

//APPConfStruct .
const APPConfStruct = `//appconf.struct#//
	//AppConf 应用程序配置
	type AppConf struct {
	}
	//#appconf.struct//`

//APPConfFunc .
const APPConfFunc = `//appconf.func#//
	//获取配置
	var conf AppConf
	if err := c.GetAppConf(&conf); err != nil {
		return err
	}
	if b, err := govalidator.ValidateStruct(&conf); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	//#appconf.func//`

//DBInit .
const DBInit = `
	//db.init#//
	//检查db配置是否正确
	if _, err := c.GetDB(); err != nil {
	return err
	}
	//#db.init//`

//CacheInit .
const CacheInit = `//cache.init#//
	//检查cache配置是否正确
	if _, err := c.GetCache(); err != nil {
		return err
	}
	//#cache.init//`

//QueueInit .
const QueueInit = `//queue.init#//
	//检查queue配置是否正确
	if _, err := c.GetQueue(); err != nil {
		return err
	}
	//#queue.init//`
