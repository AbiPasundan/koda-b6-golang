package lib

import (
	"fmt"
	"strconv"
)

func AddMenu() {
	fmt.Println("Masukan id untuk memasukan kedalam keranjang")
	var input string
	fmt.Scanln(&input)
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
