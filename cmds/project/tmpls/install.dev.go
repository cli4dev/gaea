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
		//api.port#//
			s.Conf.API.SetMainConf("{'address':'{{.port}}'}")
	//#api.port//
		{{- else -}}
			//api.port#//
			//#api.port//
		{{- end}}

		{{if fServer .serverType $api -}}
		{{if .appconf -}}
			//api.appconf#//
			s.Conf.API.SetSubConf('app', "
			{
				'appname':'app_name'
			}")
		//#api.appconf//
		{{- else -}}
		//api.appconf#//
		//#api.appconf//
		{{- end}}
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
					'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type',
					'Access-Control-Allow-Credentials': 'true'
				}")
		//#api.cros//	
			{{else}}
			//api.cros#//
			//#api.cros//	
			{{end}}
		{{- else}}
			//api.cros#//
		//#api.cros//
		{{end}}
	
	
		
		{{if fServer .serverType $api -}}
			{{if .jwt}}
				//api.jwt#//
				s.Conf.API.SetSubConf('auth', "{
					'jwt': {
						'exclude': ['/{{.projectName|lName}}/login'],
						'expireAt': 36000,
						'mode': 'HS512',
						'name': '{{.projectName|lName}}_sid',
						'secret': '12345678'
					}
				}")	
		//#api.jwt//
			{{- else -}}
			//api.jwt#//
		//#api.jwt//
			{{- end -}}
		{{- else}}
			//api.jwt#//
		//#api.jwt//
		{{end}}
	
	
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
		{{- else}}
		//api.metric#//
		//#api.metric//
		{{end}}
	

	
		

		{{if eq .db $empty }}	
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
	

		{{if .cache}}
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
	
			
	

		{{if .queue}}
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
			//cron.app#//
			s.Conf.CRON.SetSubConf('app', "{
				'appname':'app_name'
			}")
	//#cron.app//
		{{- else}}
		//cron.app#//
		//#cron.app//
		{{end}}
	
	
		{{if fServer .serverType $cron -}}
			//cron.task#//
			s.Conf.CRON.SetSubConf('task', "{
				'tasks':[
				{'cron':'@every 1m','service':'/hello'}
				]		
			}")
	//#cron.task//
		{{- else}}
		//cron.task#//
		//#cron.task//
		{{end}}
	

		{{if fServer .serverType $mqc -}}
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
		{{- else}}
		//mqc.server#//
		//#mqc.server//
		{{end}}
	

		{{if fServer .serverType $mqc -}}
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
		{{- else}}
		//mqc.queue#//
		//#mqc.queue//
		{{end}}
	
	
		{{if fServer .serverType $web -}}
			//web.port#//
			s.Conf.WEB.SetMainConf("{'address':'{{.port}}'}")
		//#web.port//
		{{- else}}
		//web.port#//
		//#web.port//
		{{end}}
	

		{{if fServer .serverType $web -}}
			//web.static#//
			s.Conf.WEB.SetSubConf('static', "{
				'dir':'./static',
				'rewriters':['*'],
				'exts':['.ttf','.woff','.woff2']			
			}")
		//#web.static//	
		{{- else}}
		//web.static#//
		//#web.static//	
		{{end}}
	

		{{if fServer .serverType $ws -}}
			//ws.app#//
			s.Conf.WS.SetSubConf('app', "{
				'appname': 'gaea'
			}")
		//#ws.app//
		{{- else}}
		//ws.app#//
		//#ws.app//
		{{end}}
	

		{{if fServer .serverType $ws -}}
			//ws.auth#//
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
		//#ws.auth//
		{{- else}}
		//ws.auth#//
		//#ws.auth//
		{{end}}
	

		{{if fServer .serverType $rpc -}}
		//rpc.port#//
			s.Conf.API.SetMainConf("{'address':':8090'}")
		//#rpc.port//
		{{- else}}
		//rpc.port#//
		//#rpc.port//
		{{end}}
	
}`
