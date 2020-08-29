package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tonfada202/gofinal/driver"
	cus "github.com/tonfada202/gofinal/handler/http"
	"github.com/tonfada202/gofinal/middleware"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Auth)
	connection := driver.ConnectSQL()
	hCus := cus.Handler{
		DB: connection,
	}
	r.POST("/customers", hCus.SaveCusHandler)
	r.GET("/customers", hCus.GetAllCusHandler)
	r.GET("/customers/:id", hCus.GetCusByIdHandler)
	r.PUT("/customers/:id", hCus.UpdateCusByIdHandler)
	r.DELETE("/customers/:id", hCus.DelCusByIdHandler)

	return r
}

func main() {
	driver.CreateTable()
	r := setupRouter()
	r.Run(":2009")
}
