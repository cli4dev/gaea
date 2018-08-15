package insert

const insertTmpl = `
//Insert{{.name|cname}} 添加{{.desc}}
const Insert{{.name|cname}} = 'insert into {{.name}}({{range $i,$c:=.columns}}{{$c.name}}{{if $c.end}},{{end}}{{end}})values({{range $i,$c:=.columns}}@{{$c.name}}{{if $c.end}},{{end}}{{end}})'
`
