package new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/api"
	"github.com/urfave/cli"
)

type projectCmd struct {
}

func newProjectCmd() cli.Command {
	p := &projectCmd{}
	return cli.Command{
		Name:   "project",
		Usage:  "创建项目",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *projectCmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "server-type,s",
		Value: "api",
		Usage: "服务器类型(api,rpc,web,mqc,cron),多个用短横线分隔",
	})
	flags = append(flags, cli.StringSliceFlag{
		Name:  "modules,m",
		Usage: "注册的模块名称",
	})
	flags = append(flags, cli.BoolFlag{
		Name:  "restful,r",
		Usage: "生成restful风格的服务代码",
	})
	return flags
}

func (p *projectCmd) action(c *cli.Context) (err error) {
	if c.NArg() == 0 {
		err = fmt.Errorf("未指定项目名称")
		fmt.Println(err)
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}
	projectName := c.Args().First()
	serverType := c.String("server-type")
	modules := c.StringSlice("modules")

	if err := p.new(projectName, serverType, modules, c.Bool("restful")); err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}

func (p *projectCmd) new(projectName string, serviceType string, modules []string, restful bool) error {
	gopath, err := cmds.GetGoPath()
	if err != nil {
		return err
	}
	projectPath := filepath.Join(gopath, "src", projectName)
	tmpls, err := api.GetTmpl(projectName, serviceType, modules, restful)
	if err != nil {
		return err
	}
	return p.createProject(projectPath, tmpls)

}

func (p *projectCmd) createProject(projectPath string, data map[string]string) error {
	for k, v := range data {
		path := filepath.Join(projectPath, k)
		dir := filepath.Dir(path)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}
		}

		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
			return err
		}
		defer srcf.Close()
		_, err = srcf.WriteString(v)
		if err != nil {
			return err
		}
		cmds.Log.Info("创建文件:", path)
	}
	return nil

}
