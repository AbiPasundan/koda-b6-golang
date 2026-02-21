package lib

import "fmt"

type History struct {
	Name        string
	Descriptiom string
	Price       int
}

var HistoryOrderProduct []History

func HistoryOrder() {
	// MenuHistoriOrder()
	var input string
	fmt.Scanln(&input)

	fmt.Println(input)
}
