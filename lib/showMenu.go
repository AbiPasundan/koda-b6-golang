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
	})
	// tumbal := slices.IndexFunc(Cart, func(i string) true {
	// 	return i.Cart == 0
	// })
	// // cartValue := Cart[2:]
	// fmt.Println(tumbal)
	// cartValue := Menu[1:2]
	// fmt.Println(cartValue)
	for _, m := range Cart {
		table.Append([]string{
			fmt.Sprintf("%d", m.Id),
			fmt.Sprintf("%s", m.Name),
			fmt.Sprintf("%s", m.Descriptiom),
			fmt.Sprintf("Rp %d", m.Price),
		})
	}
	table.Render()
	fmt.Println(Cart)
}
