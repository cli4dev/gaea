package service

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/service/tpl"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
)

//makeGetAndQueryHandle .
func (p *serviceCmd) makeGetAndQueryHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("get and query handle", tpl.GetAndQueryTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	return tmpls, nil
}

//makePostHandle .
func (p *serviceCmd) makePostHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("post handle", tpl.PostTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	return tmpls, nil
}

//makePutHandle .
func (p *serviceCmd) makePutHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("put handle", tpl.PutTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	return tmpls, nil
}

//makeDeleteHandle .
func (p *serviceCmd) makeDeleteHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("delete handle", tpl.DeleteTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	return tmpls, nil
}
