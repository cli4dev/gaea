package dev

//RPCMainPort .
const RPCMainPort = `
	//rpc.port#//
	s.Conf.API.SetMainConf("{'address':'{{.port}}'}")
	//#rpc.port//
`
