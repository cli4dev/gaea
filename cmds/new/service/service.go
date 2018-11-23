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

type serviceCmd struct {
}

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
	flags = append(flags, cli.StringSliceFlag{
		Name:  "service,s",
		Usage: "注册的模块名称",
	})
	flags = append(flags, cli.BoolFlag{
		Name:  "restful,r",
		Usage: "生成restful风格的服务代码",
	})
	return flags
}

func (p *serviceCmd) action(c *cli.Context) (err error) {
	if c.NArg() == 0 {
		err = fmt.Errorf("未指定项目名称")
		fmt.Println(err)
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	projectName := strings.Trim(c.Args().First(), "/")
	modules := c.StringSlice("service")

	if err := p.new(projectName, "", modules, c.Bool("restful")); err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("服务生成完成")
	return nil
}

func (p *serviceCmd) new(projectName string, serviceType string, modules []string, restful bool) error {
	gopath, err := cmds.GetGoPath()
	if err != nil {
		return err
	}
	projectPath := filepath.Join(gopath, "src", projectName)
	tmpls, err := tmpls.GetServiceTmpls(projectName, serviceType, modules, restful)
	if err != nil {
		return err
	}
	return p.createProject(projectPath, tmpls)

}

func (p *serviceCmd) createProject(projectPath string, data map[string]string) error {
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
