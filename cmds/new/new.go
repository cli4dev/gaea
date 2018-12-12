package new

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/service"
	"github.com/urfave/cli"
)

func NewCmd() []cli.Command {
	return []cli.Command{
		service.NewServiceCmd(),
	}
}
func init() {
	cmds.Register(NewCmd()...)
}
