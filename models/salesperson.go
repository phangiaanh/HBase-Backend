package models

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type SalesPersonResponse struct {
	Key          int64  `json:"key"`
	ID           int64  `json:"id"`
	Territory    string `json:"territory"`
	TerritoryURL string `json:"territoryURL"`
	StoreNum     int64  `json:"storeNum"`
	CustomerNum  int64  `json:"customerNum"`
	OrderNum     int64  `json:"orderNum"`
}

type SalesPersonRecommendResponse struct {
	ID   int64               `json:"id"`
	Data SalesPersonResponse `json:"data"`
}

var (
	TerritoryMap map[string]string = map[string]string{
		"1":  "Northwest-North America",
		"2":  "Northeast-North America",
		"3":  "Central-North America",
		"4":  "Southwest-North America",
		"5":  "Southeast-North America",
		"6":  "Canada",
		"7":  "France",
		"8":  "Germany",
		"9":  "Australia",
		"10": "United Kingdom",
	}

	TerritoryURL map[string]string = map[string]string{
		"1":  "https://www.discovernorthamerica.co.uk/wp-content/uploads/2018/08/beach-2737460_960_720-900x596.jpg",
		"2":  "https://a57.foxnews.com/static.foxnews.com/foxnews.com/content/uploads/2021/09/1200/675/statue-of-liberty-1.jpg?ve=1&tl=1",
		"3":  "https://www.planetware.com/photos-large/HON/honduras-copan.jpg",
		"4":  "https://www.discovernorthamerica.co.uk/wp-content/uploads/2018/08/monument-valley-1081996_960_720-900x450.jpg",
		"5":  "https://static.wixstatic.com/media/536dac_44ed5d1c16cc407dbc8fcbfa111db7e3~mv2.jpg/v1/fill/w_380,h_182,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/logo%20IMAE.jpg",
		"6":  "https://vietucnews.net/wp-content/uploads/2020/12/dinh-cu-canada-2.jpg",
		"7":  "https://www.planetware.com/wpimages/2020/02/france-in-pictures-beautiful-places-to-photograph-eiffel-tower.jpg",
		"8":  "https://media.timeout.com/images/105237824/750/422/image.jpg",
		"9":  "https://www.state.gov/wp-content/uploads/2022/02/shutterstock_1025960785-2560x1300.jpg",
		"10": "https://dynamic-media-cdn.tripadvisor.com/media/photo-o/14/10/2f/fe/united-kingdom.jpg?w=700&h=500&s=1",
	}
)

func GetSalesPersonWithID(key string) (SalesPersonResponse, error) {
	var res SalesPersonResponse = SalesPersonResponse{}
	getRequest, err := hrpc.NewGetStr(context.Background(), "SalesPerson", key)
	if err != nil {
		log.Fatalln(err)
	}
	// var res Customer
	getRsp, _ := hbaseClient.Get(getRequest)
	if len(getRsp.Cells) == 0 {
		return SalesPersonResponse{}, errors.New("not exist")
	}
	// fmt.Println(getRsp)
	i, _ := strconv.ParseInt(string(getRsp.Cells[0].Row), 10, 64)
	res.Key = 1
	res.ID = i
	for _, item := range getRsp.Cells {
		if string(item.Qualifier) == "TerritoryID" {
			res.Territory = TerritoryMap[string(item.Value)]
			res.TerritoryURL = TerritoryURL[string(item.Value)]
		}
	}

	family := map[string][]string{"Analysis": []string{"CusNum"}}
	getXRequest, _ := hrpc.NewGetStr(context.Background(), "SalesAnalysis", key,
		hrpc.Families(family))
	getXRsp, _ := hbaseClient.Get(getXRequest)

	familyY := map[string][]string{"Analysis": []string{"OrderNum"}}
	getYRequest, _ := hrpc.NewGetStr(context.Background(), "SalesAnalysis", key,
		hrpc.Families(familyY))
	getYRsp, _ := hbaseClient.Get(getYRequest)

	var storeNum int64 = 0
	pFilter := filter.NewSingleColumnValueFilter([]byte("Info"), []byte("SalesPersonID"), filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte(key))), true, true)
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "Store",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	item, err := scanRsp.Next()
	// // fmt.Println(item)
	for item != nil {
		storeNum += 1
		item, err = scanRsp.Next()
		if err != nil {
			// fmt.Println(err)
			continue
		}
		// res = append(res, string(item.Cells[0].Row))
	}

	res.StoreNum = storeNum
	res.CustomerNum, _ = strconv.ParseInt(string(getXRsp.Cells[0].Value), 10, 64)
	res.OrderNum, _ = strconv.ParseInt(string(getYRsp.Cells[0].Value), 10, 64)

	// log.Println(getRsp.Cells[0])
	// cell := getRsp.Cells[0]
	// res.CustomerID, _ = strconv.ParseInt(string(cell.Row), 10, 64)
	// res.PersonID = cell.
	return res, nil
}

func GetSalesPersonRecommendWithID(key string) (SalesPersonRecommendResponse, error) {
	var res SalesPersonResponse = SalesPersonResponse{}
	getRequest, err := hrpc.NewGetStr(context.Background(), "SalesPerson", key)
	if err != nil {
		log.Fatalln(err)
	}
	// var res Customer
	getRsp, _ := hbaseClient.Get(getRequest)
	if len(getRsp.Cells) == 0 {
		return SalesPersonRecommendResponse{}, errors.New("not exist")
	}
	// fmt.Println(getRsp)
	i, _ := strconv.ParseInt(string(getRsp.Cells[0].Row), 10, 64)
	res.Key = 1
	res.ID = i
	for _, item := range getRsp.Cells {
		if string(item.Qualifier) == "TerritoryID" {
			res.Territory = TerritoryMap[string(item.Value)]
			res.TerritoryURL = TerritoryURL[string(item.Value)]
		}
	}
	var res1 SalesPersonRecommendResponse
	res1.ID = res.ID
	res1.Data = res
	// log.Println(getRsp.Cells[0])
	// cell := getRsp.Cells[0]
	// res.CustomerID, _ = strconv.ParseInt(string(cell.Row), 10, 64)
	// res.PersonID = cell.
	return res1, nil
}

func GetNumCustomerBySalesPersonID(id string) int64 {

	// pFilter := filter.NewColumnPrefixFilter([]byte("ID")))
	pFilter := filter.NewSingleColumnValueFilter([]byte("ID"), []byte("Person"), filter.CompareType(filter.Equal), filter.NewBinaryComparator(filter.NewByteArrayComparable([]byte(id))), true, true)
	scanRequest, _ := hrpc.NewScanStr(context.Background(), "Customer",
		hrpc.Filters(pFilter))
	scanRsp := hbaseClient.Scan(scanRequest)
	var err error
	item, err := scanRsp.Next()
	// // fmt.Println(item)
	for item != nil {
		// fmt.Println(item)
		item, err = scanRsp.Next()
		if err != nil {
			// fmt.Println(err)
			continue
		}
		// res = append(res, string(item.Cells[0].Row))
	}
	return 0
}
