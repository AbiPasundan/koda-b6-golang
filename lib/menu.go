package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
		MainMenu()
	case "2":
		AddMenu()
		FoodsMenu()
		MainMenu()
	case "3":
		fmt.Println("berhasil menekan tiga")
		ShowMenu()
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

func FoodsMenu() []Menu {
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
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"ID",
		"NAME",
		"DESCRIPTION",
		"PRICE",
		"STOCK",
	})

	for _, m := range Menu {
		table.Append([]string{
			fmt.Sprintf("%d", m.Id),
			m.Name,
			m.Descriptiom,
			fmt.Sprintf("Rp %d", m.Price),
			fmt.Sprintf("%d", m.Stock),
		})
	}
	table.Render()
	return Menu
}

func AddMenu() {
	fmt.Println("Masukan id untuk memasukan kedalam keranjang")
	var input string
	fmt.Scanln(&input)
	for _, y := range FoodsMenu() {
		Cart = append(Cart, y)
	}
}

func ShowMenu() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"NAME",
		"DESCRIPTION",
		"PRICE",
		"STOCK",
	})
	fmt.Println("Masukan id untuk memasukan kedalam keranjang")
	for _, m := range Cart {
		table.Append([]string{
			fmt.Sprintf("%d", m.Id),
			m.Name,
			m.Descriptiom,
			fmt.Sprintf("Rp %d", m.Price),
			fmt.Sprintf("%d", m.Stock),
		})
	}
	table.Render()
	fmt.Println(Cart)
}
