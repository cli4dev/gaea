package new

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/module"
)

func init() {
	cmds.Registry(module.NewModuleCmd())
}
