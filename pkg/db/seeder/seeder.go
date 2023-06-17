package seeder

import (
	"context"

	"aramen-api/cmd/api/ent"
)

type Seeder struct {
	client *ent.Client
}

func New(client *ent.Client) *Seeder {
	return &Seeder{client: client}
}

func (s *Seeder) Seed(ctx context.Context) error {
	return nil
}
