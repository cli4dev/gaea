package conf

//WSSubAPP .
const WSSubAPP = `//ws.appconf#//
	s.Conf.WS.SetSubConf('app', "{
	}")
	//#ws.appconf//`

//WSSubAuth .
const WSSubAuth = `//ws.jwt#//
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
	//#ws.jwt//`

//WSSubMetric .
const WSSubMetric = `//ws.metric#//
	s.Conf.WS.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
	//#ws.metric//`
