package str

import (
	"testing"
)

// Pimplode使用案例
func TestPimplode (t *testing.T) {
	// 	切片
	test := []string{"编", "程"}
	// 数组
	// test := [3]string{"a", "b"}
	str := Pimplode(",", test)
	t.Log(str)
}

// PJoin使用案例，同Pimplode
func TestPJoin(t *testing.T) {
	// 	切片
	test := []string{"编", "程"}
	// 数组
	// test := [3]string{"a", "b"}
	str := Pjoin(",", test)
	t.Log(str)
}

// Pexplode使用案例
func TestPexplode(t *testing.T) {
	test := "hello,world"
	// limit = 0
	slice := Pexplode(",", test, 0)
	t.Log(slice)

	// limit < 0
	slice = Pexplode(",", test, -1)
	t.Log(slice)

	// limit >0
	slice = Pexplode(",", test, 6)
	t.Log(slice)
}

// Pstrlen使用案例
func TestPstrlen(t *testing.T) {
	str := "1q测试"
	byteNum := Pstrlen(str)
	t.Log(byteNum)
}

// Pmb_strlen使用案例
func TestPmb_strlen(t *testing.T) {
	str := "1q测试"
	charNum := Pmb_strlen(str)
	t.Log(charNum)
}

// Plcfirst使用案例
func TestPlcfirst(t *testing.T) {
	str := "Atrim B23"
	lowerStr := Plcfirst(str)
	t.Log(lowerStr)
}