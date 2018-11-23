package project

import (
	"fmt"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/sql/conf"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
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
		c.StringSlice("f"))
}

func (p *moduleCmd) createModules(c, r, u, d, add, cover bool, t, o string, f []string) (err error) {

	//查找*.md文件
	mdList := []string{t}
	if t == "" {
		mdList = cmds.GetMDPath()
	}
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}
	//获取modules文件路径位置
	modulePath := o
	if o == "" {
		modulePath = cmds.GetModulePath()
	}
	for _, v := range mdList {
		//获取数据表
		tables, err := md.Markdown2Table(v)
		if err != nil {
			cmds.Log.Error(err)
			continue
		}
		//生成数据表对应的sql语句
		p.makeSQL(c, r, u, d, add, cover, tables, f, modulePath)
		//生成crud函数
		p.makeCrudFunc(c, r, u, d, add, cover, tables, f, modulePath)
	}
	return nil
}

func (p *moduleCmd) makeCrudFunc(c, r, u, d, add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {
	if c {
		p.makeInsertFunc(add, cover, tables, filters, modulePath)
	}
	if r {
		p.makeSelectFunc(add, cover, tables, filters, modulePath)
	}
	if u {
		p.makeUpdateFunc(add, cover, tables, filters, modulePath)
	}
	if d {
		p.makeDeleteFunc(add, cover, tables, filters, modulePath)
	}
	return nil
}

func (p *moduleCmd) makeSQL(c, r, u, d, add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {
	if c {
		p.makeInsertSQL(add, cover, tables, filters, modulePath)
	}
	if r {
		p.makeSelectSQL(add, cover, tables, filters, modulePath)
	}
	if u {
		p.makeUpdateSQL(add, cover, tables, filters, modulePath)
	}
	if d {
		p.makeDeleteSQL(add, cover, tables, filters, modulePath)
	}
	return nil
}
