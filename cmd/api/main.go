package main

import (
	"net/http"

	"github.com/Jesusdasilva2001/learning-Golang/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// chi
	/*r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", Order)
	http.ListenAndServe(":8880", r)*/

	// echo server http
	e := echo.New()
	e.GET("/order", Order) 
	e.Logger.Fatal(e.Start(":8880"))
}

// order endpoint for echo framework
func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
