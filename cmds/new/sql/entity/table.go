package entity

type Table struct {
	Name    string
	Desc    string
	CNames  []string
	Lens    []string
	Types   []string
	Defs    []string
	IsNulls []bool
	Cons    []string
	Descs   []string
}

func NewTable(name string, desc string) *Table {
	return &Table{
		Name:    name,
		Desc:    desc,
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
