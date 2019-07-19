package subcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	datas "github.com/micro-plat/gaea/cmds/util/data"
	"github.com/micro-plat/gaea/cmds/util/path"

	"github.com/micro-plat/gaea/cmds/project/tmpls"
	"github.com/urfave/cli"
)

func create(projectPath string) (err error) {
	if path.Exists(filepath.Join(projectPath, "main.go")) {
		err = fmt.Errorf("项目%s已经存在", projectPath)
		return err
	}
	return nil
}

func writeTemplate(cover bool, projectName string, projectPath string, input map[string]interface{}) error {

	if path.Exists(filepath.Join(projectPath, "main.go")) {
		return fmt.Errorf("项目文件夹(%s)不为空", projectPath)
	}

	data, err := tmpls.GetTmpls(projectName, input)
	if err != nil {
		return err
	}
	for k, v := range data {
		fpath := filepath.Join(projectPath, k)

		f, err := path.CreatePath(fpath, cover)
		if err != nil {
			continue
		}
		_, err = f.WriteString(v)
		if err != nil {
			continue
		}
		defer f.Close()
		cmds.Log.Info("生成文件:", fpath)
	}
	return nil
}

func appendTemplate(projectPath string, block []string, input map[string]interface{}) error {

	if !path.Exists(filepath.Join(projectPath, "main.go")) ||
		!path.Exists(filepath.Join(projectPath, "init.go")) ||
		!path.Exists(filepath.Join(projectPath, "handling.go")) ||
		!path.Exists(filepath.Join(projectPath, "install.dev.go")) {
		return fmt.Errorf("项目文件(%s)不存在，请先创建", projectPath)
	}

	data, err := tmpls.GetConfTmpls(block, input)
	if err != nil {
		return err
	}
	for k, v := range data {
		for key, val := range v {
			err = datas.ReplaceFileStr(key, filepath.Join(projectPath, k), val)
			if err != nil {
				cmds.Log.Error(err)
				continue
			}
		}

		cmds.Log.Info("生成文件:", filepath.Join(projectPath, k))
	}
	return nil
}

func removeTemplate(projectPath string, block []string) error {

	if !path.Exists(filepath.Join(projectPath, "main.go")) ||
		!path.Exists(filepath.Join(projectPath, "init.go")) ||
		!path.Exists(filepath.Join(projectPath, "handling.go")) ||
		!path.Exists(filepath.Join(projectPath, "install.dev.go")) {
		return fmt.Errorf("项目文件(%s)不存在，请先创建", projectPath)
	}

	data, err := tmpls.GetEmptyTmpls(block)
	if err != nil {
		return err
	}
	for k, v := range data {
		for key, val := range v {
			err = datas.ReplaceFileStr(key, filepath.Join(projectPath, k), val)
			if err != nil {
				cmds.Log.Error(err)
				continue
			}
		}
		cmds.Log.Info("生成文件:", filepath.Join(projectPath, k))
	}
	return nil
}

func getBlock(c *cli.Context, b ...string) []string {
	blocks := make([]string, 0, 2)
	for _, v := range b {
		if c.Bool(v) {
			blocks = append(blocks, v)
		} else {
			if c.IsSet(v) {
				blocks = append(blocks, v)
			}
		}
	}
	return blocks
}

func addConf2Main(projectPath string, serviceType string) error {

	srcf, err := os.OpenFile(filepath.Join(projectPath, "./main.go"), os.O_CREATE|os.O_RDWR, os.ModePerm)
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
			return nil
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

func writeVueTemplate(projectPath, projectName string) error {

	if path.Exists(filepath.Join(projectPath, "main.go")) {
		return fmt.Errorf("项目文件夹(%s)不为空", projectPath)
	}

	data, err := tmpls.GetVueTmpls(projectName)
	if err != nil {
		return err
	}
	for k, v := range data {
		fpath := filepath.Join(projectPath, k)

		f, err := path.CreatePath(fpath, true)
		if err != nil {
			continue
		}
		_, err = f.WriteString(v)
		if err != nil {
			continue
		}
		defer f.Close()
		cmds.Log.Info("生成文件:", fpath)
	}
	return nil
}
