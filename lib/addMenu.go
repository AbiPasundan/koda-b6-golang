package lib

import "fmt"

func AddMenu() {
	fmt.Println("Masukan id untuk memasukan kedalam keranjang")
	var input string
	fmt.Scanln(&input)
	for _, y := range FoodsMenu() {
		Cart = append(Cart, y)
	}
}
