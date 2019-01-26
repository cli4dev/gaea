package subcmd

import (
	"testing"

	"github.com/micro-plat/lib4go/db"
)

func TestTable2MD_getSchema(t *testing.T) {
	type fields struct {
		db       string
		provider string
		obj      *db.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "1", fields: fields{db: "convoy:MsqlDb4567$%^&@tcp(192.168.0.36)/convoy", provider: "mysql", obj: nil}, want: "convoy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Table2MD{
				db:       tt.fields.db,
				provider: tt.fields.provider,
				obj:      tt.fields.obj,
			}
			if got := m.getSchema(); got != tt.want {
				t.Errorf("Table2MD.getSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
