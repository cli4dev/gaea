package dev

const Plat = `
	{ //plat
		s.Conf.Plat.SetVarConf('db', 'db', '{			
			'provider':'ora',
			'connString':'sso/123456@orcl136',
			'maxOpen':20,
			'maxIdle':10,
			'lifeTime':600		
		}')
		
		s.Conf.Plat.SetVarConf('cache', 'cache', '
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
		}')
	}	
}
`

//PlatVarDB .
const PlatVarDB = `
//plat.var.db#//
	s.Conf.Plat.SetVarConf('db', 'db', "{			
		'provider':'ora',
		'connString':'{{.db}}',
		'maxOpen':20,
		'maxIdle':10,
		'lifeTime':600		
	}")
//#plat.var.db//
`

//PlatVarCache .
const PlatVarCache = `
//plat.var.cache#//
	s.Conf.Plat.SetVarConf('cache', 'cache', "
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
//#plat.var.cache//			
`

//PlatVarQueue .
const PlatVarQueue = `
//plat.var.queue#//
	s.Conf.Plat.SetVarConf('queue', 'queue', "
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
//plat.var.queue#//
`