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

func FoodsMenu() {
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
	// fmt.Println(Menu)
	// fmt.Printf("%#v\n", Menu)
	// for x := range Menu {
	// 	fmt.Printf(
	// 		"Name  : %s\nId : %s\nAge   : %d\n", Menu[x].Id, Menu[x].Name, Menu[x].Descriptiom,
	// 	)
	// }

	// w := tabwriter.NewWriter(&io.OffsetWriter{}.Stdout, 0, 0, 2, ' ', 0)
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION")

	// for _, u := range Menu {
	// 	fmt.Fprintf(w, "%s\t%s\t%d\n", u.Id, u.Name, u.Descriptiom)
	// }
	// w.Flush()

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

}
