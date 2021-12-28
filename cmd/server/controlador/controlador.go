package controlador

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/PabloOvejeroML/web-server/internal/productos"
	"github.com/PabloOvejeroML/web-server/pkg/web"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre         string `json:"nombre"`
	Precio         int    `json:"precio"`
	Stock          int    `json:"stock"`
	Codigo         string `json:"codigo"`
	Publicado      bool   `json:"publicado"`
	Fecha_creacion string `json:"fecha_creacion"`
}

type Product struct {
	service productos.Service
}

func NewProduct(p productos.Service) *Product {
	return &Product{service: p}
}

func getField(v interface{}, name string, c *gin.Context) (interface{}, error) {
	rv := reflect.ValueOf(v)

	rv = rv.Elem()

	fv := rv.FieldByName(name)

	if fv.IsZero() {
		c.JSON(405, web.NewResponse(405, nil, "el campo "+name+" es requerido"))
		return nil, errors.New("error")
	}

	return fv, nil
}

//ListProducts godoc
//@Summary List of products
//@Tags Products
//@Description Get all products
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Success 200 {object} web.Response
//@Failure      401  {object}  web.Response
//@Failure      404  {object}  web.Response
//@Failure      500  {object}  web.Response
//@Router /productos [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		products, err := p.service.GetAll()

		if err != nil {
			c.JSON(500, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		if len(products) == 0 {
			c.JSON(404, web.NewResponse(401, nil, "No hay productos almacenados"))
			return
		}
		//esta mal que toda esta logica este aca.. lo hice para probar. Mas adelante si queres consultar por capos habria que hacer una query

		var filtrados []productos.Product

		for _, v := range products {
			if c.Query("nombre") == v.Nombre {
				filtrados = append(filtrados, v)
			}
		}

		if len(filtrados) > 0 {
			c.JSON(200, web.NewResponse(200, filtrados, ""))
		} else {
			c.JSON(200, web.NewResponse(200, products, ""))
		}

	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		prod, err := p.service.Get(id)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "ID not found"))
			return
		}

		c.JSON(200, web.NewResponse(200, prod, ""))
	}
}

//ListProducts godoc
//@Summary Store products
//@Tags Products
//@Description store products
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param product body request true "Product to store"
//@Success 200 {object} web.Response
//@Failure      401  {object}  web.Response
//@Failure      404  {object}  web.Response
//@Router /productos [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {

			c.JSON(404, web.NewResponse(401, nil, err.Error()))

			return
		}

		listaDeCamposRequeridos := []string{"Nombre", "Precio", "Stock", "Codigo", "Publicado", "Fecha_creacion"}

		for _, nombreDelCampo := range listaDeCamposRequeridos {
			_, err := getField(&req, nombreDelCampo, c)
			if err != nil {
				return
			}
		}

		prod, err := p.service.Store(req.Nombre, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha_creacion)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, prod, ""))
	}

}

//ListProducts godoc
//@Summary Update product
//@Tags Products
//@Description update product
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param id  path string true "Product ID"
//@Param product body request true "Product to update"
//@Success 200 {object} web.Response
//@Failure      400  {object}  web.Response
//@Failure      401  {object}  web.Response
//@Failure      404  {object}  web.Response
//@Router /productos/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid ID"))
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}

		listaDeCamposRequeridos := []string{"Nombre", "Precio", "Stock", "Codigo", "Publicado", "Fecha_creacion"}

		for _, nombreDelCampo := range listaDeCamposRequeridos {
			_, err := getField(&req, nombreDelCampo, c)
			if err != nil {
				return
			}
		}

		prod, err := p.service.Update(int(id), req.Nombre, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha_creacion)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, prod, ""))
	}

}

//ListProducts godoc
//@Summary Update fields
//@Tags Products
//@Description Update fields of product by id
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param id  path int true "Product ID"
//@Param product body request true "Fields of product to update (name & price)"
//@Success 200 {object} web.Response
//@Failure      400  {object}  web.Response
//@Failure      401  {object}  web.Response
//@Failure      404  {object}  web.Response
//@Router /productos/{id} [patch]
func (p *Product) UpdateFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("entré")
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid ID"))
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		listaDeCamposRequeridos := []string{"Nombre", "Precio"}

		for _, nombreDelCampo := range listaDeCamposRequeridos {
			_, err := getField(&req, nombreDelCampo, c)
			if err != nil {
				return
			}
		}

		prod, err := p.service.UpdateNamePrice(int(id), req.Nombre, req.Precio)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, prod, ""))
	}

}

//ListProducts godoc
//@Summary Delete product
//@Tags Products
//@Description Delete a product by id
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param id  path int true "Product ID"
//@Success 200 {object} web.Response
//@Failure      400  {object}  web.Response
//@Failure      401  {object}  web.Response
//@Failure      404  {object}  web.Response
//@Router /productos/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid ID"))

		}

		err = p.service.Delete(int(id))

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, nil, ""))

	}

}
