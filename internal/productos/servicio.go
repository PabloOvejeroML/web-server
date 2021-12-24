package productos

import (
	"strconv"
)

type Service interface {
	GetAll() ([]Product, error)
	Get(id string) (Product, error)
	Store(nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error)
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
