package module

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/module/tmpls"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
)

//makeInsertFunc .
//生成inster 函数
func (p *moduleCmd) makeInsertFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetTmples("insert func", tmpls.InsertFunc, tables, filters, true, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成 insert函数时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil

}

//makeSelectFunc .
//生成select 函数
func (p *moduleCmd) makeSelectFunc(add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {
	tplName := tmpls.SelectOracleFunc
	if db == "mysql" {
		tplName = tmpls.SelectMysqlFunc
	}

	tmpls, err := data.GetTmples("select func", tplName, tables, filters, true, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 select函数时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil

}

//makeUpdateFunc .
//生成update 函数
func (p *moduleCmd) makeUpdateFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tmpls, err := data.GetTmples("update func", tmpls.UpdateFunc, tables, filters, true, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 update函数时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil

}

//makeDeleteFunc .
//生成delete 函数
func (p *moduleCmd) makeDeleteFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tmpls, err := data.GetTmples("delete func", tmpls.DeleteFunc, tables, filters, true, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成 delete函数时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil
}
