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
###  {{.desc}}[{{.name}}]

| 字段名       | 类型       | 默认值  | 为空 |   约束    | 描述                                |
| ------------| -----------| :----: | :--: | :-------: | :---------------------------------|
{{range $i,$c:=.columns -}}
| {{$c.name}} | {{$c.type}}({{$c.len}})|{{$c.def}}|{{if $c.isnull}}是{{else}}否{{end}}| {{$c.cons}} | {{$c.desc}}|
{{end -}}
`
