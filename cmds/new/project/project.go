package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/project/tmpls"
	"github.com/micro-plat/gaea/cmds/new/util/path"
	"github.com/urfave/cli"
)

type projectCmd struct {
}

//NewProjectCmd .
func NewProjectCmd() cli.Command {
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
		Name:  "s",
		Value: "api", //默认值
		Usage: "服务器类型(api,rpc,web,mqc,cron),多个用短横线分隔",
	}, cli.StringSliceFlag{
		Name:  "n",
		Usage: "项目名称",
	}, cli.BoolFlag{
		Name:  "a",
		Usage: "是否追加服务类型（项目生成后，修改服务类型使用）",
	})

	return flags
}

func (p *projectCmd) action(c *cli.Context) (err error) {
	projectName := c.StringSlice("n")
	if !pathExists("main.go") || !pathExists("modules") {
		if len(projectName) == 0 {
			cmds.Log.Error(`"请输入项目名称，例如: gaea new project -n "myproject/apiserver"`)
			cmds.Log.Info("如需要添加配置，请到项目根目录值执行")
			return nil
		}
	}

	//获取服务类型
	serverType := c.String("s")

	//创键项目
	for _, v := range projectName {
		err := p.new(v, serverType)
		if err != nil {
			cmds.Log.Error(err)
		}
	}

	//追加配置
	if c.Bool("a") {
		err = p.addConf(serverType)
		if err != nil {
			return err
		}
	}

	cmds.Log.Info("项目生成完成")
	return nil
}

func (p *projectCmd) new(projectName string, serviceType string) error {
	gopath, err := path.GetGoPath()
	if err != nil {
		return err
	}
	projectPath := filepath.Join(gopath, "src", projectName)
	tmpls, err := tmpls.GetTmpls(projectName, serviceType)
	if err != nil {
		return err
	}
	err = p.createProject(projectPath, tmpls)
	if err != nil {
		return err
	}
	//写入配置
	return p.writeConf(projectPath, tmpls)

}
func (p *projectCmd) addConf(serviceType string) error {
	tmpls, err := tmpls.GetConfTmpls(serviceType)
	if err != nil {
		return err
	}
	//向main.go添加服务类型
	err = p.addConf2Main(serviceType)
	if err != nil {
		return err
	}
	//写入配置
	return p.writeConf(".", tmpls)

}

func (p *projectCmd) createProject(projectPath string, data map[string]string) error {
	for k, v := range data {
		path := filepath.Join(projectPath, k)
		dir := filepath.Dir(path)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}
		}
		if v == "dir" || strings.HasPrefix(k, "conf") {
			continue
		}
		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
			return err
		}

		_, err = srcf.WriteString(v)
		if err != nil {
			return err
		}
		srcf.Close()
		cmds.Log.Info("生成文件:", path)
	}
	return nil

}

func (p *projectCmd) writeConf(projectPath string, data map[string]string) error {
	for k, v := range data {
		if strings.HasPrefix(k, "conf.dev") {
			f, err := os.OpenFile(projectPath+"/install.dev.go", os.O_WRONLY, os.ModePerm)
			if err != nil {
				return fmt.Errorf("打开文件install.dev.go失败. err: " + err.Error())
			}
			n, err := f.Seek(0, os.SEEK_END)
			if err != nil {
				return err
			}
			_, err = f.WriteAt([]byte(v), n-2)
			if err != nil {
				return err
			}
			defer f.Close()
			cmds.Log.Info("写入dev配置成功")
		}
		if strings.HasPrefix(k, "conf.prod") {
			f, err := os.OpenFile(projectPath+"/install.prod.go", os.O_WRONLY, os.ModePerm)
			if err != nil {
				return fmt.Errorf("打开文件install.prod.go失败. err: " + err.Error())
			}
			n, err := f.Seek(0, os.SEEK_END)
			if err != nil {
				return err
			}
			_, err = f.WriteAt([]byte(v), n-2)
			if err != nil {
				return err
			}
			defer f.Close()
			cmds.Log.Info("写入prod配置成功")
		}
	}
	return nil
}

//addConf2Main 向main.go添加服务类型
func (p *projectCmd) addConf2Main(serviceType string) error {

	srcf, err := os.OpenFile("./main.go", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("无法打开文件:%s(err:%v)", "./main.go", err)
		return err
	}
	buf, err := ioutil.ReadAll(srcf)
	if err != nil {
		cmds.Log.Errorf("%v", err.Error())
		return err
	}

	result := string(buf)
	//解析正则表达式
	re, err := regexp.Compile("hydra.WithServerTypes(.*),") //WithServerTypes("api")
	if err != nil {
		return err
	}
	//Find函数返回匹配的第一个字符串
	srcStr := string(re.Find([]byte(result)))
	first := strings.Index(srcStr, "\"")
	last := strings.LastIndex(srcStr, "\"")
	old := srcStr[first+1 : last]
	for _, v := range strings.Split(serviceType, "-") {
		if strings.Contains(old, v) {
			cmds.Log.Errorf("服务已经存在：%s", serviceType)
			return fmt.Errorf("服务已经存在")
		}
	}
	new := old + "-" + serviceType
	newStr := srcStr[:first+1] + new + srcStr[last:]
	newContent := strings.Replace(result, srcStr, newStr, -1)
	n, _ := srcf.Seek(0, os.SEEK_SET)
	_, err = srcf.WriteAt([]byte(newContent), n)
	if err != nil {
		return err
	}
	defer srcf.Close()
	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
