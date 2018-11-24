package module

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/module/tmpls"
	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/data"
)

//makeInsertSQL .
//生成inster sql语句
func (p *moduleCmd) makeInsertSQL(add, cover bool, db string, tables []*conf.Table, filters []string, modulePath string) error {

	tplName := tmpls.InsertMysqlTmpl
	if db != "mysql" {
		tplName = tmpls.InsertOracleTmpl
	}
	tmpls, err := data.GetTmples("insert sql", tplName, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 insert sql 时未找到数据表信息")
		return nil
	}

	cmds.Log.Infof("发现%d个insert语句", len(tmpls))
	//创建文件
	if err = data.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}

	return nil
}

//makeSelectSQL .
//生成select sql语句
func (p *moduleCmd) makeSelectSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := data.GetTmples("select sql", tmpls.SelectTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 select sql 时未找到数据表信息")
		return nil
	}

	cmds.Log.Infof("发现%d个select语句", len(tmpls))
	//创建文件
	if err = data.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}

	return nil
}

//makeUpdateSQL .
//生成update sql语句
func (p *moduleCmd) makeUpdateSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := data.GetTmples("update sql", tmpls.UpdateTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 update sql 时未找到数据表信息")
		return nil
	}

	cmds.Log.Infof("发现%d个update语句", len(tmpls))
	//创建文件
	if err = data.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}

	return nil
}

//makeDeleteSQL .
//生成delete sql语句
func (p *moduleCmd) makeDeleteSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := data.GetTmples("delete sql", tmpls.DeleteTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 delete sql 时未找到数据表信息")
		return nil
	}

	cmds.Log.Infof("发现%d个delete语句", len(tmpls))
	//创建文件
	if err = data.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}

	return nil
}
