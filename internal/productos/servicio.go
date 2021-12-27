package productos

import (
	"strconv"
)

type Service interface {
	GetAll() ([]Product, error)
	Get(id string) (Product, error)
	Store(nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error)
	Update(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Get(id string) (Product, error) {
	//esto hacerlo en el controller
	idInt, _ := strconv.Atoi(id)

	product, err := s.repository.Get(idInt)

	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Store(nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, precio, stock, codigo, publicado, fecha_creacion)

	if err != nil {
		return Product{}, err
	}

	return producto, nil

}

func (s *service) Update(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error) {
	return s.repository.Update(id, nombre, precio, stock, codigo, publicado, fecha_creacion)

}

func (s *service) UpdateNamePrice(id int, name string, price int) (Product, error) {
	return s.repository.UpdateNamePrice(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
