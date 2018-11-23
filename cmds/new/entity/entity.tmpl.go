package entity

const entityTmpl = `
package {{.path|lname}}


//{{.name|cname}} {{.desc}} 
type {{.name|cname}} struct {		
		{{range $i,$c:=.columns}}
		//{{$c.name|cname}} {{$c.desc|cname}}
		{{$c.name|cname}} {{$c.type|ctype}} 'json:"{{$c.name|lower}}" form:"{{$c.name|lower}}" m2s:"{{$c.name|lower}}" {{$c.null|valid}}'
		{{end}}	
}
`
