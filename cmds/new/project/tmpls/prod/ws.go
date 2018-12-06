package prod

//WSSubAPP .
const WSSubAPP = `
//ws.sub.app#//
	s.Conf.WS.SetSubConf('app', "{
		'appname': 'gaea'
	}")
//#ws.sub.app//
`

//WSSubAuth .
const WSSubAuth = `
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
`
