package conf

type Table struct {
	Name    string   //表名
	Desc    string   //表描述
	DBLink  string   //dblink
	CNames  []string //字段名
	Lens    []string //长度
	Types   []string //类型
	Defs    []string //默认值
	IsNulls []bool   //为空
	Cons    []string //约束
	Descs   []string //描述
}

func NewTable(name, desc, dblink string) *Table {
	return &Table{
		Name:    name,
		Desc:    desc,
		DBLink:  dblink,
		CNames:  make([]string, 0, 4),
		Lens:    make([]string, 0, 4),
		Types:   make([]string, 0, 4),
		Defs:    make([]string, 0, 4),
		IsNulls: make([]bool, 0, 4),
		Cons:    make([]string, 0, 4),
		Descs:   make([]string, 0, 4),
	}
}
func (t *Table) AppendColumn(name string, tp string, len string, def string, isNull bool, cons string, des string) error {
	t.CNames = append(t.CNames, name)
	t.Lens = append(t.Lens, len)
	t.Types = append(t.Types, tp)
	t.Defs = append(t.Defs, def)
	t.IsNulls = append(t.IsNulls, isNull)
	t.Cons = append(t.Cons, cons)
	t.Descs = append(t.Descs, des)
	return nil
}
