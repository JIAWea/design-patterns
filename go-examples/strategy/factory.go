package strategy

import "fmt"

// AnimalFactory 抽象工厂接口
type AnimalFactory interface {
	Speak()
}

// Animal 动物
type Animal struct {
	Name   string
	Mater string
}

// SetInfo 给动物设置基本信息
func (a *Animal) SetInfo(Name, Mater string) {
	a.Name = Name
	a.Mater = Mater
}

type Dog struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Printf("%s dog speak\n", d.Name)
}

type Cat struct {
	Animal
}

func (c *Cat) Speak() {
	fmt.Printf("%s Cat speak\n", c.Name)
}
