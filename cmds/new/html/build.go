package html

import (
	"github.com/micro-plat/gaea/cmds"

	"github.com/micro-plat/gaea/cmds/new/html/tmpls"
	"github.com/micro-plat/gaea/cmds/new/util/conf"
	"github.com/micro-plat/gaea/cmds/new/util/data"
)

//GetHTMLData .
//获取生成 vue所需的数据
func (p *htmlCmd) GetHTMLData(tables []*conf.Table, filters []string, htmlPath string) (out map[string]map[string]string, err error) {
	//获取模板数据
	tmpls, err := data.GetHTMLTmples("html", tmpls.HTMLTpl, tables, filters, htmlPath)

	if err != nil {
		cmds.Log.Error(err)
		return nil, err
	}

	if len(tmpls) == 0 {
		cmds.Log.Error("生成html时未找到数据表信息")
		return nil, err
	}
	return tmpls, nil

}
