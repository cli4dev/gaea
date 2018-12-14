package data

import (
	"fmt"
	"testing"
)

func TestTGetType(t *testing.T) {
	fmt.Println(fGetType("number(11)"))
}

func TestFGetAName(t *testing.T) {
	fmt.Println(fGetAName("sp_id"))
}
func TestFGetPName(t *testing.T) {
	fmt.Println(fGetPName("sys_supplier_info"))
}

func Test_getPath(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{name: "sys_supplier_info"}, want: "/sys/supplier/info"},
		{name: "2", args: args{name: "supplier_info"}, want: "/supplier/info"},
		{name: "3", args: args{name: "info"}, want: "/info"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPath(tt.args.name); got != tt.want {
				t.Errorf("getPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHandleName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{name: "sys_supplier_info"}, want: ""},
		{name: "2", args: args{name: "supplier_info"}, want: ""},
		{name: "3", args: args{name: "info"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHandleName(tt.args.name); got != tt.want {
				t.Errorf("GetHandleName() = %v, want %v", got, tt.want)
			}
		})
	}
}