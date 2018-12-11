package dev

//RPCMainPort .
const RPCMainPort = `//rpc.port#//
	s.Conf.RPC.SetMainConf("{'address':'{{.port}}'}")
	//#rpc.port//`

//RPCSubMetric .
const RPCSubMetric = `//rpc.metric#//
	s.Conf.RPC.SetSubConf('metric', "{
		'host':'http://192.168.106.219:8086',
		'dataBase':'gcr',
		'cron':'@every 10s',
		'userName':'',
		'password':''
	}")
	//#rpc.metric//`
