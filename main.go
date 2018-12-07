package main

import "github.com/micro-plat/gaea/cmds"
import _ "github.com/micro-plat/gaea/cmds/new"
import _ "github.com/micro-plat/gaea/cmds/create"
import _ "github.com/micro-plat/gaea/cmds/module"

func main() {
	cmds.Start()
}
