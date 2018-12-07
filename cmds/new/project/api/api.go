package api

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

//ProjectAPICmd .
type ProjectAPICmd struct {
}

//GeStartFlags .
func (p *ProjectAPICmd) GeStartFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "name,n",
		Value: "./",
		Usage: "项目名称",
	}, cli.StringFlag{
		Name:  "p",
		Value: ":8090",
		Usage: "指定服务端口号",
	}, cli.BoolFlag{
		Name:  "jwt",
		Usage: "是否启用jwt",
	}, cli.StringFlag{
		Name:  "db",
		Value: "ora:test/123456@orcl136",
		Usage: "指定数据库类型和数据库链接串",
	}, cli.BoolFlag{
		Name:  "domain",
		Usage: "是否启用跨域设置，默认不启用",
	})

	return flags
}

//Action .
func (p *ProjectAPICmd) Action(c *cli.Context) (err error) {
	name := c.String("n")
	//服务类型
	serverType := "api"

	port := c.String("p")
	name, path, err := path.GetProjectPath(name)
	if err != nil {
		return err
	}
	//创键项目
	err = p.new(name, path, serverType, port, c.String("db"), c.Bool("jwt"), c.Bool("domain"))
	if err != nil {
		return err
	}

	//cmds.Log.Info("项目生成完成")
	return nil
}

func (p *ProjectAPICmd) new(projectName, projectPath, serviceType, port, db string, jwt, domain bool) (err error) {

	if pathExists(filepath.Join(projectPath, "main.go")) {
		err = fmt.Errorf("项目已经存在")
		cmds.Log.Error(err)
		return err
	}

	tmpls, err := tmpls.GetTmpls(projectName, serviceType, port, db, jwt, domain)
	if err != nil {
		return err
	}
	err = p.createProject(projectPath, tmpls)
	if err != nil {
		return err
	}
	return nil

}

func (p *ProjectAPICmd) createProject(projectPath string, data map[string]string) error {
	for k, v := range data {
		path := filepath.Join(projectPath, k)
		dir := filepath.Dir(path)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}

		}

		_, err = os.Stat(path)
		var srcf *os.File
		if os.IsNotExist(err) {
			// err := os.MkdirAll(dir, os.ModePerm)
			// if err != nil {
			// 	err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
			// 	return err
			// }
			srcf, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
			if err != nil {
				err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
				return err
			}

		} else {
			srcf, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, os.ModePerm)
			if err != nil {
				err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
				return err
			}
		}

		if v == "dir" || strings.HasPrefix(k, "conf") {
			continue
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

func replaceStr(name, projectPath, fileName, value string) error {

	srcf, err := os.OpenFile(filepath.Join(projectPath, fileName), os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer srcf.Close()
	if err != nil {
		err = fmt.Errorf("无法打开文件:%s(err:%v)", filepath.Join(projectPath, fileName), err)
		return err
	}
	buf, err := ioutil.ReadAll(srcf)
	if err != nil {
		cmds.Log.Errorf("%v", err.Error())
		return err
	}

	result := string(buf)

	k := fmt.Sprintf(`//%s#//[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|a-zA-Z0-9]+//#%s//`, name, name)

	if ok, _ := regexp.Match(k, buf); !ok {
		cmds.Log.Error("没有找到配置定位标识，请手动添加")
		return nil
	}

	re, _ := regexp.Compile(k)

	str := re.ReplaceAllString(result, value)

	n, _ := srcf.Seek(0, os.SEEK_SET)
	_, err = srcf.WriteAt([]byte(str), n)
	if err != nil {
		return err
	}
	return nil
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}
