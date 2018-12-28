package sql

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/md"
)

type handler func([]*conf.Table) (map[string]string, error)

func create(tpl handler, mdFilePath []string, outPath string, cover bool, filters ...string) error {

	var tables = make([]*conf.Table, 0, 2)
	for _, p := range mdFilePath {
		t1, err := md.Markdown2Table(p)
		if err != nil {
			cmds.Log.Error(err)
			return err
		}
		tables = append(tables, t1...)
	}
	if len(filters) > 0 {
		ntable := make([]*conf.Table, 0, len(filters))
		for _, t := range tables {
			for _, c := range filters {
				if strings.Contains(t.Name, c) {
					ntable = append(ntable, t)
					break
				}
			}
		}
		tables = ntable
	}

	tmpls, err := tpl(tables)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}
	if len(tmpls) == 0 {
		cmds.Log.Errorf("未找到数据表信息%v", mdFilePath)
		return nil
	}
	if err = createFile(outPath, tmpls, cover); err != nil {
		cmds.Log.Error(err)
		return err
	}
	cmds.Log.Infof("SQL生成完成,共生成%d个文件", len(tmpls))
	return nil

}
func createFile(root string, data map[string]string, cover bool) error {
	for k, v := range data {
		path := filepath.Join(root, k)
		dir := filepath.Dir(path)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
			if err != nil {
				err = fmt.Errorf("创建文件夹%s失败:%v", path, err)
				return err
			}
		}
		if _, err := os.Stat(path); !cover && (err == nil || os.IsExist(err)) {
			cmds.Log.Warn("文件已存在:", path)
			continue
		}
		srcf, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			err = fmt.Errorf("无法打开文件:%s(err:%v)", path, err)
			return err
		}
		defer srcf.Close()
		_, err = srcf.WriteString(v)
		if err != nil {
			return err
		}
		cmds.Log.Info("生成文件:", path)
	}
	return nil

}
