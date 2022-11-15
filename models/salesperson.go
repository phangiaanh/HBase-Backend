package models

import (
	"context"
	"fmt"
	"log"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

func GetSalesPersonWithID(key string) *hrpc.Result {
	getRequest, err := hrpc.NewGetStr(context.Background(), "SalesPerson", key)
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

func GetNumCustomerBySalesPersonID(id string) int64 {

	// pFilter := filter.NewColumnPrefixFilter([]byte("ID")))
	pFilter := filter.NewSingleColumnValueFilter([]byte("ID"), []byte("Person"), filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte(id))), true, true)
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "Customer",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	// fmt.Println(item)
	for item != nil {
		fmt.Println(item)
		item, err = scanRsp.Next()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// res = append(res, string(item.Cells[0].Row))
	}
	return 0
}
