package project

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/project/subcmd"

	"github.com/urfave/cli"
)

type projectCmd struct {
}

//NewProjectCmd .
func NewProjectCmd() []cli.Command {
	return []cli.Command{
		subcmd.NewAPICmd(),
		subcmd.NewCronCmd(),
		subcmd.NewWebCmd(),
		subcmd.NewMQCCmd(),
		subcmd.NewRPCCmd(),
		subcmd.NewWSCmd(),
		subcmd.NewVueCmd(),
	}
}

func init() {
	cmds.Register(NewProjectCmd()...)
}
