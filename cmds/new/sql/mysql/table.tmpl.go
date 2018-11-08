package mysql

const tableTmpl = `
	create table {{.name}}(
		{{range $i,$c:=.columns}}{{$c.name}} {{$c.type|cType}} {{$c.def}} {{$c.null}}{{$c.seqs}} COMMENT '{{$c.desc}}' {{if $c.not_end}},{{end}},
		{{if $c.has_pk}}PRIMARY KEY (.pks){{end}}
		 
		{{end}}		
  );

 



	{{range $i,$c:=.unqs}}alter table {{$.name}}
 CREATE UNIQUE INDEX {{$c.name|nName}} ON {{$.name}}({{$c.flds}});
 {{end}}
`
