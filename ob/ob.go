package ob

// Animal 父类
type Animal struct {
	Color string
	Age   int
}

// Dog 实现继承
type Dog struct {
	Animal
}

// Call 面向对象方法
func (a *Animal) Call() string {
	return `动物叫！！！`
}

// Cat 实现继承
type Cat struct {
	Animal
}

type Move interface {
	Run() string
}

func (c Cat) Run() string {
	return "猫爬"
}

func (d Dog) Run() string {
	return "狗刨"
}

func (a *Animal) Run() string {
	return "动物移动"
}
