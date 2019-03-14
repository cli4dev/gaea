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
			if got := GetRouterPath(tt.args.name); got != tt.want {
				t.Errorf("getPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAppconf(t *testing.T) {
	type args struct {
		str   string
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{str: "http://sso.100bm.cn:8089|en01kslkl233l|coupon", index: 1}, want: ""},
		{name: "2", args: args{str: "http://sso.100bm.cn:8089|en01kslkl233l|coupon", index: 2}, want: ""},
		{name: "3", args: args{str: "http://sso.100bm.cn:8089|en01kslkl233l", index: 3}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAppconf(tt.args.str, tt.args.index); got != tt.want {
				t.Errorf("getAppconf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSource(t *testing.T) {

	got := getSource("atv_id", "IQS,DD(atv_activity_info)")
	fmt.Println(got)
}

func Test_fGetLastName(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{n: "user_base_info"}, want: "info"},
		{name: "2", args: args{n: "/user/base/info"}, want: "info"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fGetLastName(tt.args.n); got != tt.want {
				t.Errorf("fGetLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}
