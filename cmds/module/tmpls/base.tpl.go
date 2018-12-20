package tmpls

//BaseTpl .
const BaseTpl = `package %s
import "github.com/micro-plat/hydra/component"
type I%s interface {
}
type %s struct {
	c component.IContainer
}
func New%s(c component.IContainer) *%s{
	return &%s{
		c: c,
	}
}
`

const DicTpl = `package sql
{{$empty := "" -}}
{{if ne .di $empty -}}
//Get{{.name|cname}}Dictionary  获取数据字典
const Get{{.name|cname}}Dictionary = 'select {{.di}} as id,{{.dn}} as name from {{.name}}'
{{- end}}
`
