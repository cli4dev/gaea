package new

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/module"
	"github.com/micro-plat/gaea/cmds/new/sql"
	"github.com/urfave/cli"
)

func init() {
	cmds.Registry(getCmd())
}

func getCmd() cli.Command {
	return cli.Command{
		Name:  "new",
		Usage: "创建项目,文件，数据库等",
		Subcommands: []cli.Command{
			//project.NewProjectCmd(),
			module.NewModuleCmd(),
			//service.NewServiceCmd(),
			sql.NewSqlCmd(),
			//sql.NewStructCmd(),
		},
	}
}
