package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds"
)

//GetGoPath 获取$GOPATH路径
func GetGoPath(v ...string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	if len(v) > 0 {
		npath := filepath.Join(v...)
		return filepath.Join(path[0], npath), nil
	}
	return path[0], nil
}

//GetProjectPath 获取项目路径
func GetProjectPath(path string) (string, string, error) {
	srcPath, err := GetGoPath("src")
	if err != nil {
		return "", "", err
	}
	npath := path
	if !strings.HasPrefix(path, "./") && !strings.HasPrefix(path, "/") && !strings.HasPrefix(path, "../") {
		npath = filepath.Join(srcPath, path)
	}
	fmt.Println("npath:", npath, path)

	root, err := filepath.Abs(npath)
	if err != nil {
		return "", "", fmt.Errorf("不是有效的项目路径:%s", root)
	}

	if !strings.Contains(root, srcPath) {
		return "", "", fmt.Errorf("项目路径无效:%s(必须位于%s目录)", root, srcPath)
	}
	if strings.Contains(root, "modules") {
		fullPath := root[:strings.Index(root, "modules")-1]
		name := strings.Replace(fullPath, srcPath, "", -1)
		return strings.Trim(name, "/"), fullPath, nil
	}
	if strings.Contains(root, "services") {
		fullPath := root[:strings.Index(root, "services")-1]
		name := strings.Replace(fullPath, srcPath, "", -1)
		return strings.Trim(name, "/"), fullPath, nil
	}
	name := strings.Replace(root, srcPath, "", -1)
	return strings.Trim(name, "/"), root, nil
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
	if strings.Contains(path, "../") {
		cmds.Log.Error("没有找到 md 文件")
		return nil
	}
	s, err := getMDFileList(path)
	if err != nil {
		cmds.Log.Error("没有找到 md 文件")
		os.Exit(1)
	}
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
	//return getModulesPath()
	return "./modules"
}

//GetHTMLPath .
func GetHTMLPath() (path string) {
	pwd, _ := os.Getwd()
	if strings.Contains(pwd, "html") {
		comma := strings.Index(pwd, "html")
		return strings.Join([]string{pwd[:comma-1], "/html"}, "")
	}

	return "./html"
}

//GetServerPath .
//path  有modules的目录   比如 /github.com/micro-plat/gaea/modules/
func GetServerPath() (path string) {
	pwd, _ := os.Getwd()
	//在services里面
	if strings.Contains(pwd, "services") {
		comma := strings.Index(pwd, "services")
		return strings.Join([]string{pwd[:comma-1], "/services"}, "")
	}
	//不在services里面

	return "./services"
}

func getModulesPath() string {
	p := getModulesPathRec(".")
	if len(p) >= 1 {
		return p[0]
	}
	return ""
}

func getModulesPathRec(path string) []string {
	if strings.Contains(path, "../../") {
		cmds.Log.Error("没有找到 modules 文件夹，请手动创建")
		return []string{}
	}
	s, err := getModulesList(path)
	if err != nil {
		cmds.Log.Error("没有找到 modules 文件夹，请手动创建")
		os.Exit(1)
	}
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
