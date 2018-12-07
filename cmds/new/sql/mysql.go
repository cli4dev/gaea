package sql

import (
	"path/filepath"

	"github.com/micro-plat/gaea/cmds/new/sql/mysql"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

type mySqlCmd struct{}

func NewMySqlCmd() cli.Command {
	p := &mySqlCmd{}
	return cli.Command{
		Name:   "mysql",
		Usage:  "生成mysql数据库的表结构",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *mySqlCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 1)
	return flags
}

func (p *mySqlCmd) action(c *cli.Context) (err error) {
	mdFilePath := path.GetMDPath()
	outPath := filepath.Join(path.GetModulePath(), "const/sql/mysql")
	if c.NArg() > 0 {
		mdFilePath = []string{c.Args().Get(0)}
	}
	if c.NArg() > 1 {
		outPath = c.Args().Get(1)
	}
	return create(mysql.GetTmples, mdFilePath, outPath)
}
