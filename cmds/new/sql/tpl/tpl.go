package tpl

//DbTpl .
const DbTpl = `
package %s
import (
	"fmt"
	"github.com/micro-plat/%s/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/hydra/component"

)

//Create{{.name|cname}} 创建{{.desc}} 
type Create{{.name|cname}} struct {		
		{{range $i,$c:=.createcolumns}}
		//{{$c.name|cname}} {{$c.desc|cname}}
		{{$c.name|cname}} {{$c.type|ctype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
		{{end}}	
}

//Update{{.name|cname}} 添加{{.desc}} 
type Update{{.name|cname}} struct {		
		{{range $i,$c:=.updatecolumns}}
		//{{$c.name|cname}} {{$c.desc|cname}}
		{{$c.name|cname}} {{$c.type|ctype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
		{{end}}	
}

//Query{{.name|cname}} 查询{{.desc}} 
type Query{{.name|cname}} struct {		
		{{range $i,$c:=.querycolumns}}
		//{{$c.name|cname}} {{$c.desc|cname}}
		{{$c.name|cname}} {{$c.type|ctype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
		{{end}}	
}
//IDb{{.name|cname}} .
type IDb{{.name|cname}} interface {
	//创建
	Create(input *Create{{.name|cname}}) (error)
	//单条查询
	Get(id string)(db.QueryRow,error)
	//列表查询
	Query(input *Query{{.name|cname}}) (db.QueryRows,error)
	//更新
	Update(input *Update{{.name|cname}}) (err error)
	//删除
	Delete(id string) (err error)
}
//Db{{.name|cname}} .
type Db{{.name|cname}} struct {
	c component.IContainer
}
//NewDb{{.name|cname}} .
func NewDb{{.name|cname}}(c component.IContainer) *Db{{.name|cname}} {
	return &Db{{.name|cname}}{
		c: c,
	}
}
`
