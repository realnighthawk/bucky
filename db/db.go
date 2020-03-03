package db

type ProductsRepository interface {
	Get(pid string) (*Products, error)
}
type CategoryRepository interface {
	Get(cat string) (*Category, error)
}
type HeaderRepository interface {
	Get() (*Header, error)
}
