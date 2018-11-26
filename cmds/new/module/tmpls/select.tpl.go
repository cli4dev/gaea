package tmpls

//SelectTmpl Select sql 模板
const SelectTmpl = `
//Get{{.name|cname}} 查询单条数据{{.desc}}
const Get{{.name|cname}} = 'select {{range $i,$c:=.selectcolumns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}} where {{range $i,$c:=.pk}}&{{$c.name}} {{end}}'

//Query{{.name|cname}} 查询{{.desc}}列表数据
const Query{{.name|cname}} = 'select {{range $i,$c:=.selectcolumns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}} where {{range $i,$c:=.querycolumns}}&{{$c.name}} {{end}}'
`

//SelectFunc select 函数模板
const SelectFunc = `
//Get 查询单条数据{{.desc}}
func(d *Db{{.name|cname}}) Get({{range $i,$c:=.pk -}}{{$c.name|aname}} {{$c.type|ctype}}{{if $c.end}},{{end}}{{end -}}) (db.QueryRow,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.Get{{.name|cname}}, map[string]interface{}{
		{{range $i,$p:=.pk -}}
		"{{$p.name}}":{{$p.name|aname}},
		{{end -}}
	})
	if err != nil {
		return nil, fmt.Errorf("获取{{.desc}}数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取{{.desc}}列表
func(d *Db{{.name|cname}}) Query(input *Query{{.name|cname}}) (db.QueryRows,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.Query{{.name|cname}}, map[string]interface{}{
		{{range $i,$c:=.querycolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
	})
	if err != nil {
		return nil, fmt.Errorf("获取{{.desc}}数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}
`
