package micservice

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/micservice/subcmd"
	"github.com/urfave/cli"
)

type mic struct {
}

//NewMicCmd 创建空的module和service函数及对象
func NewMicCmd() []cli.Command {
	return []cli.Command{
		subcmd.NewMicServiceCmd(),
	}
}

func init() {
	cmds.Register(NewMicCmd()...)
}
