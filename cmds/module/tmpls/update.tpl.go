package tmpls

//UpdateTmpl update sql 模板
const UpdateTmpl = `
//Update{{.name|cname}} 更新{{.desc}}
const Update{{.name|cname}} = 'update {{.name}} set
{{range $i,$c:=.updatecolumns}}{{$c.name}}=@{{$c.name}}{{if $c.end}},{{end}}{{end}}
where 1=1 {{range $i,$c:=.pk}}&{{$c.name}} {{end}}'
`

//UpdateFunc update 函数模板
const UpdateFunc = `
//Update 更新{{.desc}}
func(d *Db{{.name|cname}}) Update(input *Update{{.name|cname}}) error {

	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.Update{{.name|cname}}, map[string]interface{}{
		{{range $i,$c:=.pk -}}
		"{{$c.name}}": input.{{$c.name|cname}},
		{{end -}}
		{{range $i,$c:=.updatecolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
	})
	if err != nil {
		return fmt.Errorf("更新{{.desc}}数据发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}
`
