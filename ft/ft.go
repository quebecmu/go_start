// 自定义包名
package ft

// 导入标准输入输出包
import "fmt"

// Go语言规定只有首字母大写的属性和函数才可以被外部的包进行调用
func Test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 测试属性被外部包调用
const Description = "测试自定义包导入功能"
