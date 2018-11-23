package module

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/util/data"
	"github.com/micro-plat/gaea/cmds/new/util/tb"
)

//makeInsertFunc .
//生成inster 函数
func (p *moduleCmd) makeInsertFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {
	//获取模板数据

	tmpls, err := util.GetTmples("insert func", tpl.InsertFunc, tables, filters, true, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Error("生成 insert函数时未找到数据表信息")
		return nil
	}
	cmds.Log.Infof("发现%d个insert函数", len(tmpls))
	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//makeSelectFunc .
//生成select 函数
func (p *moduleCmd) makeSelectFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("select func", tpl.SelectFunc, tables, filters, true, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 select函数时未找到数据表信息")
		return nil
	}
	cmds.Log.Infof("发现%d个select函数", len(tmpls))
	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}

//makeUpdateFunc .
//生成update 函数
func (p *moduleCmd) makeUpdateFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("update func", tpl.UpdateFunc, tables, filters, true, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 update函数时未找到数据表信息")
		return nil
	}
	cmds.Log.Infof("发现%d个update函数", len(tmpls))
	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}

	return nil
}

//makeDeleteFunc .
//生成delete 函数
func (p *moduleCmd) makeDeleteFunc(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("delete func", tpl.DeleteFunc, tables, filters, true, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Error("生成 delete函数时未找到数据表信息")
		return nil
	}
	cmds.Log.Infof("发现%d个delete函数", len(tmpls))
	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	return nil
}
