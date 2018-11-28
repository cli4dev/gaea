package tmpls

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds/new/project/tmpls/dev"
	"github.com/micro-plat/gaea/cmds/new/project/tmpls/prod"
)

//GetTmpls 获取模板
func GetTmpls(projectName string, serverType string) (out map[string]string, err error) {

	input := makeParams(projectName, serverType)
	out = make(map[string]string)
	if out["main.go"], err = translate(mainTmpl, input); err != nil {
		return nil, err
	}
	if out["init.go"], err = translate(initTmpl, input); err != nil {
		return nil, err
	}
	if out["install.dev.go"], err = translate(strings.Replace(strings.Replace(installDevTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["install.prod.go"], err = translate(strings.Replace(strings.Replace(installProdTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["handling.go"], err = translate(strings.Replace(strings.Replace(handingTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out[".gitignore"], err = translate(gitignoreTmpl, input); err != nil {
		return nil, err
	}
	serverTypeSlice := strings.Split(serverType, "-")

	for _, v := range serverTypeSlice {
		switch v {
		case "api":
			out["conf.dev.api.go"], _ = translate(strings.Replace(strings.Replace(dev.Api, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.api.go"], _ = translate(strings.Replace(strings.Replace(prod.Api, "\"", "`", -1), "'", "\"", -1), input)
		case "cron":
			out["conf.dev.cron.go"], _ = translate(strings.Replace(strings.Replace(dev.Cron, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.cron.go"], _ = translate(strings.Replace(strings.Replace(prod.Cron, "\"", "`", -1), "'", "\"", -1), input)
		case "rpc":
			out["conf.dev.rpc.go"], _ = translate(strings.Replace(strings.Replace(dev.Rpc, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.rpc.go"], _ = translate(strings.Replace(strings.Replace(prod.Rpc, "\"", "`", -1), "'", "\"", -1), input)
		case "mqc":
			out["conf.dev.mqc.go"], _ = translate(strings.Replace(strings.Replace(dev.Mqc, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.mqc.go"], _ = translate(strings.Replace(strings.Replace(prod.Mqc, "\"", "`", -1), "'", "\"", -1), input)
		case "web":
			out["conf.dev.web.go"], _ = translate(strings.Replace(strings.Replace(dev.Web, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.web.go"], _ = translate(strings.Replace(strings.Replace(prod.Web, "\"", "`", -1), "'", "\"", -1), input)
		}
	}
	out["conf.dev.plat.go"], _ = translate(strings.Replace(strings.Replace(dev.Plat, "\"", "`", -1), "'", "\"", -1), input)
	out["conf.prod.plat.go"], _ = translate(strings.Replace(strings.Replace(prod.Plat, "\"", "`", -1), "'", "\"", -1), input)
	out["db.md"] = ""
	out["modules/const/sql/sql.go"] = "dir"
	out["services/server.go"] = "dir"

	return out, nil
}

//GetConfTmpls 获取配置模板
func GetConfTmpls(serverType string) (out map[string]string, err error) {

	out = make(map[string]string)
	input := make(map[string]interface{})
	serverTypeSlice := strings.Split(serverType, "-")
	for _, v := range serverTypeSlice {
		switch v {
		case "api":
			out["conf.dev.api.go"], _ = translate(strings.Replace(strings.Replace(dev.Api, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.api.go"], _ = translate(strings.Replace(strings.Replace(prod.Api, "\"", "`", -1), "'", "\"", -1), input)
		case "cron":
			out["conf.dev.cron.go"], _ = translate(strings.Replace(strings.Replace(dev.Cron, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.cron.go"], _ = translate(strings.Replace(strings.Replace(prod.Cron, "\"", "`", -1), "'", "\"", -1), input)
		case "rpc":
			out["conf.dev.rpc.go"], _ = translate(strings.Replace(strings.Replace(dev.Rpc, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.rpc.go"], _ = translate(strings.Replace(strings.Replace(prod.Rpc, "\"", "`", -1), "'", "\"", -1), input)
		case "mqc":
			out["conf.dev.mqc.go"], _ = translate(strings.Replace(strings.Replace(dev.Mqc, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.mqc.go"], _ = translate(strings.Replace(strings.Replace(prod.Mqc, "\"", "`", -1), "'", "\"", -1), input)
		case "web":
			out["conf.dev.web.go"], _ = translate(strings.Replace(strings.Replace(dev.Web, "\"", "`", -1), "'", "\"", -1), input)
			out["conf.prod.web.go"], _ = translate(strings.Replace(strings.Replace(prod.Web, "\"", "`", -1), "'", "\"", -1), input)
		}
	}

	return out, nil
}

func getServices(serverType string) []string {
	s := make([]string, 0, 2)
	if strings.Contains(serverType, "api") || strings.Contains(serverType, "rpc") {
		s = append(s, "Micro")
	}
	if strings.Contains(serverType, "mqc") || strings.Contains(serverType, "cron") {
		s = append(s, "Autoflow")
	}
	if strings.Contains(serverType, "web") {
		s = append(s, "Page")
	}
	return s
}

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("api").Funcs(makeFunc())
	np, err := tmpl.Parse(c)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBufferString("")
	if err := np.Execute(buff, input); err != nil {
		return "", err
	}
	return buff.String(), nil
}

//获取生成项目的数据
func makeParams(projectName string, serverType string) map[string]interface{} {
	return map[string]interface{}{
		"projectName": projectName,
		"serverType":  serverType,
		//	"pkgs":    gGetModulePackageName(modules),
		"rss": getServices(serverType),
	}
}

func getRModules(modules []string) []string {
	nmodule := make([]string, 0, len(modules)+1)
	for _, m := range modules {
		ms := strings.Split(m, " ")
		for _, s := range ms {
			nmodule = append(nmodule, filepath.Join("/", s))
		}
	}
	return nmodule
}

func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"humpName": fGetHumpName,           //多个单词首字符大写
		"spkgName": fGetServicePackageName, //包路径
		"mpkgName": fGetModulePackageName,  //包路径
		"lName":    fGetLastName,           //取最后一个单词
		"fName":    fGetFirstName,          //取第一个单词
	}
}

func fGetFirstName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[0]
}

func fGetHumpName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	buff := bytes.NewBufferString("")
	for _, v := range names {
		buff.WriteString(fGetLoopHumpName(v, "."))
	}
	return strings.Replace(buff.String(), ".", "", -1)
}

func fGetLoopHumpName(n string, s string) string {
	names := strings.Split(strings.Trim(n, s), s)
	buff := bytes.NewBufferString("")
	for _, v := range names {
		buff.WriteString(strings.ToUpper(v[0:1]))
		buff.WriteString(v[1:])
	}
	return strings.Replace(buff.String(), ".", "", -1)
}

func fGetServicePackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return strings.ToLower(strings.Join(names[0:len(names)-1], "/"))
}

func fGetPackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return names[0]
	}
	return strings.Join(names[0:len(names)-1], "/")
}

func fGetModulePackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "modules"
	}
	return strings.ToLower(strings.Join(names[0:len(names)-1], "/"))
}

func fGetServicePackagePath(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return strings.ToLower(filepath.Join("services", strings.Join(names[0:len(names)-1], "/")))
}

func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}

func gGetModulePackageName(module []string) []string {
	npkgs := make([]string, 0, len(module)/2)
	n := make(map[string]string)
	for _, m := range module {
		nm := fGetServicePackagePath(m)
		if _, ok := n[nm]; !ok {
			npkgs = append(npkgs, nm)
			n[nm] = nm
		}
	}
	return npkgs
}
