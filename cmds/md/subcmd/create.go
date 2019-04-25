package subcmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/micro-plat/gaea/cmds"
	"github.com/micro-plat/gaea/cmds/util/conf"
	"github.com/micro-plat/gaea/cmds/util/data"
	"github.com/micro-plat/gaea/cmds/util/path"
	"github.com/micro-plat/lib4go/db"
	"github.com/urfave/cli"
)

//Table2MD .
type Table2MD struct {
	db       []string
	provider []string
	obj      []*db.DB
}

// NewMicServiceCmd .
func NewTable2MDCmd() cli.Command {

	m := &Table2MD{}

	return cli.Command{
		Name:  "md",
		Usage: "根据数据表结构生成 md 文件",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Usage:  "创建 md 文件",
				Flags:  m.getCreateFlags(),
				Action: m.action,
			},
		},
	}
}

//mysql  convoy:MsqlDb4567$%^&@tcp(192.168.0.36)/convoy
//oracle sso/123456@orcl136
func (m *Table2MD) getCreateFlags() []cli.Flag {
	flags := make([]cli.Flag, 0, 4)
	flags = append(flags,
		cli.StringFlag{
			Name:  "db",
			Usage: "数据库链接串",
		},
	)
	return flags
}

func (m *Table2MD) action(c *cli.Context) (err error) {

	//初始化
	err = m.init(c)
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	//根据表结构创建md
	err = m.createMD()
	if err != nil {
		cmds.Log.Error(err)
		return err
	}

	cmds.Log.Info("生成完成！")
	return nil
}

func (m *Table2MD) init(c *cli.Context) (err error) {
	if c.String("db") == "" {
		return fmt.Errorf("db 参数不能为空")
	}

	m.db = strings.Split(c.String("db"), ",")

	for _, v := range m.db {
		if strings.Contains(v, "mysql") {
			m.provider = append(m.provider, "mysql")
		} else {
			m.provider = append(m.provider, "ora")
		}
	}

	return nil
}

func (m *Table2MD) createMD() (err error) {
	err = m.getDB()
	if err != nil {
		return err
	}

	for _, v := range m.provider {
		switch v {
		case "mysql":
			err = m.createMysqlMD()
		case "ora":
			err = m.createOracleMD()
		}
	}

	return err
}

func (m *Table2MD) getDB() (err error) {
	for k, v := range m.provider {
		obj, err := db.NewDB(v, m.db[k], 20, 10, 20000)
		if err != nil {
			return err
		}
		m.obj = append(m.obj, obj)
	}
	return nil
}

func (m *Table2MD) createMysqlMD() (err error) {
	for k, v := range m.provider {
		if v != "mysql" {
			continue
		}
		datas, q, a, err := m.obj[k].Query(QueryMysql, map[string]interface{}{
			"schema": m.getSchema(k),
		})
		if err != nil {
			return fmt.Errorf("mysql(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}

		d := map[string]*conf.Table{}
		for _, v := range datas {
			if _, ok := d[v["table_name"].(string)]; !ok {
				d[v["table_name"].(string)] = conf.NewTable(v["table_name"].(string), v["table_comment"].(string), "")
			}
			d[v["table_name"].(string)].AppendColumn(
				v["column_name"].(string),
				v["column_type"].(string),
				"",
				v["column_default"].(string),
				m.getBool(v["is_nullable"]),
				v["column_key"].(string),
				v["column_comment"].(string),
			)
		}
		tbs := []*conf.Table{}
		for _, v := range d {
			tbs = append(tbs, v)
		}

		out, err := data.GetMDTmples("md", mdTPL, tbs)
		if err != nil {
			return err
		}

		f, err := path.CreatePath("./db.md", true)
		defer f.Close()
		if err != nil {
			return err
		}
		for _, v := range out {
			f.WriteString(v)
		}

	}

	return nil
}

func (m *Table2MD) createOracleMD() (err error) {

	for k, v := range m.provider {
		if v != "ora" {
			continue
		}
		tableNames, q, a, err := m.obj[k].Query(GetAllTableNameInOracle, map[string]interface{}{})
		if err != nil {
			return fmt.Errorf("oracle(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}

		d := map[string]*conf.Table{}
		for _, v := range tableNames {
			datas, q, a, err := m.obj[k].Query(GetSingleTableInfoInOracle, map[string]interface{}{
				"table_name": strings.ToUpper(v.GetString("table_name")),
			})
			if err != nil {
				return fmt.Errorf("oracle(err:%v),sql:%s,输入参数:%v,", err, q, a)
			}

			if _, ok := d[v.GetString("table_name")]; !ok {
				d[v.GetString("table_name")] = conf.NewTable(strings.ToLower(v.GetString("table_name")), datas.Get(0).GetString("table_comments"), "")
			}
			for _, v2 := range datas {
				d[v.GetString("table_name")].AppendColumn(
					strings.ToLower(v2.GetString("column_name")),
					strings.ToLower(v2.GetString("data_type")),
					v2.GetString("data_length"),
					v2.GetString("data_default"),
					m.getBool(v2.GetString("nullable")),
					"",
					v2.GetString("column_comments"),
				)
			}
		}
		tbs := []*conf.Table{}
		for _, v := range d {
			tbs = append(tbs, v)
		}

		out, err := data.GetMDTmples("md", mdOracleTPL, tbs)
		if err != nil {
			return err
		}

		f, err := path.CreatePath("./db"+strconv.Itoa(k)+".md", true)
		defer f.Close()
		if err != nil {
			return err
		}
		for _, v := range out {
			f.WriteString(v)
		}
	}

	return nil
}

func (m *Table2MD) getSchema(k int) string {
	strArray := strings.Split(m.db[k], "/")
	return strArray[len(strArray)-1]
}

func (m *Table2MD) getBool(in interface{}) bool {
	switch in.(string) {
	case "NO", "no", "N", "n":
		return false
	case "YES", "yes", "Y", "y":
		return true
	}
	return false
}
