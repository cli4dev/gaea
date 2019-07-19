package subcmd

import (
	"github.com/micro-plat/gaea/cmds/util"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/path"
	"github.com/urfave/cli"
)

//MQCCmd .
type MQCCmd struct {
	cli.Command
	cover bool
}

//NewMQCCmd .
func NewMQCCmd() cli.Command {

	mqcCreator := &MQCCmd{cover: false}
	mqcCover := &MQCCmd{cover: true}

	return cli.Command{
		Name:  "mqc",
		Usage: "创建mqc项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建mqc项目",
				Flags:  mqcCreator.getStartFlags(),
				Action: mqcCreator.action,
			},
			cli.Command{
				Name:   "cover",
				Usage:  "根据指定参数覆盖已有 mqc 配置",
				Flags:  mqcCover.getStartFlags(),
				Action: mqcCover.action,
			}, cli.Command{
				Name:   "remove",
				Usage:  "移除指定mqc配置",
				Flags:  mqcCover.getRemoveStartFlags(),
				Action: mqcCover.removeAction,
			},
		},
	}
}

//getStartFlags .
func (p *MQCCmd) getStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "server,mqc.server",
		Usage: "mqc server配置",
	}, cli.BoolFlag{
		Name:  "queue,mqc.queue",
		Usage: "mqc queue 配置",
	}, cli.BoolFlag{
		Name:  "metric,mqc.metric",
		Usage: "mqc metric 配置",
	}, cli.StringFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	})
	return flags
}
func (p *MQCCmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "queue,mqc.queue",
		Usage: "mqc queue配置",
	}, cli.BoolFlag{
		Name:  "server,mqc.server",
		Usage: "mqc server 配置",
	}, cli.BoolFlag{
		Name:  "metric,mqc.metric",
		Usage: "mqc metric 配置",
	}, cli.BoolFlag{
		Name:  "db",
		Usage: "指定数据库类型和数据库链接串(ora:test/123456@orcl136)",
	}, cli.BoolFlag{
		Name:  "cache",
		Usage: "启用cache配置",
	})
	return flags
}

func (p *MQCCmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return removeTemplate(projectPath, getBlock(c, "db", "queue", "cache", "mqc.server", "mqc.queue", "mqc.metric"))
}

//Action .
func (p *MQCCmd) action(c *cli.Context) (err error) {
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
			"serverType":  "mqc",
			"projectName": name,
			"mqc.server":  c.Bool("mqc.server"),
			"mqc.queue":   c.Bool("mqc.queue"),
			"mqc.metric":  c.Bool("mqc.metric"),
			"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
			"cache":       c.Bool("cache"),
			"queue":       c.Bool("queue"),
			"login":       "",
		})
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("项目生成完成")
		return nil
	}

	//追加项目代码
	err = appendTemplate(projectPath, getBlock(c, "db", "queue", "cache", "mqc.server", "mqc.queue", "mqc.metric"), map[string]interface{}{
		"serverType":  "mqc",
		"projectName": name,
		"dbname":      util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"db":          util.GetRightString(types.GetString(c.String("db")), ":", ""),
		"mqc.server":  c.Bool("mqc.server"),
		"mqc.queue":   c.Bool("mqc.queue"),
		"mqc.metric":  c.Bool("mqc.metric"),
		"queue":       c.Bool("queue"),
		"cache":       c.Bool("cache"),
	})
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	err = addConf2Main(projectPath, "mqc")
	if err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}
