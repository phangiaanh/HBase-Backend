package controller

import (
	"hbase-processor/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetSalesPersonWithID(c echo.Context) error {
	id := c.Param("id")
	res, err := models.GetSalesPersonWithID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "not exist")
	}
	// return c.String(http.StatusOK, id)
	return c.JSON(http.StatusOK, res)
}

func GetSalesPersonRecommendWithID(c echo.Context) error {
	id := c.Param("id")
	res, err := models.GetSalesPersonRecommendWithID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "not exist")
	}
	// return c.String(http.StatusOK, id)
	return c.JSON(http.StatusOK, res)
}

func GetNumCustomerBySalesPersonID(c echo.Context) error {
	id := c.Param("id")
	num := models.GetNumCustomerBySalesPersonID(id)
	return c.String(http.StatusOK, strconv.FormatInt(num, 10))
}
