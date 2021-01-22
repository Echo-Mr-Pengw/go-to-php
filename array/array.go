// 封装PHP常用的数组

package array

import "strings"

// 等价于PHP array_change_key_case
func Parray_change_key_case(m map[string]interface{}) (map[string]interface{}){
	tmp := make(map[string]interface{})
	for i, v := range m {
		tmp[strings.ToUpper(i)] = v
	}
	return tmp
}
