// PHP输出相关函数

package output

import "fmt"

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