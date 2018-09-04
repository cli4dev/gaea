package entity

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/micro-plat/gaea/cmds/new/sql/conf"
	"github.com/micro-plat/lib4go/types"
)

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("entity").Funcs(makeFunc())
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
func GetTmples(tbs []*conf.Table, path string, filters []string) (out map[string]string, err error) {
	out = make(map[string]string, len(tbs))
	for _, tb := range tbs {
		if len(filters) > 0 {
			e := false
			for _, f := range filters {
				if strings.Contains(tb.Name, f) {
					e = true
					break
				}
			}
			if !e {
				continue
			}
		}
		columns := make([]map[string]interface{}, 0, len(tb.CNames))
		for i, v := range tb.CNames {
			row := map[string]interface{}{
				"name": cSpecialName(v),
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"null": tb.IsNulls[i],
			}
			columns = append(columns, row)
		}
		input := map[string]interface{}{
			"name":    tb.Name,
			"desc":    tb.Desc,
			"columns": columns,
			"path":    path,
		}
		content, err := translate(entityTmpl, input)
		if err != nil {
			return nil, err
		}
		out[fmt.Sprintf("%s.go", strings.Replace(tb.Name, "_", ".", -1))] = strings.Replace(strings.Replace(content, "'", "`", -1), "&#34;", `"`, -1)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"cname": fGetCName,
		"ctype": fGetType,
		"lname": fGetLastName,
		"valid": fValidName,
	}
}
func fGetCName(n string) string {
	items := strings.Split(n, "_")
	nitems := make([]string, 0, len(items))
	for _, i := range items {
		nitems = append(nitems, strings.ToUpper(i[0:1])+i[1:])
	}
	return strings.Join(nitems, "")
}
func fGetType(n string) string {
	reg := regexp.MustCompile(`[\w]+`)
	tps := reg.FindAllString(n, -1)
	switch len(tps) {
	case 1:
		switch tps[0] {
		case "date":
			return "time.Time"
		}
	case 2:
		switch tps[0] {
		case "nvarchar2", "varchar2":
			return "string"
		case "number":
			num := types.GetInt(tps[1], -1)
			if num <= 10 {
				return "int"
			}
			return "int64"
		}
	case 3:
		switch tps[0] {
		case "number":
			num := types.GetInt(tps[1], -1)
			if num <= 10 {
				return "float32"
			}
			return "float64"
		}
	}
	panic("未处理的类型:" + n)
}
func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}
func cSpecialName(v string) string {
	if strings.HasSuffix(v, "id") {
		return v[0:len(v)-2] + "ID"
	}
	return v
}
func fValidName(b bool) string {
	if b {
		return `valid:"required"`
	}
	return ""
}
