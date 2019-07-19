package sql

import (
	"fmt"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds/create/sql/mysql"
	"github.com/micro-plat/gaea/cmds/util/path"
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
	flags = append(flags, cli.StringSliceFlag{
		Name:  "filter,f",
		Usage: "过滤表名",
	})
	flags = append(flags, cli.BoolFlag{
		Name:  "cover,c",
		Usage: "覆盖已存在的文件",
	}, cli.StringFlag{
		Name:  "t",
		Usage: "指定文件",
	})
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
	return create(mysql.GetTmples, mdFilePath, outPath, c.Bool("cover"), c.StringSlice("filter")...)
}

func getMdList(t string) (mdList []string, err error) {
	if t == "" {
		mdList = path.GetMDPath()
	}
	mdList = append(mdList, t)
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return nil, err
	}
	return mdList, nil
}
