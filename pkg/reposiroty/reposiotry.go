package reposiroty

type Token interface {
}

type Repository struct {
	Token
}

func NewRepository() *Repository {
	return &Repository{}
}
