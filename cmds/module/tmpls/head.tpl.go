package tmpls

//DbHeadTpl Db文件头模板
const DbHeadTpl = `
package %s

import (
	"fmt"
	"%smodules/const/sql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Create{{.name|cname}} 创建{{.desc}} 
type Create{{.name|cname}} struct {		
	{{range $i,$c:=.createcolumns}}
	//{{$c.name|cname}} {{$c.desc|cname}}
	{{$c.name|cname}} {{$c.type|cstype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
	{{end}}	
}

//Update{{.name|cname}} 添加{{.desc}} 
type Update{{.name|cname}} struct {	
	{{range $i,$c:=.pk}}
	//{{$c.name|cname}} {{$c.desc|cname}}
	{{$c.name|cname}} {{$c.type|cstype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
	{{end -}}	
	{{range $i,$c:=.updatecolumns}}
	//{{$c.name|cname}} {{$c.desc|cname}}
	{{$c.name|cname}} {{$c.type|cstype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" valid:"required"'
	{{end}}	
}

//Query{{.name|cname}} 查询{{.desc}} 
type Query{{.name|cname}} struct {		
	{{range $i,$c:=.querycolumns}}
	//{{$c.name|cname}} {{$c.desc|cname}}
	{{$c.name|cname}} {{$c.type|cstype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}"'
	{{end}}
	Pi string 'json:"pi" form:"pi" m2s:"pi" valid:"required"'
	Ps string 'json:"ps" form:"ps" m2s:"ps" valid:"required"'
}

//IDb{{.name|cname}}  {{.desc}}接口
type IDb{{.name|cname}} interface {

	//Create 创建
	Create(input *Create{{.name|cname}}) (error)

	//Get 单条查询
	Get({{range $i,$c:=.pk -}}{{$c.name|aname}} {{$c.type|ctype}}{{if $c.end}},{{end}}{{end -}})(db.QueryRow,error)

	//Query 列表查询
	Query(input *Query{{.name|cname}})  (data db.QueryRows, count int, err error)

	//Update 更新
	Update(input *Update{{.name|cname}}) (err error)

	//Delete 删除
	Delete({{range $i,$c:=.pk -}}{{$c.name|aname}} {{$c.type|ctype}}{{if $c.end}},{{end}}{{end -}}) (err error)
}

//Db{{.name|cname}} {{.desc}}对象
type Db{{.name|cname}} struct {
	c component.IContainer
}

//NewDb{{.name|cname}} 创建{{.desc}}对象
func NewDb{{.name|cname}}(c component.IContainer) *Db{{.name|cname}} {
	return &Db{{.name|cname}}{
		c: c,
	}
}
`
