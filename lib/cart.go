package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var HistoryOrders []Orders

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

func CheckoutRemove() {
	// empty := []Menu{}
	// fmt.Println("empty")
	// fmt.Println(empty)
	Cart = []Menu{}
}

type Orders struct {
	Id      int
	product []Menu
}

// func CheckoutProduct() {
// 	// HistoryOrderProduct = []History{
// 	// 	{
// 	// 		Name: "Nama",
// 	// 	},
// 	// }

// 	// HistoryOrders := []Orders{
// 	// }

// 	result := []Orders{
// 		{
// 			Id:      1,
// 			product: Cart,
// 		},
// 	}

// 	HistoryOrders = append(HistoryOrders, result...)
// }

func Checkout() {
	CheckoutMenu()

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)

	switch input {
	case "0":
		MainMenu()
	case "1":
		result := []Orders{
			{
				Id:      1,
				product: Cart,
			},
		}

		HistoryOrders = append(HistoryOrders, result...)
		fmt.Println("Berhasil Checkout")
		CheckoutRemove()
		MainMenu()
	case "2":
		CheckoutRemove()
		ClearScreen()
		fmt.Println("Keranjang Berhasil Dikosongkan")
		MainMenu()
	default:
		ClearScreen()
		fmt.Println("Mohon Masukan input yang benar ya:)")
		MainMenu()
	}
}
