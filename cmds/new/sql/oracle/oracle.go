package oracle

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/micro-plat/gaea/cmds/new/util/conf"
)

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("oracle").Funcs(makeFunc())
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
func getDef(n string) string {
	if n == "" {
		return ""
	}
	if strings.TrimSpace(n) == "-" {
		return "default '-'"
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

func getPks(tb *conf.Table) []string {
	out := make([]string, 0, 1)
	for i, v := range tb.Cons {
		if strings.Contains(v, "PK") {
			out = append(out, tb.CNames[i])
		}
	}
	return out
}
func getSeqs(tb *conf.Table) (out []map[string]interface{}) {
	out = make([]map[string]interface{}, 0, 1)
	seqs := make(map[string]string)
	for i, v := range tb.Cons {
		if strings.Contains(v, "SEQ") {
			seqs[tb.CNames[i]] = tb.CNames[i]
		}
	}
	for _, v := range seqs {
		out = append(out, map[string]interface{}{
			"name": fmt.Sprintf("seq_%s_%s", fGetNName(tb.Name), getFilterName(tb.Name, v)),
			"min":  100,
			"max":  99999999999,
		})
	}
	return out
}

func GetTmples(tbs []*conf.Table) (out map[string]string, err error) {
	out = make(map[string]string, len(tbs))
	for _, tb := range tbs {
		columns := make([]map[string]interface{}, 0, len(tb.CNames))
		for i, v := range tb.CNames {
			row := map[string]interface{}{
				"name": v,
				"desc": strings.Replace(tb.Descs[i], ";", " ", -1),
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"def":  getDef(tb.Defs[i]),
				"null": getNull(tb.IsNulls[i]),
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}
		input := map[string]interface{}{
			"name":    tb.Name,
			"desc":    tb.Desc,
			"columns": columns,
			"seqs":    getSeqs(tb),
			"pks":     getPks(tb),
			"unqs":    getUnqs(tb),
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
	if strings.HasPrefix(n, "nvarchar") {
		return "string"
	} else if strings.HasPrefix(n, "number") {
		if strings.Contains(n, ",") {
			return "float64"
		}
		return "int64"
	} else if strings.HasPrefix(n, "date") {
		return "time.Time"
	}
	return "string"
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
