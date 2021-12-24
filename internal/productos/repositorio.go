package productos

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
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) Get(id int) (Product, error) {

	for _, v := range ps {
		if id == v.Id {
			return v, nil
		}
	}

	return Product{}, nil
}

func (r *repository) LastID() (int, error) {
	lastID := 0
	if len(ps) > 0 {
		lastProduct := ps[len(ps)-1]
		lastID = lastProduct.Id
	}
	return lastID, nil
}

func (r *repository) Store(id int, nombre string, precio int, stock int, codigo string, publicado bool, fecha_creacion string) (Product, error) {
	p := Product{id, nombre, precio, stock, codigo, publicado, fecha_creacion}
	ps = append(ps, p)
	return p, nil
}
