package prod

const Web = `
	{ //web
	
		s.Conf.WEB.SetMainConf("{'address':':8080'}")
		s.Conf.WEB.SetSubConf('static', "{
			'dir':'./static',
			'rewriters':['*'],
			'exts':['.ttf','.woff','.woff2']			
		}")
	}	
}
`

//WebMainPort .
const WebMainPort = `
//web.main.port#//
s.Conf.WEB.SetMainConf("{'address':':8080'}")
//#web.main.port//
`

//WebSubStatic .
const WebSubStatic = `
//web.sub.static#//
	s.Conf.WEB.SetSubConf('static', "{
		'dir':'./static',
		'rewriters':['*'],
		'exts':['.ttf','.woff','.woff2']			
	}")
//#web.sub.static//	
`
