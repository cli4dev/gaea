package dev

//APIMainPort .
const APIMainPort = `
	//api.port#//
	s.Conf.API.SetMainConf("{'address':'{{.port}}'}")
	//#api.port//
`

//APISubHeaderDomain 设置跨域的头
const APISubCros = `
	//api.cros#//
	s.Conf.API.SetSubConf('header', "
	{
		'Access-Control-Allow-Origin': '*', 
		'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
		'Access-Control-Allow-Headers': 'X-Requested-With,Content-Type',
		'Access-Control-Allow-Credentials': 'true'
	}")
	//#api.cros//	
`

//APISubAuth .
const APISubAuth = `
	//api.jwt#//
	s.Conf.API.SetSubConf('auth', "{
		'jwt': {
			'exclude': ['/{{.projectName|lName}}/login'],
			'expireAt': 36000,
			'mode': 'HS512',
			'name': '{{.projectName|lName}}_sid',
			'secret': '{{.devSecret}}'
		}
	}")	
	//#api.jwt//
`

//APISubMetric .
const APISubMetric = `
	//api.metric#//
	s.Conf.API.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")	
	//#api.metric//
`

//HandlingJWT .
const HandlingJWT = `
	//handling.jwt#//
	jwt, err := ctx.Request.GetJWTConfig() //获取jwt配置
	if err != nil {
		return err
	}
	for _, u := range jwt.Exclude { //排除指定请求
		if u == ctx.Service {
			return nil
		}
	}
	//#handling.jwt//
`

const APIApp = `
	//api.appconf#//
	s.Conf.API.SetSubConf('app', "
	{
	}")
	//#api.appconf//
`
