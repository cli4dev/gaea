package tmpls

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds/project/tmpls/vue"
	"github.com/micro-plat/gaea/cmds/util"
	"github.com/micro-plat/gaea/cmds/util/data"

	"github.com/micro-plat/lib4go/utility"

	"github.com/micro-plat/gaea/cmds/project/tmpls/dev"
)

var templateFiles map[string][]string
var names map[string]string
var templates map[string]string
var vueTemplates map[string]string

func init() {
	templateFiles = make(map[string][]string)
	templateFiles["install.dev.go"] = []string{"api.port", "api.jwt", "db", "cache", "api.cros", "queue",
		"api.appconf", "api.metric", "cron.appconf", "cron.task", "cron.metric",
		"web.port", "web.static", "web.metric",
		"mqc.server", "mqc.queue", "mqc.metric",
		"rpc.port", "rpc.metric",
		"ws.appconf", "ws.jwt", "ws.metric",
	}
	templateFiles["install.prod.go"] = []string{"api.port", "api.jwt.prod", "db", "cache", "api.cros", "queue",
		"api.appconf", "api.metric", "cron.appconf", "cron.task", "cron.metric",
		"web.port", "web.static", "web.metric",
		"mqc.server", "mqc.queue", "mqc.metric",
		"rpc.port", "rpc.metric",
		"ws.appconf", "ws.jwt", "ws.metric",
	}
	templateFiles["init.go"] = []string{"db.init", "queue.init", "cache.init", "appconf.func"}
	templateFiles["handling.go"] = []string{"handling.jwt"}

	names = make(map[string]string)
	names["api.port"] = apiPort
	names["api.cros"] = apiCros
	names["api.jwt"] = apiJWT
	names["api.jwt.prod"] = apiJWTProd
	names["api.metric"] = apiMetric
	names["db"] = db
	names["cache"] = cache
	names["queue"] = queue
	names["api.appconf"] = apiAppConf
	names["db.init"] = dbInit
	names["queue.init"] = queueInit
	names["cache.init"] = cacheInit
	names["appconf.func"] = appconfFunc
	names["handling.jwt"] = handlingJWT
	names["cron.appconf"] = cronApp
	names["cron.task"] = cronTask
	names["cron.metric"] = cronMetric
	names["web.port"] = webPort
	names["web.static"] = webStatic
	names["web.metric"] = webMetric
	names["mqc.server"] = mqcServer
	names["mqc.queue"] = mqcQueue
	names["mqc.metric"] = mqcMetric
	names["rpc.port"] = rpcPort
	names["rpc.metric"] = rpcMetric
	names["ws.metric"] = wsMetric
	names["ws.appconf"] = wsApp
	names["ws.jwt"] = wsAuth

	templates = make(map[string]string)
	templates["api.port"] = dev.APIMainPort
	templates["api.jwt"] = dev.APISubAuth
	templates["api.jwt.prod"] = dev.APISubAuthProd
	templates["db"] = dev.PlatVarDB
	templates["cache"] = dev.PlatVarCache
	templates["queue"] = dev.PlatVarQueue
	templates["api.cros"] = dev.APISubCros
	templates["api.metric"] = dev.APISubMetric
	templates["api.appconf"] = dev.APIApp
	templates["db.init"] = DBInit
	templates["queue.init"] = QueueInit
	templates["cache.init"] = CacheInit
	templates["appconf.func"] = APPConfFunc
	templates["handling.jwt"] = dev.HandlingJWT
	templates["cron.appconf"] = dev.CronSubApp
	templates["cron.task"] = dev.CronSubTask
	templates["cron.metric"] = dev.CronSubMetric
	templates["web.port"] = dev.WebMainPort
	templates["web.static"] = dev.WebSubStatic
	templates["web.metric"] = dev.WebSubMetric
	templates["mqc.server"] = dev.MqcSubServer
	templates["mqc.queue"] = dev.MqcSubQueue
	templates["mqc.metric"] = dev.MqcSubQueue
	templates["rpc.metric"] = dev.RPCSubMetric
	templates["rpc.port"] = dev.RPCMainPort
	templates["ws.metric"] = dev.WSSubMetric
	templates["ws.appconf"] = dev.WSSubAPP
	templates["ws.jwt"] = dev.WSSubAuth

	vueTemplates = make(map[string]string)
	vueTemplates["public/index.html"] = vue.IndexHTML
	vueTemplates["package.json"] = vue.PackAgeJSON
	vueTemplates["package-lock.json"] = vue.PackAgeLock
	vueTemplates["postcss.config.js"] = vue.PostCss
	vueTemplates["babel.config.js"] = vue.Babel
	vueTemplates["src/main.js"] = vue.MainJS
	vueTemplates["src/router.js"] = vue.Router
	vueTemplates["src/App.vue"] = vue.AppVue
	vueTemplates["src/store.js"] = vue.Store
	vueTemplates["src/pages/menu/menu.vue"] = vue.MenuTpl
	vueTemplates["src/pages/login/login.vue"] = vue.LoginTpl
	vueTemplates["src/util/http.js"] = vue.HTTPTpl
	vueTemplates[".gitignore"] = gitignoreTmpl
	vueTemplates[".env.prod"] = vue.EnvProd
	vueTemplates[".env.dev"] = vue.EnvDev
	vueTemplates["vue.config.js"] = vue.VueConfig
	vueTemplates["main.go"] = strings.Replace(strings.Replace(vue.MainGo, "\"", "`", -1), "'", "\"", -1)
	vueTemplates["build.sh"] = vue.Build
}

const (
	apiPort     = `api.port`
	apiCros     = `api.cros`
	apiJWT      = `api.jwt`
	apiJWTProd  = `api.jwt.prod`
	apiAppConf  = `api.appconf`
	apiMetric   = `api.metric`
	db          = `db`
	cache       = `cache`
	queue       = `queue`
	cronApp     = `cron.appconf`
	cronTask    = `cron.task`
	cronMetric  = `cron.metric`
	mqcServer   = `mqc.server`
	mqcQueue    = `mqc.queue`
	mqcMetric   = `mqc.metric`
	webPort     = `web.port`
	webStatic   = `web.static`
	webMetric   = `web.metric`
	wsApp       = `ws.appconf`
	wsAuth      = `ws.jwt`
	wsMetric    = `ws.metric`
	rpcPort     = `rpc.port`
	rpcMetric   = `rpc.metric`
	dbInit      = `db.init`
	queueInit   = "queue.init"
	cacheInit   = "cache.init"
	appconfFunc = "appconf.func"
	handlingJWT = "handling.jwt"
)

//GetTmpls 获取模板
func GetTmpls(projectName string, input map[string]interface{}) (out map[string]string, err error) {
	input = makeParams(input)
	out = make(map[string]string)
	if out["main.go"], err = data.Translate("main.go.tpl", mainTmpl, input); err != nil {
		return nil, err
	}
	if out["init.go"], err = data.Translate("init.go.tpl", initTmpl, input); err != nil {
		return nil, err
	}
	if out["install.dev.go"], err = data.Translate("tpl", strings.Replace(strings.Replace(installDevTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["install.prod.go"], err = data.Translate("tpl", strings.Replace(strings.Replace(installProdTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["handling.go"], err = data.Translate("tpl", strings.Replace(strings.Replace(handlingTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out[".gitignore"], err = data.Translate("tpl", gitignoreTmpl, input); err != nil {
		return nil, err
	}
	out["modules/const/sql/sql.go"] = "package sql"
	out["services/server.go"] = "package server"
	if input["login"] != "" {
		out["services/member/info.go"], _ = data.Translate("tpl", dev.Info, input)
		out["services/member/login.go"], _ = data.Translate("tpl", dev.Login, input)
		out["services/member/menu.go"], _ = data.Translate("tpl", dev.Menu, input)
		out["services/member/changepwd.go"], _ = data.Translate("tpl", dev.ChangePwd, input)
		out["services/member/user.go"], _ = data.Translate("tpl", dev.User, input)

		out["modules/app/conf.go"], _ = data.Translate("tpl", strings.Replace(strings.Replace(dev.AppModuleConf, "\"", "`", -1), "'", "\"", -1), input)
		out["modules/member/user.go"], _ = data.Translate("tpl", dev.AppModuleMemberUser, input)
		out["modules/member/menu.go"], _ = data.Translate("tpl", strings.Replace(strings.Replace(dev.AppModuleMemberMenu, "\"", "`", -1), "'", "\"", -1), input)
		out["modules/member/state.go"], _ = data.Translate("tpl", strings.Replace(strings.Replace(dev.AppModuleMemberState, "\"", "`", -1), "'", "\"", -1), input)
		out["modules/util/util.go"], _ = data.Translate("tpl", dev.Util, input)
	}

	return out, nil
}

//GetVueTmpls 获取vue项目模板
func GetVueTmpls(projectName string) (out map[string]string, err error) {

	out = make(map[string]string)
	for k, v := range vueTemplates {
		out[k], err = data.Translate(k, v, map[string]interface{}{
			"projectName": projectName,
			"ip":          util.GetLocalhostIP(),
		})
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

//GetConfTmpls 获取配置模板
func GetConfTmpls(blocks []string, input map[string]interface{}) (out map[string]map[string]string, err error) {
	input = makeParams(input)
	out = make(map[string]map[string]string)
	for fname, f := range templateFiles {
		for _, n := range f {
			for _, name := range blocks {

				if !strings.Contains(n, name) {
					continue
				}
				if _, ok := out[fname]; !ok {
					out[fname] = make(map[string]string)
				}
				out[fname][names[n]], err = data.Translate("conf", strings.Replace(strings.Replace(templates[n], "\"", "`", -1), "'", "\"", -1), input)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return out, err
}

//GetEmptyTmpls .
func GetEmptyTmpls(blocks []string) (out map[string]map[string]string, err error) {
	out = make(map[string]map[string]string)
	for fname, f := range templateFiles {
		for _, n := range f {
			for _, name := range blocks {
				if !strings.Contains(n, name) {
					continue
				}
				if _, ok := out[fname]; !ok {
					out[fname] = make(map[string]string)
				}
				out[fname][names[n]] = fmt.Sprintf(`//%s#//
	//#%s//`, n, n)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return out, err
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

//获取生成项目的数据
func makeParams(input map[string]interface{}) map[string]interface{} {

	input["devSecret"] = utility.GetGUID()
	input["prodSecret"] = utility.GetGUID()

	return input
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
