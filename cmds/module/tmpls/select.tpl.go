package tmpls

//SelectOracleTmpl Select oracle sql 模板
const SelectOracleTmpl = `
//Get{{.name|cname}} 查询单条数据{{.desc}}
const Get{{.name|cname}} = 'select {{range $i,$c:=.getcolumns}}{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from {{.name}} where 1=1 {{range $i,$c:=.pk}}&{{$c.name}} {{end}}'

//Query{{.name|cname}}Count 获取{{.desc}}列表条数
const Query{{.name|cname}}Count = 'select count(1)
from {{.name}} where 1=1 {{range $i,$c:=.querycolumns}}&{{$c.name}} {{end}}'


//Query{{.name|cname}} 查询{{.desc}}列表数据
const Query{{.name|cname}} = '
select {{range $i,$c:=.selectcolumns}}TAB1.{{$c.name}}{{if $c.end}},{{end}}{{end}} 
from (select L.*  
	from (select rownum as rn,R.* 
		from (select t.* from {{.name}} t 
			where 1=1 {{range $i,$c:=.querycolumns}}&t.{{$c.name}} {{end}}) R 
	where rownum <= @pi * @ps) L 
where L.rn > (@pi - 1) * @ps) TAB1'
`

//SelectMysqlTmpl Select mysql sql 模板
const SelectMysqlTmpl = `
//Get{{.name|cname}} 查询单条数据{{.desc}}
{{$tbname := .name -}}
const Get{{.name|cname}} = 'select {{range $i,$c:=.getcolumns}}{{$tbname}}.{{$c.pname}}{{if $c.end}},{{end}}{{end}} 
{{if ne (.joinField|len) 0 -}}
,{{range $i,$c:=.joinField -}}
{{$c}}
{{- end}}
{{- end}}
from {{.name}} 
{{range $i,$c:=.joinCondition}}
{{$c}}
{{end}}
where 1=1 
{{if ne (.joinWhere|len) 0 -}}
{{range $i,$c:=.joinWhere -}}
{{$c}}
{{- end}}
{{- end}}
{{range $i,$c:=.pk}}&{{$c.name}} {{end}}'

//Query{{.name|cname}}Count 获取{{.desc}}列表条数
const Query{{.name|cname}}Count = 'select count(1)
from {{.name}} 
{{range $i,$c:=.joinCondition}}
{{$c}}
{{end}}
where 1=1 
{{if ne (.joinWhere|len) 0 -}}
{{range $i,$c:=.joinWhere -}}
{{$c}}
{{- end}}
{{- end}}
{{range $i,$c:=.querycolumns}} and if(@{{$c.name}} <> "",{{$tbname}}.{{$c.name}}=@{{$c.name}},1=1) {{end}}'

//Query{{.name|cname}} 查询{{.desc}}列表数据
const Query{{.name|cname}} = 'select {{range $i,$c:=.selectcolumns}}{{$tbname}}.{{$c.pname}}{{if $c.end}},{{end}}{{end}} 
{{if ne (.joinField|len) 0 -}}
,{{range $i,$c:=.joinField -}}
{{$c}}
{{- end}}
{{- end}}
from {{.name}} 
{{range $i,$c:=.joinCondition}}
{{$c}}
{{end}}
where 1=1 
{{if ne (.joinWhere|len) 0 -}}
{{range $i,$c:=.joinWhere -}}
{{$c}}
{{- end}}
{{- end}}
{{range $i,$c:=.querycolumns}} and if(@{{$c.name}} <> "",{{$tbname}}.{{$c.name}}=@{{$c.name}},1=1) {{end}} 
order by {{range $i,$c:=.pk}}{{$tbname}}.{{$c.name}} desc {{if $c.end}},{{end}}{{end}}
limit #pageSize offset #currentPage'
`

//SelectOracleFunc select 函数模板
const SelectOracleFunc = `
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
func(d *Db{{.name|cname}}) Query(input *Query{{.name|cname}}) (data db.QueryRows, count int, err error){

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.Query{{.name|cname}}Count, map[string]interface{}{
		{{range $i,$c:=.querycolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取{{.desc}}列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.Query{{.name|cname}}, map[string]interface{}{
		{{range $i,$c:=.querycolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
		"pi": input.Pi,
		"ps": input.Ps,
	})
	if err != nil {
		return nil,0, fmt.Errorf("获取{{.desc}}数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
`

//SelectMysqlFunc select 函数模板
const SelectMysqlFunc = `
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
func(d *Db{{.name|cname}}) Query(input *Query{{.name|cname}}) (data db.QueryRows, count int, err error){

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.Query{{.name|cname}}Count, map[string]interface{}{
		{{range $i,$c:=.querycolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取{{.desc}}列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.Query{{.name|cname}}, map[string]interface{}{
		{{range $i,$c:=.querycolumns -}}
		"{{$c.name}}":input.{{$c.name|cname}},
		{{end -}}
		"currentPage":  (types.GetInt(input.Pi) - 1) * types.GetInt(input.Ps),
		"pageSize": input.Ps,
	})
	if err != nil {
		return nil,0, fmt.Errorf("获取{{.desc}}数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
`
