package dev

//CronSubApp .
const CronSubApp = `
	//cron.app#//
	s.Conf.CRON.SetSubConf('app', "{
		'appname':'app_name'
	}")
	//#cron.app//
`

//CronSubTask .
const CronSubTask = `
	//cron.task#//
	s.Conf.CRON.SetSubConf('task', "{
		'tasks':[
		{'cron':'@every 1m','service':'/hello'}
		]		
	}")
	//#cron.task//
`
