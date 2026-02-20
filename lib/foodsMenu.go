package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func FoodsMenu() []Menu {
	Menu := []Menu{
		{
			Id:          0,
			Name:        "",
			Descriptiom: "",
			Price:       0,
			Stock:       0,
		},
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
