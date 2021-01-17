// 通过Go封装PHP开发者常用的字符串函数

package str

import (
	"crypto/md5"
	"crypto/sha1"
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

	if len(str) == 0 {
		return ""
	}

	s := []rune(str)
	if s[0] >= 65 && s[0] <= 90 {
		s[0] += 32
	}
	return string(s)
}

// 等价于PHP函数ucfirst()
// 将字符串中首字母为英文的字母 转成大写
func Pucfirst(str string) string {

	if len(str) == 0 {
		return ""
	}

	s := []rune(str)
	if s[0] >= 97 && s[0] <= 122 {
		s[0] -= 32
	}
	return string(s)
}

// 等价于PHP函数strtoupper()
// 将字符串变成大写
func Pstrtoupper(str string) string {
	return strings.ToLower(str)
}

// 等价于PHP函数strtolower()
// 将字符串变成小写
func Pstrtolower(str string) string {
	return strings.ToUpper(str)
}

// 等价于PHP函数ucfirst()
func Pucword(str string) string {
	return strings.Title(str)
}

// 等价于PHP函数trim()
func Ptrim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.Trim(str, character_mask)
}

// 等价于PHP函数ltrim()
func Pltrim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.TrimLeft(str, character_mask)
}

// 等价于PHP函数rtrim()
func Prtrim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.TrimRight(str, character_mask)
}

// 等价于PHP函数chop()
func Pchop(str, character_mask string) string {
	return Prtrim(str, character_mask)
}

// 等价于PHP函数md5()
func Pmd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// 等价于PHP函数sha1()
func Psha1(str string) string {
	return (fmt.Sprintf("%x", sha1.Sum([]byte(str))))
}

// 等价于PHP函数ord
func Pord(str string) int {

	if str == "" {
		return 0
	}

	s := []rune(str)
	return int(s[0])
}

// ord函数的升级版，转化全部字符
func Pallord(str string) interface{}{
	return []rune(str)
}

// 等价于PHP函数chr
func Pchr(ascii int32) string {
	return string(ascii)
}

// 等价于PHP函数echo
func Pecho(a ...interface{}) {
	fmt.Print(a...)
}

// 等价于PHP函数var_dump()
func Pvar_dump(a ...interface{}) {
	for _, v := range a {
		fmt.Printf("%T ", v)
		fmt.Println(v)
	}
}

// 等价于PHP函数print()
func Pprint(a interface{}) {
	fmt.Println(a)
}

// 等价于PHP函数str_repeat()
func Pstr_repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

// 等价于PHP函数str_shuffle()
//func Pstr_shuffle(str string) string {
//
//	l := len(str)
//	if l < 0 {
//		return ""
//	}
//}

// 等价于PHP函数str_split
//func Pstr_split(str string, splitLen int) []string {
//	var s []string
//	var l = len(str)
//
//	if l == 0 || splitLen == 0{
//		return s
//	}
//
//	r := []rune(str)
//	if splitLen == 1 || splitLen >= l {
//		return r
//	}
//
//	for _, v := range r {
//		s = append(s, string(v))
//	}
//	return s
//}