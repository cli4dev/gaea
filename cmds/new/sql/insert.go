package sql

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/sql/insert"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
	"github.com/urfave/cli"
)

type insertCmd struct {
}

func NewInsertCmd() cli.Command {
	p := &insertCmd{}
	return cli.Command{
		Name:   "insert",
		Usage:  "创建基于md的insert into sql语句",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *insertCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 0)
	flags = append(flags, cli.StringSliceFlag{
		Name:  "filter,f",
		Usage: "关键字过滤",
	})
	return flags
}

func (p *insertCmd) action(c *cli.Context) (err error) {
	if c.NArg() < 1 {
		err = fmt.Errorf("请指定md文件路径和输出路径")
		fmt.Println(err)
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	filePath := c.Args().First()
	//	outPath := c.Args().Get(1)
	filters := c.StringSlice("filter")
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(tables, "", filters)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	for k, v := range tmpls {
		fmt.Println(k)
		fmt.Println(v)
	}
	// if err = p.createFile(outPath, tmpls); err != nil {
	// 	cmds.Log.Error(err)
	// 	return err
	// }
	cmds.Log.Infof("生成完成,共生成%d个语句", len(tmpls))
	return nil
}

func (p *insertCmd) createFile(root string, data map[string]string) error {
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
