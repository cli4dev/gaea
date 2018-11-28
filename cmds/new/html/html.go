package html

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

type htmlCmd struct {
}

//NewHTMLCmd .
func NewHTMLCmd() cli.Command {
	p := &htmlCmd{}
	return cli.Command{
		Name:   "html",
		Usage:  "创建基于element-ui的前端页面",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *htmlCmd) geStartFlags() []cli.Flag {

	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringSliceFlag{
		Name:  "f",
		Usage: "过滤器，指定表明或关键字",
	}, cli.StringFlag{
		Name:  "t",
		Usage: "指定数据表 md 文件",
	}, cli.StringFlag{
		Name:  "o",
		Usage: "生成的文件输出路径",
	})
	return flags
}

func (p *htmlCmd) action(c *cli.Context) (err error) {

	return p.createHTML(
		c.String("t"),
		c.String("o"),
		c.StringSlice("f"),
	)
}

func (p *htmlCmd) createHTML(t, o string, f []string) (err error) {

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
	//获取生成的vue文件输出路径
	htmlPath := o
	if o == "" {
		htmlPath = path.GetHTMLPath()
	}
	if htmlPath == "" || !strings.Contains(htmlPath, "html") {
		htmlPath = "html/" + o
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
		err = p.makeHTML(tables, f, htmlPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *htmlCmd) makeHTML(tables []*conf.Table, filters []string, htmlPath string) (err error) {

	tmpls, err := p.GetHTMLData(tables, filters, htmlPath)
	if err != nil {
		return err
	}

	err = data.CreateHTMLFile(tmpls)
	if err != nil {
		return err
	}
	return nil
}
