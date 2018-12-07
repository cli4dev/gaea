package tmpls

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds/project/tmpls/dev"
)

var templateFiles map[string][]string
var names map[string]string
var templates map[string]string

func init() {
	templateFiles = make(map[string][]string)
	templateFiles["install.dev"] = []string{"port"}

	names = make(map[string]string)
	names["port"] = APIMainPort

	templates = make(map[string]string)
	templates["port"] = dev.APIMainPort
}

const (
	APIMainPort  = `api.main.port`
	APISubHeader = `api.sub.header`
	APISubAuth   = `api.sub.auth`
	APISubMetric = `api.sub.metric`
	PlatVarDB    = `plat.var.db`
	PlatVarCache = `plat.var.cache`
	PlatVarQueue = `plat.var.queue`
	CronSubApp   = `cron.sub.app`
	CronSubTask  = `cron.sub.task`
	MQCSubServer = `mqc.sub.server`
	MQCSubQueue  = `mqc.sub.queue`
	WebMainPort  = `web.main.port`
	WebSubstatic = `web.sub.static`
	WSSubApp     = `ws.sub.app`
	WSSubAuth    = `ws.sub.auth`
	RPCMainPort  = `rpc.main.port`
)

//GetTmpls 获取模板
func GetTmpls(projectName string, input map[string]interface{}) (out map[string]string, err error) {
	fmt.Println("input:", input)
	//input := makeParams(projectName, serverType, port, db, jwt, domain)
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

	out["db.md"] = ""
	out["modules/const/sql/sql.go"] = "dir"
	out["services/server.go"] = "dir"

	return out, nil
}

//GetConfTmpls 获取配置模板
func GetConfTmpls(blocks []string, input map[string]interface{}) (out map[string]map[string]string, err error) {

	out = make(map[string]map[string]string)
	for fname, f := range templateFiles {
		for _, n := range f {
			for _, name := range blocks {
				if name != n {
					break
				}

				if _, ok := out[fname]; !ok {
					out[fname] = make(map[string]string)
				}

				out[fname][names[name]], err = translate(strings.Replace(strings.Replace(templates[name], "\"", "`", -1), "'", "\"", -1), input)
				if err != nil {
					return nil, err
				}
			}
		}

	}

	return out, err

	// for _, v := range serverTypeSlice {
	// 	switch v {
	// 	case "api":

	// 		out[APIMainPort], _ = translate(strings.Replace(strings.Replace(dev.APIMainPort, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[APISubHeader], _ = translate(strings.Replace(strings.Replace(dev.APISubHeader, "\"", "`", -1), "'", "\"", -1), input)
	// 		if domain {
	// 			out[APISubHeader], _ = translate(strings.Replace(strings.Replace(dev.APISubHeaderDomain, "\"", "`", -1), "'", "\"", -1), input)
	// 		}
	// 		if jwt {
	// 			out[APISubAuth], _ = translate(strings.Replace(strings.Replace(dev.APISubAuth, "\"", "`", -1), "'", "\"", -1), input)
	// 		}

	// 		out[APISubMetric], _ = translate(strings.Replace(strings.Replace(dev.APISubMetric, "\"", "`", -1), "'", "\"", -1), input)
	// 	case "cron":
	// 		out[CronSubApp], _ = translate(strings.Replace(strings.Replace(dev.CronSubApp, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[CronSubTask], _ = translate(strings.Replace(strings.Replace(dev.CronSubTask, "\"", "`", -1), "'", "\"", -1), input)
	// 	case "rpc":
	// 		out[RPCMainPort], _ = translate(strings.Replace(strings.Replace(dev.RPCMainPort, "\"", "`", -1), "'", "\"", -1), input)
	// 	case "mqc":
	// 		out[MQCSubQueue], _ = translate(strings.Replace(strings.Replace(dev.MqcSubQueue, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[MQCSubServer], _ = translate(strings.Replace(strings.Replace(dev.MqcSubServer, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[PlatVarQueue], _ = translate(strings.Replace(strings.Replace(dev.PlatVarQueue, "\"", "`", -1), "'", "\"", -1), input)
	// 	case "web":
	// 		out[WebMainPort], _ = translate(strings.Replace(strings.Replace(dev.WebMainPort, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[WebSubstatic], _ = translate(strings.Replace(strings.Replace(dev.WebSubStatic, "\"", "`", -1), "'", "\"", -1), input)
	// 	case "ws":
	// 		out[WSSubApp], _ = translate(strings.Replace(strings.Replace(dev.WSSubAPP, "\"", "`", -1), "'", "\"", -1), input)
	// 		out[WSSubAuth], _ = translate(strings.Replace(strings.Replace(dev.WSSubAuth, "\"", "`", -1), "'", "\"", -1), input)
	// 	}
	// }

	// return out, nil
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
func makeParams(projectName, serverType, port, db string, jwt, domain bool) map[string]interface{} {
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	return map[string]interface{}{
		"projectName": projectName,
		"serverType":  serverType,
		//	"pkgs":    gGetModulePackageName(modules),
		"rss":    getServices(serverType),
		"port":   port,
		"dbname": strings.Split(db, ":")[0],
		"db":     db,
		"jwt":    jwt,
		"domain": domain,
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
		"fServer":  fServer,                //判断是否有这个服务
	}
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
