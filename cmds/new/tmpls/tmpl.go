package tmpls

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"
)

//GetTmpl 获取模板
func GetTmpl(projectName string, serverType string, modules []string, restful bool) (out map[string]string, err error) {
	rmodules := getRModules(modules)
	input := makeParams(projectName, serverType, rmodules, restful)
	out = make(map[string]string)
	if out["main.go"], err = translate(mainTmpl, input); err != nil {
		return nil, err
	}
	if out["bind.go"], err = translate(bindTmpl, input); err != nil {
		return nil, err
	}
	for _, m := range rmodules {
		input["moduleName"] = m
		if out[filepath.Join("services", m+".go")], err = translate(serviceTmpl, input); err != nil {
			return nil, err
		}
		if out[filepath.Join("modules", m+".go")], err = translate(moduleTmpl, input); err != nil {
			return nil, err
		}
		if out[filepath.Join("modules", "sql", strings.Replace(fGetPackageName(m), "/", ".", -1)+".go")], err = translate(sqlTmpl, input); err != nil {
			return nil, err
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
func makeParams(projectName string, serverType string, modules []string, restful bool) map[string]interface{} {
	return map[string]interface{}{
		"projectName": projectName,
		"serverType":  serverType,
		"modules":     modules,
		"pkgs":        gGetModulePackageName(modules),
		"rss":         getServices(serverType),
		"restful":     restful,
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
		buff.WriteString(strings.ToUpper(v[0:1]))
		buff.WriteString(v[1:])
	}
	return buff.String()
}
func fGetServicePackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return strings.Join(names[0:len(names)-1], "/")
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
	return strings.Join(names[0:len(names)-1], "/")
}
func fGetServicePackagePath(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return filepath.Join("services", strings.Join(names[0:len(names)-1], "/"))
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
