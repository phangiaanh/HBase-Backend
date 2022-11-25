package controller

import (
	"hbase-processor/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// func GetTop10SalesPersonByCustomer(c echo.Context) {
// 	// f := filter.NewColumnCountGetFilter()
// }

func GetAllCategory(c echo.Context) error {
	a := models.GetAllCategory()
	// j, _ := json.Marshal(a)
	return c.JSON(http.StatusOK, a)
}

func GetCategoryByID(c echo.Context) error {
	id := c.Param("id")
	a := models.GetCategoryBySalesPersonID(id)
	return c.JSON(http.StatusOK, a)
}

func GetRankByCus(c echo.Context) error {
	x := time.Now().UnixNano()
	a := models.GetRankByCus()
	// fmt.Println(time.Now().UnixNano() - x)
	return c.JSON(http.StatusOK, a)
}

func GetRankByOrder(c echo.Context) error {
	x := time.Now().UnixNano()
	a := models.GetRankByOrd()
	// fmt.Println(time.Now().UnixNano() - x)
	return c.JSON(http.StatusOK, a)
}

func GetQuotaByOrd(c echo.Context) error {
	res := models.GetSalesQuotaByOrd()

	return c.JSON(http.StatusOK, res)
}

func GetQuotaByOrdAndSalesID(c echo.Context) error {

	var res models.SalesQuotaResponse
	var data []models.DataArray = make([]models.DataArray, 0)
	res = models.SalesQuotaResponse{
		Labels: []string{"2021/1", "2021/2", "2021/3", "2021/4", "2021/5", "2021/6", "2021/7", "2021/8", "2021/9", "2021/10", "2021/11", "2021/12", "2022/1", "2022/2", "2022/3", "2022/4", "2022/5", "2022/6"},
	}
	id := c.Param("id")
	a := models.GetSalesQuotaByID(id)
	var r models.DataArray = make(models.DataArray, 0)
	for _, item := range res.Labels {
		if data, ok := a[item]; ok {
			r = append(r, int64(data))
		}

	}
	data = append(data, r)
	res.Data = data
	res.SalesID = append(res.SalesID, id)
	return c.JSON(http.StatusOK, res)
}
