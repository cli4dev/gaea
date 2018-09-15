package mysql

const tableTmpl = `
	create table {{.name}}(
		{{range $i,$c:=.columns}}{{$c.name}} {{$c.type}} {{$c.def}} {{$c.null}} {{if $c.end}},{{end}}
		{{end}}		
  );

	comment on table {{.name}} is '{{.desc}}';
	{{range $i,$c:=.columns}}comment on column {{$.name}}.{{$c.name}} is '{{$c.desc}}';	
	{{end}}
 
	{{range $i,$c:=.pks}}alter table {{$.name}}
	add constraint pk_{{$.name|nName}} primary key({{$c}});{{end}}

	{{range $i,$c:=.unqs}}alter table {{$.name}}
  add constraint {{$c.name|nName}} unique({{$c.flds}});{{end}}

	{{range $i,$c:=.seqs}}
	create sequence {{$c.name}}
	minvalue {{$c.min}}
	maxvalue {{$c.max}}
	start with {{$c.min}}
	cache 20;
	{{end}}
`
