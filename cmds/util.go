package cmds

import (
	"fmt"
	"os"
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
