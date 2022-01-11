package productos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubFileStore struct {
}

type spyFileStore struct {
	products  []Product
	spyCalled bool
}

func (fs *stubFileStore) Write(data interface{}) error {
	return nil
}

func (fs *stubFileStore) Read(data interface{}) error {

	products :=
		[]Product{
			{
				Id:             1,
				Nombre:         "nintendo switch lite",
				Precio:         2233,
				Stock:          121212,
				Codigo:         "string",
				Publicado:      true,
				Fecha_creacion: "22/12/2021",
			},
			{
				Id:             2,
				Nombre:         "PS6",
				Precio:         150000,
				Stock:          1000,
				Codigo:         "asdad1212dsda",
				Publicado:      true,
				Fecha_creacion: "22/12/2021",
			},
		}

	prodJson, _ := json.Marshal(products)

	json.Unmarshal(prodJson, &data)

	return nil
}

func (fs *spyFileStore) Write(data interface{}) error {
	return nil
}

func (fs *spyFileStore) Read(data interface{}) error {

	prodJson, _ := json.Marshal(fs.products)

	json.Unmarshal(prodJson, &data)

	return nil
}

func TestGetAll(t *testing.T) {

	repo := NewRepository(&stubFileStore{})

	request, err := repo.GetAll()
	assert.Equal(t, 2, len(request))
	assert.Equal(t, nil, err)

}

func TestUpdate(t *testing.T) {

	spy := &spyFileStore{[]Product{
		{
			Id:     1,
			Nombre: "Before Update",
		},
	}, false}
	repo := NewRepository(spy)

	updProd, err := repo.Update(1, "After Update", 23, 2, "sss", true, "jjjj")
	assert.Equal(t, "After Update", updProd.Nombre)
	assert.Equal(t, nil, err)

}
