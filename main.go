package main

import (
	"github.com/PabloOvejeroML/web-server/cmd/server/controlador"
	"github.com/PabloOvejeroML/web-server/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := productos.NewRepository()
	service := productos.NewService(repo)
	product := controlador.NewProduct(service)

	router := gin.Default()

	rp := router.Group("/productos")

	rp.GET("/", product.GetAll())
	rp.GET("/:id", product.Get())
	rp.POST("/", product.Store())

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
