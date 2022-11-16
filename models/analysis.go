package models

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type CategoryResponse struct {
	Labels []string `json:"label"`
	Data   []int64  `json:"data"`
}

type RankByCusResponse struct {
	Labels []string `json:"label"`
	Data   []int64  `json:"data"`
}

type KV struct {
	key   string
	value int64
}

func GetAllCategory() *CategoryResponse {
	var res *CategoryResponse

	res = &CategoryResponse{
		Labels: []string{
			"A", "N", "H",
		},
		Data: []int64{
			1, 2, 3,
		},
	}

	return res
}

func GetRankByCus() *RankByCusResponse {
	var res *RankByCusResponse
	var tmp []KV = make([]KV, 0)
	pFilter := filter.NewFirstKeyOnlyFilter()
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "SalesAnalysis",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	fmt.Println(item)
	for item != nil {
		i, _ := strconv.ParseInt(string(item.Cells[0].Value), 10, 64)
		fmt.Println(i)
		tmp = append(tmp, KV{
			key:   string(item.Cells[0].Row),
			value: i,
		})
		item, err = scanRsp.Next()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].value > tmp[j].value
	})

	fmt.Println(tmp)

	return res
}
