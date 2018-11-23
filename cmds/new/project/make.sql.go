package project

import (
	"github.com/micro-plat/gaea/cmds"
	"github.com/urfave/cli"

	"github.com/micro-plat/gaea/cmds/new/sql/delete"
	"github.com/micro-plat/gaea/cmds/new/sql/insert"
	"github.com/micro-plat/gaea/cmds/new/sql/md"
	"github.com/micro-plat/gaea/cmds/new/sql/read"
	"github.com/micro-plat/gaea/cmds/new/sql/update"
)

//makeInsertSQL .
//生成inster sql语句
//@filePath md文件路径
//@filters 表名
//@outPath 输出路径，为空则根据默认规则输出
func (p *moduleCmd) makeInsertSQL(c *cli.Context, filePath string, filters []string, outPath string) error {
	//获取默认输出路径
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := insert.GetTmples(insert.InsertTmpl, tables, "", filters, false)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}

	if err = createFile(c, outPath, tmpls); err != nil {
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
func (p *moduleCmd) makeSelectSQL(c *cli.Context, filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := read.GetTmples(read.SelectTmpl, tables, "", filters, false)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	// for k, v := range tmpls {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	if err = createFile(c, outPath, tmpls); err != nil {
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
func (p *moduleCmd) makeUpdateSQL(c *cli.Context, filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := update.GetTmples(update.UpdateTmpl, tables, "", filters, false)

	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}
	// for k, v := range tmpls {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	if err = createFile(c, outPath, tmpls); err != nil {
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
func (p *moduleCmd) makeDeleteSQL(c *cli.Context, filePath string, filters []string, outPath string) error {
	if outPath == "" {
		outPath = cmds.GetLocation() + "/const/sql"
	}
	//获取默认输出路径
	tables, err := md.Markdown2Table(filePath)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	tmpls, err := delete.GetTmples(delete.DeleteTmpl, tables, "", filters, false)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("%s中未找到数据表信息", filePath)
		return nil
	}

	if err = createFile(c, outPath, tmpls); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("发现%d个delete语句", len(tmpls))
	return nil
}
