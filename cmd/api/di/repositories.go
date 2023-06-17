package di

import (
	"aramen-api/cmd/api/ent"
	"aramen-api/srv/repositories"
	authRepo "aramen-api/srv/repositories/auth"
	branchRepo "aramen-api/srv/repositories/branch"
	newsRepo "aramen-api/srv/repositories/news"
	qrcodeRepo "aramen-api/srv/repositories/qrcode"
)

func newRepositories(db *ent.Client) *repositories.Repositories {
	return &repositories.Repositories{
		AuthRepo:   authRepo.NewAuthRepo(db),
		BranchRepo: branchRepo.NewBranchRepo(db),
		NewsRepo:   newsRepo.NewNewsRepo(db),
		QrcodeRepo: qrcodeRepo.NewQrcodeRepo(db),
	}
}
