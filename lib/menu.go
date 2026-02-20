package lib

import (
	"fmt"
	"os"
)

type Menu struct {
	Id          int
	Name        string
	Descriptiom string
	Price       int
	Stock       int
}

var Cart []Menu

func mainMenuList() {
	defer fmt.Println("0. Keluar")
	fmt.Println("1. Lihat Menu")
	fmt.Println("2. Pesan Menu")
	fmt.Println("3. Lihat Keranjang")
	fmt.Println("4. Lihat Histori Order")
}

func MainMenu() {
	mainMenuList()
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearScreen()
		FoodsMenu()
		MainMenu()
	case "2":
		defer MainMenu()
		AddMenu()
		FoodsMenu()
	case "3":
		ClearScreen()
		CartFunc()
		MainMenu()
	case "4":
		fmt.Println("berhasil menekan empat")
		MainMenu()
	case "0":
		ClearScreen()
		fmt.Println("Sampai Jumpa Lagi")
		os.Exit(0)
	default:
		fmt.Println("Mohon Masukan input yang benar ya:)")
		MainMenu()
	}

}
