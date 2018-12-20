package data

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/module/tmpls"
	"github.com/micro-plat/gaea/cmds/service/tpl"
	"github.com/micro-plat/gaea/cmds/util/conf"
)

//GetInputData 获取模板数据
func getInputData(tb *conf.Table) map[string]interface{} {
	input := map[string]interface{}{
		"name":          tb.Name,
		"desc":          tb.Desc,
		"createcolumns": getCreateColumns(tb), //创建数据必需的字段
		"querycolumns":  getQueryColumns(tb),  //查询字段
		"updatecolumns": getUpdateColumns(tb), //可更新的字段
		"selectcolumns": getSelectColumns(tb), //要显示的字段
		"pk":            getPks(tb),
		"seqs":          getSeqs(tb),
		"path":          GetRouterPath(tb.Name),
		"di":            getDictionariesID(tb),
		"dn":            getDictionariesName(tb),
	}

	return input
}

func getDictionariesID(tb *conf.Table) string {
	for i, v := range tb.CNames {
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic(tb.Cons[i] + "数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "DI") {
			return v
		}
	}
	return ""
}

func getDictionariesName(tb *conf.Table) string {
	for i, v := range tb.CNames {
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic(tb.Cons[i] + "数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "DN") {
			return v
		}
	}
	return ""
}

//GetRouterPath .
func GetRouterPath(tabName string) string {
	return "/" + strings.Replace(tabName, "_", "/", -1)
}

//GetHandleName .
func GetHandleName(name string, ident string) string {
	strArray := strings.Split(name, ident)
	var strPre string
	if len(strArray) > 1 {
		strPre = strArray[len(strArray)-2]
	} else {
		strPre = strArray[0]
	}
	return strPre + ".New" + fGetCNameByIdent(name, ident) + "Handler"
}
func fGetCNameByIdent(name, ident string) string {
	items := strings.Split(name, ident)
	nitems := make([]string, 0, len(items))
	for _, i := range items {
		nitems = append(nitems, strings.ToUpper(i[0:1])+i[1:])
	}
	return strings.Join(nitems, "")
}

func getCreateColumns(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic("数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "I") && !strings.Contains(tb.Cons[i], "SEQ") {
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
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
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic("数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "Q") && !strings.Contains(tb.Cons[i], "SEQ") {
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			var domType string
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
				"domType":    domType,
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
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic("数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "U") && !strings.Contains(tb.Cons[i], "SEQ") && !strings.Contains(tb.Cons[i], "PK") {
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
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
		if tb.Cons[i] == "" || tb.Cons[i] == "-" {
			panic("数据表没有指定约束")
		}
		if strings.Contains(tb.Cons[i], "S") {
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
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
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}
	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func getSeqs(tb *conf.Table) []map[string]interface{} {
	columns := make([]map[string]interface{}, 0, len(tb.CNames))

	for i, v := range tb.CNames {
		if strings.Contains(tb.Cons[i], "SEQ") {
			descsimple := tb.Descs[i]
			if strings.Contains(tb.Descs[i], "(") {
				descsimple = tb.Descs[i][:strings.Index(tb.Descs[i], "(")]
			}
			row := map[string]interface{}{
				"name":       v,
				"descsimple": descsimple,
				"seqname":    fmt.Sprintf("seq_%s_%s", fGetNName(tb.Name), getFilterName(tb.Name, v)),
				"desc":       tb.Descs[i],
				"type":       tb.Types[i],
				"len":        tb.Lens[i],
				"end":        i != len(tb.CNames)-1,
			}
			columns = append(columns, row)
		}
	}
	if len(columns) > 0 {
		columns[len(columns)-1]["end"] = false
	}
	return columns
}

func fGetNName(n string) string {
	items := strings.Split(n, "_")
	if len(items) <= 1 {
		return n
	}
	return strings.Join(items[1:], "_")
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

func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"aname":      fGetAName,
		"cname":      fGetCName,
		"pname":      fGetPName,
		"ctype":      fGetType,
		"cstype":     fsGetType,
		"lname":      fGetLastName,
		"lower":      fToLower,
		"vname":      vName,
		"humpName":   fGetHumpName,           //多个单词首字符大写
		"spkgName":   fGetServicePackageName, //包路径
		"mpkgName":   fGetModulePackageName,  //包路径
		"lName":      fGetLastName,           //取最后一个单词
		"fName":      fGetFirstName,          //取第一个单词
		"fServer":    fServer,                //判断是否有这个服务
		"getAppconf": getAppconf,
	}
}
func getAppconf(str string, index int) string {
	strArray := strings.Split(str, "|")
	if len(strArray) < index {
		return ""
	}
	if ok := strArray[index-1]; ok != "" {
		return ok
	}
	return ""
}

func vName(v string) string {
	items := strings.Split(v, "/")
	nitems := make([]string, 0, len(items))
	for k, i := range items {
		if k == 0 {
			nitems = append(nitems, i)
		}
		if k > 0 {
			nitems = append(nitems, strings.ToUpper(i[0:1])+i[1:])
		}

	}
	return strings.Join(nitems, "")
}

func fGetAName(n string) string {
	items := strings.Split(n, "_")
	nitems := make([]string, 0, len(items))
	for k, i := range items {
		if k == 0 {
			nitems = append(nitems, i)
		}
		if k > 0 {
			nitems = append(nitems, strings.ToUpper(i[0:1])+i[1:])
		}

	}
	return strings.Join(nitems, "")
}

func fGetCName(n string) string {
	items := strings.Split(n, "_")
	nitems := make([]string, 0, len(items))
	for _, i := range items {
		nitems = append(nitems, strings.ToUpper(i[0:1])+i[1:])
	}
	return strings.Join(nitems, "")
}

func fGetPName(n string) string {
	items := strings.Split(n, "_")
	if len(items) >= 2 {
		return items[len(items)-2]
	}
	return items[0]
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
func fsGetType(n string) string {
	t := fGetType(n)
	if t == "time.Time" {
		return "string"
	}
	return t
}

func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}

func fToLower(s string) string {
	return strings.ToLower(s)
}

//Translate  .
func Translate(tag string, tplName string, input interface{}) (string, error) {
	var tmpl = template.New(tag).Funcs(makeFunc())
	np, err := tmpl.Parse(tplName)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBufferString("")
	if err := np.Execute(buff, input); err != nil {
		return "", err
	}
	return buff.String(), nil
}

//GetTmples 获取模板
//@tag 模板标签
//@tplName 模板名字
//@tbs 表结构体
//@filters 过滤字段
//@makeFunc 是否生成函数
//@return out 返回数据
func GetTmples(tag, tplName string, tbs []*conf.Table, filters []string, makeFunc bool, modulePath string) (out map[string]map[string]string, err error) {
	out = map[string]map[string]string{}
	for _, tb := range tbs {
		if len(filters) > 0 {
			e := false
			for _, f := range filters {
				if strings.EqualFold(tb.Name, f) {
					e = true
					break
				}
			}
			if !e {
				continue
			}
		}
		//获取模板数据
		input := getInputData(tb)
		//翻译模板
		content, err := Translate(tag, tplName, input)

		if err != nil {
			return nil, err
		}
		if makeFunc { //生成函数
			c := make(map[string]string)
			if strings.Contains(modulePath, "sql") || !strings.Contains(modulePath, "modules") {
				modulePath = "modules"
			}
			c[fmt.Sprintf(modulePath+"/%s.go", strings.Replace(tb.Name, "_", "/", -1))] = strings.Replace(content, "'", "`", -1)
			head, err := Translate("head", tmpls.DbHeadTpl, input)
			if err != nil {
				return nil, err
			}
			c["head"] = strings.Replace(head, "'", "`", -1)
			out[fmt.Sprintf(modulePath+"/%s.go", strings.Replace(tb.Name, "_", "/", -1))] = c
		} else { //生成sql
			c := make(map[string]string)

			modulePath = "modules/const/sql"

			c[fmt.Sprintf(modulePath+"/%s.go", strings.Replace(tb.Name, "_", ".", -1))] = strings.Replace(content, "'", "`", -1)
			sq, err := Translate("head sql", tmpls.DicTpl, input)
			if err != nil {
				return nil, err
			}
			c["sql"] = strings.Replace(sq, "'", "`", -1)
			out[fmt.Sprintf(modulePath+"/%s.go", strings.Replace(tb.Name, "_", ".", -1))] = c
		}
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

//GetServerTmples 获取模板
//@tag 模板标签
//@tplName 模板名字
//@tbs 表结构体
//@filters 过滤字段
//@makeFunc 是否生成函数
//@return out 返回数据
func GetServerTmples(tag, tplName string, tbs []*conf.Table, filters []string, serverPath string) (out map[string]map[string]string, err error) {
	out = map[string]map[string]string{}
	for _, tb := range tbs {
		if len(filters) > 0 {
			e := false
			for _, f := range filters {
				if strings.EqualFold(tb.Name, f) {
					e = true
					break
				}
			}
			if !e {
				continue
			}
		}
		//获取模板数据
		input := getInputData(tb)
		//翻译模板
		content, err := Translate(tag, tplName, input)

		if err != nil {
			return nil, err
		}

		c := make(map[string]string)
		if !strings.Contains(serverPath, "services") {
			serverPath = "services"
		}
		c[fmt.Sprintf(serverPath+"/%s.go", strings.Replace(tb.Name, "_", "/", -1))] = strings.Replace(content, "'", "`", -1)
		head, err := Translate("head", tpl.HeadTpl, input)
		if err != nil {
			return nil, err
		}
		c["head"] = strings.Replace(head, "'", "`", -1)
		out[fmt.Sprintf(serverPath+"/%s.go", strings.Replace(tb.Name, "_", "/", -1))] = c

		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

//GetHTMLTmples 获取模板
//@tag 模板标签
//@tplName 模板名字
//@tbs 表结构体
//@filters 过滤字段
//@return out 返回数据
func GetHTMLTmples(tag, tplName string, tbs []*conf.Table, filters []string, projectPath string) (out map[string]map[string]string, err error) {
	out = map[string]map[string]string{}
	for _, tb := range tbs {
		if len(filters) > 0 {
			e := false
			for _, f := range filters {
				if strings.EqualFold(tb.Name, f) {
					e = true
					break
				}
			}
			if !e {
				continue
			}
		}
		//获取模板数据
		input := getInputData(tb)
		//翻译模板
		content, err := Translate(tag, tplName, input)
		if err != nil {
			return nil, err
		}
		c := make(map[string]string)
		if !strings.Contains(projectPath, "src/pages") {
			projectPath = filepath.Join(projectPath, "/src/pages")
		}
		c[fmt.Sprintf(projectPath+"/%s.vue", strings.Replace(tb.Name, "_", "/", -1))] = strings.Replace(content, "'", "`", -1)
		out[fmt.Sprintf(projectPath+"/%s.vue", strings.Replace(tb.Name, "_", "/", -1))] = c
	}
	return out, nil
}

//createFile
//创建并生成文件
func createFile(add bool, ms string, data map[string]map[string]string) error {
	for k, v := range data {
		_, ok := v["head"]
		if ok { //生成函数文件头
			_, err := os.Stat(k)
			//不存在则创建
			if err != nil {
				add = true
				os.MkdirAll(path.Dir(k), os.ModePerm)
				f, err := os.Create(k)
				if err != nil {
					err = fmt.Errorf("无法创建文件:%s(err:%v)", k, err)
					return err
				}
				defer f.Close()
				m := strings.Split(k, "/")
				absPath, _ := filepath.Abs(k)
				i := strings.Index(absPath, "src")
				j := strings.Index(absPath, ms)
				var packageName string
				if ms == "modules" {
					packageName = absPath[i+4 : j]
				} else {
					str := strings.Replace(absPath[i+4:], "services", "modules", -1)
					strSlice := strings.Split(str, "/")
					packageName = strings.Join(strSlice[:len(strSlice)-1], "/")
				}
				_, err = f.WriteString(fmt.Sprintf(v["head"], m[len(m)-2], packageName))
				if err != nil {
					return err
				}
				cmds.Log.Info("写入crud函数头部文件成功:", k)
			}
		} else { //生成sql文件头
			_, err := os.Stat(k)
			//不存在则创建
			if err != nil {
				add = true
				os.MkdirAll(path.Dir(k), os.ModePerm)
				f, err := os.Create(k)
				if err != nil {
					err = fmt.Errorf("无法创建文件:%s(err:%v)", k, err)
					return err
				}
				defer f.Close()
				headStr, ok := v["sql"]
				if !ok {
					headStr = "package sql"
					if strings.Contains(k, "vue") {
						headStr = ""
					}
				}
				_, err = f.WriteString(headStr)
				if err != nil {
					return err
				}
			}
		}
		if !add {
			return fmt.Errorf("文件已经存在：%s，未输入 -add 不执行任何操作", k)
		}
		srcf, err := os.OpenFile(k, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			cmds.Log.Errorf("无法打开文件:%s", k)
			return err
		}
		defer srcf.Close()
		_, err = srcf.WriteString(v[k])
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		cmds.Log.Info("写入文件成功:", k)

	}
	return nil

}

//CreateModulesFile 创建 modules 文件
func CreateModulesFile(add, cover bool, tmpls map[string]map[string]string) (err error) {
	if cover {
		for k := range tmpls {
			cmds.Log.Warnf("覆盖文件：%s", k)
			os.Remove(k)
			break
		}
	}
	//创建文件
	if err = createFile(add, "modules", tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//CreateHTMLFile 创建 vue 文件
func CreateHTMLFile(tmpls map[string]map[string]string) (err error) {
	// for k := range tmpls {
	// 	cmds.Log.Warnf("覆盖文件：%s", k)
	// 	os.Remove(k)
	// }
	//创建文件
	if err = createFile(true, "html", tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//CreateServicesFile .
func CreateServicesFile(add, cover bool, tmpls map[string]map[string]string) (err error) {
	if cover {
		for k := range tmpls {
			cmds.Log.Warnf("覆盖文件：%s", k)
			os.Remove(k)
		}
	}
	//创建文件
	if err = createFile(add, "services", tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//ReplaceFileStr .
//替换文件内容
func ReplaceFileStr(name, filePath, value string) error {

	srcf, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer srcf.Close()
	if err != nil {
		err = fmt.Errorf("无法打开文件:%s(err:%v)", filePath, err)
		return err
	}
	buf, err := ioutil.ReadAll(srcf)
	if err != nil {
		cmds.Log.Errorf("%v", err.Error())
		return err
	}
	result := string(buf)

	k := fmt.Sprintf(`//%s#//[\s\S]+//#%s//`, name, name)
	if ok, _ := regexp.Match(k, buf); !ok {
		cmds.Log.Errorf("没有找到配置定位标识符:%s//%s#//....//#%s//", filePath, name, name)
		return nil
	}
	re, _ := regexp.Compile(k)
	str := re.ReplaceAllString(result, value)
	//清空文件数据
	srcf.Truncate(0)
	n, _ := srcf.Seek(0, os.SEEK_SET)
	_, err = srcf.WriteAt([]byte(str), n)
	if err != nil {
		return err
	}
	return nil
}

func fServer(s, substr string) bool {
	return strings.Contains(s, substr)
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
