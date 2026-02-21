package lib

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type History struct {
	Name        string
	Descriptiom string
	Price       int
}

var HistoryOrderProduct []History

func HistoryOrder() {

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"ID",
		"NAME",
		"PRICE",
		"TOTAL",
	})
	var total int
	for x, item := range HistoryOrders {
		table.Append([]string{
			fmt.Sprintf("%d", item.Id),
			fmt.Sprintf("%s", item.product[x].Name),
			fmt.Sprintf("%d", item.Total),
		})
	}

	fmt.Printf("\nTOTAL: Rp %d\n", total)

	table.Render()
}
