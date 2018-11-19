package read

const SelectTmpl = `
//Select{{.name|cname}} 查询{{.desc}}
const Select{{.name|cname}} = 'select {{range $i,$c:=.columns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}}
where
{{range $i,$c:=.pk}}{{$c}}=@{{$c}}{{end}}'
`
const SelectFunc = `
//Select 查询{{.desc}}
func(d *Db{{.name|cname}}) Select{{.name|cname}}(input {{.name|cname}}) (db.QueryRows,error){

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.Select{{.name|cname}}, map[string]interface{}{
		{{range $i,$c:=.pk}}"{{$c}}":input.{{$c|cname}},{{end}}
	})
	if err != nil {
		return nil, fmt.Errorf("获取{{.desc}}数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}
`
