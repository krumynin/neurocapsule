package repositories

type Repository struct {
	Product Product
}

func New() *Repository {
	return &Repository{
		Product: NewProduct(),
	}
}
