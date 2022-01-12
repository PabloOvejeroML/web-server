package productos

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/PabloOvejeroML/web-server/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {

	input := []Product{
		{
			Id:             1,
			Nombre:         "nintendo switch lite",
			Precio:         2233,
			Stock:          121212,
			Codigo:         "string",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		}, {
			Id:             2,
			Nombre:         "PS6",
			Precio:         150000,
			Stock:          1000,
			Codigo:         "asdad1212dsda",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)

	myService := NewService(myRepo)
	// Test Execution
	resp, err := myService.GetAll()
	// Validation
	assert.Equal(t, input, resp)
	assert.Nil(t, err)

}

func TestStore(t *testing.T) {
	testProduct := Product{
		Id:             1,
		Nombre:         "nintendo switch lite",
		Precio:         2233,
		Stock:          121212,
		Codigo:         "string",
		Publicado:      true,
		Fecha_creacion: "22/12/2021",
	}
	dbMock := store.Mock{}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, _ := myService.Store(testProduct.Nombre, testProduct.Precio, testProduct.Stock, testProduct.Codigo, testProduct.Publicado, testProduct.Fecha_creacion)
	assert.Equal(t, testProduct.Nombre, result.Nombre)
	assert.Equal(t, testProduct.Codigo, result.Codigo)
	assert.Equal(t, testProduct.Precio, result.Precio)
	assert.Equal(t, 1, result.Id)
}

func TestServiceGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestServiceUpdate(t *testing.T) {

	input := []Product{
		{
			Id:             1,
			Nombre:         "nintendo switch lite",
			Precio:         2233,
			Stock:          121212,
			Codigo:         "string",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		}, {
			Id:             2,
			Nombre:         "PS6",
			Precio:         150000,
			Stock:          1000,
			Codigo:         "asdad1212dsda",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data:      dataJson,
		SpyCalled: false,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)

	myService := NewService(myRepo)

	prod, err := myService.Update(2, "PS8", 160000, 23, "jujuu", false, "año 3000")
	// Validation
	assert.Equal(t, prod.Nombre, "PS8")
	assert.Equal(t, prod.Id, 2)
	assert.Nil(t, err)
	assert.Equal(t, true, storeStub.Mock.SpyCalled)

}

func TestServiceDelete(t *testing.T) {

	input := []Product{
		{
			Id:             1,
			Nombre:         "nintendo switch lite",
			Precio:         2233,
			Stock:          121212,
			Codigo:         "string",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		}, {
			Id:             2,
			Nombre:         "PS6",
			Precio:         150000,
			Stock:          1000,
			Codigo:         "asdad1212dsda",
			Publicado:      true,
			Fecha_creacion: "22/12/2021",
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data:      dataJson,
		SpyCalled: false,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)

	myService := NewService(myRepo)

	err := myService.Delete(1)
	result, _ := myService.GetAll()
	assert.Equal(t, 1, len(result))
	assert.Nil(t, err)

}
