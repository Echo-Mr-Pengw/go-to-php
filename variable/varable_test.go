package variable

import (
	"testing"
)

// Pgettype使用案例
func TestPgettype(t *testing.T) {
	var i int
	var str string
	var arr [3]int
	var slice []int

	t.Log(Pgettype(i), Pgettype(str), Pgettype(arr), Pgettype(slice)) // int string array slice
}

// Pis_array使用案例
func TestPis_array(t *testing.T) {
	var i int
	var str string
	var arr [3]int
	var slice []int

	t.Log(Pis_array(i), Pis_array(str), Pis_array(arr), Pis_array(slice)) // false false true false
}

// Pis_bool使用案例
func TestPis_bool(t *testing.T) {
	var i int
	var b bool
	var str string
	var arr [3]int

	t.Log(Pis_bool(i), Pis_bool(b), Pis_bool(str), Pis_bool(arr)) // false true false false
}

// Pis_double使用案例
func TestPis_double(t *testing.T) {
	var i int = 1
	var b bool
	var str string
	var d1 float32 = 1
	var d2 float64 = 2

	t.Log(Pis_double(i), Pis_double(b), Pis_double(str), Pis_double(d1), Pis_double(d2)) // false false false true true
}

// Pis_float使用案例
func TestPis_float(t *testing.T) {
	var i int = 1
	var b bool
	var str string
	var d1 float32 = 1
	var d2 float64 = 2

	t.Log(Pis_float(i), Pis_float(b), Pis_float(str), Pis_float(d1), Pis_float(d2)) // false false false true true
}

// Pis_int使用案例
func TestPis_int(t *testing.T) {
	var i1 int
	var i2 uint
	var i3 uint8
	var f1 float64

	t.Log(Pis_int(i1), Pis_int(i2), Pis_int(i3), Pis_int(f1))  // true true true false
}

// Pis_integer使用案例同Pis_int
// Pis_long使用案例同Pis_int

// Pis_string使用案例
func TestPis_string(t *testing.T) {
	var str string = "1"
	var b bool

	t.Log(Pis_string(str), Pis_string(b)) // true false
}

// Pis_numeric
func TestPis_numeric(t *testing.T) {
	var i1 int = 12
	var i2 uint = 34
	var str1 string = "1212"
	var str2 string = "h1234"
	//var str3 string = "1.2" // 注意⚠️
	t.Log(Pis_numeric(i1), Pis_numeric(i2), Pis_numeric(str1), Pis_numeric(str2)) // true true true false
}