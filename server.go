package main

import (
	"hbase-processor/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/salesperson/:id", controller.GetSalesPersonWithID)
	e.GET("/salesperson/numcus/:id", controller.GetNumCustomerBySalesPersonID)

	e.GET("/customer/all", controller.GetAllCustomers)
	e.GET("/customer/:id", controller.GetCustomerWithID)
	// e.GET("/salesperson/:id", controller.GetUser)

	e.Logger.Fatal(e.Start(":1323"))

}
