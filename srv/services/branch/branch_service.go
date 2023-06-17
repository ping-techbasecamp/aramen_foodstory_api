package branchService

import (
	"aramen-api/cmd/api/ent"
	"aramen-api/pkg/config"
	branchRepo "aramen-api/srv/repositories/branch"
)

type BranchServiceInterface interface {
	GetBranches() ([]*ent.Branches, error)
	GetBranchDetail() (*ent.Branches, error)
}

type BranchService struct {
	cfg        *config.AppConfig
	branchRepo branchRepo.BranchRepo
}

func NewService(cfg *config.AppConfig, branchRepo branchRepo.BranchRepo) BranchService {
	return BranchService{cfg: cfg, branchRepo: branchRepo}
}

func (s BranchService) GetBranches() ([]*ent.Branches, error) {
	return s.branchRepo.GetBranches()
}

func (s BranchService) GetBranchDetail(branchId int) (*ent.Branches, error) {
	return s.branchRepo.GetBranchDetail(branchId)
}
