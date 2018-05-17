package cmds

import (
	"os"
	"path/filepath"

	"github.com/micro-plat/lib4go/logger"
	"github.com/urfave/cli"
)

var cmds []cli.Command

//Log  日志组件
var Log *logger.Logger

func init() {
	cmds = make([]cli.Command, 0, 4)
	Log = logger.GetSession("gaea", logger.CreateSession())
}

//Registry 注册子
func Registry(c cli.Command) {
	cmds = append(cmds, c)
}

//VERSION 版本号
var VERSION = "1.0.0"

//Start 启动应用程序
func Start() {
	defer logger.Close()
	app := NewCliApp()
	if err := app.Run(os.Args); err != nil {
		return
	}
}

//NewCliApp 创建app
func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = VERSION
	app.Usage = "hydra框架辅助工具，用于管理基于hydra的项目"
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help,h",
		Usage: "查看帮助信息",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version,v",
		Usage: "查看版本信息",
	}
	app.Commands = cmds
	return app
}
