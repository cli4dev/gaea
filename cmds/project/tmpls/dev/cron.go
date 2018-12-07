package dev

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

//CronSubApp .
const CronSubApp = `
//cron.sub.app#//
	s.Conf.CRON.SetSubConf('app', "{
		'appname':'app_name'
	}")
//#cron.sub.app//
`

//CronSubTask .
const CronSubTask = `
//cron.sub.task#//
	s.Conf.CRON.SetSubConf('task', "{
		'tasks':[
		{'cron':'@every 1m','service':'/hello'}
		]		
	}")
//#cron.sub.task//
`
