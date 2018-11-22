package delete

const DeleteTmpl = `
//Delete{{.name|cname}} 删除{{.desc}}
const Delete{{.name|cname}} = 'delete from {{.name}} where {{range $i,$c:=.pk}}&{{$c.name}} {{end}}'
`
const DeleteFunc = `
//Delete 删除{{.desc}}
func(d *Db{{.name|cname}}) Delete({{range $i,$c:=.pk -}}{{$c.name|cname}} {{$c.type|ctype}}{{if $c.end}},{{end}}{{end -}}) error {

	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.Delete{{.name|cname}}, map[string]interface{}{
		  {{range $i,$c:=.pk -}}
		  "{{$c.name}}":{{$c.name|cname}},
		  {{end -}}
	})
	if err != nil {
		return fmt.Errorf("删除{{.desc}}数据发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	return nil
}
`