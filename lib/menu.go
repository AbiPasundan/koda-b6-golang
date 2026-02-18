package lib

import "fmt"

type Menu struct {
	Id          int
	Name        string
	Descriptiom string
	Price       int
	Stock       int
}

func FoodsMenu() {
	Menu := []Menu{
		{
			Id:          1,
			Name:        "Sensasi 1",
			Descriptiom: "Spaghetti Ayam + Drink",
			Price:       10000,
			Stock:       13,
		},
		{
			Id:          2,
			Name:        "California Burger",
			Descriptiom: "California Burger",
			Price:       20000,
			Stock:       15,
		},
		{
			Id:          3,
			Name:        "French Fries Large",
			Descriptiom: "French Fries Large",
			Price:       15000,
			Stock:       20,
		},
		{
			Id:          4,
			Name:        "Coca-Cola",
			Descriptiom: "Coca-Cola pet 250ml",
			Price:       5000,
			Stock:       30,
		},
	}
	fmt.Println(Menu)
}
