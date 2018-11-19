package tpl

const DbTpl = `
package %s
import (
	"fmt"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/hydra/component"

)

//{{.name|cname}} {{.desc}} 
type {{.name|cname}} struct {		
		{{range $i,$c:=.columnsOriginal}}
		//{{$c.name|cname}} {{$c.desc|cname}}
		{{$c.name|cname}} {{$c.type|ctype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}"'
		{{end}}	
}

type IDb{{.name|cname}} interface {
	Insert{{.name|cname}}(input {{.name|cname}}) (error)
	Select{{.name|cname}}(input {{.name|cname}}) (db.QueryRows,error)
	Update{{.name|cname}}(input {{.name|cname}}) (err error)
	Delete{{.name|cname}}(input {{.name|cname}}) (err error)
}

type Db{{.name|cname}} struct {
	c component.IContainer
}

func NewDb{{.name|cname}}(c component.IContainer) *Db{{.name|cname}} {
	return &Db{{.name|cname}}{
		c: c,
	}
}
`
