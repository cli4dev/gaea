package mysql

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/micro-plat/gaea/cmds/new/sql/conf"
	"github.com/micro-plat/lib4go/types"
)

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("mysql").Funcs(makeFunc())
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
func getNull(n bool) string {
	if n {
		return "not null"
	}
	return ""
}
func getDef(n string, c string) string {
	if strings.Contains(c, "SEQ") {
		return ""
	}
	if n == "" {
		return ""
	}
	if strings.TrimSpace(n) == "-" {
		return "default '-'"
	}
	if strings.TrimSpace(n) == "sysdate" {
		return "default current_timestamp"
	}
	return "default " + n
}

func getUnqs(tb *conf.Table) (out []map[string]interface{}) {
	out = make([]map[string]interface{}, 0, 1)
	seqs := make([]string, 0, 1)
	for i, v := range tb.Cons {
		if strings.Contains(v, "UNQ") {
			seqs = append(seqs, tb.CNames[i])
		}
	}
	if len(seqs) == 0 {
		return nil
	}
	out = append(out, map[string]interface{}{
		"name": fmt.Sprintf("unq_%s_%s", tb.Name, seqs[0]),
		"flds": strings.Join(seqs, ","),
	})
	return out
}

func getPk(c string) string {
	if strings.Contains(c, "PK") {
		return "PRIMARY KEY"
	}
	return ""
}
func getSeqs(v string) string {
	if strings.Contains(v, "SEQ") {
		return "AUTO_INCREMENT"
	}
	return ""
}

func getAutoIncrement(tbs *conf.Table) string {
	for i, v := range tbs.Cons {
		if strings.Contains(v, "SEQ") {
			if tbs.Defs[i] != "" {
				return fmt.Sprintf("AUTO_INCREMENT=%s,", strings.TrimSpace(strings.TrimLeft(tbs.Defs[i], "default")))
			}
		}
	}
	return ""
}

func GetTmples(tbs []*conf.Table) (out map[string]string, err error) {
	out = make(map[string]string, len(tbs))
	for _, tb := range tbs {
		columns := make([]map[string]interface{}, 0, len(tb.CNames))
		for i, v := range tb.CNames {
			row := map[string]interface{}{
				"name":    v,
				"desc":    strings.Replace(tb.Descs[i], ";", " ", -1),
				"type":    tb.Types[i],
				"len":     tb.Lens[i],
				"def":     getDef(tb.Defs[i], tb.Cons[i]),
				"seqs":    getSeqs(tb.Cons[i]),
				"null":    getNull(tb.IsNulls[i]),
				"pk":      getPk(tb.Cons[i]),
				"not_end": i < len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}
		input := map[string]interface{}{
			"name":      tb.Name,
			"desc":      tb.Desc,
			"columns":   columns,
			"unqs":      getUnqs(tb),
			"seq_value": getAutoIncrement(tb),
		}
		out[fmt.Sprintf("%s.sql", tb.Name)], err = translate(tableTmpl, input)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"cName": fGetCName,
		"nName": fGetNName,
		"cType": fGetType,
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
func fGetNName(n string) string {
	items := strings.Split(n, "_")
	if len(items) <= 1 {
		return n
	}
	return strings.Join(items[1:], "_")
}

func fGetType(n string) string {
	reg := regexp.MustCompile(`[\w]+`)
	tps := reg.FindAllString(n, -1)
	switch len(tps) {
	case 1:
		switch tps[0] {
		case "date":
			return "datetime"
		}
	case 2:
		switch tps[0] {
		case "nvarchar2", "varchar2", "varchar":
			return fmt.Sprintf("varchar(%s)", tps[1])
		case "number":
			num := types.GetInt(tps[1], -1)
			if num <= 10 {
				return "int"
			}
			return "bigint"
		}
	case 3:
		switch tps[0] {
		case "number":
			return fmt.Sprintf("decimal(%s,%s)", tps[1], tps[2])
		}
	}
	panic("未处理的类型:" + n)
}

func getFilterName(t string, f string) string {
	text := make([]string, 0, 1)
	tb := strings.Split(t, "_")
	fs := strings.Split(f, "_")
	for _, v := range fs {
		ex := false
		for _, k := range tb {
			if v == k {
				ex = true
				break
			}
		}
		if !ex {
			text = append(text, v)
		}
	}
	if len(text) == 0 {
		return "id"
	}
	return strings.Join(text, "_")
}
