package module

import (
	"fmt"
	"os"
	paths "path"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/module/tmpls"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
	"github.com/micro-plat/gaea/cmds/util/md"
	"github.com/micro-plat/gaea/cmds/util/path"

	"github.com/urfave/cli"
)

type moduleCmd struct {
}

//NewModuleCmd .
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
	}, cli.BoolFlag{
		Name:  "crud",
		Usage: "根据表结构生成 insert select update delete 函数,SQL语句，输入参数实体等文件",
	}, cli.BoolFlag{
		Name:  "cru",
		Usage: "根据表结构生成 insert select update 函数,SQL语句，输入参数实体等文件",
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
	}, cli.BoolFlag{
		Name:  "m",
		Usage: "是否生成一个基础的module",
	}, cli.StringFlag{
		Name:  "n",
		Usage: "生成基础module的名字",
	},
	)
	return flags
}

func (p *moduleCmd) action(c *cli.Context) (err error) {

	return p.createModules(
		c.Bool("c") || c.Bool("crud") || c.Bool("cru"),
		c.Bool("r") || c.Bool("crud") || c.Bool("cru"),
		c.Bool("u") || c.Bool("crud") || c.Bool("cru"),
		c.Bool("d") || c.Bool("crud"),
		c.Bool("add"),
		c.Bool("cover"),
		c.String("t"),
		c.String("o"),
		c.StringSlice("f"),
		c.String("db"),
		c.Bool("m"),
		c.String("n"),
	)
}

func (p *moduleCmd) createModules(c, r, u, d, add, cover bool, t, o string, f []string, db string, m bool, n string) (err error) {

	if m {
		if n == "" {
			cmds.Log.Error("请输入module名称，比如: -n user/info")
			return nil
		}

		modulePath := "modules/" + n + ".go"

		err = p.createBaseModule(modulePath, n)
		if err != nil {
			cmds.Log.Error(err)
		}
		return nil
	}

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
		modulePath = "modules/" + o
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
		tmpls, err := p.makeInsertFunc(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}

		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if r {
		tmpls, err := p.makeSelectFunc(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}
		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if u {
		tmpls, err := p.makeUpdateFunc(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}
		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if d {
		tmpls, err := p.makeDeleteFunc(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}
		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	return err
}

func (p *moduleCmd) makeSQL(c, r, u, d, add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) (err error) {

	if c {
		tmpls, err := p.makeInsertSQL(add, cover, db, tables, filters, modulePath)
		if err != nil {
			return err
		}

		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if r {
		tmpls, err := p.makeSelectSQL(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}

		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if u {
		tmpls, err := p.makeUpdateSQL(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}

		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	if d {
		tmpls, err := p.makeDeleteSQL(add, cover, tables, filters, modulePath)
		if err != nil {
			return err
		}

		err = data.CreateModulesFile(add, cover, tmpls)
		if err != nil {
			return err
		}

		cover = false
	}
	return err
}

func (p *moduleCmd) createBaseModule(modulePath, name string) error {
	_, err := os.Stat(modulePath)
	if err == nil {
		cmds.Log.Error("module已经存在")
		return nil
	}
	err = os.MkdirAll(paths.Dir(modulePath), os.ModePerm)
	f, err := os.Create(modulePath)
	if err != nil {
		err = fmt.Errorf("无法创建文件:%s(err:%v)", modulePath, err)
		return err
	}
	defer f.Close()
	m := strings.Split(modulePath, "/")
	_, err = f.WriteString(fmt.Sprintf(tmpls.BaseTpl, m[len(m)-2], m[len(m)-2], m[len(m)-2], m[len(m)-2], m[len(m)-2], m[len(m)-2]))
	if err != nil {
		return err
	}
	cmds.Log.Info("module创建成功")
	return nil
}
func init() {
	cmds.Register(NewModuleCmd())
}
