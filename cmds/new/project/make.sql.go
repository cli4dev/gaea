package project

import (
	"fmt"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/new/sql/insert"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
)

//makeInsertSQL .
//生成inster sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeInsertSQL(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(tables, "", filters)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	for k, v := range tmpls {
		fmt.Println(k)
		fmt.Println(v)
	}
	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个语句", len(tmpls))
	return nil
}

//makeSelectSQL .
//生成select sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeSelectSQL(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(tables, "", filters)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	for k, v := range tmpls {
		fmt.Println(k)
		fmt.Println(v)
	}
	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个语句", len(tmpls))
	return nil
}

//makeUpdateSQL .
//生成update sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeUpdateSQL(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(tables, "", filters)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	for k, v := range tmpls {
		fmt.Println(k)
		fmt.Println(v)
	}
	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个语句", len(tmpls))
	return nil
}

//makeDeleteSQL .
//生成delete sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeDeleteSQL(filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(tables, "", filters)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	for k, v := range tmpls {
		fmt.Println(k)
		fmt.Println(v)
	}
	if err = createFile(outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("生成完成,共生成%d个语句", len(tmpls))
	return nil
}
