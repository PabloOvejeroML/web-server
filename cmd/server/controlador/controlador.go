package controlador

import (
	"errors"
	"reflect"

	"github.com/PabloOvejeroML/web-server/internal/productos"
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
		c.JSON(405, gin.H{
			"error": "el campo " + name + " es requerido",
		})
		return nil, errors.New("error")
	}

	return fv, nil
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		products, err := p.service.GetAll()

		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
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
			c.JSON(200, filtrados)
		} else {
			c.JSON(200, products)
		}

	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		prod, err := p.service.Get(id)

		if err != nil {
			c.JSON(404, gin.H{
				"error": "ID not found",
			})
			return
		}
		c.JSON(200, prod)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
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
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, prod)
	}

}
