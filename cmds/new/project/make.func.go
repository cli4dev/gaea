package project

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/sql/delete"
	"github.com/micro-plat/gaea/cmds/new/sql/insert"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
	"github.com/micro-plat/gaea/cmds/new/sql/read"
	"github.com/micro-plat/gaea/cmds/new/sql/update"
)

//makeInsertFunc .
//生成inster 函数
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeInsertFunc(filePath string, filters []string, outPath string) error {
	//获取默认输出路径
	if outPath == "" {
		outPath = cmds.GetLocation()
	}
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(insert.InsertFunc, tables, "", filters, true)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}

	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个insert函数", len(tmpls))
	return nil
}

//makeSelectFunc .
//生成select 函数
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeSelectFunc(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation()
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := read.GetTmples(read.SelectFunc, tables, "", filters, true)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}

	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个select函数", len(tmpls))
	return nil
}

//makeUpdateFunc .
//生成update 函数
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeUpdateFunc(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation()
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := update.GetTmples(update.UpdateFunc, tables, "", filters, true)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}

	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个update函数", len(tmpls))
	return nil
}

//makeDeleteFunc .
//生成delete 函数
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeDeleteFunc(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation()
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := delete.GetTmples(delete.DeleteFunc, tables, "", filters, true)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个delete函数", len(tmpls))
	return nil
}
