package subcmd

import (
	"github.com/micro-plat/gaea/cmds/new/util"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

//WSCmd .
type WSCmd struct {
	cli.Command
	cover bool
}

//NewWSCmd .
func NewWSCmd() cli.Command {

	wsCreator := &WSCmd{cover: false}
	wsCover := &WSCmd{cover: true}

	return cli.Command{
		Name:  "ws",
		Usage: "创建ws项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建ws项目",
				Flags:  wsCreator.getStartFlags(),
				Action: wsCreator.action,
			},
			cli.Command{
				Name:   "cover",
				Usage:  "根据指定参数覆盖已有 ws 配置",
				Flags:  wsCover.getStartFlags(),
				Action: wsCover.action,
			}, cli.Command{
				Name:   "remove",
				Usage:  "移除指定ws配置",
				Flags:  wsCover.getRemoveStartFlags(),
				Action: wsCover.removeAction,
			},
		},
	}
}

//getStartFlags .
func (p *WSCmd) getStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "appconf,ws.appconf",
		Usage: "ws app配置",
	}, cli.BoolFlag{
		Name:  "jwt,ws.jwt",
		Usage: "ws jwt 配置",
	}, cli.BoolFlag{
		Name:  "metric,ws.metric",
		Usage: "ws metric 配置",
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
func (p *WSCmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "appconf,ws.appconf",
		Usage: "ws port配置",
	}, cli.BoolFlag{
		Name:  "jwt,ws.jwt",
		Usage: "ws jwt 配置",
	}, cli.BoolFlag{
		Name:  "metric,ws.metric",
		Usage: "ws metric 配置",
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

func (p *WSCmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return removeTemplate(projectPath, getBlock(c, "db", "queue", "cache", "ws.appconf", "ws.jwt", "ws.metric"))
}

//Action .
func (p *WSCmd) action(c *cli.Context) (err error) {
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
			"serverType":  "ws",
			"projectName": name,
			"ws.appconf":  c.Bool("appconf"),
			"ws.jwt":      c.Bool("jwt"),
			"ws.metric":   c.Bool("metric"),
			"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"db":          types.GetString(c.String("db")),
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
	err = appendTemplate(projectPath, getBlock(c, "db", "queue", "cache", "ws.appconf", "ws.jwt", "ws.metric"), map[string]interface{}{
		"serverType":  "ws",
		"projectName": name,
		"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"db":          types.GetString(c.String("db")),
		"ws.port":     c.Bool("appconf"),
		"ws.jwt":      c.Bool("jwt"),
		"ws.metric":   c.Bool("metric"),
	})
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	err = addConf2Main(projectPath, "ws")
	if err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}
