package branchRepo

import (
	"aramen-api/cmd/api/ent"
	"aramen-api/cmd/api/ent/branches"
	"context"
)

type BranchRepoInterface interface {
	GetBranches() ([]*ent.Branches, error)
	GetBranchDetail() (*ent.Branches, error)
}

type BranchRepo struct {
	db *ent.Client
}

func NewBranchRepo(db *ent.Client) BranchRepo {
	return BranchRepo{db}
}

func (r BranchRepo) GetBranches() ([]*ent.Branches, error) {
	return r.db.Branches.Query().All(context.Background())
}

func (r BranchRepo) GetBranchDetail(branchId int) (*ent.Branches, error) {
	return r.db.Branches.Query().Where(branches.IDEQ(branchId)).First(context.Background())
}
