package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

		err = p.writeRouter(tables, f)
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

func (p *serviceCmd) writeRouter(tables []*conf.Table, filters []string) error {
	out := map[string]string{}
	for _, tb := range tables {
		if len(filters) > 0 {
			e := false
			for _, f := range filters {
				if strings.EqualFold(tb.Name, f) {
					e = true
					break
				}
			}
			if !e {
				continue
			}
		}
		out[data.GetPath(tb.Name)] = data.GetHandleName(tb.Name)
	}

	srcf, err := os.OpenFile("./init.go", os.O_RDWR, os.ModePerm)
	defer srcf.Close()
	if err != nil {
		err = fmt.Errorf("无法打开文件:%s(err:%v)", ".init.go", err)
		return err
	}
	buf, err := ioutil.ReadAll(srcf)
	if err != nil {
		cmds.Log.Errorf("%v", err.Error())
		return err
	}
	result := string(buf)

	k := fmt.Sprintf(`//%s#//[\s\S]+//#%s//`, "service.router", "service.router")
	if ok, _ := regexp.Match(k, buf); !ok {
		cmds.Log.Errorf("没有找到配置定位标识符:%s//%s#//....//#%s//", "./init.go", "service.router", "service.router")
		return nil
	}
	re, _ := regexp.Compile(k)
	str := re.Find(buf)
	re.FindStringSubmatch(result)
	s := strings.Replace(string(str), "//service.router#//", "", -1)
	s = strings.Replace(s, "//#service.router//", "", -1)

	for k, v := range out {
		if !strings.Contains(s, k) {
			s += fmt.Sprintf(`r.Micro("%s",%s,"*")`, k, v) + "\n\t\t"
		}
	}
	s = "//service.router#//" + s + "//#service.router//"

	fstr := re.ReplaceAllString(result, s)
	//清空文件数据
	srcf.Truncate(0)
	n, _ := srcf.Seek(0, os.SEEK_SET)
	_, err = srcf.WriteAt([]byte(fstr), n)
	if err != nil {
		return err
	}
	return nil
}
