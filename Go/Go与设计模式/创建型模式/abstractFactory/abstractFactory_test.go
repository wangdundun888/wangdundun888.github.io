package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var a ItemFactory
	a = &A{}
	GetItem(a)
}

func GetItem(itemFactory ItemFactory) {
	meal := itemFactory.CreateMeal()
	meal.SetMeal()
	drink := itemFactory.CreateDrink()
	drink.SetDrink()
}
