package models

import (
	"context"
	"log"
	"sort"
	"strconv"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type DataArray []int64

type CategoryResponse struct {
	Labels []string `json:"label"`
	Data   []int64  `json:"data"`
}

type RankByCusResponse struct {
	Labels []string `json:"label"`
	Data   []int64  `json:"data"`
}

type SalesQuotaResponse struct {
	Labels  []string    `json:"label"`
	SalesID []string    `json:"salesid"`
	Data    []DataArray `json:"data"`
}

type KV struct {
	key   string
	value int64
}

func GetAllCategory() *CategoryResponse {
	var res *CategoryResponse

	res = &CategoryResponse{
		Labels: []string{
			"Price", "On Promotion", "Magazine Advertisement", "Television  Advertisement", "Manufacturer", "Review", "Demo Event", "Sponsorship", "Quality", "Other",
		},
		Data: []int64{
			52666, 71867, 56585, 53399, 53864, 55060, 54468, 53118, 53250, 54388,
		},
	}

	var dataRes []int64 = make([]int64, 0)
	for i := 1; i <= 10; i++ {
		var sum int64 = 0
		pFilter := filter.NewQualifierFilter(filter.NewCompareFilter(filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte(strconv.FormatInt(int64(i), 10))))))
		scanRequest, _ := hrpc.NewScanStr(context.Background(), "SalesAnalysis",
			hrpc.Filters(pFilter))
		scanRsp := hbaseClient.Scan(scanRequest)
		var err error
		item, err := scanRsp.Next()
		// // fmt.Println(item)
		for item != nil {
			k, _ := strconv.ParseInt(string(item.Cells[0].Value), 10, 64)
			sum += k
			item, err = scanRsp.Next()
			if err != nil {
				// fmt.Println(err)
				continue
			}
		}
		dataRes = append(dataRes, sum)
	}

	// fmt.Println(dataRes)
	res.Data = dataRes

	return res
}

func GetRankByCus() RankByCusResponse {
	var res RankByCusResponse = RankByCusResponse{}
	var tmp []KV = make([]KV, 0)
	pFilter := filter.NewFirstKeyOnlyFilter()
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "SalesAnalysis",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	// // fmt.Println(item)
	for item != nil {
		i, _ := strconv.ParseInt(string(item.Cells[0].Value), 10, 64)
		// // fmt.Println(i)
		tmp = append(tmp, KV{
			key:   string(item.Cells[0].Row),
			value: i,
		})
		item, err = scanRsp.Next()
		if err != nil {
			// fmt.Println(err)
			continue
		}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].value > tmp[j].value
	})

	// // fmt.Println(tmp)
	for i, item := range tmp {
		if i >= 10 {
			break
		}
		res.Labels = append(res.Labels, item.key)
		res.Data = append(res.Data, item.value)
	}

	return res
}

func GetRankByOrd() RankByCusResponse {
	var res RankByCusResponse = RankByCusResponse{}
	var tmp []KV = make([]KV, 0)
	pFilter := filter.NewQualifierFilter(filter.NewCompareFilter(filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte("OrderNum")))))
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "SalesAnalysis",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	// // fmt.Println(item)
	for item != nil {
		i, _ := strconv.ParseInt(string(item.Cells[0].Value), 10, 64)
		// // fmt.Println(item)
		tmp = append(tmp, KV{
			key:   string(item.Cells[0].Row),
			value: i,
		})
		item, err = scanRsp.Next()
		if err != nil {
			// fmt.Println(err)
			continue
		}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].value > tmp[j].value
	})

	// // fmt.Println(tmp)
	for i, item := range tmp {
		if i >= 10 {
			break
		}
		res.Labels = append(res.Labels, item.key)
		res.Data = append(res.Data, item.value)
	}

	return res
}

func GetCategoryBySalesPersonID(id string) *CategoryResponse {
	var res *CategoryResponse = &CategoryResponse{
		Labels: []string{
			"Price", "On Promotion", "Magazine Advertisement", "Television  Advertisement", "Manufacturer", "Review", "Demo Event", "Sponsorship", "Quality", "Other",
		},
	}
	getRequest, err := hrpc.NewGetStr(context.Background(), "SalesAnalysis", id)
	if err != nil {
		log.Fatalln(err)
	}
	// var res Customer
	getRsp, err := hbaseClient.Get(getRequest)

	var res1 map[string]int64 = make(map[string]int64)
	for _, item := range getRsp.Cells {
		if string(item.Family) == "Category" && string(item.Qualifier) != "0" {
			res1[string(item.Qualifier)], _ = strconv.ParseInt(string(item.Value), 10, 64)
		}
	}
	// log.Println(getRsp.Cells[0])
	// cell := getRsp.Cells[0]
	// res.CustomerID, _ = strconv.ParseInt(string(cell.Row), 10, 64)
	// res.PersonID = cell.
	var dataRes []int64 = make([]int64, 0)
	for i := 1; i <= 10; i++ {
		if key, ok := res1[strconv.FormatInt(int64(i), 10)]; ok {
			dataRes = append(dataRes, key)
		} else {
			dataRes = append(dataRes, 0)
		}
	}
	res.Data = dataRes
	return res
}

func GetSalesQuotaByOrd() SalesQuotaResponse {
	var res SalesQuotaResponse

	res = SalesQuotaResponse{
		Labels: []string{"2021/1", "2021/2", "2021/3", "2021/4", "2021/5", "2021/6", "2021/7", "2021/8", "2021/9", "2021/10", "2021/11", "2021/12", "2022/1", "2022/2", "2022/3", "2022/4", "2022/5", "2022/6"},
	}

	var data []DataArray = make([]DataArray, 0)

	rank10 := GetRankByOrd()
	res.SalesID = rank10.Labels
	for _, item := range rank10.Labels {
		var r DataArray = make(DataArray, 0)
		a := GetSalesQuotaByID(item)
		for _, item := range res.Labels {
			if data, ok := a[item]; ok {
				r = append(r, int64(data))
			}
		}

		data = append(data, r)
	}

	res.Data = data
	// fmt.Println(res)
	return res
}

func GetSalesQuotaByID(id string) map[string]float64 {
	var res map[string]float64 = make(map[string]float64)

	pFilter := filter.NewSingleColumnValueFilter([]byte("ID"), []byte("BusinessEntityID"), filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte(id))), true, true)
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "SalesPersonQuotaHistory",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	// // fmt.Println(item)
	for item != nil {

		f, _ := strconv.ParseFloat(string(item.Cells[2].Value), 8)
		res[string(item.Cells[1].Value)] = f

		item, err = scanRsp.Next()
		if err != nil {
			// fmt.Println(err)
			continue
		}
		// res = append(res, string(item.Cells[0].Row))
	}

	// // fmt.Println(res)

	return res
}
