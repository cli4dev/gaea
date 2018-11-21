package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/project/tmpls"
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
		Name:  "add",
		Usage: "是否执行追加crud函数和sql语句操作",
	}, cli.BoolFlag{
		Name:  "cover",
		Usage: "是否执行覆盖crud函数和sql语句操作",
	})
	return flags
}

func (p *moduleCmd) action(c *cli.Context) (err error) {
	//没有指定md路径，和生成文件输出路径
	if c.String("t") == "" && c.String("o") == "" {
		return p.creatDefautModule(c)
	}
	if c.String("t") != "" && c.String("o") == "" {
		return p.createMdModule(c)
	}
	if c.String("t") == "" && c.String("o") != "" {
		return p.createOutPutModule(c)
	}
	if c.String("t") != "" && c.String("o") != "" {
		return p.createModule(c)
	}
	return nil
}

//creatDefautModule .
//gaea new module -c -r -u -d
//基于当前目录,根据trd_order_info表结构生成crud函数,SQL语句，输入参数实体等文件
func (p *moduleCmd) creatDefautModule(c *cli.Context) (err error) {
	//得到表名
	tables := c.StringSlice("f")
	fmt.Println(tables)
	//查找*.md文件
	mdList := cmds.GetMDPath()
	fmt.Println("md 文件", mdList)
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取modules文件路径位置
	modulPath := cmds.GetLocation()
	fmt.Println("modules path:", modulPath)
	//生成数据表对应的sql语句函数
	for _, v := range mdList {
		p.makeSQL(c, v, tables, "")
	}
	//生成crud函数
	for _, v := range mdList {
		p.makeCrudFunc(c, v, tables, "")
	}
	return nil
}

//createMdModule .
//gaea new module -c -t ./db.md
//指定mdwen文件,根据trd_order_info表结构生成crud函数,SQL语句，输入参数实体等文件
func (p *moduleCmd) createMdModule(c *cli.Context) (err error) {
	//得到表名
	tables := c.StringSlice("f")
	fmt.Println(tables)
	//查找*.md文件
	mdList := []string{c.String("t")}
	fmt.Println("md 文件", mdList)
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取modules文件路径位置
	modulPath := cmds.GetLocation()
	fmt.Println("modules path:", modulPath)
	//生成数据表对应的sql语句函数
	for _, v := range mdList {
		p.makeSQL(c, v, tables, "")
	}

	//生成crud函数
	for _, v := range mdList {
		p.makeCrudFunc(c, v, tables, "")
	}

	return nil
}

//createOutPutModule .
//gaea new module -c -o ./modules
//指定输出文件,根据trd_order_info表结构生成crud函数,SQL语句，输入参数实体等文件
func (p *moduleCmd) createOutPutModule(c *cli.Context) (err error) {
	//得到表名
	tables := c.StringSlice("f")
	fmt.Println(tables)
	//查找*.md文件
	mdList := cmds.GetMDPath()
	fmt.Println("md 文件", mdList)
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取modules文件输出路径位置
	modulPath := c.String("o")
	fmt.Println("modules path:", modulPath)
	//生成数据表对应的sql语句函数
	for _, v := range mdList {
		p.makeSQL(c, v, tables, modulPath)
	}

	//生成crud函数
	for _, v := range mdList {
		p.makeCrudFunc(c, v, tables, modulPath)
	}

	return nil
}

//createModule .
//gaea new module -c -t ./db.md -o ./modules
//指定输出文件,根据trd_order_info表结构生成crud函数,SQL语句，输入参数实体等文件
func (p *moduleCmd) createModule(c *cli.Context) (err error) {
	//得到表名
	tables := c.StringSlice("f")
	fmt.Println(tables)
	//查找*.md文件
	mdList := []string{c.String("t")}
	fmt.Println("md 文件", mdList)
	//判断是否有文件
	if len(mdList) == 0 {
		err = fmt.Errorf("未找到任何 *.md 文件")
		return err
	}

	//获取modules文件输出路径位置
	modulPath := c.String("o")
	fmt.Println("modules path:", modulPath)
	//生成数据表对应的sql语句函数
	for _, v := range mdList {
		p.makeSQL(c, v, tables, modulPath)
	}

	//生成crud函数
	for _, v := range mdList {
		p.makeCrudFunc(c, v, tables, modulPath)
	}

	return nil
}

func (p *moduleCmd) makeCrudFunc(c *cli.Context, filePath string, filters []string, outPath string) error {
	if c.Bool("c") {
		p.makeInsertFunc(c, filePath, filters, outPath)
	}
	if c.Bool("r") {
		p.makeSelectFunc(c, filePath, filters, outPath)
	}
	if c.Bool("u") {
		p.makeUpdateFunc(c, filePath, filters, outPath)
	}
	if c.Bool("d") {
		p.makeDeleteFunc(c, filePath, filters, outPath)
	}
	return nil
}

func (p *moduleCmd) makeSQL(c *cli.Context, filePath string, filters []string, outPath string) error {
	if c.Bool("c") {
		p.makeInsertSQL(c, filePath, filters, outPath)
	}
	if c.Bool("r") {
		p.makeSelectSQL(c, filePath, filters, outPath)
	}
	if c.Bool("u") {
		p.makeUpdateSQL(c, filePath, filters, outPath)
	}
	if c.Bool("d") {
		p.makeDeleteSQL(c, filePath, filters, outPath)
	}
	return nil
}

//createFile .
//创建并生成文件
func createFile(c *cli.Context, root string, data map[string]map[string]string) error {
	for k, v := range data {

		path := filepath.Join(root, k)
		dir := filepath.Dir(path)
		fmt.Println("path: ", path)
		if c.Bool("cover") {
			os.Remove(path)
			cmds.Log.Warn("将对文件进行覆盖操作", path)
		}
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			if strings.Contains(path, "sql") || strings.Contains(path, "_") {
				if !strings.Contains(path, "sql") {

					p := strings.Split(path, "/")
					s := strings.Join(p[:len(p)-1], "/")
					path = strings.Join([]string{s, "/const/sql/"}, "")
					err := os.MkdirAll(path, os.ModeDir|os.ModePerm)
					if err != nil {
						err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
						return err
					}
					path = strings.Join([]string{path, p[len(p)-1]}, "")
				} else {

					err := os.MkdirAll(root, os.ModeDir|os.ModePerm)
					if err != nil {
						err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
						return err
					}
				}
				f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
				if err != nil {
					err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
					return err
				}
				_, err = f.WriteString("package sql")
				if err != nil {
					return err
				}
				f.Close()
				srcf, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
				if err != nil {
					err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
					return err
				}
				defer srcf.Close()
				_, err = srcf.WriteString(v[k])
				if err != nil {
					return err
				}
				cmds.Log.Info("生成文件:", path)
				continue
			} else {
				err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
				if err != nil {
					err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
					return err
				}
				m := strings.Split(path, "/")
				f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
				if err != nil {
					err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
					return err
				}
				absPath, _ := filepath.Abs(path)
				fmt.Println("absPath", absPath)
				absPathArr := strings.Split(absPath, "/")
				var projectName string
				for i := 0; i < len(absPathArr); i++ {
					if absPathArr[i] == "modules" {
						projectName = absPathArr[i-1]
					}
				}
				fmt.Println("projectName", projectName)
				_, err = f.WriteString(fmt.Sprintf(v["head"], m[len(m)-2], projectName))
				if err != nil {
					return err
				}
				f.Close()
				srcf, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
				if err != nil {
					err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
					return err
				}
				defer srcf.Close()
				_, err = srcf.WriteString(v[k])
				if err != nil {
					return err
				}
				cmds.Log.Info("生成文件:", path)
				continue
			}

		}
		if _, err := os.Stat(path); err == nil || os.IsExist(err) {
			if c.Bool("add") {
				//继续执行追加操作
			} else {
				cmds.Log.Warn("文件已存在:", path, "未输入 -add 不执行追加操作")
				continue
			}
		}
		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
			return err
		}
		defer srcf.Close()
		_, err = srcf.WriteString(v[k])
		if err != nil {
			return err
		}
		cmds.Log.Info("生成文件:", path)
	}
	return nil

}

func (p *moduleCmd) new(projectName string, serviceType string, modules []string, restful bool) error {
	gopath, err := cmds.GetGoPath()
	if err != nil {
		return err
	}
	projectPath := filepath.Join(gopath, "src", projectName)
	tmpls, err := tmpls.GetModuleTmpls(projectName, serviceType, modules, restful)
	if err != nil {
		return err
	}
	return p.createProject(projectPath, tmpls)

}

func (p *moduleCmd) createProject(projectPath string, data map[string]string) error {
	for k, v := range data {
		path := filepath.Join(projectPath, k)
		dir := filepath.Dir(path)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}
		}
		if _, err := os.Stat(path); err == nil || os.IsExist(err) {
			cmds.Log.Warn("文件已存在:", path)
			continue
		}
		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
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
