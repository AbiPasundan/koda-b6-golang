package lib

import (
	"fmt"
	"strconv"
)

func AddMenu() {
	fmt.Println("Masukan id untuk memasukan kedalam keranjang")
	var input string
	fmt.Scanln(&input)
	// for range FoodsMenu() {
	// 	// fmt.Println(y.Id)
	// 	// fmt.Println(y.Name)
	// 	// input, _ = strconv.ParseInt(input)

	// 	i, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	foods := FoodsMenu()

	// 	fmt.Println(foods[i].Name)
	// 	result := Menu{
	// 		Id:          foods[i].Id,
	// 		Name:        foods[i].Name,
	// 		Descriptiom: foods[i].Descriptiom,
	// 		Price:       foods[i].Price,
	// 		Stock:       foods[i].Stock,
	// 	}

	// 	Cart = append(Cart, result)
	// }

	i, err := strconv.Atoi(input)

	if err != nil {
		ClearScreen()
		fmt.Println("Input harus berupa angka")
	}

	foods := FoodsMenu()

	if i <= 0 || i >= len(foods) {
		ClearScreen()
		fmt.Println("Id tidak ditemukan")
	}

	selected := foods[i+1]

	if selected.Stock <= 0 {
		ClearScreen()
		fmt.Println("Stok habis")
	}

	Cart = append(Cart, selected)
	ClearScreen()
	fmt.Println("berhasil menambahkan ", selected.Name)

}
