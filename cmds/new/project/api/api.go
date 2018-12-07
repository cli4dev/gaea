package api

import (
	"fmt"
	"path/filepath"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/project/tmpls"
	"github.com/micro-plat/gaea/cmds/new/util"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//APICmd .
type APICmd struct {
	cli.Command
	cover bool
}

func NewAPICmd(cover bool) *APICmd {
	cmd := &APICmd{cover: cover}
	cmd.Command = cli.Command{
		Name:   "api",
		Usage:  "创建api项目",
		Flags:  cmd.geStartFlags(),
		Action: cmd.action,
	}
	return cmd
}

//GeStartFlags .
func (p *APICmd) geStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.StringFlag{
		Name:  "p",
		Value: ":8090",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt",
		Usage: "是否启用jwt",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "domain",
		Usage: "是否启用跨域设置，默认不启用",
	})

	return flags
}

//Action .
func (p *APICmd) action(c *cli.Context) (err error) {
	name, path, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		return err
	}
	//创键项目
	err = p.create(path)
	if err != nil && !p.cover {
		return err
	}

	err = p.writeTemplate(name, path, map[string]interface{}{
		"port":       util.GetPrefixString(types.GetString(c.String("p"), "8090"), ":"),
		"serverType": "api",
		"dbname":     util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"jwt":        c.String("jwt"),
		"db":         types.GetString(c.String("db")),
		"domain":     c.String("domain"),
	})
	if err != nil {
		return err
	}

	cmds.Log.Info("项目生成完成")
	return nil
}

func (p *APICmd) create(projectPath string) (err error) {
	if path.Exists(filepath.Join(projectPath, "main.go")) {
		err = fmt.Errorf("项目%s已经存在", projectPath)
		return err
	}
	return nil
}

func (p *APICmd) writeTemplate(projectName string, projectPath string, input map[string]interface{}) error {
	data, err := tmpls.GetTmpls(projectName, input)
	if err != nil {
		return err
	}
	for k, v := range data {
		fpath := filepath.Join(projectPath, k)

		f, err := path.CreatePath(fpath, p.cover)
		if err != nil {
			continue
		}
		_, err = f.WriteString(v)
		if err != nil {
			continue
		}
		defer f.Close()
		cmds.Log.Info("生成文件:", p)
	}
	return nil
}
