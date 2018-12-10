package dev

//WSSubAPP .
const WSSubAPP = `
	//ws.app#//
	s.Conf.WS.SetSubConf('app', "{
	}")
	//#ws.app//
`

//WSSubAuth .
const WSSubAuth = `
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
`
