package project

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/project/subcmd"

	"github.com/urfave/cli"
)

type projectCmd struct {
}

//NewProjectCmd .
func NewProjectCmd() []cli.Command {
	return []cli.Command{
		subcmd.NewAPICmd(),
		subcmd.NewCronCmd(),
		subcmd.NewWebCmd(),
		subcmd.NewMQCCmd(),
		subcmd.NewRPCCmd(),
		subcmd.NewWSCmd(),
		subcmd.NewVueCmd(),
	}
}

func init() {
	cmds.Register(NewProjectCmd()...)
}

// func (p *projectCmd) geStartFlags() []cli.Flag {
// 	flags := make([]cli.Flag, 0, 4)
// 	flags = append(flags, cli.StringFlag{
// 		Name:  "s",
// 		Value: "api", //默认值
// 		Usage: "服务器类型(api,rpc,web,mqc,cron),多个用短横线分隔",
// 	}, cli.StringFlag{
// 		Name:  "name,n",
// 		Value: "./",
// 		Usage: "项目名称",
// 	}, cli.BoolFlag{
// 		Name:  "a",
// 		Usage: "是否追加服务类型（项目生成后，修改服务类型使用）",
// 	}, cli.StringFlag{
// 		Name:  "p",
// 		Value: ":9090",
// 		Usage: "指定服务端口号",
// 	}, cli.BoolFlag{
// 		Name:  "jwt",
// 		Usage: "是否启用jwt",
// 	}, cli.StringFlag{
// 		Name:  "db",
// 		Value: "ora:test/123456@orcl136",
// 		Usage: "指定数据库类型和数据库链接串",
// 	}, cli.BoolFlag{
// 		Name:  "domain",
// 		Usage: "是否启用跨域设置，默认不启用",
// 	})

// 	return flags
// }

// func (p *projectCmd) action(c *cli.Context) (err error) {
// 	name := c.String("n")

// 	//获取服务类型
// 	serverType := c.String("s")

// 	port := c.String("p")
// 	name, path, err := path.GetProjectPath(name)
// 	if err != nil {
// 		return err
// 	}

// 	//创键项目
// 	err = p.new(name, path, serverType, port, c.String("db"), c.Bool("jwt"), c.Bool("domain"), c.Bool("a"))
// 	if err != nil {
// 		cmds.Log.Error(err)
// 		return err
// 	}

// 	//追加配置
// 	if c.Bool("a") {
// 		err = p.addConf(path, serverType, port, c.String("db"), c.Bool("jwt"), c.Bool("domain"))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	//cmds.Log.Info("项目生成完成")
// 	return nil
// }

// func (p *projectCmd) new(projectName, projectPath, serviceType, port, db string, jwt, domain bool, append bool) error {

// 	if pathExists(filepath.Join(projectPath, "main.go")) {
// 		return nil
// 	}

// 	tmpls, err := tmpls.GetTmpls(projectName, serviceType, port, db, jwt, domain)
// 	if err != nil {
// 		return err
// 	}
// 	err = p.createProject(projectPath, tmpls, append)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
// func (p *projectCmd) addConf(projectPath string, serviceType, port, db string, jwt, domain bool) error {
// 	tmpls, err := tmpls.GetConfTmpls(serviceType, port, db, jwt, domain)
// 	if err != nil {
// 		return err
// 	}
// 	//向main.go添加服务类型
// 	err = p.addConf2Main(projectPath, serviceType)
// 	if err != nil {
// 		return err
// 	}
// 	//写入配置
// 	return p.writeConf(".", tmpls)

// }

// func (p *projectCmd) createProject(projectPath string, data map[string]string, append bool) error {
// 	for k, v := range data {
// 		path := filepath.Join(projectPath, k)
// 		dir := filepath.Dir(path)
// 		_, err := os.Stat(dir)
// 		if os.IsNotExist(err) {
// 			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
// 				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
// 				return err
// 			}

// 		}

// 		_, err = os.Stat(path)
// 		var srcf *os.File
// 		if os.IsNotExist(err) {
// 			// err := os.MkdirAll(dir, os.ModePerm)
// 			// if err != nil {
// 			// 	err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
// 			// 	return err
// 			// }
// 			srcf, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
// 			if err != nil {
// 				err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
// 				return err
// 			}

// 		} else if !append {
// 			cmds.Log.Errorf("文件%s已存在", path)
// 			continue
// 		} else {
// 			srcf, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, os.ModePerm)
// 			if err != nil {
// 				err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
// 				return err
// 			}
// 		}

// 		if v == "dir" || strings.HasPrefix(k, "conf") {
// 			continue
// 		}
// 		_, err = srcf.WriteString(v)
// 		if err != nil {
// 			return err
// 		}
// 		srcf.Close()
// 		cmds.Log.Info("生成文件:", path)
// 	}
// 	return nil

// }

// func (p *projectCmd) writeConf(projectPath string, data map[string]string) error {
// 	for k, v := range data {

// 		err := conf(k, v, "dev", projectPath)
// 		if err != nil {
// 			return err
// 		}
// 		err = conf(k, v, "prod", projectPath)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func conf(k, v, model, projectPath string) error {

// 	srcf, err := os.OpenFile(projectPath+"/install."+model+".go", os.O_CREATE|os.O_RDWR, os.ModePerm)
// 	defer srcf.Close()
// 	if err != nil {
// 		err = fmt.Errorf("无法打开文件:%s(err:%v)", "./install.dev.go", err)
// 		return err
// 	}
// 	buf, err := ioutil.ReadAll(srcf)
// 	if err != nil {
// 		cmds.Log.Errorf("%v", err.Error())
// 		return err
// 	}

// 	result := string(buf)

// 	if ok, _ := regexp.Match(k, buf); !ok {
// 		cmds.Log.Error("没有找到配置定位标识，请手动添加")
// 		return nil
// 	}

// 	re, _ := regexp.Compile(k)

// 	str := re.ReplaceAllString(result, v)

// 	n, _ := srcf.Seek(0, os.SEEK_SET)
// 	_, err = srcf.WriteAt([]byte(str), n)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //addConf2Main 向main.go添加服务类型
// func (p *projectCmd) addConf2Main(projectPath string, serviceType string) error {

// 	srcf, err := os.OpenFile(filepath.Join(projectPath, "./main.go"), os.O_CREATE|os.O_RDWR, os.ModePerm)
// 	if err != nil {
// 		err = fmt.Errorf("无法打开文件:%s(err:%v)", "./main.go", err)
// 		return err
// 	}
// 	buf, err := ioutil.ReadAll(srcf)
// 	if err != nil {
// 		cmds.Log.Errorf("%v", err.Error())
// 		return err
// 	}

// 	result := string(buf)
// 	//解析正则表达式
// 	re, err := regexp.Compile("hydra.WithServerTypes(.*),") //WithServerTypes("api")
// 	if err != nil {
// 		return err
// 	}
// 	//Find函数返回匹配的第一个字符串
// 	srcStr := string(re.Find([]byte(result)))
// 	first := strings.Index(srcStr, "\"")
// 	last := strings.LastIndex(srcStr, "\"")
// 	old := srcStr[first+1 : last]
// 	for _, v := range strings.Split(serviceType, "-") {
// 		if strings.Contains(old, v) {
// 			cmds.Log.Errorf("main.go 中服务已经存在：%s", serviceType)
// 			return fmt.Errorf("服务已经存在")
// 		}
// 	}
// 	new := old + "-" + serviceType
// 	newStr := srcStr[:first+1] + new + srcStr[last:]
// 	newContent := strings.Replace(result, srcStr, newStr, -1)
// 	n, _ := srcf.Seek(0, os.SEEK_SET)
// 	_, err = srcf.WriteAt([]byte(newContent), n)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcf.Close()
// 	return nil
// }

// func pathExists(path string) bool {
// 	if _, err := os.Stat(path); err != nil {
// 		return !os.IsNotExist(err)
// 	}
// 	return true
// }

// func replaceStr(name, projectPath, fileName, value string) error {

// 	srcf, err := os.OpenFile(filepath.Join(projectPath, fileName), os.O_CREATE|os.O_RDWR, os.ModePerm)
// 	defer srcf.Close()
// 	if err != nil {
// 		err = fmt.Errorf("无法打开文件:%s(err:%v)", filepath.Join(projectPath, fileName), err)
// 		return err
// 	}
// 	buf, err := ioutil.ReadAll(srcf)
// 	if err != nil {
// 		cmds.Log.Errorf("%v", err.Error())
// 		return err
// 	}

// 	result := string(buf)

// 	k := fmt.Sprintf(`//%s#//[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|a-zA-Z0-9]+//#%s//`, name, name)

// 	if ok, _ := regexp.Match(k, buf); !ok {
// 		cmds.Log.Error("没有找到配置定位标识，请手动添加")
// 		return nil
// 	}

// 	re, _ := regexp.Compile(k)

// 	str := re.ReplaceAllString(result, value)

// 	n, _ := srcf.Seek(0, os.SEEK_SET)
// 	_, err = srcf.WriteAt([]byte(str), n)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
