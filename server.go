package main

import (
	"hbase-processor/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/salesperson/:id", controller.GetSalesPersonWithID)
	e.GET("/salesperson/numcus/:id", controller.GetNumCustomerBySalesPersonID)

	e.GET("/customer/all", controller.GetAllCustomers)
	e.GET("/customer/:id", controller.GetCustomerWithID)
	// e.GET("/salesperson/:id", controller.GetUser)

	e.GET("/analysis/category/all", controller.GetAllCategory)
	e.GET("/analysis/category/:id", controller.GetCategoryByID)
	e.GET("/analysis/rank/cus/", controller.GetRankByCus)

	e.Logger.Fatal(e.Start(":1323"))

}
