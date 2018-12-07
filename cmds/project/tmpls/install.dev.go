package tmpls

const installDevTmpl = `// +build !prod

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
		//api.main.port#//
			s.Conf.API.SetMainConf("{'address':'{{.port}}'}")
	//#api.main.port//
		{{- else -}}
			//api.main.port#//
			//#api.main.port//
		{{- end}}
	
	
	
		{{if fServer .serverType $api -}}
			{{if .cros -}}
			//api.sub.header#//
				s.Conf.API.SetSubConf('header', "
				{
					'Access-Control-Allow-Origin': '*', 
					'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
					'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type',
					'Access-Control-Allow-Credentials': 'true'
				}")
		//#api.sub.header//	
			{{else}}
			//api.sub.header#//
			//#api.sub.header//	
			{{end}}
		{{- else}}
			//api.sub.header#//
		//#api.sub.header//
		{{end}}
	
	
		
		{{if fServer .serverType $api -}}
			{{if .jwt}}
				//api.sub.auth#//
				s.Conf.API.SetSubConf('auth', "{
					'jwt': {
						'exclude': ['/{{.projectName|lName}}/login'],
						'expireAt': 36000,
						'mode': 'HS512',
						'name': '{{.projectName|lName}}_sid',
						'secret': '12345678'
					}
				}")	
		//#api.sub.auth//
			{{- else -}}
			//api.sub.auth#//
		//#api.sub.auth//
			{{- end -}}
		{{- else}}
			//api.sub.auth#//
		//#api.sub.auth//
		{{end}}
	
	
		{{if fServer .serverType $api -}}
		{{if .metric -}}
		//api.sub.metric#//
		s.Conf.API.SetSubConf('metric', "{
			'host':'http://192.168.106.219:8086',
			'dataBase':'gcr',
			'cron':'@every 10s',
			'userName':'',
			'password':''
		}")	
	//#api.sub.metric//
		{{- else -}}
		//api.sub.metric#//
		//#api.sub.metric//
		{{- end -}}
		{{- else}}
		//api.sub.metric#//
		//#api.sub.metric//
		{{end}}
	

	
		

		{{if eq .db $empty }}	
		//plat.var.db#//
		//#plat.var.db//
		{{- else -}}
		//plat.var.db#//
		s.Conf.Plat.SetVarConf('db', 'db', "{			
			'provider':'{{.dbname}}',
			'connString':'{{.db}}',
			'maxOpen':20,
			'maxIdle':10,
			'lifeTime':600		
		}")
	//#plat.var.db//
		{{- end}}
	

		{{if .cache}}
		//plat.var.cache#//
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
	//#plat.var.cache//
		{{- else -}}
		//plat.var.cache#//
		//#plat.var.cache//	
		{{- end}}
	
			
	

		{{if .queue}}
		//plat.var.queue#//
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
	//#plat.var.queue//
		{{- else -}}
		//plat.var.queue#//
		//#plat.var.queue//
		{{- end}}
	
	
	
		{{if fServer .serverType $cron -}}
			//cron.sub.app#//
			s.Conf.CRON.SetSubConf('app', "{
				'appname':'app_name'
			}")
	//#cron.sub.app//
		{{- else}}
		//cron.sub.app#//
		//#cron.sub.app//
		{{end}}
	
	
		{{if fServer .serverType $cron -}}
			//cron.sub.task#//
			s.Conf.CRON.SetSubConf('task', "{
				'tasks':[
				{'cron':'@every 1m','service':'/hello'}
				]		
			}")
	//#cron.sub.task//
		{{- else}}
		//cron.sub.task#//
		//#cron.sub.task//
		{{end}}
	

		{{if fServer .serverType $mqc -}}
			//mqc.sub.server#//
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
	//#mqc.sub.server//
		{{- else}}
		//mqc.sub.server#//
		//#mqc.sub.server//
		{{end}}
	

		{{if fServer .serverType $mqc -}}
			//mqc.sub.queue#//
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
	//#mqc.sub.queue//
		{{- else}}
		//mqc.sub.queue#//
		//#mqc.sub.queue//
		{{end}}
	
	
		{{if fServer .serverType $web -}}
			//web.main.port#//
			s.Conf.WEB.SetMainConf("{'address':'{{.port}}'}")
		//#web.main.port//
		{{- else}}
		//web.main.port#//
		//#web.main.port//
		{{end}}
	

		{{if fServer .serverType $web -}}
			//web.sub.static#//
			s.Conf.WEB.SetSubConf('static', "{
				'dir':'./static',
				'rewriters':['*'],
				'exts':['.ttf','.woff','.woff2']			
			}")
		//#web.sub.static//	
		{{- else}}
		//web.sub.static#//
		//#web.sub.static//	
		{{end}}
	

		{{if fServer .serverType $ws -}}
			//ws.sub.app#//
			s.Conf.WS.SetSubConf('app', "{
				'appname': 'gaea'
			}")
		//#ws.sub.app//
		{{- else}}
		//ws.sub.app#//
		//#ws.sub.app//
		{{end}}
	

		{{if fServer .serverType $ws -}}
			//ws.sub.auth#//
			s.Conf.WS.SetSubConf('auth', "{
				'jwt': {
					'exclude': [],
					'source':'header',
					'expireAt': 36000,
					'mode': 'HS512',
					'name': '__jwt__',
					'secret': '12345678'
				}
			}")
		//#ws.sub.auth//
		{{- else}}
		//ws.sub.auth#//
		//#ws.sub.auth//
		{{end}}
	

		{{if fServer .serverType $rpc -}}
		//rpc.main.port#//
			s.Conf.API.SetMainConf("{'address':':8090'}")
		//#rpc.main.port//
		{{- else}}
		//rpc.main.port#//
		//#rpc.main.port//
		{{end}}
	
}`
