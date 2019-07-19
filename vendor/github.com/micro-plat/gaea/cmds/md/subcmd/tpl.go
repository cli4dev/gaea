package subcmd

const mdTPL = `
###  {{.desc}}[{{.name}}]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
{{range $i,$c:=.columns -}}
| {{$c.name}} | {{$c.type}}|{{$c.def}}|{{if $c.isnull}}是{{else}}否{{end}}| {{$c.cons}} | {{$c.desc}}|
{{end -}}
`

const mdOracleTPL = `
{{ $empty := "" -}}
###  {{.desc}}[{{.name}}]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
{{range $i,$c:=.columns -}}
| {{$c.name}} | {{$c.type}}{{if eq $c.len $empty}}{{else}}({{$c.len}}){{end}}|{{$c.def}}|{{if $c.isnull}}是{{else}}否{{end}}| {{$c.cons}} | {{$c.desc}}|
{{end -}}
`
