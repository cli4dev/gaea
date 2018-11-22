package delete

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/micro-plat/gaea/cmds/new/sql/conf"
	"github.com/micro-plat/gaea/cmds/new/sql/tpl"
	"github.com/micro-plat/gaea/cmds/new/sql/util"
)

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("delete").Funcs(util.MakeFunc())
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

//GetTmples .
func GetTmples(tplName string, tbs []*conf.Table, path string, filters []string, makeFunc bool) (out map[string]map[string]string, err error) {
	out = make(map[string]map[string]string, len(tbs))
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
		input := util.GetInputData(tb, path)
		content, err := translate(tplName, input)
		if err != nil {
			return nil, err
		}
		if makeFunc {
			c := make(map[string]string)
			c[fmt.Sprintf("%s.go", strings.Replace(tb.Name, "_", "/", -1))] = strings.Replace(content, "'", "`", -1)

			head, err := translate(tpl.DbTpl, input)
			if err != nil {
				return nil, err
			}
			c["head"] = strings.Replace(head, "'", "`", -1)
			out[fmt.Sprintf("%s.go", strings.Replace(tb.Name, "_", "/", -1))] = c
		} else {
			c := make(map[string]string)
			c[fmt.Sprintf("%s.go", strings.Replace(tb.Name, "_", ".", -1))] = strings.Replace(content, "'", "`", -1)
			out[fmt.Sprintf("%s.go", strings.Replace(tb.Name, "_", ".", -1))] = c

		}

		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
