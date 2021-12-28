package main

import (
	"log"
	"os"

	"github.com/PabloOvejeroML/web-server/cmd/server/controlador"
	"github.com/PabloOvejeroML/web-server/docs"
	"github.com/PabloOvejeroML/web-server/internal/productos"
	"github.com/PabloOvejeroML/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title PRODUCTS
//@version 1.0
//@description La API permite crear, leer, actualizar y borrar productos (CRUD)
//@termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name PabloOvejeroML

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rp := router.Group("/productos")

	rp.GET("/", product.GetAll())
	rp.GET("/:id", product.Get())
	rp.POST("/", product.Store())
	rp.PUT("/:id", product.Update())
	rp.DELETE("/:id", product.Delete())
	rp.PATCH("/:id", product.UpdateFields())
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
