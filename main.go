package main

import "github.com/micro-plat/gaea/cmds"
import _ "github.com/micro-plat/gaea/cmds/create"
import _ "github.com/micro-plat/gaea/cmds/service"
import _ "github.com/micro-plat/gaea/cmds/module"
import _ "github.com/micro-plat/gaea/cmds/project"
import _ "github.com/micro-plat/gaea/cmds/md"
import _ "github.com/micro-plat/gaea/cmds/micservice"
import _ "github.com/go-sql-driver/mysql"

func main() {
	cmds.Start()
}
