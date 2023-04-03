// 包声明语句
package main

// 导入包语句
import (
	"fmt"

	ft "github.com/xxandjg/go_start/ft"
	ob "github.com/xxandjg/go_start/ob"
)

// 系统入口主函数
func main() {
	fmt.Println("hello world!!!")
	ft.Test()
	fmt.Println(ft.Description)

	a := ob.Animal{Color: "white", Age: 15}
	fmt.Println(a.Call())

	// 继承属性
	b := ob.Dog{
		Animal: ob.Animal{
			Color: "red",
			Age:   18,
		}}
	fmt.Println(b.Age)
	fmt.Println(b.Color)
	// 继承方法
	fmt.Println(b.Call())

	// 多态与重写
	c := ob.Cat{Animal: ob.Animal{
		Color: "blue",
		Age:   12,
	}}
	fmt.Println(a.Run())
	fmt.Println(b.Run())
	fmt.Println(c.Run())

}
