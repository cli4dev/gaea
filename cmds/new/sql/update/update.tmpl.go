package update

const UpdateTmpl = `
//Update{{.name|cname}} 更新{{.desc}}
const Update{{.name|cname}} = 'update {{.name}} set
{{range $i,$c:=.columns}}{{$c.name}}=@{{$c.name}}{{if $c.end}},{{end}}{{end}}
where
{{range $i,$c:=.pk}}{{$c}}=@{{$c}}{{end}}'
`
const UpdateFunc = `
//Update 更新{{.desc}}
func(d *Db{{.name|cname}}) Update{{.name|cname}}(input {{.name|cname}}) error {

	db := d.c.GetRegularDB()
	params := map[string]interface{}{
		{{range $i,$c:=.columns}}"{{$c.name}}":input.{{$c.name|cname}}{{if $c.end}},{{end}}{{end}},
	}
	_, q, a, err := db.Execute(sql.Update{{.name|cname}}, params)
	if err != nil {
		return fmt.Errorf("更新{{.desc}}数据发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}
`
