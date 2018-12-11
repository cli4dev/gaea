package dev

//PlatVarDB .
const PlatVarDB = `//db#//
	s.Conf.Plat.SetVarConf('db', 'db', "{			
		'provider':'{{.dbname}}',
		'connString':'{{.db}}',
		'maxOpen':20,
		'maxIdle':10,
		'lifeTime':600		
	}")
	//#db//`

//PlatVarCache .
const PlatVarCache = `//cache#//
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
	//#cache//`

//PlatVarQueue .
const PlatVarQueue = `//queue#//
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
	//#queue//`
