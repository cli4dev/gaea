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
	cmds = append(cmds, sql.NewSqlCmd())
	cmds = append(cmds, module.NewModuleCmd())
	cmds = append(cmds, project.NewProjectCmd()...)
	return cmds
	// return []cli.Command{
	// 	cli.Command{
	// 		Name:  "new",
	// 		Usage: "创建项目,文件，数据库等",
	// 		Subcommands: []cli.Command{
	// 			project.NewProjectCmd(),
	// 			module.NewModuleCmd(),
	// 			service.NewServiceCmd(),
	// 			sql.NewSqlCmd(),
	// 			key.NewKeyCmd(),
	// 			html.NewHTMLCmd(),
	// 			//sql.NewStructCmd(),
	// 		},
	// 		cli.Command{
	// 		Name:  "project",
	// 		Usage: "创建项目,文件，数据库等",
	// 		Subcommands: project.NewProjectCmd(),
	// 		,
	// 	}}
}
