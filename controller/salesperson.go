package controller

import (
	"hbase-processor/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetSalesPersonWithID(c echo.Context) error {
	id := c.Param("id")
	res := models.GetSalesPersonWithID(id)
	// return c.String(http.StatusOK, id)
	return c.JSON(http.StatusOK, res.String())
}

func GetNumCustomerBySalesPersonID(c echo.Context) error {
	id := c.Param("id")
	num := models.GetNumCustomerBySalesPersonID(id)
	return c.String(http.StatusOK, strconv.FormatInt(num, 10))
}
