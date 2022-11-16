package controller

import (
	"hbase-processor/models"
	"net/http"

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

	return c.JSON(http.StatusOK, nil)
}

func GetRankByCus(c echo.Context) error {
	a := models.GetRankByCus()
	return c.JSON(http.StatusOK, a)
}
