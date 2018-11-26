package service

import (
	"fmt"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/data"
	"github.com/micro-plat/gaea/cmds/new/util/md"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

type serviceCmd struct {
}

//NewServiceCmd .
func NewServiceCmd() cli.Command {
	p := &serviceCmd{}
	return cli.Command{
		Name:   "service",
		Usage:  "创建服务",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *serviceCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.BoolFlag{
		Name:  "c",
		Usage: "根据表结构生成 post handle函数",
	}, cli.BoolFlag{
		Name:  "r",
		Usage: "根据表结构生成 get and query handle 函数",
	}, cli.BoolFlag{
		Name:  "u",
		Usage: "根据表结构生成 update handle函数",
	}, cli.BoolFlag{
		Name:  "d",
		Usage: "根据表结构生成 delete handle函数",
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
		Usage: "是否执行追加 handle 函数",
	}, cli.BoolFlag{
		Name:  "cover",
		Usage: "是否执行覆盖 handle 函数操作",
	})
	return flags
}

func (p *serviceCmd) action(c *cli.Context) (err error) {

	return p.createServices(
		c.Bool("c"),
		c.Bool("r"),
		c.Bool("u"),
		c.Bool("d"),
		c.Bool("add"),
		c.Bool("cover"),
		c.String("t"),
		c.String("o"),
		c.StringSlice("f"),
	)
}

func (p *serviceCmd) createServices(c, r, u, d, add, cover bool, t, o string, f []string) (err error) {

	//查找*.md文件
	mdList := []string{t}
	if t == "" {
		mdList = path.GetMDPath()
	}

	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取services文件路径位置
	serverPath := o
	if o == "" {
		serverPath = path.GetServerPath()
	}
	if serverPath == "" || !strings.Contains(serverPath, "services") {
		serverPath = "services/" + o
	}
	if cover {
		add = true
	}
	//创建文件
	for _, v := range mdList {
		//获取数据表

		tables, err := md.Markdown2Table(v)
		if err != nil {
			cmds.Log.Error(err)
			continue
		}

		err = p.makeServices(c, r, u, d, add, cover, tables, f, serverPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *serviceCmd) makeServices(c, r, u, d, add, cover bool, tables []*conf.Table, filters []string, serverPath string) error {
	if c {
		tmpls, err := p.makePostHandle(add, cover, tables, filters, serverPath)
		if err != nil {
			return err
		}

		err = data.CreateServicesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if r {
		tmpls, err := p.makeGetAndQueryHandle(add, cover, tables, filters, serverPath)
		if err != nil {
			return err
		}

		err = data.CreateServicesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if u {
		tmpls, err := p.makePutHandle(add, cover, tables, filters, serverPath)
		if err != nil {
			return err
		}

		err = data.CreateServicesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if d {
		tmpls, err := p.makeDeleteHandle(add, cover, tables, filters, serverPath)
		if err != nil {
			return err
		}

		err = data.CreateServicesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}

	return nil
}
