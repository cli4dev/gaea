package tmpls

const mainTmpl = `
package main

import "github.com/micro-plat/hydra/hydra"
{{$mysql := "mysql" -}}
{{if eq .dbname $mysql -}}
import _ "github.com/go-sql-driver/mysql"
{{- end}}

type {{.projectName|lName}} struct {
	*hydra.MicroApp
}

func main() {
	app := &{{.projectName|lName}}{
		hydra.NewApp(
			hydra.WithPlatName("{{.projectName|fName}}"),
			hydra.WithSystemName("{{.projectName|lName}}"),
			hydra.WithServerTypes("{{.serverType}}"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
`
