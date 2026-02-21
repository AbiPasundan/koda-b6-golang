package lib

import "fmt"

type History struct {
	Name        string
	Descriptiom string
	Price       int
}

var HistoryOrderProduct []History

func MenuHistoriOrder() {
	defer fmt.Println("0. ")
	fmt.Println("1. ")
	fmt.Println("2. ")
	fmt.Println("3. ")
}

func HistoryOrder() {
	MenuHistoriOrder()
	fmt.Println("test")
}
