package tmpls

const installDevTmpl = `// +build !prod

package main

//install 用于配置开发环境参数，执行 install 命令时执行此函数进行配置安装。此处可以执行服务端口，跨域配置，数据库，缓存，消息队列，应用程序等配置
func (s *{{.projectName|lName}}) install() {
	s.Conf.API.SetMainConf("{'address':':9091'}")
	s.Conf.API.SetSubConf('header', "
				{
					'Access-Control-Allow-Origin': '*', 
					'Access-Control-Allow-Methods': 'GET,POST,PUT,DELETE,PATCH,OPTIONS', 
					'Access-Control-Allow-Credentials': 'true'
				}
			")

	s.Conf.API.SetSubConf('auth', "
		{
			'jwt': {
				'exclude': ['/{{.projectName|lName}}/login'],
				'expireAt': 36000,
				'mode': 'HS512',
				'name': '{{.projectName|lName}}_sid',
				'secret': '12345678'
			}
		}
		")
	s.Conf.Plat.SetVarConf('db', 'db', "{			
			'provider':'ora',
			'connString':'sso/123456@orcl136',
			'maxOpen':20,
			'maxIdle':10,
			'lifeTime':600		
	}")

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
	}
		
		")
}


`
