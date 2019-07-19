package md

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/md/subcmd"
	"github.com/urfave/cli"
)

type md struct {
}

//NewMDCmd .
func NewMDCmd() []cli.Command {
	return []cli.Command{
		subcmd.NewTable2MDCmd(),
	}
}

func init() {
	cmds.Register(NewMDCmd()...)
}
