package subcmd

import (
	"github.com/micro-plat/gaea/cmds/util"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/path"
	"github.com/urfave/cli"
)

//CronCmd .
type CronCmd struct {
	cli.Command
	cover bool
}

//NewCronCmd .
func NewCronCmd() cli.Command {

	cronCreator := &CronCmd{cover: false}
	cronCover := &CronCmd{cover: true}

	return cli.Command{
		Name:  "cron",
		Usage: "创建cron项目",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建cron项目",
				Flags:  cronCreator.getStartFlags(),
				Action: cronCreator.action,
			},
			cli.Command{
				Name:   "cover",
				Usage:  "根据指定参数覆盖已有 cron 配置",
				Flags:  cronCover.getStartFlags(),
				Action: cronCover.action,
			}, cli.Command{
				Name:   "remove",
				Usage:  "移除指定cron配置",
				Flags:  cronCover.getRemoveStartFlags(),
				Action: cronCover.removeAction,
			},
		},
	}
}

//getStartFlags .
func (p *CronCmd) getStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "appconf,cron.appconf",
		Usage: "cron app配置",
	}, cli.BoolFlag{
		Name:  "task,cron.task",
		Usage: "cron task 配置",
	}, cli.BoolFlag{
		Name:  "metric,cron.metric",
		Usage: "cron metric 配置",
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
func (p *CronCmd) getRemoveStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "appconf,cron.appconf",
		Usage: "cron app配置",
	}, cli.BoolFlag{
		Name:  "task,cron.task",
		Usage: "cron task 配置",
	}, cli.BoolFlag{
		Name:  "metric,cron.metric",
		Usage: "cron metric 配置",
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

func (p *CronCmd) removeAction(c *cli.Context) (err error) {
	_, projectPath, err := path.GetProjectPath(c.String("n"))
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	return removeTemplate(projectPath, getBlock(c, "db", "queue", "cache", "cron.appconf", "cron.task", "cron.metric"))
}

//Action .
func (p *CronCmd) action(c *cli.Context) (err error) {
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
			"serverType":   "cron",
			"projectName":  name,
			"cron.appconf": c.Bool("cron.appconf"),
			"cron.task":    c.Bool("cron.task"),
			"cron.metric":  c.Bool("cron.metric"),
			"dbname":       util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
			"db":           util.GetRightString(types.GetString(c.String("db")), ":", ""),
			"cache":        c.Bool("cache"),
			"queue":        c.Bool("queue"),
			"login":        "",
		})
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("项目生成完成")
		return nil
	}

	//追加项目代码
	err = appendTemplate(projectPath, getBlock(c, "db", "queue", "cache", "cron.appconf", "cron.task", "cron.metric"), map[string]interface{}{
		"serverType":   "cron",
		"projectName":  name,
		"dbname":       util.GetLeftString(types.GetString(c.String("db")), ":", "mysql"),
		"db":           util.GetRightString(types.GetString(c.String("db")), ":", ""),
		"cron.appconf": c.Bool("appconf"),
		"cron.task":    c.Bool("cron.task"),
		"cron.metric":  c.Bool("cron.metric"),
	})
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	err = addConf2Main(projectPath, "cron")
	if err != nil {
		cmds.Log.Error(err)
	}
	cmds.Log.Info("项目生成完成")
	return nil
}
