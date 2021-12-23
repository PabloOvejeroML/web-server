package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Precio         int    `json:"precio"`
	Stock          int    `json:"stock"`
	Codigo         string `json:"codigo"`
	Publicado      bool   `json:"publicado"`
	Fecha_creacion string `json:"fecha_creacion"`
}

var products []Product

func inicializarProducts() {
	file, _ := ioutil.ReadFile("products.json")

	//hmmm.. concatenar con los de memoria
	prods := Products{}

	_ = json.Unmarshal([]byte(file), &prods)

	for _, v := range prods.Products {
		products = append(products, v)
	}

}

func handlerProduct(c *gin.Context) {

	var filtrados []*Product
	//break porque el id deberia ser unico entiendo..
	for _, v := range products {
		if c.Param("id") == strconv.FormatInt(int64(v.Id), 10) {
			filtrados = append(filtrados, &v)
			break
		}
	}

	if filtrados != nil {
		c.JSON(200, filtrados)
	} else {
		c.JSON(404, "Error: ID not found")
	}

}

func handlerProducts(c *gin.Context) {

	//	out, _ := json.MarshalIndent(prods, "", " ")

	//	response, _ := json.Marshal(out)

	//	fmt.Println(string(out))

	//me parece que te parsea la estructura a json

	//	queries := c.Request.URL.Query()

	var filtrados []*Product

	//por las dudas: los query params primero ? y despues se separan por &
	//Mi solucion actual contempla el caso que te llegue un query param de cualquier cosa, pero no mas de uno
	for _, v := range products {
		if c.Query("id") == strconv.FormatInt(int64(v.Id), 10) {
			filtrados = append(filtrados, &v)
			break
		}
		if c.Query("nombre") == v.Nombre {
			filtrados = append(filtrados, &v)
			break
		}
		if c.Query("precio") == strconv.FormatInt(int64(v.Precio), 10) {
			filtrados = append(filtrados, &v)
			break
		}
		if c.Query("stock") == strconv.FormatInt(int64(v.Stock), 10) {
			filtrados = append(filtrados, &v)
			break
		}
		if c.Query("codigo") == v.Codigo {
			filtrados = append(filtrados, &v)
			break
		}
		if c.Query("publicado") == strconv.FormatBool(v.Publicado) {
			filtrados = append(filtrados, &v)
			continue
		}
		if c.Query("fecha_creacion") == v.Fecha_creacion {
			filtrados = append(filtrados, &v)
			break
		}

	}
	//header := c.Request.Header

	if filtrados != nil {
		c.JSON(200, filtrados)
	} else {
		c.JSON(200, products)
	}

}

func validarToken(c *gin.Context) error {
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(401, gin.H{
			"error": "token invalido",
		})
		return errors.New("error")
	}
	return nil
}

func getField(v interface{}, name string, c *gin.Context) (interface{}, error) {
	rv := reflect.ValueOf(v)

	rv = rv.Elem()

	fv := rv.FieldByName(name)

	if fv.IsZero() {
		c.JSON(405, gin.H{
			"error": "el campo " + name + " es requerido",
		})
		return nil, errors.New("error")
	}

	return fv, nil
}

func validarCampos(p Product, c *gin.Context) error {

	listaDeCamposRequeridos := []string{"Id", "Nombre", "Precio", "Stock", "Codigo", "Publicado", "Fecha_creacion"}

	for _, nombreDelCampo := range listaDeCamposRequeridos {
		_, err := getField(&p, nombreDelCampo, c)
		if err != nil {
			return errors.New("error")
		}
	}
	return nil
}

// `json:"name" binding:"required"`
func handlerSaveProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if validarToken(c) != nil {
			return
		}
		var prod Product
		if err := c.ShouldBindJSON(&prod); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastProduct := products[len(products)-1]
		prod.Id = lastProduct.Id + 1
		err := validarCampos(prod, c)
		if err != nil {
			return
		}
		products = append(products, prod)
		c.JSON(200, prod)
	}
}

func main() {
	router := gin.Default()
	inicializarProducts()

	router.GET("/message/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hola " + name,
		})
	})

	rp := router.Group("/productos")

	rp.GET("/", handlerProducts)
	rp.GET("/:id", handlerProduct)
	rp.POST("/", handlerSaveProduct())

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
