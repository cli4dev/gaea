package entity

import (
	"testing"

	"github.com/micro-plat/lib4go/ut"
)

func TestGetType(t *testing.T) {
	ut.Expect(t, fGetType("number(20)"), "int64")
	ut.Expect(t, fGetType("varchar2(32)"), "string")
	ut.Expect(t, fGetType("number(2)"), "int")
}
