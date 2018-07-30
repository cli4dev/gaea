package tmpls

const initTmpl = `
package main
import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	{{range $i,$m:=.pkgs}}"{{$.projectName}}/{{$m}}"
	{{end}}
)

//AppConf 应用程序配置
type AppConf struct {
}
//init 检查应用程序配置文件，并根据配置初始化服务
func (r *{{.projectName|lName}}) init() {
	r.Initializing(func(c component.IContainer) error {
	//获取配置
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

		{{range $i,$m:=.modules}}{{range $x,$s:=$.rss}}r.{{$s}}("{{$m}}", {{$m|spkgName|lName}}.New{{$m|lName|humpName}}Handler)
		{{end}}
		{{end}}
		return nil
	})
}


`
