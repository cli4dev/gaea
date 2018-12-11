package dev

//CronSubApp .
const CronSubApp = `//cron.appconf#//
	s.Conf.CRON.SetSubConf('app', "{
	}")
	//#cron.appconf//`

//CronSubTask .
const CronSubTask = `//cron.task#//
	s.Conf.CRON.SetSubConf('task', "{
		'tasks':[
		{'cron':'@every 1m','service':'/hello'}
		]		
	}")
	//#cron.task//`

//CronSubMetric .
const CronSubMetric = `//cron.metric#//
	s.Conf.CRON.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")	
	//#cron.metric//`
