package api

const bindTmpl = `
package main
import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
	{{range $i,$m:=.modules}}
	"{{$.projectName}}/services/{{pkgName $m "service"}}"
	{{end}}

)

//AppConf 应用程序配置
type AppConf struct {
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func bindConf(app *hydra.MicroApp) {
}

//bind 检查应用程序配置文件，并根据配置初始化服务
func bind(r *hydra.MicroApp) {
	bindConf(r)
	r.Initializing(func(c component.IContainer) error {

		//获取微信配置
		var conf AppConf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if b, err := govalidator.ValidateStruct(&conf); !b {
			return fmt.Errorf("app 配置文件有误:%v", err)
		}

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}
		{{range $i,$m:=.modules}}
		r.Micro("{{$m}}", {{$m|pkgName|lName}}.New{{$m|lName|humpName}}Handler) //接口,用于添加微信公众号基础参数
		{{end}}
		return nil
	})
}


`