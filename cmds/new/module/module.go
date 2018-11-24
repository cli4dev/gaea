package module

import (
	"fmt"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/md"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

type moduleCmd struct {
}

func NewModuleCmd() cli.Command {
	p := &moduleCmd{}
	return cli.Command{
		Name:   "module",
		Usage:  "创建模块",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *moduleCmd) geStartFlags() []cli.Flag {

	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.BoolFlag{
		Name:  "c",
		Usage: "根据表结构生成 insert 函数,insert SQL语句，输入参数实体等文件",
	}, cli.BoolFlag{
		Name:  "r",
		Usage: "根据表结构生成 select 函数,select SQL语句，输入参数实体等文件",
	}, cli.BoolFlag{
		Name:  "u",
		Usage: "根据表结构生成 update 函数,update SQL语句，输入参数实体等文件",
	}, cli.BoolFlag{
		Name:  "d",
		Usage: "根据表结构生成 delete 函数,delet SQL语句，输入参数实体等文件",
	}, cli.StringSliceFlag{
		Name:  "f",
		Usage: "过滤器，指定表明或关键字",
	}, cli.StringFlag{
		Name:  "t",
		Usage: "指定数据表 md 文件",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定生成 mysql 或 oracle 的函数和sql,默认为 mysql",
	}, cli.StringFlag{
		Name:  "o",
		Usage: "生成的文件输出路径",
	}, cli.BoolFlag{
		Name:  "add,a",
		Usage: "是否执行追加crud函数和sql语句操作",
	}, cli.BoolFlag{
		Name:  "cover",
		Usage: "是否执行覆盖crud函数和sql语句操作",
	})
	return flags
}

func (p *moduleCmd) action(c *cli.Context) (err error) {

	return p.createModules(
		c.Bool("c"),
		c.Bool("r"),
		c.Bool("u"),
		c.Bool("d"),
		c.Bool("add"),
		c.Bool("cover"),
		c.String("t"),
		c.String("o"),
		c.StringSlice("f"),
		c.String("db"),
	)
}

func (p *moduleCmd) createModules(c, r, u, d, add, cover bool, t, o string, f []string, db string) (err error) {

	//查找*.md文件
	mdList := []string{t}
	if t == "" {
		mdList = path.GetMDPath()
	}

	//获取db 类型
	if db == "" || !strings.Contains(db, "oracle") {
		db = "mysql"
	}

	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取modules文件路径位置
	modulePath := o
	if o == "" {
		modulePath = path.GetModulePath()
	}
	if modulePath == "" || !strings.Contains(modulePath, "modules") {
		cmds.Log.Error("没有指定 modules,或 'modules' 输入错误")
		return nil
	}

	//创建文件
	for _, v := range mdList {
		//获取数据表
		tables, err := md.Markdown2Table(v)
		if err != nil {
			cmds.Log.Error(err)
			continue
		}
		//生成数据表对应的sql语句
		err = p.makeSQL(c, r, u, d, add, cover, db, tables, f, modulePath)
		if err != nil {
			return err
		}
		//生成crud函数
		err = p.makeCrudFunc(c, r, u, d, add, cover, tables, f, modulePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *moduleCmd) makeCrudFunc(c, r, u, d, add, cover bool, tables []*conf.Table, filters []string, modulePath string) (err error) {

	if c {
		err = p.makeInsertFunc(add, cover, tables, filters, modulePath)
	}
	if r {
		err = p.makeSelectFunc(add, cover, tables, filters, modulePath)
	}
	if u {
		err = p.makeUpdateFunc(add, cover, tables, filters, modulePath)
	}
	if d {
		err = p.makeDeleteFunc(add, cover, tables, filters, modulePath)
	}
	return err
}

func (p *moduleCmd) makeSQL(c, r, u, d, add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) (err error) {

	if c {
		err = p.makeInsertSQL(add, cover, db, tables, filters, modulePath)
	}
	if r {
		err = p.makeSelectSQL(add, cover, tables, filters, modulePath)
	}
	if u {
		err = p.makeUpdateSQL(add, cover, tables, filters, modulePath)
	}
	if d {
		err = p.makeDeleteSQL(add, cover, tables, filters, modulePath)
	}
	return err
}
