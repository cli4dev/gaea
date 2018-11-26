package service

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/service/tmpls"
	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/data"
)

//makeGetAndQueryHandle .
func (p *serviceCmd) makeGetAndQueryHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("get and query handle", tmpls.GetAndQueryTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成get and query handle时未找到数据表信息")
		return nil, err
	}

	return tmpls, nil
}

//makePostHandle .
func (p *serviceCmd) makePostHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("post handle", tmpls.PostTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成post handle时未找到数据表信息")
		return nil, err
	}

	return tmpls, nil
}

//makePutHandle .
func (p *serviceCmd) makePutHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("put handle", tmpls.PutTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成put handle时未找到数据表信息")
		return nil, err
	}

	return tmpls, nil
}

//makeDeleteHandle .
func (p *serviceCmd) makeDeleteHandle(add, cover bool, tables []*conf.Table, filters []string, servicePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetServerTmples("delete handle", tmpls.DeleteTpl, tables, filters, servicePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成delete handle时未找到数据表信息")
		return nil, err
	}

	return tmpls, nil
}
