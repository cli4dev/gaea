package conf

//WebMainPort .
const WebMainPort = `//web.port#//
	s.Conf.WEB.SetMainConf("{'address':'{{.port}}'}")
	//#web.port//`

//WebSubStatic .
const WebSubStatic = `//web.static#//
	s.Conf.WEB.SetSubConf('static', "{
		'dir':'./static',
		'rewriters':['*'],
		'exts':['.ttf','.woff','.woff2']			
	}")
	//#web.static//`

//WebSubMetric .
const WebSubMetric = `//web.metric#//
	s.Conf.WEB.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
	//#web.metric//`
