package sql

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
	"github.com/micro-plat/gaea/cmds/new/sql/oracle"
	"github.com/urfave/cli"
)

type sqlCmd struct {
}

func NewSqlCmd() cli.Command {
	p := &sqlCmd{}
	return cli.Command{
		Name:   "sql",
		Usage:  "创建SQL语言",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *sqlCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	// flags = append(flags, cli.StringSliceFlag{
	// 	Name:  "dbt,t",
	// 	Usage: "数据库类型",
	// })
	// flags = append(flags, cli.BoolFlag{
	// 	Name:  "restful,r",
	// 	Usage: "生成restful风格的服务代码",
	// })
	return flags
}

func (p *sqlCmd) action(c *cli.Context) (err error) {
	if c.NArg() < 2 {
		err = fmt.Errorf("请指定md文件路径和输出路径")
		fmt.Println(err)
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	filePath := c.Args().First()
	outPath := c.Args().Get(1)
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := oracle.GetTmples(tables)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	if err = p.createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Info("SQL生成完成")
	return nil
}

func (p *sqlCmd) createFile(root string, data map[string]string) error {
	for k, v := range data {
		path := filepath.Join(root, k)
		dir := filepath.Dir(path)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
			if err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}
		}
		if _, err := os.Stat(path); err == nil || os.IsExist(err) {
			cmds.Log.Warn("文件已存在:", path)
			continue
		}
		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
			return err
		}
		defer srcf.Close()
		_, err = srcf.WriteString(v)
		if err != nil {
			return err
		}
		cmds.Log.Info("生成文件:", path)
	}
	return nil

}
