package productos

import (
	"fmt"

	"github.com/PabloOvejeroML/web-server/pkg/store"
)

type Product struct {
	Id             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Precio         int    `json:"precio"`
	Stock          int    `json:"stock"`
	Codigo         string `json:"codigo"`
	Publicado      bool   `json:"publicado"`
	Fecha_creacion string `json:"fecha_creacion"`
}

var ps []Product

type Repository interface {
	GetAll() ([]Product, error)
	Get(id int) (Product, error)
	Store(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error)
	LastID() (int, error)
	Update(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error)
	UpdateNamePrice(id int, name string, price int) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) Get(id int) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for _, v := range ps {
		if id == v.Id {
			return v, nil
		}
	}

	return Product{}, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, nombre, precio, stock, codigo, publicado, fecha_creacion}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) Update(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error) {
	var ps []Product
	r.db.Read(&ps)

	p := Product{id, nombre, precio, stock, codigo, publicado, fecha_creacion}
	updated := false

	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}

	if err := r.db.Write(ps); err != nil {
		return p, err
	}
	return p, nil
}

func (r *repository) UpdateNamePrice(id int, nombre string, price int) (Product, error) {
	var ps []Product
	r.db.Read(&ps)

	var p Product
	updated := false

	for i := range ps {
		if ps[i].Id == id {
			ps[i].Nombre = nombre
			ps[i].Precio = price
			updated = true
			p = ps[i]
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}

	if err := r.db.Write(ps); err != nil {
		return p, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product
	r.db.Read(&ps)

	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			fmt.Println(ps[i].Id)
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil

}
