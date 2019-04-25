package module

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/module/tmpls"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
)

// makeInsertSQL 生成inster sql语句
func (p *moduleCmd) makeInsertSQL(add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tplName := tmpls.InsertMysqlTmpl
	if db != "mysql" {
		tplName = tmpls.InsertOracleTmpl
	}
	tmpls, err := data.GetTmples("insert sql", tplName, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 insert sql 时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil
}

// makeSelectSQL 生成select sql语句
func (p *moduleCmd) makeSelectSQL(add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tplName := tmpls.SelectOracleTmpl
	if db == "mysql" {
		tplName = tmpls.SelectMysqlTmpl
	}

	tmpls, err := data.GetTmples("select sql", tplName, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 select sql 时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil
}

// makeUpdateSQL 生成update sql语句
func (p *moduleCmd) makeUpdateSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tmpls, err := data.GetTmples("update sql", tmpls.UpdateTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 update sql 时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil
}

// makeDeleteSQL 生成delete sql语句
func (p *moduleCmd) makeDeleteSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) (out map[string]map[string]string, err error) {

	tmpls, err := data.GetTmples("delete sql", tmpls.DeleteTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 delete sql 时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil
}
