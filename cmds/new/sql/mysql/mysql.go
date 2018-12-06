package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/sql/mysql"
	"github.com/micro-plat/gaea/cmds/new/util/md"
	"github.com/urfave/cli"
)

type mysqlCmd struct {
}

func NewMysqlCmd() cli.Command {
	p := &mysqlCmd{}
	return cli.Command{
		Name:   "mysql",
		Usage:  "创建SQL语句",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}
func (p *mysqlCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 0)
	return flags
}
func (p *mysqlCmd) action(c *cli.Context) (err error) {
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
	tmpls, err := mysql.GetTmples(tables)
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
	cmds.Log.Infof("SQL生成完成,共生成%d个文件", len(tmpls))
	return nil
}

func (p *mysqlCmd) createFile(root string, data map[string]string) error {
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

func init() {
	sql.Register("mysql", GetTmples)
}
