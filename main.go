package main

import (
	"log"

	"github.com/PabloOvejeroML/web-server/cmd/server/controlador"
	"github.com/PabloOvejeroML/web-server/internal/productos"
	"github.com/PabloOvejeroML/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no se pudo")
	}
	db := store.New(store.FileType, "./products.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	product := controlador.NewProduct(service)

	router := gin.Default()

	rp := router.Group("/productos")

	rp.GET("/", product.GetAll())
	rp.GET("/:id", product.Get())
	rp.POST("/", product.Store())
	rp.PUT("/:id", product.Update())
	rp.DELETE("/:id", product.Delete())
	rp.PATCH("/:id", product.UpdateFields())
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
