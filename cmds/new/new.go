package new

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/module"
	"github.com/micro-plat/gaea/cmds/new/project"
	"github.com/micro-plat/gaea/cmds/new/sql"
	_ "github.com/micro-plat/gaea/cmds/new/sql/mysql"
	_ "github.com/micro-plat/gaea/cmds/new/sql/oracle"
	"github.com/urfave/cli"
)

func init() {
	cmds.Registry(getCmd()...)
}

func getCmd() []cli.Command {
	cmds := make([]cli.Command, 0, 4)
	cmds = append(cmds, module.NewModuleCmd())
	cmds = append(cmds, project.NewProjectCmd()...)
	cmds = append(cmds, cli.Command{
		Name:        "create",
		Usage:       "创建数据库表结构",
		Subcommands: sql.NewSqlCmd(),
	})
	return cmds
}
