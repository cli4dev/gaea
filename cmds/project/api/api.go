package api

import (
	"fmt"
	"path/filepath"

	datas "github.com/micro-plat/gaea/cmds/new/util/data"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/util"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/micro-plat/gaea/cmds/project/tmpls"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//APICmd .
type APICmd struct {
	cli.Command
	cover bool
}

func NewAPICmd() cli.Command {

	apiCreator := &APICmd{cover: false}
	apiCover := &APICmd{cover: true}

	return cli.Command{
		Name:  "api",
		Usage: "创建api项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建api项目",
				Flags:  apiCreator.getStartFlags(),
				Action: apiCreator.action,
			},
			cli.Command{
				Name:   "cover",
				Usage:  "根据指定参数覆盖已有配置",
				Flags:  apiCover.getStartFlags(),
				Action: apiCover.action,
			}, cli.Command{
				Name:   "remove",
				Usage:  "移除指定配置",
				Flags:  apiCover.getRemoveStartFlags(),
				Action: apiCover.removeAction,
			},
		},
	}

}

//getStartFlags .
func (p *APICmd) getStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.StringFlag{
		Name:  "port,p",
		Value: ":8090",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt",
		Usage: "是否启用jwt",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cros",
		Usage: "是否启用跨域设置，默认不启用",
	}, cli.BoolFlag{
		Name:  "metric",
		Usage: "启用metric配置",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	}, cli.BoolFlag{
		Name:  "appconf",
		Usage: "启用appconf配置",
	},
	)
	return flags
}
func (p *APICmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "port,p",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt",
		Usage: "是否启用jwt",
	}, cli.BoolFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cros",
		Usage: "是否启用跨域设置，默认不启用",
	}, cli.BoolFlag{
		Name:  "metric",
		Usage: "启用metric配置",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	}, cli.BoolFlag{
		Name:  "appconf",
		Usage: "启用appconf配置",
	},
	)
	return flags
}
func (p *APICmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return p.removeTemplate(projectPath, getBlock(c, "port", "jwt", "db", "cros", "metric", "cache", "queue", "appconf"))
}

//Action .
func (p *APICmd) action(c *cli.Context) (err error) {
	name, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//创键项目
	err = p.create(projectPath)
	if err != nil && !p.cover {
		cmds.Log.Error(err)
		return err
	}
	//创建项目
	if !p.cover {
		err = p.writeTemplate(name, projectPath, map[string]interface{}{
			"port":        util.GetPrefixString(types.GetString(c.String("p"), "8090"), ":"),
			"serverType":  "api",
			"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"jwt":         c.Bool("jwt"),
			"db":          types.GetString(c.String("db")),
			"cros":        c.Bool("cros"),
			"projectName": name,
			"cache":       c.Bool("cache"),
			"queue":       c.Bool("queue"),
			"appconf":     c.Bool("appconf"),
			"metric":      c.Bool("metric"),
		})
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("项目生成完成")
		return nil
	}

	//追加项目代码
	err = p.appendTemplate(projectPath, getBlock(c, "port", "jwt", "db", "cros", "cache", "queue", "appconf", "metric"), map[string]interface{}{
		"port":        util.GetPrefixString(types.GetString(c.String("p"), "8090"), ":"),
		"serverType":  "api",
		"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"jwt":         c.Bool("jwt"),
		"db":          types.GetString(c.String("db")),
		"cros":        c.Bool("cros"),
		"projectName": name,
		"cache":       c.Bool("cache"),
		"queue":       c.Bool("queue"),
		"appconf":     c.Bool("appconf"),
		"metric":      c.Bool("metric"),
	})
	if err != nil {
		cmds.Log.Error(err)
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

	if path.Exists(filepath.Join(projectPath, "main.go")) {
		return fmt.Errorf("项目文件夹(%s)不为空", projectPath)
	}

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
		cmds.Log.Info("生成文件:", fpath)
	}
	return nil
}

func (p *APICmd) appendTemplate(projectPath string, block []string, input map[string]interface{}) error {

	if !path.Exists(filepath.Join(projectPath, "main.go")) ||
		!path.Exists(filepath.Join(projectPath, "init.go")) ||
		!path.Exists(filepath.Join(projectPath, "handling.go")) ||
		!path.Exists(filepath.Join(projectPath, "install.dev.go")) {
		return fmt.Errorf("项目文件(%s)不存在，请先创建", projectPath)
	}

	data, err := tmpls.GetConfTmpls(block, input)
	if err != nil {
		return err
	}
	for k, v := range data {
		for key, val := range v {
			err = datas.ReplaceFileStr(key, filepath.Join(projectPath, k), val)
			if err != nil {
				cmds.Log.Error(err)
				continue
			}
		}

		cmds.Log.Info("生成文件:", filepath.Join(projectPath, k))
	}
	return nil
}

func (p *APICmd) removeTemplate(projectPath string, block []string) error {

	if !path.Exists(filepath.Join(projectPath, "main.go")) ||
		!path.Exists(filepath.Join(projectPath, "init.go")) ||
		!path.Exists(filepath.Join(projectPath, "handling.go")) ||
		!path.Exists(filepath.Join(projectPath, "install.dev.go")) {
		return fmt.Errorf("项目文件(%s)不存在，请先创建", projectPath)
	}

	data, err := tmpls.GetEmptyTmpls(block)
	if err != nil {
		return err
	}
	for k, v := range data {
		for key, val := range v {
			err = datas.ReplaceFileStr(key, filepath.Join(projectPath, k), val)
			if err != nil {
				cmds.Log.Error(err)
				continue
			}
		}

		cmds.Log.Info("生成文件:", filepath.Join(projectPath, k))
	}
	return nil
}
func getBlock(c *cli.Context, b ...string) []string {
	blocks := make([]string, 0, 2)
	for _, v := range b {
		if c.Bool(v) {
			blocks = append(blocks, v)
		} else {
			if c.IsSet(v) {
				blocks = append(blocks, v)
			}

			// if v == "port" && c.String("port") != ":8090" {
			// 	blocks = append(blocks, v)
			// }
			// if v == "db" && c.String("db") != "" {
			// 	blocks = append(blocks, v)
			// }
		}

	}
	return blocks
}
