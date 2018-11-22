package read

const SelectTmpl = `
//Get{{.name|cname}} 查询单条数据{{.desc}}
const Get{{.name|cname}} = 'select {{range $i,$c:=.selectcolumns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}} where {{range $i,$c:=.pk}}{{$c.name}}=@{{$c.name}}{{end}}'

//Query{{.name|cname}} 查询{{.desc}}列表数据
const Query{{.name|cname}} = 'select {{range $i,$c:=.selectcolumns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}} where 1=1 {{range $i,$c:=.querycolumns}}&{{$c.name}}{{end}}'
`
const SelectFunc = `
//Get 查询单条数据{{.desc}}
func(d *Db{{.name|cname}}) Get(id string) (db.QueryRow,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.Get{{.name|cname}}, map[string]interface{}{
		"{{range $i,$p:=.pk}}{{$p.name}}{{end}}":id,
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
