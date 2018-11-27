package tmpls

const installProdTmpl = `// +build prod

package main

import (
	'fmt'
	'os'
	'path/filepath'
	'strings'

	'github.com/micro-plat/hydra/component'
)

//getSQLPath 获取getSQLPath
func getSQLPath() (string, error) {
	gopath := os.Getenv('GOPATH')
	if gopath == '' {
		return '', fmt.Errorf('未配置环境变量GOPATH')
	}
	path := strings.Split(gopath, ';')
	if len(path) == 0 {
		return '', fmt.Errorf('环境变量GOPATH配置的路径为空')
	}
	return filepath.Join(path[0], 'src/{{.projectName}}/modules/const/sql'), nil
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *{{.projectName|lName}}) install() {
	s.IsDebug = false	

	
	//@install
	//自定义安装程序
	s.Conf.API.Installer(func(c component.IContainer) error {
		if !s.Conf.Confirm('创建数据库表结构,添加基础数据?') {
			return nil
		}
		path, err := getSQLPath()
		if err != nil {
			return err
		}
		sqls, err := s.Conf.GetSQL(path)
		if err != nil {
			return err
		}
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		for _, sql := range sqls {
			if sql != '' {
				if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
					if !strings.Contains(err.Error(), 'ORA-00942') {
						s.Conf.Log.Errorf('执行SQL失败： %v %s\n', err, q)
					}
				}
			}
		}
		return nil
	})

}`
