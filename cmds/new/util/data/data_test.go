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
