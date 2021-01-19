// 封装PHP公共的处理变量的函数

package variable

import (
	"reflect"
	"strconv"
)

// 等价于PHP函数gettype()
func Pgettype(variable interface{}) string {
	return reflect.TypeOf(variable).Kind().String()
}

// 等价于PHP函数is_array()
func Pis_array(variable interface{}) bool {
	var b bool
	if Pgettype(variable) == "array" {
		b = true
	}
	return b
}

// 等价于PHP函数is_bool
func Pis_bool(variable interface{}) bool {
	var b bool
	if Pgettype(variable) == "bool" {
		b = true
	}
	return b
}

// 等价于PHP函数is_double
func Pis_double(variable interface{}) bool {
	var b bool
	variableType := Pgettype(variable)
	if variableType == "float64" || variableType == "float32" {
		b = true
	}
	return b
}

// is_double的别名
func Pis_float(variable interface{}) bool {
	return Pis_double(variable)
}

// 等价于PHP函数is_int
func Pis_int(variable interface{}) bool {
	var b bool
	variableType := Pgettype(variable)
	if variableType == "int" || variableType == "uint" || variableType == "uint64" ||
		variableType == "int64" || variableType == "uint32" || variableType == "int32" ||
		variableType == "int8" || variableType == "uint8"{
		b = true
	}
	return b
}

// is_integer是is_int的别名
func Pis_integer(variable interface{}) bool {
	return Pis_int(variable)
}

// is_long是is_int的别名
func Pis_long(variable interface{}) bool {
	return Pis_int(variable)
}

// 等价于PHP函数is_string
func Pis_string(variable interface{}) bool {
	var b bool
	if Pgettype(variable) == "string" {
		b = true
	}
	return b
}

// 等价于PHP函数is_numeric
func Pis_numeric(variable interface{}) bool {

	var b bool
	if Pis_int(variable) || Pis_float(variable) {
		b = true
	}

	if Pis_string(variable) {
		// 接口转string
		tmpStr, err := variable.(string)
		if err {
			_, err := strconv.Atoi(tmpStr)
			if err == nil {
				b = true
			}
		}
	}
	return b
}

// 等价于PHP函数empty
//func Pempty(variable interface{}) bool {
//	var b bool
//	variableType := Pgettype(variable)
//	if variableType == "string" || variable == "" {
//		b = false
//	}
//
//	if variableType == "int" || variable == 0 {
//		b = false
//	}
//
//	if variableType == "bool" || variable == false {
//		b = false
//	}
//}



