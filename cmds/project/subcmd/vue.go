package subcmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/data"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/micro-plat/gaea/cmds/project/tmpls/vue"

	"github.com/micro-plat/gaea/cmds/new/util/md"

	"github.com/urfave/cli"
)

type VueCmd struct {
}

//NewVueCmd .
func NewVueCmd() cli.Command {
	p := &VueCmd{}
	return cli.Command{
		Name:  "vue",
		Usage: "创建vue项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建vue项目",
				Flags:  p.geStartFlags(),
				Action: p.action,
			},
		},
	}
}

func (p *VueCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "n",
		Value: "./",
		Usage: "vue项目名称",
	}, cli.StringSliceFlag{
		Name:  "f",
		Usage: "过滤器，指定表明或关键字",
	}, cli.StringFlag{
		Name:  "t",
		Usage: "指定数据表 md 文件",
	})
	return flags
}

func (p *VueCmd) action(c *cli.Context) (err error) {

	f := c.StringSlice("f")
	projectPath, err := getVueProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	err = p.createVue(projectPath, c.String("t"), f)
	if err != nil {
		return err
	}

	return nil
}

func (p *VueCmd) createVue(projectPath, t string, f []string) error {
	//生成项目模板
	err := writeVueTemplate(projectPath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//查找*.md文件
	mdList, err := getMdList(t)
	if err != nil {
		return err
	}
	//生成页面模板
	for _, v := range mdList {
		//获取数据表
		tables, err := md.Markdown2Table(v)
		if err != nil {
			cmds.Log.Error(err)
			continue
		}
		//生成数据表对应的 vue 文件
		err = p.makeHTML(tables, f, projectPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *VueCmd) makeHTML(tables []*conf.Table, filters []string, projectPath string) (err error) {

	tmpls, err := p.getHTMLData(tables, filters, projectPath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	err = data.CreateHTMLFile(tmpls)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	err = writeRouter(projectPath, tmpls)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//GetHTMLData .
//获取生成 vue所需的数据
func (p *VueCmd) getHTMLData(tables []*conf.Table, filters []string, projectPath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetHTMLTmples("html", vue.HTMLTpl, tables, filters, projectPath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成html时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil

}

func getVueProjectPath(n string) (projectPath string, err error) {
	if path.Exists(filepath.Join(n, "package.json")) {
		err = fmt.Errorf("项目%s已经存在", n)
		return "", err
	}
	path, _ := os.Getwd()
	return filepath.Join(path, n), nil
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

func writeRouter(projectPath string, d map[string]map[string]string) error {
	router := make(map[string]string)
	for k := range d {
		i := strings.Index(k, "pages")
		dot := strings.LastIndex(k, ".")
		router[k[i+5:dot]] = k[i+6:]
	}
	str, err := data.Translate("router", vue.RouterString, map[string]interface{}{
		"router": router,
	})
	if err != nil {
		return err
	}
	err = data.ReplaceFileStr("page.router", filepath.Join(projectPath, "/src/router.js"), str)
	if err != nil {
		return err
	}
	return nil
}
