package mysql

const tableTmpl = `
	create table {{.name}}(
		{{range $i,$c:=.columns}}{{$c.name}}  {{$c.type|cType}} {{$c.def}} {{$c.null}} {{$c.pk}} {{$c.seqs}}  comment '{{$c.desc}}' {{if $c.not_end}},{{end}}
		{{end}}		
  )COMMENT='{{.desc}}';

 



	{{range $i,$c:=.unqs}}alter table {{$.name}}
 CREATE UNIQUE INDEX {{$c.name|nName}} ON {{$.name}}({{$c.flds}});
 {{end}}
`
