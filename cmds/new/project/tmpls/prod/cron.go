package prod

const Cron = `
	{ //cron
		s.Conf.CRON.SetSubConf('app', "{
			'appname':'app_name'
		}")
		
		s.Conf.CRON.SetSubConf('task', "{
			'tasks':[
			{'cron':'@every 1m','service':'/hello'}
			]		
		}")
	}
}
`
