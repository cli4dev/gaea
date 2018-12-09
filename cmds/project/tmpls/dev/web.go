package dev

//WebMainPort .
const WebMainPort = `
	//web.port#//
	s.Conf.WEB.SetMainConf("{'address':':8080'}")
	//#web.port//
`

//WebSubStatic .
const WebSubStatic = `
	//web.static#//
	s.Conf.WEB.SetSubConf('static', "{
		'dir':'./static',
		'rewriters':['*'],
		'exts':['.ttf','.woff','.woff2']			
	}")
	//#web.static//	
`
