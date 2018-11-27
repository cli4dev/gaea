package prod

const Web = `
	{//web
		s.Conf.WEB.SetMainConf("{'address':':8080'}")
		s.Conf.WEB.SetSubConf('static', "{
				'dir':'./static',
				'rewriters':['*'],
				'exts':[".ttf','.woff','.woff2']			
		}")
	}
}
`
