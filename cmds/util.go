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

var listfile []string    //获取文件列表
var modulesfile []string //获取文件列表
//Listfunc .
func Listfunc(path string, f os.FileInfo, err error) error {
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
		listfile = append(listfile, path)
	}
	return nil
}

func getFileList(path string) string {
	//var strRet string
	err := filepath.Walk(path, Listfunc) //

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return " "
}

//ListFileFunc .
func ListFileFunc(p []string) []string {
	s := []string{}
	for _, value := range p {
		s = append(s, value)
	}
	return s
}

//getMDPathRec .
func getMDPathRec(path string) []string {
	getFileList(path)
	s := ListFileFunc(listfile)
	if len(s) == 0 {
		return getMDPathRec("../" + path)
	}
	return s
}

//GetMDPath 获取*.md文件
func GetMDPath() []string {
	return getMDPathRec(".")
}

//GetLocation .
//path  有modules的目录   比如 /github.com/micro-plat/gaea/modules/
func GetLocation() (path string) {
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
	getModulesList(path)
	s := ListFileFunc(modulesfile)
	if len(s) == 0 {
		return getModulesPathRec("../" + path)
	}
	return s
}

func getModulesList(path string) string {
	//var strRet string
	err := filepath.Walk(path, ListModulesfunc) //

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return " "
}

//ListModulesfunc .
func ListModulesfunc(path string, f os.FileInfo, err error) error {

	if f == nil {
		return err
	}
	//排除指定文件夹
	if !strings.Contains(path, "modules") {
		return nil
	}

	//将目录push到listfile []string中
	modulesfile = append(modulesfile, path)

	return nil
}
