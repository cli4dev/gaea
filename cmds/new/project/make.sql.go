package project

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/sql/conf"
	"github.com/micro-plat/gaea/cmds/new/sql/tpl"
	"github.com/micro-plat/gaea/cmds/new/sql/util"
)

//makeInsertSQL .
//生成inster sql语句
//@filePath md文件路径
//@filters 表名
//@modulePath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeInsertSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("insert sql", tpl.InsertTmpl, tables, filters, false, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 insert sql 时未找到数据表信息")
		return nil
	}

	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("发现%d个insert语句", len(tmpls))
	return nil
}

//makeSelectSQL .
//生成select sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeSelectSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("select sql", tpl.SelectTmpl, tables, filters, false, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 select sql 时未找到数据表信息")
		return nil
	}

	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("发现%d个select语句", len(tmpls))
	return nil
}

//makeUpdateSQL .
//生成update sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeUpdateSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("update sql", tpl.UpdateTmpl, tables, filters, false, modulePath)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 update sql 时未找到数据表信息")
		return nil
	}

	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("发现%d个update语句", len(tmpls))
	return nil
}

//makeDeleteSQL .
//生成delete sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeDeleteSQL(add, cover bool, tables []*conf.Table, filters []string, modulePath string) error {

	tmpls, err := util.GetTmples("delete sql", tpl.DeleteTmpl, tables, filters, false, modulePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("生成 delete sql 时未找到数据表信息")
		return nil
	}

	if err = util.CreateFile(add, cover, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("发现%d个delete语句", len(tmpls))
	return nil
}
