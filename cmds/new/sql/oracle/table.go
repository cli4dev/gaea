package oracle

const tableTmpl = `
	create table {{.name}}(
		{{range $i,$c:=.columns}}
		{{$c.name}} {{$c.type}} {{$c.def}} {{$c.null}},
		{{end}}
  );

	comment on table {{.name}} is '{{.desc}}';

	{{range $i,$c:=.columns}}
	{{$c.name}} {{$c.type}} {{$c.def}} {{$c.null}},
	comment on column {{$.name}}.{{$c.name}} is '{{$c.desc}}';
	{{end}}
 

	{{if .pkname!=""}}
		alter table {{.name}}
		add constraint PK_{{.name}} primary key({{.pkname}});
	{{end}}

	{{range $i,$c:=.seqs}}
	create sequence {{$c.name}}
	minvalue {{$c.min}}
	maxvalue {{$c.max}}
	start with {{$c.min}}
	cache 20;
	{{end}}


`
