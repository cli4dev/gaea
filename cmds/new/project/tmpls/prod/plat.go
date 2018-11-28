package prod

const Plat = `
	{ //plat
		s.Conf.Plat.SetVarConf('db', 'db', "{			
			'provider':'ora',
			'connString':'#db_string',
			'maxOpen':20,
			'maxIdle':10,
			'lifeTime':600		
		}")
		
		s.Conf.Plat.SetVarConf('cache', 'cache', "
		{
			'proto':'redis',
			'addrs':[
					#redis
			],
			'db':1,
			'dial_timeout':10,
			'read_timeout':10,
			'write_timeout':10,
			'pool_size':10
		}")
	}
}
`
