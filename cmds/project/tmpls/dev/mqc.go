package dev

//MqcSubServer .
const MqcSubServer = `//mqc.server#//
	s.Conf.MQC.SetSubConf('server', "
		{
			'proto':'redis',
			'addrs':[
					'192.168.0.111:6379',
					'192.168.0.112:6379',
					'192.168.0.113:6379',
					'192.168.0.114:6379',
					'192.168.0.115:6379',
					'192.168.0.116:6379'
			],
			'db':1,
			'dial_timeout':10,
			'read_timeout':10,
			'write_timeout':10,
			'pool_size':10
	}")
	//#mqc.server//`

//MqcSubQueue .
const MqcSubQueue = `//mqc.queue#//
	s.Conf.MQC.SetSubConf('queue', "{
		'queues':[
			{
				'queue':'cnp_make_coupon',
				'service':'/coupon/produce'
			},
			{
				'queue':'payment_mq',
				'service':'/order/pay'
			}
		]
	}")
	//#mqc.queue//`
