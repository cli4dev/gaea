package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//GetGoPath 获取$GOPATH路径
func GetGoPath() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	return path[0], nil
}

func getMDFileList(path string) (mdListfile []string, err error) {

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		//排除指定文件夹
		if strings.Contains(path, ".git") || strings.Contains(path, "vendor") || strings.Contains(path, "README") {
			return nil
		}
		//用strings.HasSuffix(src, suffix)
		//判断src中是否包含 suffix结尾
		ok := strings.HasSuffix(path, ".md")
		if ok {
			//将目录push到listfile []string中
			mdListfile = append(mdListfile, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return mdListfile, nil
}

//getMDPathRec .
func getMDPathRec(path string) []string {
	s, _ := getMDFileList(path)
	if len(s) == 0 {
		return getMDPathRec("../" + path)
	}
	return s
}

//GetMDPath 获取*.md文件
func GetMDPath() []string {
	return getMDPathRec(".")
}

//GetModulePath .
//path  有modules的目录   比如 /github.com/micro-plat/gaea/modules/
func GetModulePath() (path string) {
	pwd, _ := os.Getwd()
	//在modules里面
	if strings.Contains(pwd, "modules") {
		comma := strings.Index(pwd, "modules")
		return strings.Join([]string{pwd[:comma-1], "/modules"}, "")
	}
	//不在modules里面
	return getModulesPath()
}

func getModulesPath() string {
	p := getModulesPathRec(".")
	return p[0]
}

func getModulesPathRec(path string) []string {
	s, _ := getModulesList(path)
	if len(s) == 0 {
		return getModulesPathRec("../" + path)
	}
	return s
}

func getModulesList(path string) (modulesfile []string, err error) {

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {

		if f == nil {
			return err
		}
		//排除指定文件夹
		if !strings.Contains(path, "modules") {
			return nil
		}
		modulesfile = append(modulesfile, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return modulesfile, nil
}
