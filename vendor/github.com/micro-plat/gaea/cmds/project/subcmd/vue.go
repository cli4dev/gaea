package subcmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/project/tmpls/vue/src"
	"github.com/micro-plat/gaea/cmds/project/tmpls/vue/src/pages"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
	"github.com/micro-plat/gaea/cmds/util/md"
	"github.com/micro-plat/gaea/cmds/util/path"

	"github.com/urfave/cli"
)

type VueCmd struct {
	projectName string
	projectPath string
	filters     []string
	t           string
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

	p.filters = c.StringSlice("f")
	p.projectName = c.String("n")
	p.t = c.String("t")

	err = p.setVueProjectPathAndName()
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	err = p.createVue()
	if err != nil {
		return err
	}

	return nil
}

func (p *VueCmd) createVue() error {
	//生成项目模板
	err := writeVueTemplate(p.projectPath, p.projectName)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//查找*.md文件
	mdList, err := getMdList(p.t)
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
		err = p.makeHTML(tables)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *VueCmd) makeHTML(tables []*conf.Table) (err error) {
	//获取模板数据
	tmpls, err := p.getHTMLData(tables)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//创建vue文件
	err = data.CreateHTMLFile(tmpls)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//写路由
	err = p.writeRouter(tmpls)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//生成配置文件
	err = p.writeMenuConf(tables)
	if err != nil {
		return err
	}

	return nil
}

//GetHTMLData .
//获取生成 vue所需的数据
func (p *VueCmd) getHTMLData(tables []*conf.Table) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetHTMLTmples("html", pages.HTMLTpl, tables, p.filters, p.projectPath)

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

func (p *VueCmd) setVueProjectPathAndName() (err error) {
	if path.Exists(filepath.Join(p.projectName, "package.json")) {
		err = fmt.Errorf("项目%s已经存在", p.projectName)
		return err
	}

	if p.projectName == "./" {
		p.projectName = "vue"
	}
	path, _ := os.Getwd()
	p.projectPath = filepath.Join(path, p.projectName)

	return nil
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

func (p *VueCmd) writeRouter(d map[string]map[string]string) error {
	router := make(map[string]string)
	for k := range d {
		i := strings.Index(k, "pages")
		dot := strings.LastIndex(k, ".")
		router[k[i+6:dot]] = k[i+6:]
	}
	str, err := data.Translate("router", src.RouterString, map[string]interface{}{
		"router": router,
	})
	if err != nil {
		return err
	}
	err = data.ReplaceFileStr("page.router", filepath.Join(p.projectPath, "/src/router.js"), str)
	if err != nil {
		return err
	}
	return nil
}

func (p *VueCmd) writeMenuConf(tables []*conf.Table) error {
	f, err := path.CreatePath(filepath.Join(p.projectName, "/menuConf.js"), true)
	defer f.Close()
	if err != nil {
		return err
	}
	var s string
	for _, v := range tables {
		s += v.Desc + "  " + data.GetRouterPath(v.Name) + "\n"
	}
	_, err = f.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}
