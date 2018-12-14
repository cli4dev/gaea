package subcmd

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util"
	"github.com/micro-plat/gaea/cmds/util/path"

	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//APICmd .
type APICmd struct {
	cli.Command
	cover bool
}

//NewAPICmd .
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
		Name:  "port,p,api.port",
		Value: ":9090",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt,api.jwt,handling.jwt",
		Usage: "是否启用jwt",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cros,api.cros",
		Usage: "是否启用跨域设置，默认不启用",
	}, cli.BoolFlag{
		Name:  "metric,api.metric",
		Usage: "启用metric配置",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	}, cli.BoolFlag{
		Name:  "appconf,api.appconf",
		Usage: "启用appconf配置",
	}, cli.StringFlag{
		Name:  "login",
		Usage: "启用 login 模块,需要输入 sso 地址",
	}, cli.BoolFlag{
		Name:  "menu",
		Usage: "启用 menu 模块",
	})
	return flags
}
func (p *APICmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "port,p,api.port",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt,api.jwt,handling.jwt",
		Usage: "是否启用jwt",
	}, cli.BoolFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cros,api.cros",
		Usage: "是否启用跨域设置，默认不启用",
	}, cli.BoolFlag{
		Name:  "metric,api.metric",
		Usage: "启用metric配置",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	}, cli.BoolFlag{
		Name:  "appconf,api.appconf",
		Usage: "启用appconf配置",
	}, cli.BoolFlag{
		Name:  "login",
		Usage: "移除登录模块",
	}, cli.BoolFlag{
		Name:  "menu",
		Usage: "移除菜单模块",
	})
	return flags
}

func (p *APICmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return removeTemplate(projectPath, getBlock(c, "api.port", "api.jwt", "handling.jwt", "db", "api.cros", "api.metric", "cache", "queue", "api.appconf"))
}

//Action .
func (p *APICmd) action(c *cli.Context) (err error) {
	name, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//创键项目
	err = create(projectPath)
	if err != nil && !p.cover {
		cmds.Log.Error(err)
		return err
	}
	//创建项目
	if !p.cover {
		err = writeTemplate(p.cover, name, projectPath, map[string]interface{}{
			"port":        util.GetPrefixString(types.GetString(c.String("p"), "9090"), ":"),
			"serverType":  "api",
			"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"jwt":         c.Bool("jwt"),
			"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
			"cros":        c.Bool("cros"),
			"projectName": name,
			"cache":       c.Bool("cache"),
			"queue":       c.Bool("queue"),
			"appconf":     c.Bool("appconf"),
			"metric":      c.Bool("metric"),
			"login":       c.String("login"),
		})
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("项目生成完成")
		return nil
	}

	//追加项目代码
	err = appendTemplate(projectPath, getBlock(c, "api.port", "api.jwt", "handling.jwt", "db", "api.cros", "api.metric", "cache", "queue", "api.appconf"), map[string]interface{}{
		"port":        util.GetPrefixString(types.GetString(c.String("p"), "9090"), ":"),
		"serverType":  "api",
		"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"jwt":         c.Bool("jwt"),
		"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
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
	err = addConf2Main(projectPath, "api")
	if err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}
