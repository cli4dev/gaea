package sql

import (
	"path/filepath"

	"github.com/micro-plat/gaea/cmds/create/sql/oracle"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

type oracleCmd struct{}

func NewOracleCmd() cli.Command {
	p := &oracleCmd{}
	return cli.Command{
		Name:   "oracle",
		Usage:  "生成oracle数据的表结构，索引，序列等",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *oracleCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 1)
	flags = append(flags, cli.StringSliceFlag{
		Name:  "filter,f",
		Usage: "过滤表名",
	})
	flags = append(flags, cli.BoolFlag{
		Name:  "cover,c",
		Usage: "覆盖已存在的文件",
	})
	return flags
}

func (p *oracleCmd) action(c *cli.Context) (err error) {

	mdFilePath := path.GetMDPath()
	outPath := filepath.Join(path.GetModulePath(), "const/sql/oracle")
	if c.NArg() > 0 {
		mdFilePath = []string{c.Args().Get(0)}
	}
	if c.NArg() > 1 {
		outPath = c.Args().Get(1)
	}
	return create(oracle.GetTmples, mdFilePath, outPath, c.Bool("cover"), c.StringSlice("filter")...)
}
