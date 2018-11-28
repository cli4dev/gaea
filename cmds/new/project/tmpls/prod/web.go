package prod

const Web = `
	{ //web
		s.Conf.WEB.SetMainConf("{'address':':#web_port'}")
		s.Conf.WEB.SetSubConf('static', "{
				'dir':'./static',
				'rewriters':['*'],
				'exts':['.ttf','.woff','.woff2']			
		}")
	}
}
`
