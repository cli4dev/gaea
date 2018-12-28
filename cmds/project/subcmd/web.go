package subcmd

import (
	"github.com/micro-plat/gaea/cmds/util"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/path"
	"github.com/urfave/cli"
)

//WebCmd .
type WebCmd struct {
	cli.Command
	cover bool
}

//NewWebCmd .
func NewWebCmd() cli.Command {

	webCreator := &WebCmd{cover: false}
	webCover := &WebCmd{cover: true}

	return cli.Command{
		Name:  "web",
		Usage: "创建web项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建web项目",
				Flags:  webCreator.getStartFlags(),
				Action: webCreator.action,
			},
			cli.Command{
				Name:   "cover",
				Usage:  "根据指定参数覆盖已有 web 配置",
				Flags:  webCover.getStartFlags(),
				Action: webCover.action,
			}, cli.Command{
				Name:   "remove",
				Usage:  "移除指定web配置",
				Flags:  webCover.getRemoveStartFlags(),
				Action: webCover.removeAction,
			},
		},
	}
}

//getStartFlags .
func (p *WebCmd) getStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.StringFlag{
		Name:  "port,p,web.port",
		Usage: "web port配置",
	}, cli.BoolFlag{
		Name:  "static,web.static",
		Usage: "web static 配置",
	}, cli.BoolFlag{
		Name:  "metric,web.metric",
		Usage: "web metric 配置",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	})
	return flags
}
func (p *WebCmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "port,p,web.port",
		Usage: "web port配置",
	}, cli.BoolFlag{
		Name:  "static,web.static",
		Usage: "web static 配置",
	}, cli.BoolFlag{
		Name:  "metric,web.metric",
		Usage: "web metric 配置",
	}, cli.BoolFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	}, cli.BoolFlag{
		Name:  "queue",
		Usage: "启用queue配置",
	})
	return flags
}

func (p *WebCmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return removeTemplate(projectPath, getBlock(c, "db", "queue", "cache", "web.port", "web.static", "web.metric"))
}

//Action .
func (p *WebCmd) action(c *cli.Context) (err error) {
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
			"serverType":  "web",
			"projectName": name,
			"port":        util.GetPrefixString(types.GetString(c.String("p"), "9090"), ":"),
			"web.static":  c.Bool("web.static"),
			"web.metric":  c.Bool("web.metric"),
			"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
			"cache":       c.Bool("cache"),
			"queue":       c.Bool("queue"),
		})
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("项目生成完成")
		return nil
	}

	//追加项目代码
	err = appendTemplate(projectPath, getBlock(c, "db", "queue", "cache", "web.port", "web.static", "web.metric"), map[string]interface{}{
		"serverType":  "web",
		"projectName": name,
		"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
		"port":        util.GetPrefixString(types.GetString(c.String("p"), "9090"), ":"),
		"web.static":  c.Bool("web.static"),
		"web.metric":  c.Bool("web.metric"),
	})
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	err = addConf2Main(projectPath, "web")
	if err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}
