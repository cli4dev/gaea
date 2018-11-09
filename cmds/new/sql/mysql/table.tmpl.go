package mysql

const tableTmpl = `
 drop table {{.name}};

	create table {{.name}}(
		{{range $i,$c:=.columns}}{{$c.name}}  {{$c.type|cType}} {{$c.def}} {{$c.null}} {{$c.pk}} {{$c.seqs}}  comment '{{$c.desc}}' {{if $c.not_end}},{{end}}
		{{end}}		
  )COMMENT='{{.desc}}';

 



{{range $i,$c:=.unqs}}
	drop index {{$c.name|nName}} ON {{$.name}};
 create unique index {{$c.name|nName}} ON {{$.name}}({{$c.flds}});
 {{end}}
`
