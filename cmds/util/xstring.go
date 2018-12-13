package util

import (
	"fmt"
	"strings"
)

//GetPrefixString .
func GetPrefixString(str string, s string) string {
	if strings.HasPrefix(str, s) {
		return str
	}
	return fmt.Sprintf("%s%s", s, str)
}

//GetLeftString .
func GetLeftString(str string, s string, def ...string) string {
	index := strings.LastIndex(str, s)
	if index >= 0 {
		return str[:index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

//GetRightString .
func GetRightString(str string, s string, def ...string) string {
	index := strings.LastIndex(str, s)
	if index >= 0 {
		return str[index+1:]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
