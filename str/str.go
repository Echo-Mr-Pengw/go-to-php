// 通过Go封装PHP开发者常用的字符串函数

package str

import (
	"fmt"
	"reflect"
	"strings"
)

// 等价于PHP中的implde()
// implode(string $glue, array $pieces): string
// glue   字符串
// picces 数组(array)或者切片(slice)
func Pimplode (glue string, pieces interface{}) string {

	// 获取数值
	v := reflect.ValueOf(pieces)
	// 获取含有的元素个数
	len := v.Len()
	if len == 0 {
		return ""
	}

	piecesMap := map[string]string {
		"array": "array",
		"slice": "slice",
	}

	varType := fmt.Sprintf("%s", reflect.TypeOf(pieces).Kind())

	if _, exist := piecesMap[varType]; !exist {
		panic("pieces参数类型错误，只能是array或者slice")
	}

	// interface 变成切片类型
	var slice []string
	for i := 0; i < len; i++ {
		slice = append(slice, v.Index(i).String())
	}
	return strings.Join(slice, glue)
}

// 别名 Pimplode(),等价于PHP函数join()
func Pjoin (glue string, pieces interface{}) string {
	return Pimplode(glue, pieces)
}

// 等价于PHP函数explode()
// delimiter: 分割符
// str: 待分割的字符串
// limit: 分割后的总元素，limit请参考PHP explode()
func Pexplode(delimiter, str string, limit int) []string {

	if limit == 0 || limit == 1 {
		return strings.Split(str, delimiter)
	}

	if limit < 0 {
		// 负数变正数
		limit = -limit
		s := strings.Split(str, delimiter)
		count := len(s)
		if limit >= count {
			return s[0:0]
		}
		return s[0:count - limit]
	}
	return strings.SplitN(str, delimiter, limit)
}

// 等价于PHP strlen()
// 获取字符串中的总字节数
func Pstrlen(str string) int {
	return len(str)
}

// 等价于PHP mb_strlen()
// 获取字符串中的总字符数
func Pmb_strlen(str string) int {

	// s := []byte(str)
	// return utf8.RuneCount(s)
	// return utf8.RuneCountInString(str)
	return len([]rune(str))
}

// 等价于PHP函数lcfirst()
// 将字符串中首字母为英文的字母 转成小写
func Plcfirst(str string) string {

	s := []rune(str)
	if s[0] >= 65 && s[0] <= 90 {
		s[0] += 32
	}
	return string(s)
}