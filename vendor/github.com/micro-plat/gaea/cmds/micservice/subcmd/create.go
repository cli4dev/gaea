package subcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/micro-plat/gaea/cmds/micservice/tpl"
	"github.com/micro-plat/gaea/cmds/util/data"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/path"
	"github.com/urfave/cli"
)

//MicService .
type MicService struct {
	serverName     string
	moduleMethods  []string
	serverHandlers []string
	servicesPath   string
	modulesPath    string
}

//NewMicServiceCmd .
func NewMicServiceCmd() cli.Command {

	m := &MicService{}

	return cli.Command{
		Name:  "micservice",
		Usage: "创建 micservice",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建 micservice",
				Flags:  m.getCreateFlags(),
				Action: m.action,
			},
		},
	}
}

func (m *MicService) getCreateFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "m,methods",
		Usage: `生成modules中获得函数,参数示例："Query,Save"`,
	}, cli.StringFlag{
		Name:  "s",
		Usage: `生成服务层的函数，参数示例："Get,Delete"`,
	})
	return flags
}

func (m *MicService) action(c *cli.Context) (err error) {

	//初始化参数
	err = m.initMicServiceNameAndParams(c)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	//创建服务
	err = m.create()
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	cmds.Log.Info("生成完成！")
	return nil
}

func (m *MicService) initMicServiceNameAndParams(c *cli.Context) (err error) {
	//获取参数
	if c.NArg() == 0 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return fmt.Errorf("未指定服务名称")
	}
	m.serverName = strings.Trim(c.Args().First(), "/")
	m.moduleMethods = strings.Split(c.String("m"), ",")
	m.serverHandlers = strings.Split(c.String("s"), ",")
	m.servicesPath = path.GetServerPath()
	m.modulesPath = path.GetModulePath()
	return nil
}

func (m *MicService) create() (err error) {
	//创建文件
	err = m.createFile()
	if err != nil {
		return err
	}

	//注册路由
	err = m.registerRouter()
	if err != nil {
		return err
	}
	//创建moduleFunc
	if m.moduleMethods[0] != "" {
		err = m.createModulesFunc()
		if err != nil {
			return err
		}
	}
	//创建serverHandlerFunc
	if m.serverHandlers[0] != "" {
		err = m.createServiceFunc()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MicService) createFile() error {
	if b := path.Exists(filepath.Join(m.modulesPath, m.serverName+".go")); b {
		return fmt.Errorf("文件 %s 已存在", filepath.Join(m.modulesPath, m.serverName+".go"))
	}
	mf, err := path.CreatePath(filepath.Join(m.modulesPath, m.serverName+".go"), true)
	defer mf.Close()
	if err != nil {
		return err
	}
	mStr, err := data.Translate("module file", tpl.ModuleTmpl, map[string]interface{}{
		"moduleName": m.serverName,
	})
	if err != nil {
		return err
	}
	_, err = mf.WriteString(mStr)
	if err != nil {
		return err
	}

	sf, err := path.CreatePath(filepath.Join(m.servicesPath, m.serverName+".go"), true)
	defer sf.Close()
	if err != nil {
		return err
	}

	sStr, err := data.Translate("service file", tpl.ServiceTmpl, map[string]interface{}{
		"moduleName": m.serverName,
	})
	if err != nil {
		return err
	}
	_, err = sf.WriteString(sStr)
	if err != nil {
		return err
	}
	return nil
}

func (m *MicService) createModulesFunc() error {
	for _, v := range m.moduleMethods {
		f, err := path.CreatePath(filepath.Join(m.modulesPath, m.serverName+".go"), true)
		defer f.Close()
		if err != nil {
			return err
		}
		fStr, err := data.Translate("moduleFunc file", tpl.ModuleFuncTmpl, map[string]interface{}{
			"funcName":   v,
			"moduleName": m.serverName,
		})
		if err != nil {
			return err
		}
		_, err = f.WriteString(fStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MicService) createServiceFunc() error {
	for _, v := range m.serverHandlers {
		f, err := path.CreatePath(filepath.Join(m.servicesPath, m.serverName+".go"), true)
		defer f.Close()
		if err != nil {
			return err
		}
		fStr, err := data.Translate("serviceFunc file", tpl.ServiceFuncTmpl, map[string]interface{}{
			"funcName":   v,
			"moduleName": m.serverName,
		})
		if err != nil {
			return err
		}
		_, err = f.WriteString(fStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MicService) registerRouter() error {

	srcf, err := os.OpenFile("./init.go", os.O_RDWR, os.ModePerm)
	defer srcf.Close()
	if err != nil {
		err = fmt.Errorf("无法打开文件:%s(err:%v)", ".init.go", err)
		return err
	}
	buf, err := ioutil.ReadAll(srcf)
	if err != nil {
		cmds.Log.Errorf("%v", err.Error())
		return err
	}
	result := string(buf)

	k := fmt.Sprintf(`//%s#//[\s\S]+//#%s//`, "service.router", "service.router")
	if ok, _ := regexp.Match(k, buf); !ok {
		cmds.Log.Errorf("没有找到配置定位标识符:%s//%s#//....//#%s//", "./init.go", "service.router", "service.router")
		return nil
	}
	re, _ := regexp.Compile(k)
	str := re.Find(buf)
	re.FindStringSubmatch(result)
	s := strings.Replace(string(str), "//service.router#//", "", -1)
	s = strings.Replace(s, "//#service.router//", "", -1)

	if !strings.Contains(s, m.serverName) {

		routerStr, _ := data.Translate("router file", `r.Micro("{{.routerPath}}",{{.head}}.New{{.moduleName|lName|humpName}}Handler,"*")`, map[string]interface{}{
			"moduleName": m.serverName,
			"routerPath": "/" + m.serverName,
			"head":       strings.Split(m.serverName, "/")[0],
		})
		s += routerStr + "\n\t\t"
	}

	s = "//service.router#//" + s + "//#service.router//"

	fstr := re.ReplaceAllString(result, s)
	//清空文件数据
	srcf.Truncate(0)
	n, _ := srcf.Seek(0, os.SEEK_SET)
	_, err = srcf.WriteAt([]byte(fstr), n)
	if err != nil {
		return err
	}
	return nil
}
