package prod

//Api .
const Api = `
	{ //api
		s.Conf.API.SetMainConf("{'address':':#api_port'}")
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
