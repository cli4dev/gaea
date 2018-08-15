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
	flags = append(flags, cli.StringSliceFlag{
		Name:  "modules,m",
		Usage: "注册的模块名称",
	})
	return flags
}

func (p *moduleCmd) action(c *cli.Context) (err error) {
	if c.NArg() == 0 {
		err = fmt.Errorf("未指定项目名称")
		fmt.Println(err)
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	projectName := strings.Trim(c.Args().First(), "/")
	modules := c.StringSlice("modules")

	if err := p.new(projectName, "", modules, false); err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("模块生成完成")
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
