package insert

const updateTmpl = `
//Update{{.name|cname}} 添加{{.desc}}
const Update{{.name|cname}} = 'update {{.name}} set
{{range $i,$c:=.columns}}{{$c.name}}=@{{$c.name}}{{if $c.end}},{{end}}{{end}}
where
{{range $i,$c:=.pk}}{{$c}}=@{{$c}}{{end}}'
`
