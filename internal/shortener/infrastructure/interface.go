package infrastructure

type Repository interface {
	Save(url string, shortURL string) error
	Load(shortURL string) (string, error)
}
