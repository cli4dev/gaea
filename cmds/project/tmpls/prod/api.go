package prod

//Api .
const Api = `
	{ //api
		s.Conf.API.SetMainConf("{'address':':9091'}")
		s.Conf.API.SetSubConf('header', "
				{
					'Access-Control-Allow-Origin': '*', 
					'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
					'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type',
					'Access-Control-Allow-Credentials': 'true'
				}
			")

		s.Conf.API.SetSubConf('auth', "{
			'jwt': {
				'exclude': ['/{{.projectName|lName}}/login'],
				'expireAt': 36000,
				'mode': 'HS512',
				'name': '{{.projectName|lName}}_sid',
				'secret': '12345678'
			}
		}")
	}
}
`

//APIMainPort .
const APIMainPort = `
//api.main.port#//
	s.Conf.API.SetMainConf("{'address':':9091'}")
//#api.main.port//
`

//APISubHeader .
const APISubHeader = `
//api.sub.header#//
	s.Conf.API.SetSubConf('header', "
	{
		'Access-Control-Allow-Origin': '*', 
		'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
		'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type',
		'Access-Control-Allow-Credentials': 'true'
	}")
//#api.sub.header//
`

//APISubAuth .
const APISubAuth = `
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
`

//APISubMetric .
const APISubMetric = `
//api.sub.metric#//
	s.Conf.API.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName":'',
		'password":''
	}")	
//#api.sub.metric//
`
