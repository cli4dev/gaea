package util

import (
	"strconv"
	"strings"

	"github.com/micro-plat/gaea/cmds/new/sql/conf"
)

//GetInputData 获取模板数据
func GetInputData(tb *conf.Table, path string) map[string]interface{} {
	input := map[string]interface{}{
		"name":          tb.Name,
		"desc":          tb.Desc,
		"createcolumns": getCreateColumns(tb),
		"querycolumns":  getQueryColumns(tb),
		"updatecolumns": getUpdateColumns(tb),
		"selectcolumns": getSelectColumns(tb),
		"pk":            getPks(tb),
		"path":          path,
	}

	return input
}

func getCreateColumns(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "I") && !strings.Contains(tb.Cons[i], "SEQ") {
			row := map[string]interface{}{
				"name": v,
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}

	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func getQueryColumns(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "Q") && !strings.Contains(tb.Cons[i], "SEQ") {
			row := map[string]interface{}{
				"name": v,
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}

	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func getUpdateColumns(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "U") && !strings.Contains(tb.Cons[i], "SEQ") && !strings.Contains(tb.Cons[i], "PK") {
			row := map[string]interface{}{
				"name": v,
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}

	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func getSelectColumns(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "S") {
			row := map[string]interface{}{
				"name": v,
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}

	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}
func getPks(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "PK") {
			row := map[string]interface{}{
				"name": v,
				"desc": tb.Descs[i],
				"type": tb.Types[i],
				"len":  tb.Lens[i],
				"end":  i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}

	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func MakeFunc() map[string]interface{} {
	return map[string]interface{}{
		"cname": fGetCName,
		"ctype": fGetType,
		"lname": fGetLastName,
		"lower": fToLower,
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
	switch {
	case strings.Contains(n, "varchar"):
		return "string"
	case strings.Contains(n, "number"):
		if strings.Contains(n, ",") {
			return "float64"
		}
		var i, j int
		for k, v := range n {
			if v == '(' {
				i = k
			}
			if v == ')' {
				j = k
			}
		}
		ii, _ := strconv.Atoi(n[i+1 : j])
		if ii < 10 {
			return "int"
		}
		return "int64"

	case strings.Contains(n, "date"):
		return "time.Time"
	default:
		return "string"
	}
}
func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}

func fToLower(s string) string {
	return strings.ToLower(s)
}
