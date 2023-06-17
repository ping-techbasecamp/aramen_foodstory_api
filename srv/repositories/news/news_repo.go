package newsRepo

import "aramen-api/cmd/api/ent"

type NewsRepoInterface interface{}

type NewsRepo struct {
	db *ent.Client
}

func NewNewsRepo(db *ent.Client) NewsRepo {
	return NewsRepo{db}
}
