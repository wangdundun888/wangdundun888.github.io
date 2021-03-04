package abstractFactory

import "fmt"

/*
	抽象工厂模式一般用来生成有关联的对象,如果对象之间没有关联则退化成简单工厂模式
*/

//假设套餐中包含一个主食+一个饮料
type Meal interface {
	SetMeal()
}

type Drink interface {
	SetDrink()
}

type ItemFactory interface {
	CreateMeal() Meal
	CreateDrink() Drink
}

type MainMeal struct{}

type OtherDrink struct{}

//套餐A : 鸡肉汉堡+可乐
//func (ch *MainMeal)SetMeal(){
//	fmt.Println("one ChickenHamburger")
//}
//func (c *OtherDrink)SetDrink(){
//	fmt.Println("one Cole")
//}

//套餐B : 薯条 + 牛奶
func (ch *MainMeal) SetMeal() {
	fmt.Println("one FrenchFrice")
}
func (c *OtherDrink) SetDrink() {
	fmt.Println("one milk")
}

type A struct{}

func (a *A) CreateMeal() Meal {
	return &MainMeal{}
}

func (a A) CreateDrink() Drink {
	return &OtherDrink{}
}
