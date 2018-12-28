package tmpls

const installProdTmpl = `// +build prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *{{.projectName|lName}}) install() {
	{{$api := "api" -}}
	{{$cron := "cron" -}}
	{{$web := "web" -}}
	{{$mqc := "mqc" -}}
	{{$rpc := "rpc" -}}
	{{$ws := "ws" -}}
	{{$empty := "" -}}
	{{if fServer .serverType $api -}}
	//api.port#//
	s.Conf.API.SetMainConf("{'address':'{{.port}}'}")
	//#api.port//
	{{- else -}}
	//api.port#//
	//#api.port//
	{{- end}}

	{{if fServer .serverType $api -}}
	{{if or .api.appconf .appconf -}}
	//api.appconf#//
	s.Conf.API.SetSubConf('app', "
	{
		{{if ne .login $empty -}}
		'sso-host': '{{ getAppconf .login 1}}',
		'secret' : '{{getAppconf .login 2}}',
		'ident' : '{{getAppconf .login 3}}'
		{{- end}}
	}")
	//#api.appconf//
	{{- else -}}
	//api.appconf#//
	//#api.appconf//
	{{- end -}}
	{{- else -}}
	//api.appconf#//
	//#api.appconf//
	{{- end}}
	
	{{if fServer .serverType $api -}}
		{{if .cros -}}
		//api.cros#//
		s.Conf.API.SetSubConf('header', "
		{
			'Access-Control-Allow-Origin': '*', 
			'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
			'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type,__jwt__',
			'Access-Control-Allow-Credentials': 'true',
			'Access-Control-Expose-Headers':'__jwt__'
		}")
	//#api.cros//	
		{{- else -}}
	//api.cros#//
	//#api.cros//	
		{{- end -}}
	{{- else -}}
	//api.cros#//
	//#api.cros//
	{{- end}}
		
	{{if fServer .serverType $api -}}
		{{if .jwt}}
		//api.jwt.prod#//
		s.Conf.API.SetSubConf('auth', "{
			'jwt': {
				'exclude': ['/{{.projectName|lName}}/login'],
				'expireAt': 36000,
				'source':'H',
				'mode': 'HS512',
				'name': '__jwt__',
				'secret': '{{.prodSecret}}',
				'domain':'#domain'
			}
		}")	
	//#api.jwt.prod//
		{{- else -}}
		//api.jwt.prod#//
	//#api.jwt.prod//
		{{- end -}}
	{{- else -}}
		//api.jwt.prod#//
	//#api.jwt.prod//
	{{- end}}
	
	{{if fServer .serverType $api -}}
	{{if .metric -}}
	//api.metric#//
	s.Conf.API.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")	
//#api.metric//
	{{- else -}}
	//api.metric#//
	//#api.metric//
	{{- end -}}
	{{- else -}}
	//api.metric#//
	//#api.metric//
	{{- end}}

	{{if eq .db $empty -}}	
	//db#//
	//#db//
	{{- else -}}
	//db#//
	s.Conf.Plat.SetVarConf('db', 'db', "{			
		'provider':'{{.dbname}}',
		'connString':'{{.db}}',
		'maxOpen':20,
		'maxIdle':10,
		'lifeTime':600		
	}")
	//#db//
	{{- end}}

	{{if .cache -}}
	//cache#//
	s.Conf.Plat.SetVarConf('cache', 'cache', "
			{
				'proto':'redis',
				'addrs':[
						'192.168.0.111:6379',
						'192.168.0.112:6379',
						'192.168.0.113:6379',
						'192.168.0.114:6379',
						'192.168.0.115:6379',
						'192.168.0.116:6379'
				],
				'db':1,
				'dial_timeout':10,
				'read_timeout':10,
				'write_timeout':10,
				'pool_size':10
			}")
//#cache//
	{{- else -}}
	//cache#//
	//#cache//	
	{{- end}}
	
	{{if .queue -}}
	//queue#//
		s.Conf.Plat.SetVarConf('queue', 'queue', "
		{
			'proto':'redis',
			'addrs':[
					'192.168.0.111:6379',
					'192.168.0.112:6379',
					'192.168.0.113:6379',
					'192.168.0.114:6379',
					'192.168.0.115:6379',
					'192.168.0.116:6379'
			],
			'db':1,
			'dial_timeout':10,
			'read_timeout':10,
			'write_timeout':10,
			'pool_size':10
		}")
//#queue//
	{{- else -}}
	//queue#//
	//#queue//
	{{- end}}
	
	{{if fServer .serverType $cron -}}
	{{if .cron.appconf -}}
	//cron.appconf#//
		s.Conf.CRON.SetSubConf('app', "{
			'appname':'app_name'
		}")
//#cron.appconf//
	{{- else -}}
	//cron.appconf#//
	//#cron.appconf//
	{{- end -}}
	{{- else -}}
	//cron.appconf#//
	//#cron.appconf//
	{{- end}}
	
	{{if fServer .serverType $cron -}}
	{{if .cron.task -}}
	//cron.task#//
	s.Conf.CRON.SetSubConf('task', "{
		'tasks':[
		{'cron':'@every 1m','service':'/hello'}
		]		
	}")
//#cron.task//
	{{- else -}}
	//cron.task#//
	//#cron.task//
	{{- end -}}
	{{- else -}}
	//cron.task#//
	//#cron.task//
	{{- end}}

	{{if fServer .serverType $cron -}}
	{{if .cron.metric -}}
	//cron.metric#//
	s.Conf.CRON.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")	
//#cron.metric//
	{{- else -}}
	//cron.metric#//
	//#cron.metric//
	{{- end -}}
	{{- else -}}
	//cron.metric#//
	//#cron.metric//
	{{- end}}

	
	{{if fServer .serverType $mqc -}}
	{{if .mqc.server -}}
	//mqc.server#//
	s.Conf.MQC.SetSubConf("server", "
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.111:6379",
					"192.168.0.112:6379",
					"192.168.0.113:6379",
					"192.168.0.114:6379",
					"192.168.0.115:6379",
					"192.168.0.116:6379"
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}")
//#mqc.server//
	{{- else -}}
	//mqc.server#//
	//#mqc.server//
	{{- end -}}
	{{- else -}}
	//mqc.server#//
	//#mqc.server//
	{{- end}}

	{{if fServer .serverType $mqc -}}
	{{if .mqc.queue -}}
	//mqc.queue#//
	s.Conf.MQC.SetSubConf("queue", "{
		"queues":[
			{
				"queue":"cnp_make_coupon",
				"service":"/coupon/produce"
			},
			{
				"queue":"payment_mq",
				"service":"/order/pay"
			}
		]
	}")
//#mqc.queue//
	{{- else -}}
	//mqc.queue#//
	//#mqc.queue//
	{{- end -}}
	{{- else -}}
	//mqc.queue#//
	//#mqc.queue//
	{{- end}}

	{{if fServer .serverType $mqc -}}
	{{if .mqc.metric -}}
	//mqc.metric#//
	s.Conf.MQC.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
//#mqc.metric//
	{{- else -}}
	//mqc.metric#//
	//#mqc.metric//
	{{- end -}}
	{{- else -}}
	//mqc.metric#//
	//#mqc.metric//
	{{- end}}
	
	{{if fServer .serverType $web -}}
	//web.port#//
	s.Conf.WEB.SetMainConf("{'address':'{{.web.port}}'}")
//#web.port//
	{{- else -}}
	//web.port#//
	//#web.port//
	{{- end}}

	{{if fServer .serverType $web -}}
	{{if .web.static -}}
	//web.static#//
	s.Conf.WEB.SetSubConf('static', "{
		'dir':'./static',
		'rewriters':['*'],
		'exts':['.ttf','.woff','.woff2']			
	}")
//#web.static//	
	{{- else -}}
	//web.static#//
	//#web.static//	
	{{- end -}}
	{{- else -}}
	//web.static#//
	//#web.static//	
	{{- end}}

	{{if fServer .serverType $web -}}
	{{if .web.metric -}}
	//web.metric#//
	s.Conf.WEB.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
//#web.metric//	
	{{- else -}}
	//web.metric#//
	//#web.metric//	
	{{- end -}}
	{{- else -}}
	//web.metric#//
	//#web.metric//	
	{{- end}}

	{{if fServer .serverType $ws -}}
	{{if .ws.appconf -}}
	//ws.appconf#//
	s.Conf.WS.SetSubConf('app', "{
		'appname': 'gaea'
	}")
//#ws.appconf//
	{{- else -}}
	//ws.appconf#//
	//#ws.appconf//
	{{- end -}}
	{{- else -}}
	//ws.appconf#//
	//#ws.appconf//
	{{- end}}

	{{if fServer .serverType $ws -}}
	{{if .ws.jwt -}}
	//ws.jwt#//
	s.Conf.WS.SetSubConf('auth', "{
		'jwt': {
			'exclude': [],
			'source':'header',
			'expireAt': 36000,
			'mode': 'HS512',
			'name': '__jwt__',
			'secret': '{{.devSecret}}'
		}
	}")
//#ws.jwt//
	{{- else -}}
	//ws.jwt#//
	//#ws.jwt//
	{{- end -}}
	{{- else -}}
	//ws.jwt#//
	//#ws.jwt//
	{{- end}}

	{{if fServer .serverType $ws -}}
	{{if .ws.metric -}}
	//ws.metric#//
	s.Conf.WS.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
//#ws.metric//
	{{- else -}}
	//ws.metric#//
	//#ws.metric//
	{{- end -}}
	{{- else -}}
	//ws.metric#//
	//#ws.metric//
	{{- end}}

	{{if fServer .serverType $rpc -}}
	//rpc.port#//		
		s.Conf.RPC.SetMainConf("{'address':'{{.rpc.port}}'}")
	//#rpc.port//
	{{- else -}}
	//rpc.port#//
	//#rpc.port//
	{{- end}}
	
	{{if fServer .serverType $rpc -}}
	{{if .rpc.metric -}}
	//rpc.metric#//		
	s.Conf.RPC.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
//#rpc.metric//
	{{- else -}}
	//rpc.metric#//
	//#rpc.metric//
	{{- end -}}
	{{- else -}}
	//rpc.metric#//
	//#rpc.metric//
	{{- end}}
	
	//自定义安装程序
	s.Conf.API.Installer(func(c component.IContainer) error {
		if !s.Conf.Confirm('创建数据库表结构,添加基础数据?') {
			return nil
		}
		path, err := getSQLPath()
		if err != nil {
			return err
		}
		sqls, err := s.Conf.GetSQL(path)
		if err != nil {
			return err
		}
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		for _, sql := range sqls {
			if sql != '' {
				if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
					if !strings.Contains(err.Error(), 'ORA-00942') {
						s.Conf.Log.Errorf('执行SQL失败： %v %s\n', err, q)
					}
				}
			}
		}
		return nil
	})
}

//getSQLPath 获取getSQLPath
func getSQLPath() (string, error) {
	gopath := os.Getenv('GOPATH')
	if gopath == '' {
		return '', fmt.Errorf('未配置环境变量GOPATH')
	}
	path := strings.Split(gopath, ';')
	if len(path) == 0 {
		return '', fmt.Errorf('环境变量GOPATH配置的路径为空')
	}
	return filepath.Join(path[0], 'src/{{.projectName}}/modules/const/sql'), nil
}`
