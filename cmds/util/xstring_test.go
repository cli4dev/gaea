package util

import (
	"testing"
)

func TestGetLeftString(t *testing.T) {
	type args struct {
		str string
		s   string
		def []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{str: "ora:sso/123456@orcl136", s: ":", def: []string{"mysql"}}, want: "ora"},
		{name: "2", args: args{str: "mysql:vts:vts45678@tcp(192.168.0.36)/vts", s: ":", def: []string{"mysql"}}, want: "mysql"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLeftString(tt.args.str, tt.args.s, tt.args.def...); got != tt.want {
				t.Errorf("GetLeftString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRightString(t *testing.T) {
	type args struct {
		str string
		s   string
		def []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "1", args: args{str: "ora:sso/123456@orcl136", s: ":", def: []string{"mysql"}}, want: "sso/123456@orcl136"},
		{name: "1", args: args{str: "mysql:vts:vts45678@tcp(192.168.0.36)/vts", s: ":", def: []string{"mysql"}}, want: "vts:vts45678@tcp(192.168.0.36)/vts"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRightString(tt.args.str, tt.args.s, tt.args.def...); got != tt.want {
				t.Errorf("GetRightString() = %v, want %v", got, tt.want)
			}
		})
	}
}
