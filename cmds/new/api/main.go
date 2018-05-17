package api

const mainTmpl = `

package main

import "github.com/micro-plat/hydra/hydra"

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("{{.projectName|fName}}"),
		hydra.WithSystemName("{{.projectName|lName}}"),
		hydra.WithServerTypes("{{.serverType}}"))
	bind(app)
	app.Start()
}


`
