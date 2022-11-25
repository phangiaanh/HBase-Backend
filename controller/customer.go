package controller

import (
	"hbase-processor/models"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func GetAllCustomers(c echo.Context) error {
	res := models.GetAllCustomers()

	return c.String(http.StatusOK, strings.Join(res, " "))
	// return c.JSON(http.StatusOK, JSON.stringt)
}

func GetCustomerWithID(c echo.Context) error {
	// User ID from path `users/:id`
	x := time.Now().UnixNano()
	id := c.Param("id")
	a := models.GetCustomerWithKey(id)
	// fmt.Println(time.Now().UnixNano() - x)
	return c.String(http.StatusOK, a.String())
}
