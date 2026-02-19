package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

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
