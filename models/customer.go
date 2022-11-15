package models

import (
	"context"
	"fmt"
	"log"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type Customer struct {
	CustomerID    int64  `json:"CustomerID"`
	PersonID      int64  `json:"PersonID"`
	StoreID       int64  `json:"StoreID"`
	TerritoryID   int64  `json:"TerritoryID"`
	AccountNumber string `json:"AccountNumber"`
}

func GetAllCustomers() []string {
	var res []string
	pFilter := filter.NewFirstKeyOnlyFilter()
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "Customer",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	fmt.Println(item)
	for item != nil {
		// fmt.Println(string(item.Cells[0].Row))
		item, err = scanRsp.Next()
		if err != nil {
			fmt.Println(err)
			continue
		}
		res = append(res, string(item.Cells[0].Row))
	}
	return res

}

func GetCustomerWithKey(key string) *hrpc.Result {
	getRequest, err := hrpc.NewGetStr(context.Background(), "Customer", key)
	if err != nil {
		log.Fatalln(err)
	}
	// var res Customer
	getRsp, err := hbaseClient.Get(getRequest)
	// log.Println(getRsp.Cells[0])
	// cell := getRsp.Cells[0]
	// res.CustomerID, _ = strconv.ParseInt(string(cell.Row), 10, 64)
	// res.PersonID = cell.
	return getRsp
}
