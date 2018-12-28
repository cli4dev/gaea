package create

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/create/sql"
	"github.com/urfave/cli"
)

func init() {
	cmds.Register(cli.Command{
		Name:  "create",
		Usage: "创建数据库表结构",
		Subcommands: []cli.Command{
			sql.NewMySqlCmd(),
			sql.NewOracleCmd(),
		},
	})
}
