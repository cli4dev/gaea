package api

import (
	"bytes"
	"html/template"
	"strings"
)

//GetTmpl 获取模板
func GetTmpl(projectName string, serverType string, modules []string) (out map[string]string, err error) {
	input := makeParams(projectName, serverType, modules)
	out = make(map[string]string)
	if out["main.go"], err = translate(mainTmpl, input); err != nil {
		return nil, err
	}
	if out["bind.go"], err = translate(bindTmpl, input); err != nil {
		return nil, err
	}
	return out, nil
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
func makeParams(projectName string, serverType string, modules []string) map[string]interface{} {
	nmodule := make([]string, 0, len(modules)+1)
	for _, m := range modules {
		ms := strings.Split(m, " ")
		nmodule = append(nmodule, ms...)
	}
	return map[string]interface{}{
		"projectName": projectName,
		"serverType":  serverType,
		"modules":     nmodule,
	}
}
func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"humpName": fGetHumpName,
		"pkgName":  fGetPackageName,
		"lName":    fGetLastName,
		"fName":    fGetFirstName,
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

func fGetPackageName(n string, d ...string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	buff := bytes.NewBufferString("")
	if len(names) == 1 {
		if len(d) > 0 {
			return d[0]
		}
		return ""
	}
	for i := 0; i < len(names)-1; i++ {
		buff.WriteString(names[i])
	}
	return buff.String()
}

func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}
