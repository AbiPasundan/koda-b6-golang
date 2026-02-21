package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func CartFunc() {
	if len(Cart) == 0 {
		fmt.Println("Keranjang masih kosong.")
		return
	}

	fmt.Println("Keranjang.")

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"ID",
		"NAME",
		"DESCRIPTION",
		"PRICE (Rp)",
		"STOCK",
	})
	var total int
	for _, item := range Cart {
		table.Append([]string{
			fmt.Sprintf("%d", item.Id),
			item.Name,
			item.Descriptiom,
			fmt.Sprintf("Rp %d", item.Price),
			fmt.Sprintf("%d", item.Stock),
		})
		total += item.Price
	}

	fmt.Printf("\nTOTAL: Rp %d\n", total)

	table.Render()

	Checkout()
}

func CheckoutMenu() {
	defer fmt.Println("0. Kembali")
	fmt.Println("1. Checkout")
	fmt.Println("2. Hapus Keranjang")
}

func Checkout() {
	CheckoutMenu()

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)

	switch input {
	case "0":
		MainMenu()
	case "1":
		MainMenu()
	case "2":
		MainMenu()
	default:
		ClearScreen()
		fmt.Println("Mohon Masukan input yang benar ya:)")
		MainMenu()
	}
}
